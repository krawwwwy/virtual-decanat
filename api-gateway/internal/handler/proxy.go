package handler

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// ServiceConfig содержит конфигурацию для сервиса
type ServiceConfig struct {
	URL     string
	Timeout time.Duration
}

// ProxyHandler обрабатывает проксирование запросов к микросервисам
type ProxyHandler struct {
	logger          *zap.Logger
	serviceConfigs  map[string]ServiceConfig
	defaultTimeout  time.Duration
}

// NewProxyHandler создает новый обработчик прокси
func NewProxyHandler(logger *zap.Logger) *ProxyHandler {
	serviceConfigs := make(map[string]ServiceConfig)
	
	// Загрузка конфигурации сервисов из Viper
	for service, config := range viper.GetStringMap("services") {
		serviceConfig := config.(map[string]interface{})
		serviceURL := serviceConfig["url"].(string)
		timeout := viper.GetDuration("services." + service + ".timeout")
		
		serviceConfigs[service] = ServiceConfig{
			URL:     serviceURL,
			Timeout: timeout,
		}
	}
	
	defaultTimeout := viper.GetDuration("server.timeout")
	
	return &ProxyHandler{
		logger:          logger,
		serviceConfigs:  serviceConfigs,
		defaultTimeout:  defaultTimeout,
	}
}

// ProxyRequest проксирует запрос к указанному сервису
func (h *ProxyHandler) ProxyRequest(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceConfig, exists := h.serviceConfigs[serviceName]
		if !exists {
			h.logger.Error("Service configuration not found", zap.String("service", serviceName))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Service configuration not found"})
			return
		}

		// Создаем URL для проксирования
		targetURL, err := url.Parse(serviceConfig.URL)
		if err != nil {
			h.logger.Error("Failed to parse service URL", zap.String("url", serviceConfig.URL), zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid service URL"})
			return
		}
		
		targetURL.Path = c.Request.URL.Path
		targetURL.RawQuery = c.Request.URL.RawQuery

		// Читаем тело запроса
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, err = io.ReadAll(c.Request.Body)
			if err != nil {
				h.logger.Error("Failed to read request body", zap.Error(err))
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
				return
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// Создаем новый запрос
		req, err := http.NewRequest(c.Request.Method, targetURL.String(), bytes.NewBuffer(bodyBytes))
		if err != nil {
			h.logger.Error("Failed to create proxy request", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create proxy request"})
			return
		}

		// Копируем заголовки
		for key, values := range c.Request.Header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}

		// Устанавливаем таймаут
		timeout := serviceConfig.Timeout
		if timeout == 0 {
			timeout = h.defaultTimeout
		}
		
		client := &http.Client{
			Timeout: timeout,
		}

		// Выполняем запрос
		resp, err := client.Do(req)
		if err != nil {
			h.logger.Error("Failed to proxy request", zap.String("service", serviceName), zap.Error(err))
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service unavailable"})
			return
		}
		defer resp.Body.Close()

		// Копируем заголовки ответа
		for key, values := range resp.Header {
			for _, value := range values {
				c.Header(key, value)
			}
		}

		// Читаем тело ответа
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			h.logger.Error("Failed to read response body", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
			return
		}

		// Устанавливаем статус и отправляем тело ответа
		c.Status(resp.StatusCode)
		c.Writer.Write(respBody)
	}
} 