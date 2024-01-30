package routes

import (
	"ImageCreation/controller/creation"
	"github.com/gin-gonic/gin"
)

// CreationRoute 创作路由
func CreationRoute(CreationRoute *gin.RouterGroup) {
	// 上传创作图片
	CreationRoute.POST("/upload/image", creation.UploadImage)
	// 通用物体和场景识别
	CreationRoute.POST("/general/recognition", creation.UploadImage)
	// 图像清晰度增强
	CreationRoute.POST("/clarity/enhancement", creation.ClarityEnhancement)
	// 图像色彩增强
	CreationRoute.POST("/color/enhancement", creation.ColorEnhancement)
	// 图像对比度增强
	CreationRoute.POST("/contrast/enhancement", creation.ContrastEnhancement)
	// 图像无损放大
	CreationRoute.POST("/zoom/pro", creation.ZoomPro)
	// 图像去雾
	CreationRoute.POST("/defogging", creation.Defogging)
	// 人像动漫化
	CreationRoute.POST("/animated/portraits", creation.UploadImage)
	// 图像风格转换
	CreationRoute.POST("/style/conversion", creation.UploadImage)
	// 自定义图像风格
	CreationRoute.POST("/custom/style", creation.UploadImage)
}
