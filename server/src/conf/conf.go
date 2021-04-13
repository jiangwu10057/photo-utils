package conf

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"server/utils/logger"
	"server/cache"
)

// Init 初始化配置项
func Init() {
	confDir := os.Getenv("CONF_DIR")

	// 从本地读取环境变量
	envErr := godotenv.Load(confDir + ".env")

	if envErr != nil {
		logger.Error(envErr)
		panic(envErr)
	}

	gin.SetMode(os.Getenv("GIN_MODE"))

	// 读取翻译文件
	if err := LoadLocales(confDir + "locales/zh-cn.yaml"); err != nil {
		logger.Error(err)
		panic(err)
	}

	cache.Redis()
}
