package routes

import (
	"ImageCreation/logger"
	"ImageCreation/pkg/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strconv"
)

func Setup() *gin.Engine {
	r := gin.New() //创立新的路由
	store := cookie.NewStore([]byte("ImageCreation"))
	r.Static("/assets", "./templates/assets")
	r.Static("/static", "./static")
	r.LoadHTMLFiles("templates/index.html", "templates/contact.html", "templates/gallery.html", "templates/gallery-single.html", "templates/mine.html",
		"templates/sample-inner-page.html", "templates/author.html", "templates/login.html", "templates/reminder-login.html", "templates/creation.html",
		"templates/search.html", "templates/errors.html", "templates/modify.html")
	r.Use(cors.Cors())
	r.Use(logger.GinLogger(), logger.GinRecovery(true), sessions.Sessions("ImageCreation", store)) //使用日志记录路由信息
	pprof.Register(r)                                                                              //注册pprof相关路由
	InitSwagger(r)                                                                                 //注册Swagger文档路由

	PageRoute(r.Group("/"))                         //页面模块路由
	AccountRoute(r.Group("/account"))               //账户模块路由
	ImageRoute(r.Group("/image"))                   //账户模块路由
	r.Run(":" + strconv.Itoa(viper.GetInt("port"))) //运行路由
	return r
}
