package service

import (
	"net/http"
	"server/serializer"
	"server/utils"
	"server/utils/exception"
	"server/utils/logger"
)

// AuthService 媒体列表服务
type AuthService struct {
	Name string `form:"name" json:"name"`
}

// Login 用户登录函数
func (service *AuthService) Auth() serializer.Response {
	code := http.StatusOK
	token, err := utils.GenerateToken(service.Name)
	if err != nil {
		logger.Error(err)
		code = http.StatusUnauthorized
		return serializer.Response{
			Status: code,
			Msg:    exception.GetMsg(code),
		}
	}
	return serializer.Response{
		Data:   serializer.TokenData{Token: token},
		Status: code,
		Msg:    exception.GetMsg(code),
	}
}
