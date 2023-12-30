package mine

import (
	"ImageCreation/logic/mine"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Pages struct {
	Page   int
	UserID string
}

// ShowImageLike 显示我的点赞图片
// @Summary 显示我的点赞图片
// @Description 用于显示我的点赞图片
// @Tags 显示我的点赞图片
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "获取我的点赞图片成功"
// @failure 401 {object}  response.Information "获取我的点赞图片失败"
// @Router /like [GET]
func ShowImageLike(c *gin.Context) {
	userID := c.Query("id")
	page := c.Query("page")
	if page == "" || len(page) == 0 {
		page = "1"
	}
	var mii mine.MeImageInfo
	if imageInfo, total, err := mii.MeLikeImage(page); err != nil {
		c.HTML(http.StatusOK, "errors.html", err)
	} else {
		var pages []Pages
		for i := 1; i < total/20/20+1; i++ {
			pages = append(pages, Pages{Page: i, UserID: userID})
		}
		c.HTML(http.StatusOK, "like.html", gin.H{"ImageInfo": imageInfo, "Pages": pages})
	}
	return
}

// ShowImageCollect 显示我的收藏图片
// @Summary 显示我的收藏图片
// @Description 用于显示我的收藏图片
// @Tags 显示我的收藏图片
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "获取我的收藏图片成功"
// @failure 401 {object}  response.Information "获取我的收藏图片失败"
// @Router /collect [GET]
func ShowImageCollect(c *gin.Context) {
	userID := c.Query("id")
	page := c.Query("page")
	if page == "" || len(page) == 0 {
		page = "1"
	}
	var mii mine.MeImageInfo
	if imageInfo, total, err := mii.MeCollectImage(page); err != nil {
		c.HTML(http.StatusOK, "errors.html", err)
	} else {
		var pages []Pages
		for i := 1; i < total/20/20+1; i++ {
			pages = append(pages, Pages{Page: i, UserID: userID})
		}
		c.HTML(http.StatusOK, "collect.html", gin.H{"ImageInfo": imageInfo, "Pages": pages})
	}
	return
}

// ShowImageBrowse 显示我的浏览图片
// @Summary 显示我的浏览图片
// @Description 用于显示我的浏览图片
// @Tags 显示我的浏览图片
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "获取我的浏览图片成功"
// @failure 401 {object}  response.Information "获取我的浏览图片失败"
// @Router /browse [GET]
func ShowImageBrowse(c *gin.Context) {
	userID := c.Query("id")
	page := c.Query("page")
	if page == "" || len(page) == 0 {
		page = "1"
	}
	var mii mine.MeImageInfo
	if imageInfo, total, err := mii.MeBrowseImage(page); err != nil {
		c.HTML(http.StatusOK, "errors.html", err)
	} else {
		var pages []Pages
		for i := 1; i < total/20/20+1; i++ {
			pages = append(pages, Pages{Page: i, UserID: userID})
		}
		c.HTML(http.StatusOK, "browse.html", gin.H{"ImageInfo": imageInfo, "Pages": pages})
	}
	return
}

// ShowImageScore 显示我的评分图片
// @Summary 显示我的评分图片
// @Description 用于显示我的评分图片
// @Tags 显示我的评分图片
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "获取我的评分图片成功"
// @failure 401 {object}  response.Information "获取我的评分图片失败"
// @Router /score [GET]
func ShowImageScore(c *gin.Context) {
	userID := c.Query("id")
	page := c.Query("page")
	if page == "" || len(page) == 0 {
		page = "1"
	}
	var mii mine.MeImageInfo
	if imageInfo, total, err := mii.MeScoreImage(page); err != nil {
		c.HTML(http.StatusOK, "errors.html", err)
	} else {
		var pages []Pages
		for i := 1; i < total/20/20+1; i++ {
			pages = append(pages, Pages{Page: i, UserID: userID})
		}
		c.HTML(http.StatusOK, "score.html", gin.H{"ImageInfo": imageInfo, "Pages": pages})
	}
	return
}
