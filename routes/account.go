package routes

import (
	"ImageCreation/controller/account"
	"github.com/gin-gonic/gin"
)

// AccountRoute 账户路由
func AccountRoute(AccountGroup *gin.RouterGroup) {
	//用户登录
	AccountGroup.POST("/login", account.UserLogin)
	//用户注册
	AccountGroup.POST("/register", account.UserRegister)
	//获得验证码
	AccountGroup.POST("/mail", account.GetMail)
	//发送联系邮件
	AccountGroup.POST("/contact/mail", account.SendContactMail)
	//用户注销
	AccountGroup.GET("/logout", account.UserLogout)
	//删除账号
	AccountGroup.GET("/delete", account.UserDelete)
	//修改用户账号信息
	AccountGroup.POST("/modify", account.UserModify)
	//修改用户详细信息
	AccountGroup.POST("/modify/information", account.UserModifyInformation)
}
