package mysql

import "ImageCreation/models"

// GetAuthors 获取所有摄影师信息
func GetAuthors() ([]models.UserInformation, error) {
	var image []models.UserInformation
	err := db.Table("image").Where("is_active = 1").Order("total_likes DESC").Find(&image).Error
	return image, err
}

// GetAuthorInfo 获取摄影师详细信息
func GetAuthorInfo(id int64) (models.UserInformation, error) {
	var userInfo models.UserInformation
	err := db.Table("user_information").Where("is_active = 1").Where("user_id = ? ", id).Find(&userInfo).Error
	return userInfo, err
}
