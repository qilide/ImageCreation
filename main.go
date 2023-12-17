package main

import (
	"ImageCreation/dao/mysql"
	"ImageCreation/dao/redis"
	"ImageCreation/logger"
	"ImageCreation/middlewares"
	"ImageCreation/routes"
	"ImageCreation/settings"
	"fmt"
	"go.uber.org/zap"
)

// @title 图片摄影创作社交网站
// @version 1.0
// @description 在这里你可以获取想要的照片，并对照片进行二次创作，快来开始使用吧！
func main() {
	//1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("Init settings failed, err: %v\n", err)
		return
	}
	//2.初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("Init logger failed, err: %v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//3.初始化Mysql连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("Init mysql failed, err: %v\n", err)
		return
	}
	defer mysql.Close()
	//4.初始化Redis连接
	if err := redis.Init(); err != nil {
		fmt.Printf("Init redis failed, err: %v\n", err)
		return
	}
	defer redis.Close()
	//5.注册路由
	r := routes.Setup()
	//6.启动服务(优雅关机)
	if err := middlewares.GracefulShutdown(r); err != nil {
		fmt.Printf("Graceful Shutdown failed, err: %v\n", err)
		return
	}
}
