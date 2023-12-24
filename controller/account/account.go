package account

import (
	"ImageCreation/controller/response"
	"ImageCreation/dao/redis"
	"ImageCreation/logic/account"
	"ImageCreation/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// UserLogin 用户登录
// @Summary 用户登录
// @Description 用于用户登录
// @Tags 登录
// @Accept application/json
// @Produce application/json
// @Param object body LoginBinder true "登录参数"
// @Security ApiKeyAuth
// @Success 200 {object}  response.SuccessLogin "登陆成功"
// @failure 401 {object}  response.Information "账号或者密码错误"
// @failure 402 {object}  response.Information "请输入账号或者密码"
// @failure 403 {object}  response.Information "该用户未注册"
// @failure 404 {object}  response.Information "更新时间失败"
// @Router /account/login [POST]
func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if len(username) == 0 || len(password) == 0 {
		var reb LoginBinder
		if err := c.ShouldBind(&reb); err != nil {
			response.LoginJson(c, 200, "请输入账号或者密码", "", "")
			return
		} else {
			username = reb.Username
			password = reb.Password
		}
	}
	var la account.Account
	user, loginErr := la.Login(username, password)
	if loginErr == account.PasswordErr {
		response.LoginJson(c, 200, "账号或者密码错误", "", "")
		return
	} else if loginErr == account.NotRegister {
		response.LoginJson(c, 200, "该用户未注册", "", "")
		return
	} else {
		token, _ := jwt.GenToken(user.ID, user.Email)
		var tr redis.TokenRedis
		tr.SetToken(user.Email, token)
		loc, _ := time.LoadLocation("Asia/Shanghai")
		currentTime := time.Now().In(loc)
		if err := la.SetLastLogin(username, currentTime); err != nil {
			response.LoginJson(c, 200, "更新时间失败", "", "")
			return
		}
		response.LoginJson(c, http.StatusOK, "登陆成功", user, token)
		return
	}
}

// UserRegister 新用户注册
// @Summary 新用户注册
// @Description 用于新用户注册账号使用
// @Tags 注册
// @Accept application/json
// @Produce application/json
// @Param object body RegisterBinder false "注册参数"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "注册成功"
// @failure 401 {object}  response.Information "账号已注册"
// @failure 402 {object}  response.Information "验证码错误"
// @failure 403 {object}  response.Information "请输入完整的信息"
// @Router /account/register [POST]
func UserRegister(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	code := c.PostForm("code")
	if len(username) == 0 || len(email) == 0 || len(password) == 0 || len(code) == 0 {
		var reb RegisterBinder
		if err := c.ShouldBind(&reb); err != nil {
			response.Json(c, 200, "请输入完整的信息", 0)
			return
		} else {
			username = reb.Username
			email = reb.Email
			password = reb.Password
			code = reb.Code
		}
	}
	var ra account.Account
	user, err := ra.Register(username, email, password, code)
	if err == account.ExistUser {
		response.Json(c, 200, "账号已注册", 0)
		return
	} else if err == account.CodeErr {
		response.Json(c, 200, "验证码错误", 0)
		return
	} else {
		var mr redis.MailRedis
		mr.DelMail(email)
		response.Json(c, 200, "注册成功", user)
		return
	}
}

// GetMail 发送邮件
// @Summary 发送验证码邮件
// @Description 新用户发送验证码用于注册账号
// @Tags 验证码
// @Accept application/json
// @Produce application/json
// @Param object body MailBinder false "发送邮件参数"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "发送验证码成功"
// @failure 401 {object}  response.Information "邮件发送失败"
// @failure 402 {object}  response.Information "请输入邮箱或者密码"
// @Router /account/mail [POST]
func GetMail(c *gin.Context) {
	email := c.PostForm("email")
	if len(email) == 0 {
		var reb MailBinder
		if err := c.ShouldBind(&reb); err != nil {
			response.Json(c, 200, "请输入完整的信息", 0)
			return
		} else {
			email = reb.Email
		}
	}
	var gm account.Account
	code, err := gm.Mail(email)
	if err != nil {
		response.Json(c, 200, "邮件发送失败", err)
		return
	} else {
		var mr redis.MailRedis
		mr.SetMail(email, code)
		response.Json(c, 200, "邮件发送成功", code)
		return
	}
}

