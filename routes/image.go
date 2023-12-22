package routes

import (
	"ImageCreation/controller/image"
	"github.com/gin-gonic/gin"
)

// ImageRoute 图片路由
func ImageRoute(ImageGroup *gin.RouterGroup) {
	//获取主页图片
	ImageGroup.GET("/index", image.ShowIndexImage)
}
