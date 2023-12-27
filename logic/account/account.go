package account

import (
	"ImageCreation/dao/mysql"
	"ImageCreation/dao/redis"
	"ImageCreation/middlewares"
	"ImageCreation/models"
	email2 "ImageCreation/pkg/email"
	"ImageCreation/pkg/snowflake"
	"errors"
	"strconv"
	"time"
)

type Account struct {
}

var (
	NotRegister = errors.New("NotRegister") //该用户未注册
	PasswordErr = errors.New("PasswordErr") //账号或者密码错误
	ExistUser   = errors.New("ExistUser")   //账号已注册
	CodeErr     = errors.New("CodeErr")     //验证码错误
)

// Login 登录逻辑处理函数
func (la Account) Login(username string, password string) (models.User, error) {
	user, err := mysql.UserNameVerity(username)
	if err != nil {
		return user, NotRegister
	} else {
		correct := middlewares.Check(password, user.Password)
		if correct == false {
			return user, PasswordErr
		} else {
			return user, nil
		}
	}
}

// Register 注册逻辑处理函数
func (la Account) Register(username string, email string, password string, code string) (models.User, error) {
	var mr redis.MailRedis
	user, err := mysql.UserNameVerity(username)
	code1 := mr.GetMail(email)
	if err == nil {
		return user, ExistUser
	} else {
		if code1 != code {
			return user, CodeErr
		} else {
			var sf snowflake.Snowflake
			id := sf.NextVal()
			strInt64 := strconv.FormatInt(id, 10)
			id16, _ := strconv.Atoi(strInt64)
			loc, err := time.LoadLocation("Asia/Shanghai")
			currentTime := time.Now().In(loc)
			pwd := middlewares.Encode(password)
			CurrentUser, err := mysql.CreateUser(id16, username, pwd, email, currentTime)
			err = mysql.CreateUserInformation(id16)
			return CurrentUser, err
		}
	}
}

// Mail 接收注册验证码逻辑处理函数
func (gm Account) Mail(email string) (string, error) {
	code, err := email2.SendMail(email)
	if err != nil {
		return "", err
	}
	return code, nil
}

// ContactMail 接收注册验证码逻辑处理函数
func (gm Account) ContactMail(name string, email string, subject string, message string) error {
	return email2.SendContactMail(name, email, subject, message)
}

// Logout 注销逻辑处理函数
func (la Account) Logout(username string) interface{} {
	var tr redis.TokenRedis
	username1 := tr.GetToken(username)
	return username1
}

// SetLastLogin 修改最后登录时间
func (la Account) SetLastLogin(username string, lastLogin time.Time) error {
	return mysql.UpdateLastLogin(username, lastLogin)
}

// SetAccount 修改用户账号信息
func (la Account) SetAccount(information map[string]string) (err error) {
	var user models.User
	if information["id"] != "" {
		id, _ := strconv.ParseInt(information["id"], 10, 64)
		user, err = mysql.UserIDVerity(id)
		if err != nil {
			return errors.New("该用户不存在")
		}
	} else {
		return errors.New("该用户不存在")
	}
	for key, info := range information {
		if key == "username" {
			user.Username = info
		}
		if key == "password" {
			pwd := middlewares.Encode(info)
			user.Password = pwd
		}
		if key == "email" {
			user.Email = info
		}
		if key == "phone" {
			user.Phone = info
		}
	}
	return mysql.UpdateAccount(user)
}

// SetUserInformation 修改用户详细信息
func (la Account) SetUserInformation(information map[string]string) (err error) {
	var user models.UserInformation
	if information["id"] != "" {
		id, _ := strconv.ParseInt(information["id"], 10, 64)
		user, err = mysql.UserIDInfoVerity(id)
		if err != nil {
			return errors.New("该用户不存在")
		}
	} else {
		return errors.New("该用户不存在")
	}
	for key, info := range information {
		if info != "" || len(info) != 0 {
			if key == "nickname" {
				user.Nickname = info
			}
			if key == "age" {
				num, _ := strconv.Atoi(info)
				user.Age = num
			}
			if key == "sex" {
				user.Sex = info
			}
			if key == "brithDate" {
				layout := "2006-01-02" // 对应的时间格式
				resultTime, _ := time.Parse(layout, info)
				user.BrithDate = resultTime
			}
			if key == "avatar" { //修改

				user.Avatar = info
			}
			if key == "biography" {
				user.Biography = info
			}
			if key == "address" {
				user.Address = info
			}
		}
	}
	return mysql.UpdateUserInformation(user)
}

// DeleteAccount 删除账号
func (la Account) DeleteAccount(username string) error {
	return mysql.DeleteAccount(username)
}

// UserAccountInfo 获取用户账号和详细信息
func (la Account) UserAccountInfo(id int64) (models.User, models.UserInformation, error) {
	return mysql.GetUserAccountInfo(id)
}
