package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

// ServiceProxy обеспечивает перенаправление запросов к микросервисам
type ServiceProxy struct {
	serviceURLs  map[string]string
	reverseProxy map[string]*httputil.ReverseProxy
}

// NewServiceProxy создает новый объект ServiceProxy
func NewServiceProxy() *ServiceProxy {
	proxy := &ServiceProxy{
		serviceURLs:  make(map[string]string),
		reverseProxy: make(map[string]*httputil.ReverseProxy),
	}

	// Инициализация URL сервисов из переменных окружения
	proxy.initServiceURLs()

	return proxy
}

// initServiceURLs инициализирует URL-адреса микросервисов из переменных окружения
func (p *ServiceProxy) initServiceURLs() {
	// Сопоставление имен сервисов с переменными окружения
	serviceEnvMap := map[string]string{
		"auth":        "AUTH_SERVICE_URL",
		"applicant":   "APPLICANT_SERVICE_URL",
		"schedule":    "SCHEDULE_SERVICE_URL",
		"club":        "CLUB_SERVICE_URL",
		"performance": "PERFORMANCE_SERVICE_URL",
		"support":     "SUPPORT_SERVICE_URL",
	}

	// Заполнение URL-адресов из переменных окружения
	for service, envVar := range serviceEnvMap {
		serviceURL := os.Getenv(envVar)
		if serviceURL == "" {
			// Если переменная окружения не определена, используем стандартный URL
			serviceURL = fmt.Sprintf("http://%s-service:8080", service)
			log.Printf("Warning: %s not set, using default: %s", envVar, serviceURL)
		}
		p.serviceURLs[service] = serviceURL

		// Создаем ReverseProxy для каждого сервиса
		url, err := url.Parse(serviceURL)
		if err != nil {
			log.Printf("Error parsing URL %s: %v", serviceURL, err)
			continue
		}

		p.reverseProxy[service] = httputil.NewSingleHostReverseProxy(url)
		
		// Настройка обработки ошибок
		p.reverseProxy[service].ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			log.Printf("Error proxying request to %s: %v", serviceURL, err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(fmt.Sprintf(`{"error": "Service %s is unavailable"}`, service)))
		}
	}
}

// ProxyRequest возвращает обработчик запросов, который перенаправляет запросы к нужному микросервису
func (p *ServiceProxy) ProxyRequest(serviceName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверка доступности сервиса
		if _, ok := p.reverseProxy[serviceName]; !ok {
			log.Printf("Service %s not configured", serviceName)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"error": "Service %s not configured"}`, serviceName)))
			return
		}

		// Добавляем заголовок с временем запроса
		r.Header.Add("X-Request-Time", time.Now().Format(time.RFC3339))

		// Изменение пути запроса для передачи микросервису
		// Удаляем префикс /api/v1/{serviceName} из пути
		r.URL.Path = strings.TrimPrefix(r.URL.Path, fmt.Sprintf("/api/v1/%s", serviceName))
		if r.URL.Path == "" {
			r.URL.Path = "/"
		}

		// Логируем проксирование запроса
		log.Printf("Proxying request %s %s to %s%s", r.Method, r.RequestURI, p.serviceURLs[serviceName], r.URL.Path)

		// Передаем запрос микросервису
		p.reverseProxy[serviceName].ServeHTTP(w, r)
	}
} 