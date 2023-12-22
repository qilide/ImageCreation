package image

import (
	"ImageCreation/controller/response"
	"ImageCreation/logic/image"
	"ImageCreation/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

type Info struct {
	ImageInfo     models.Image
	ImageUserInfo models.UserInformation
}

// ShowImageInfo 显示图片详细信息
// @Summary 显示图片详细信息
// @Description 用于显示图片详细信息
// @Tags 显示图片详细信息
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "获取图片详细信息成功"
// @failure 401 {object}  response.Information "获取图片详细信息失败"
// @Router /image/index [GET]
func ShowImageInfo(c *gin.Context) {
	imageId := c.Query("id")
	id, _ := strconv.ParseInt(imageId, 10, 64)
	var si image.ShowImage
	if imageInfo, imageUserInfo, err := si.ImageInfo(id); err != nil {
		c.HTML(http.StatusOK, "gallery-single.html", Info{imageInfo, imageUserInfo})
		//response.Json(c, 200, "获取图片详细信息成功", imageInfo)
	} else {
		c.HTML(http.StatusOK, "gallery-single.html", Info{imageInfo, imageUserInfo})
	}
	return
}
