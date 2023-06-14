# 使用 golang 作为基础镜像
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 复制项目源代码到容器中
COPY . /app

# 下载和安装依赖包
RUN go mod download

# 构建项目
RUN go build -o onedrive-proxy ./main.go

# 暴露端口
EXPOSE 8080

VOLUME ["/app"]

# 启动项目
ENTRYPOINT ["./onedrive-proxy"]
