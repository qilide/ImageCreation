package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//页面路由

func PageRoute(Page *gin.RouterGroup) {
	//主页
	Page.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	//关于
	Page.GET("about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", gin.H{})
	})
	//联系
	Page.GET("contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", gin.H{})
	})
	//画廊
	Page.GET("gallery", func(c *gin.Context) {
		c.HTML(http.StatusOK, "gallery.html", gin.H{})
	})
	//画廊单体
	Page.GET("gallery-single", func(c *gin.Context) {
		c.HTML(http.StatusOK, "gallery-single.html", gin.H{})
	})
	//内置页
	Page.GET("sample-inner-page", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sample-inner-page.html", gin.H{})
	})
	//服务
	Page.GET("services", func(c *gin.Context) {
		c.HTML(http.StatusOK, "services.html", gin.H{})
	})
	//登录页面
	Page.GET("login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
}
