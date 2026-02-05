package middleware

import (
	"go-mini-admin/internal/infrastructure/response"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	UserIDKey   = "userID"
	UsernameKey = "username"
)

func (m *Middleware) JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "authorization header is required")
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.Unauthorized(c, "invalid authorizaton header format")
			c.Abort()
			return
		}

		claims, err := m.JWTManager.ParseToken(authHeader)
		if err != nil {
			response.Unauthorized(c, "invalid or expired token")
			c.Abort()
			return
		}

		// 存储用户信息到gin上下文中
		c.Set(UserIDKey, claims.UserID)
		c.Set(UsernameKey, claims.Username)
		c.Next()
	}
}

// GetUserID gets user ID from context
func GetUserID(c *gin.Context) uint {
	if userID, exists := c.Get(UserIDKey); exists {
		return userID.(uint)
	}
	return 0
}

// GetUsername gets username from context
func GetUsername(c *gin.Context) string {
	if username, exists := c.Get(UsernameKey); exists {
		return username.(string)
	}
	return ""
}