// UserLogout 用户注销
// @Summary 用户注销
// @Description 用于登录用户注销
// @Tags 注销
// @Accept application/json
// @Produce application/json
// @Param email query string false "注销参数"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "注销成功"
// @failure 401 {object}  response.Information "注销失败"
// @failure 402 {object}  response.Information "您还未登录"
// @failure 403 {object}  response.Information "请输入邮箱账号"
// @Router /account/logout [GET]
func UserLogout(c *gin.Context) {
	username := c.Query("username")
	if username == "" || len(username) == 0 {
		response.Json(c, 200, "请输入邮箱账号", 0)
		return
	}
	var la account.Account
	if err := la.Logout(username); err == nil {
		response.Json(c, 200, "注销失败", 0)
		return
	} else {
		var tr redis.TokenRedis
		tr.DelToken(username)
		response.Json(c, 200, "注销成功", 0)
		return
	}
}

// UserDelete 删除账号
// @Summary 删除账号
// @Description 用于删除账号
// @Tags 删除账号
// @Accept application/json
// @Produce application/json
// @Param username query string false "删除账号参数"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "删除账号成功"
// @failure 401 {object}  response.Information "删除账号失败"
// @failure 402 {object}  response.Information "请输入账号"
// @Router /account/delete [GET]
func UserDelete(c *gin.Context) {
	username := c.Query("username")
	if username == "" || len(username) == 0 {
		response.Json(c, 402, "请输入账号", 0)
		return
	}
	var la account.Account
	if err := la.DeleteAccount(username); err != nil {
		response.Json(c, 401, "删除账号失败", 0)
		return
	} else {
		response.Json(c, 200, "删除账号成功", 0)
		return
	}
}

// UserModify 修改用户账号信息
// @Summary 修改用户账号信息
// @Description 用于修改用户账号信息
// @Tags 修改用户账号信息
// @Accept application/json
// @Produce application/json
// @Param object body ModifyBinder false "修改用户账号信息参数"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "修改信息成功"
// @failure 401 {object}  response.Information "修改信息失败"
// @failure 402 {object}  response.Information "请输入信息"
// @Router /account/modify [POST]
func UserModify(c *gin.Context) {
	information := c.PostFormMap("information")
	if len(information) == 0 {
		response.Json(c, 402, "请输入信息", 0)
		return
	}
	var la account.Account
	if err := la.SetAccount(information); err != nil {
		response.Json(c, 401, "修改信息失败", 0)
		return
	} else {
		response.Json(c, 200, "修改信息成功", 0)
		return
	}
}

// UserModifyInformation 修改用户详细信息
// @Summary 修改用户详细信息
// @Description 用于修改用户详细信息
// @Tags 修改用户详细信息
// @Accept application/json
// @Produce application/json
// @Param object body ModifyInformationBinder false "修改用户详细信息参数"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "修改信息成功"
// @failure 401 {object}  response.Information "修改信息失败"
// @failure 402 {object}  response.Information "请输入信息"
// @Router /account/modify/information [POST]
func UserModifyInformation(c *gin.Context) {
	information := c.PostFormMap("information")
	if len(information) == 0 {
		response.Json(c, 402, "请输入信息", 0)
		return
	}
	var la account.Account
	if err := la.SetUserInformation(information); err != nil {
		response.Json(c, 401, "修改信息失败", 0)
		return
	} else {
		response.Json(c, 200, "修改信息成功", 0)
		return
	}
}

// SendContactMail 发送联系邮件
// @Summary 发送联系邮件
// @Description 用户发送联系邮件
// @Tags 发送联系邮件
// @Accept application/json
// @Produce application/json
// @Param object body ContactMailBinder false "发送联系邮件参数"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "发送联系邮件成功"
// @failure 401 {object}  response.Information "发送联系邮件失败"
// @Router /account/contact/mail [POST]
func SendContactMail(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	subject := c.PostForm("subject")
	message := c.PostForm("message")
	var gm account.Account
	err := gm.ContactMail(name, email, subject, message)
	if err != nil {
		response.Json(c, 200, "信息发送失败，请稍后重试！", err)
	} else {
		response.Json(c, 200, "您的信息已发送。非常感谢。", "")
	}
	return

}
