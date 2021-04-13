# 说明
```
本项目是将图片转换成卡通、素描风格的工具
```

# 目录结构
```
|- mini 小程序
|- h5 h5页面
|- script python脚本
  |- api flask接口，实际不会运行
  |- images 测试图片、视频目录
  |- src 脚本地址
    - Imageutils.py 图片转换脚本
|- server Golang接口
  |- cli redis队列消费程序，调用python脚本实现图片转换
     - main.go 入口文件
  |- conf 配置
    |- locales i18n文件
      - zh-cn.yaml 中文简体翻译
    - .env 环境变量
  |- src 程序目录
  |- upload 图片上传目录
  .air.toml 开发环境代码改动自动编译配置文件
.gitignore 
docker-compose.yml docker编排文件
Dockerfile 镜像描述
README.md 
```

# 启动
## 环境初始化
```
docker-composer up -d
```
## 启动接口
```
cd /workspace/server/src
air -c ../air.toml
```
## 启动队列消费程序
```
cd /workspace/cli/
go run main.go
```

# 接口
## 图片上传
```
api/v1/upload
```
## 查询结果
```
api/v1/query?id=
```