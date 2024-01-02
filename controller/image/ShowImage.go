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

// ImageLike 图片进行点赞操作
// @Summary 图片进行点赞操作
// @Description 用于图片进行点赞操作
// @Tags 图片进行点赞操作
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "点赞成功"
// @failure 401 {object}  response.Information "点赞失败"
// @Router /image/like [POST]
func ImageLike(c *gin.Context) {
	userID := c.PostForm("userId")
	imageId := c.PostForm("imageId")
	isLike := c.PostForm("isLike")
	var si image.ShowImage
	if imageInfo, err := si.ImageToLike(userID, imageId, isLike); err != nil {
		response.Json(c, 200, "点赞失败", err)
	} else {
		var msg string
		if isLike == "1" {
			msg = "点赞成功"
		} else {
			msg = "取消点赞成功"
		}
		response.Json(c, 200, msg, imageInfo)
	}
	return
}

// ImageCollect 图片进行收藏操作
// @Summary 图片进行收藏操作
// @Description 用于图片进行收藏操作
// @Tags 图片进行收藏操作
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "收藏成功"
// @failure 401 {object}  response.Information "收藏失败"
// @Router /image/collect [POST]
func ImageCollect(c *gin.Context) {
	userID := c.PostForm("userId")
	imageId := c.PostForm("imageId")
	isCollect := c.PostForm("isCollect")
	var si image.ShowImage
	if imageInfo, err := si.ImageToCollect(userID, imageId, isCollect); err != nil {
		response.Json(c, 200, "收藏失败", err)
	} else {
		var msg string
		if isCollect == "1" {
			msg = "收藏成功"
		} else {
			msg = "取消收藏成功"
		}
		response.Json(c, 200, msg, imageInfo)
	}
	return
}

//// ImageScore 图片进行评分操作
//// @Summary 图片进行评分操作
//// @Description 用于图片进行评分操作
//// @Tags 图片进行评分操作
//// @Accept application/json
//// @Produce application/json
//// @Security ApiKeyAuth
//// @Success 200 {object}  response.Information "评分成功"
//// @failure 401 {object}  response.Information "评分失败"
//// @Router /image/score [POST]
//func ImageScore(c *gin.Context) {
//	userID := c.PostForm("userId")
//	imageId := c.PostForm("imageId")
//	isScore := c.PostForm("isScore")
//	fmt.Println(userID)
//	var si image.ShowImage
//	if imageInfo, err := si.ImageToScore(userID, imageId, isScore); err != nil {
//		response.Json(c, 200, "评分失败", err)
//	} else {
//		var msg string
//		if isScore == "1" {
//			msg = "评分成功"
//		} else {
//			msg = "取消评分成功"
//		}
//		response.Json(c, 200, msg, imageInfo)
//	}
//	return
//}

// ImageOperation 查询当前用户对图片的操作
// @Summary 查询当前用户对图片的操作
// @Description 用于查询当前用户对图片的操作
// @Tags 查询当前用户对图片的操作
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "查询成功"
// @failure 401 {object}  response.Information "查询失败"
// @Router /image/collect [POST]
func ImageOperation(c *gin.Context) {
	userID := c.PostForm("userId")
	imageId := c.PostForm("imageId")
	var si image.ShowImage
	isLike, isCollect, isScore := si.ImageToOperation(userID, imageId)
	response.Json(c, 200, "查询成功", gin.H{"isLike": isLike, "isCollect": isCollect, "isScore": isScore})
	return
}

// ImageBrowse 图片进行浏览操作
// @Summary 图片进行浏览操作
// @Description 用于图片进行浏览操作
// @Tags 图片进行浏览操作
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "浏览成功"
// @failure 401 {object}  response.Information "浏览失败"
// @Router /image/browse [POST]
func ImageBrowse(c *gin.Context) {
	userID := c.PostForm("userId")
	imageId := c.PostForm("imageId")
	var si image.ShowImage
	if err := si.ImageToBrowse(userID, imageId); err != nil {
		response.Json(c, 200, "浏览失败", err)
	} else {
		response.Json(c, 200, "浏览成功", "")
	}
	return
}
