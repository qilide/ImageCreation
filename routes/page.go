package routes

import (
	"ImageCreation/controller/author"
	"ImageCreation/controller/image"
	"github.com/gin-gonic/gin"
	"net/http"
)

//页面路由

func PageRoute(Page *gin.RouterGroup) {
	//主页
	Page.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	//摄影师
	Page.GET("author", author.ShowAuthors)
	//我的
	Page.GET("mine", author.ShowAuthorInfo)
	//画廊
	Page.GET("gallery", image.ShowGalleryImage)
	//画廊单体
	Page.GET("gallery-single", image.ShowImageInfo)
	//搜索
	Page.GET("search", image.SearchImage)
	//内置页
	Page.GET("sample-inner-page", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sample-inner-page.html", gin.H{})
	})
	//联系
	Page.GET("contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", gin.H{})
	})
	//创作
	Page.GET("creation", func(c *gin.Context) {
		c.HTML(http.StatusOK, "creation.html", gin.H{})
	})
	//登录页面
	Page.GET("login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	//登录页面
	Page.GET("modify", func(c *gin.Context) {
		c.HTML(http.StatusOK, "modify.html", gin.H{})
	})
}
