package router

import (
	"os"

	"server/api"
	apiv1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.Use(middleware.Logger())

	//静态资源
	r.Static("/upload", os.Getenv("SAVE_PATH"))

	//授权
	r.POST("auth", api.Auth)

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("upload", apiv1.Upload)
		v1.GET("query", apiv1.QueryCoverStatus)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			//验证token
			authed.GET("ping", api.CheckToken)
		}
	}
	return r
}
