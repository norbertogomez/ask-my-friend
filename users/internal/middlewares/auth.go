package middlewares

import (
	"strings"

	"users/internal/config"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.GetConfig()
		reqKey := c.Request.Header.Get("X-Auth-Key")
		reqSecret := c.Request.Header.Get("X-Auth-Secret")

		var key string
		var secret string
		if key = cfg.GetString("http.auth.key"); len(strings.TrimSpace(key)) == 0 {
			c.AbortWithStatus(500)
		}
		if secret = cfg.GetString("http.auth.secret"); len(strings.TrimSpace(secret)) == 0 {
			c.AbortWithStatus(401)
		}
		if key != reqKey || secret != reqSecret {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}
