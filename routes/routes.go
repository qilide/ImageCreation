package routes

import (
	"ImageCreation/logger"
	"ImageCreation/pkg/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
)

func Setup() *gin.Engine {
	r := gin.New() //创立新的路由
	store := cookie.NewStore([]byte("ImageCreation"))
	r.Static("/assets", "./templates/assets")
	r.Static("/static", "./static")
	r.LoadHTMLFiles("templates/index.html", "templates/contact.html", "templates/gallery.html", "templates/gallery-single.html", "templates/mine.html",
		"templates/sample-inner-page.html", "templates/author.html", "templates/login.html", "templates/reminder-login.html", "templates/creation.html",
		"templates/search.html", "templates/errors.html", "templates/modify.html", "templates/like.html", "templates/collect.html",
		"templates/score.html", "templates/browse.html", "templates/404-not-found.html")
	r.Use(cors.Cors())
	r.Use(logger.GinLogger(), logger.GinRecovery(true), sessions.Sessions("ImageCreation", store)) //使用日志记录路由信息
	pprof.Register(r)                                                                              //注册pprof相关路由
	InitSwagger(r)                                                                                 //注册Swagger文档路由
	// 加载404错误页面
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404-not-found.html", gin.H{})
	})
	// 设置500提示中间件
	r.Use(errorHttp)

	PageRoute(r.Group("/"))                         //页面模块路由
	AccountRoute(r.Group("/account"))               //账户模块路由
	ImageRoute(r.Group("/image"))                   //账户模块路由
	CreationRoute(r.Group("/creation"))             //账户模块路由
	r.Run(":" + strconv.Itoa(viper.GetInt("port"))) //运行路由
	return r
}

// errorHttp 统一500错误处理函数
func errorHttp(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//封装通用json返回
			c.HTML(200, "errors.html", gin.H{"errors": r})
		}
	}()
	c.Next()
}
