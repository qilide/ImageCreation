package routes

import (
	"ImageCreation/controller/creation"
	"github.com/gin-gonic/gin"
)

// CreationRoute 创作路由
func CreationRoute(CreationRoute *gin.RouterGroup) {
	// 上传创作图片
	CreationRoute.POST("/upload/image", creation.UploadImage)
}
