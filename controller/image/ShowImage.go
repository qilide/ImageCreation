package image

import (
	"ImageCreation/controller/response"
	"ImageCreation/logic/image"
	"github.com/gin-gonic/gin"
)

// ShowIndexImage 显示主页图片
// @Summary 显示主页图片
// @Description 用于显示主页图片
// @Tags 显示主页图片
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "获取主页图片成功"
// @failure 401 {object}  response.Information "获取主页图片失败"
// @Router /image/index [GET]
func ShowIndexImage(c *gin.Context) {
	var si image.ShowImage
	if imageInfos, err := si.IndexImage(); err != nil {
		response.Json(c, 401, "获取主页图片失败", imageInfos)
	} else {
		response.Json(c, 200, "获取主页图片成功", imageInfos)
	}
	return
}
