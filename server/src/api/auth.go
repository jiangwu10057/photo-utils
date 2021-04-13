package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/service"
	"server/utils/logger"
)

// Auth 授权
func Auth(c *gin.Context) {
	var service service.AuthService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Auth()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusUnauthorized, ErrorResponse(err))
		logger.Error(err)
	}
}
