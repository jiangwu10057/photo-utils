package v1

import (
	"fmt"
	"os"
	"time"
	"path"

	"net/http"

	"github.com/gin-gonic/gin"

	"server/api"
	"server/utils"
	"server/serializer"
	"server/cache"
	"server/service"
	"server/utils/logger"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("upload")
	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    "请求失败",
			Error:  "",
		})
		return
	}

	if !utils.IsImage(file) {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    "文件格式错误",
			Error:  "",
		})
		return
	}

	_type :=  c.DefaultPostForm("type", "1")

	now := time.Now()
	dir := now.Format("2006/01/02")
	savePath := fmt.Sprintf("%s%s", os.Getenv("SAVE_PATH"), dir)

	_, fileErr := os.Stat(savePath)
	switch {
	case os.IsNotExist(fileErr):
		utils.Mkdir(savePath)
	}

	ext := path.Ext(file.Filename)

	id := now.Unix()

	filePath := fmt.Sprintf("%s/%d%s", savePath, id, ext)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    fmt.Sprintf("保存失败 Error:%s", err.Error()),
			Error:  err.Error(),
		})
		return
	}

	key := fmt.Sprintf("photo:upload:%d", id)
	data := map[string]interface{}{"dir": dir, "ext":ext,"type":_type, "isFinished": false}
	redisErr := cache.RedisClient.HMSet(key, data).Err()

	if redisErr != nil {
		logger.Error(redisErr)
		panic(redisErr)
	}

	redisErr = cache.RedisClient.LPush("photo:upload:images_waiting", id).Err()
	if redisErr != nil {
		logger.Error(redisErr)
		panic(redisErr)
	}

	c.JSON(http.StatusOK, serializer.Response{
		Status: http.StatusOK,
		Msg:    "上传文件成功",
	})
}

func MultiUpload(c *gin.Context) {
	//获取解析后表单
	form, _ := c.MultipartForm()
	//这里是多文件上传 在之前单文件upload上传的基础上加 [] 变成upload[] 类似文件数组的意思
	files := form.File["upload[]"]
	//循环存文件到服务器本地
	for _, file := range files {
		c.SaveUploadedFile(file, file.Filename)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d 个文件被上传成功!", len(files)))
}

func QueryCoverStatus(c *gin.Context) {
	service := service.FileCoverQueryService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Query()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(err))
		logger.Error(err)
	}
}
