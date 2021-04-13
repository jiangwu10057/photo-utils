package service

import (
	"fmt"
	"net/http"

	"server/serializer"
	"server/cache"

	"github.com/gin-gonic/gin"
)

type FileCoverQueryService struct {
	Id string `form:"id" json:"id"`
}

func (service *FileCoverQueryService) Query() serializer.Response {

	if service.Id == "" {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    "参数缺失",
		}
	}

	key := fmt.Sprintf("photo:upload:%s", service.Id)

	result, err := cache.RedisClient.HGetAll(key).Result()

	if err != nil {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:  "失败",
			Data: "",
		}
	}


	originFilePath := fmt.Sprintf("upload/%s/%s%s", result["dir"], service.Id, result["ext"])
	filePath := fmt.Sprintf("upload/%s/%s_%s%s", result["dir"], service.Id, result["type"], result["ext"])

	return serializer.Response{
		Status: http.StatusOK,
		Msg:    "成功",
		Data:  gin.H{
			"origin":originFilePath,
			"dist": filePath,
		},
	}
}
