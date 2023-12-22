package mysql

import (
	"ImageCreation/models"
	"strconv"
	"time"
)

// UserNameVerity 用户名参数数据库查询
func UserNameVerity(username string) (models.User, error) {
	var user models.User
	err := db.Table("user").Where("is_active = ?", 1).Where("username = ?", username).Find(&user).Error
	return user, err
}

// CreateUser 创建数据库数据
func CreateUser(id int, username string, password string, email string, createTime time.Time) (models.User, error) {
	user := models.User{
		ID:          id,
		Username:    username,
		Password:    password,
		Email:       email,
		CreateTime:  createTime,
		UpdateTime:  createTime,
		LastLogin:   createTime,
		Phone:       "",
		IsActive:    1,
		IsSuperuser: 0,
	}
	err := db.Table("user").Create(&user).Error
	return user, err
}

// CreateUserInformation 创建用户详细数据
func CreateUserInformation(userID int) error {
	user := models.UserInformation{
		UserID:    userID,
		IsActive:  1,
		Nickname:  "匿名用户" + strconv.Itoa(userID),
		Biography: "生活是一部令人着迷的故事，而你是其中独一无二的篇章。在这个广袤的舞台上，你是自己故事的作者，用坚韧和温暖编织出独特的篇章，将每一刻都演绎成无比精彩的章节",
		Avatar:    "assets/img/testimonials/testimonials-2.jpg",
	}
	return db.Table("user_information").Create(&user).Error
}

// UpdateLastLogin 修改最后登录时间
func UpdateLastLogin(username string, lastLogin time.Time) error {
	var user models.User
	db.Table("user").Where("is_active = ?", 1).Where("username = ?", username).Find(&user)
	user.LastLogin = lastLogin
	return db.Table("user").Save(&user).Error
}

// UpdateAccount 修改用户账户信息
func UpdateAccount(user models.User) error {
	return db.Table("user").Save(&user).Error
}

// UpdateUserInformation 修改用户详细信息
func UpdateUserInformation(userInformation models.UserInformation) error {
	return db.Table("user_information").Save(&userInformation).Error
}

// DeleteAccount 删除用户信息
func DeleteAccount(username string) error {
	var user models.User
	db.Table("user").Where("username=?", username).Find(&user)
	user.IsActive = 0
	if err := db.Table("user").Save(&user).Error; err != nil {
		return err
	}
	var userInformation models.UserInformation
	db.Table("user_information").Where("user_id=?", user.ID).Find(&userInformation)
	userInformation.IsActive = 0
	return db.Table("user_information").Save(&userInformation).Error
}
