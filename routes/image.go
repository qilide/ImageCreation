package routes

import (
	"ImageCreation/controller/image"
	"github.com/gin-gonic/gin"
)

// ImageRoute 图片路由
func ImageRoute(ImageGroup *gin.RouterGroup) {
	//获取主页图片
	ImageGroup.GET("/index", image.ShowIndexImage)
	//进行点赞功能
	ImageGroup.POST("/like", image.ImageLike)
	//进行收藏功能
	ImageGroup.POST("/collect", image.ImageCollect)
	//进行评分功能
	ImageGroup.POST("/score", image.ImageScore)
	//进行浏览功能
	ImageGroup.POST("/browse", image.ImageBrowse)
	//查询当前用户对图片的操作
	ImageGroup.POST("/operation", image.ImageOperation)
}
