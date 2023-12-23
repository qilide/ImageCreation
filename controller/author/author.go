package author

import (
	"ImageCreation/logic/author"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ShowAuthors 显示所有摄影师信息
// @Summary 显示所有摄影师信息
// @Description 用于所有摄影师信息
// @Tags 显示所有摄影师信息
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "获取所有摄影师信息成功"
// @failure 401 {object}  response.Information "获取所有摄影师信息失败"
// @Router /image/index [GET]
func ShowAuthors(c *gin.Context) {
	var sa author.ShowAuthor
	if AuthorsInfo, err := sa.AllAuthors(); err != nil {
		c.HTML(http.StatusOK, "author.html", err)
		//response.Json(c, 200, "获取图片详细信息成功", imageInfo)
	} else {
		c.HTML(http.StatusOK, "author.html", AuthorsInfo)
	}
	return
}

// ShowAuthorInfo 显示摄影师详细信息
// @Summary 显示摄影师详细信息
// @Description 用于摄影师详细信息
// @Tags 显示摄影师详细信息
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "获取摄影师详细信息成功"
// @failure 401 {object}  response.Information "获取摄影师详细信息失败"
// @Router /image/index [GET]
func ShowAuthorInfo(c *gin.Context) {
	userId := c.Query("id")
	id, _ := strconv.ParseInt(userId, 10, 64)
	var sa author.ShowAuthor
	if AuthorInfo, err := sa.AuthorInfo(id); err != nil {
		c.HTML(http.StatusOK, "mine.html", err)
	} else {
		c.HTML(http.StatusOK, "mine.html", AuthorInfo)
	}
	return
}
