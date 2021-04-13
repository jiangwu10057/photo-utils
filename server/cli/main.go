package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

var logger *log.Logger

func intRedis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB:       int(db),
	})

	_, err := client.Ping().Result()

	if err != nil {
		logger.Println(err)
		panic(err)
	}

	RedisClient = client
}

func initConfig() {
	pwd,_ := os.Getwd()
	confDir := pwd + "/../conf/"

	// 从本地读取环境变量
	envErr := godotenv.Load(confDir + ".env")

	if envErr != nil {
		logger.Println(envErr)
		panic(envErr)
	}
}

func initLogger() {
	filePath := fmt.Sprintf("logs/%s.%s", time.Now().Format("20060102"), "log")
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	logger = log.New(handle, "", log.LstdFlags)
}

func readInfo() (map[string]string, bool) {
	content, err := RedisClient.BRPop(0, "photo:upload:images_waiting").Result();

	if err != nil {
		logger.Println(err)
		return nil, false
	}

	id := string(content[1])
	key := fmt.Sprintf("photo:upload:%s", id)
	
	info, getErr := RedisClient.HGetAll(key).Result()

	if getErr != nil {
		logger.Println(err)
		RedisClient.LPush("photo:upload:images_waiting", id)
		return nil, false
	}
	info["id"] = id
	return info, true
}

func checkFile(info map[string]string) (string, bool) {
	filePath := fmt.Sprintf("%s%s/%s%s", os.Getenv("SAVE_PATH"), info["dir"], info["id"], info["ext"])

	_, fileErr := os.Stat(filePath)
	switch {
	case os.IsNotExist(fileErr):
		logger.Println(fileErr)
		RedisClient.LPush("photo:upload:images_waiting", info["id"])
		return "", false
	}

	return filePath, true
}

func execCover(info map[string]string, filePath string) (bool) {
	src := filePath
	dist := fmt.Sprintf("%s%s/%s_%s%s", os.Getenv("SAVE_PATH"), info["dir"], info["id"], info["type"], info["ext"])
	cmd := exec.Command("python3", "/workspace/script/src/Imageutils.py", fmt.Sprintf("--src=%s", src), fmt.Sprintf("--dist=%s", dist), fmt.Sprintf("--type=%s", info["type"]))
	buf, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	}
	
	key := fmt.Sprintf("photo:upload:%s", info["id"])

	result := strings.Replace(string(buf), "\n", "", -1)
	
	//buf是byte[]类型
	if result == "True" {
		RedisClient.HSet(key, "isFinished", 1)
		return true
	}else{
		_, ok := info["retry"]
		
		if ok{
			RedisClient.HSet(key, "retry", 1)
		}else{
			retry,_ := strconv.Atoi(info["retry"])
			RedisClient.HSet(key, "retry",  retry + 1)
		}
		
		return false
	}
}
func main() {
	initLogger()
	initConfig()
	intRedis()
	count := 1000
	for i := 0; i < count; {
		info, _ok := readInfo()
		if !_ok {
			continue
		}
		
		filePath, _checkResult := checkFile(info)

		if !_checkResult{
			continue
		}
		
		/**
		* todo 长阻塞操作用协程来改写提高处理能力
		*/
		execCover(info, filePath)
	}
}