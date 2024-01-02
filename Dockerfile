# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.20

# 设置工作目录
WORKDIR /app

# 将当前目录中的所有文件复制到容器的工作目录中
COPY . .

# 编译 Go 项目
RUN go build -o main .

# 暴露容器的端口（Gin 默认端口为 8080）
EXPOSE 4001

# 运行构建好的 Go 应用程序
CMD ["./main"]
