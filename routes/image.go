package routes

import (
	"ImageCreation/controller/image"
	"github.com/gin-gonic/gin"
)

// ImageRoute 图片路由
func ImageRoute(ImageGroup *gin.RouterGroup) {
	//用户注销
	ImageGroup.GET("/index", image.ShowIndexImage)
}
