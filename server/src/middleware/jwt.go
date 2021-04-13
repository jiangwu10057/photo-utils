package middleware

import (
	"server/utils"
	"server/utils/exception"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = http.StatusOK
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusUnauthorized
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = exception.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = exception.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": code,
				"msg":    exception.GetMsg(code),
				"data":   data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

// JWTAdmin token验证中间件
func JWTAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = http.StatusOK
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusBadRequest
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = exception.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = exception.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": code,
				"msg":    exception.GetMsg(code),
				"data":   data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
