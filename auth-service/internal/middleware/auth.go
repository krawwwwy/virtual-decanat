package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// JWTAuthMiddleware проверяет JWT токен в заголовке Authorization
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		// Проверяем формат заголовка "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer <token>"})
			return
		}

		tokenString := parts[1]
		claims := &Claims{}

		// Парсим токен
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Проверяем алгоритм подписи
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(viper.GetString("jwt.secret")), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Проверяем срок действия токена
		if claims.ExpiresAt.Time.Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			return
		}

		// Проверяем тип токена
		if claims.TokenType != "access" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token type"})
			return
		}

		// Сохраняем данные пользователя в контексте
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// Claims определяет структуру JWT токена
type Claims struct {
	UserID    uint   `json:"user_id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	TokenType string `json:"type"`
	jwt.RegisteredClaims
}

// RoleAuthMiddleware проверяет роль пользователя
func RoleAuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User role not found"})
			return
		}

		userRole := role.(string)
		authorized := false
		for _, r := range roles {
			if r == userRole {
				authorized = true
				break
			}
		}

		if !authorized {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			return
		}

		c.Next()
	}
} 