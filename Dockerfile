# 使用 golang 作为基础镜像
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 复制项目源代码到容器中
COPY . .

# 下载和安装依赖包
RUN go mod download

# 构建项目
RUN go build -o onedrive-proxy

# 暴露端口
EXPOSE 8080

# 启动项目
CMD ["./onedrive-proxy"]
