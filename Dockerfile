FROM python:3.9.2 as python
LABEL maintainer="chenwu <jiangwu10057@qq.com>" version="0.1"

FROM cosmtrek/air:v1.15.1

ENV GOPROXY=https://goproxy.io

WORKDIR /workspace

RUN apt update && apt clean && apt install -y python3.6 python3-pip libgl1-mesa-glx