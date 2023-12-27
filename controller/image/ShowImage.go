package image

import (
	"ImageCreation/controller/response"
	"ImageCreation/logic/image"
	"ImageCreation/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Info struct {
	ImageInfo     models.Image
	ImageUserInfo models.UserInformation
}

type Pages struct {
	Page  int
	Label string
}

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
		response.Json(c, 401, "获取主页图片失败", err)
	} else {
		response.Json(c, 200, "获取主页图片成功", imageInfos)
	}
	return
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
// @Router /gallery-single [GET]
func ShowImageInfo(c *gin.Context) {
	imageId := c.Query("id")
	id, _ := strconv.ParseInt(imageId, 10, 64)
	var si image.ShowImage
	if imageInfo, imageUserInfo, err := si.ImageInfo(id); err != nil {
		c.HTML(http.StatusOK, "errors.html", err)
	} else {
		//response.Json(c, 200, "获取图片详细信息成功", imageInfo)
		c.HTML(http.StatusOK, "gallery-single.html", Info{imageInfo, imageUserInfo})
	}
	return
}

// ShowGalleryImage 显示主题图片信息
// @Summary 显示主题图片信息
// @Description 用于显示主题图片信息
// @Tags 显示图片详细信息
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "获取主题图片信息成功"
// @failure 401 {object}  response.Information "获取主题图片信息失败"
// @Router /gallery [GET]
func ShowGalleryImage(c *gin.Context) {
	label := c.Query("label")
	page := c.Query("page")
	if page == "" || len(page) == 0 {
		page = "1"
	}
	var si image.ShowImage
	if imageInfo, total, err := si.GalleryImage(label, page); err != nil {
		c.HTML(http.StatusOK, "errors.html", err)
	} else {
		var pages []Pages
		for i := 1; i < total/20+1; i++ {
			pages = append(pages, Pages{Page: i, Label: label})
		}
		c.HTML(http.StatusOK, "gallery.html", gin.H{"ImageInfo": imageInfo, "Pages": pages})
	}
	return
}

// SearchImage 搜索图片
// @Summary 搜索图片
// @Description 用于搜索图片
// @Tags 搜索图片
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "搜索图片成功"
// @failure 401 {object}  response.Information "搜索图片失败"
// @Router /search [POST]
func SearchImage(c *gin.Context) {
	label := c.Query("search")
	var si image.ShowImage
	if imageInfo, err := si.GetSearchImage(label); err != nil {
		c.HTML(http.StatusOK, "errors.html", err)
	} else {
		//response.Json(c, 200, "获取图片详细信息成功", imageInfo)
		c.HTML(http.StatusOK, "search.html", imageInfo)
	}
	return
}
