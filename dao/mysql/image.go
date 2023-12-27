package mysql

import (
	"ImageCreation/models"
)

// GetIndexImage 获取所有图片信息
func GetIndexImage() ([]models.Image, error) {
	var image []models.Image
	err := db.Table("image").Where("is_active = 1").Order("score DESC").Limit(20).Find(&image).Error
	return image, err
}

// GetImageInfo 获取图片详细信息
func GetImageInfo(id int64) (models.Image, models.UserInformation, error) {
	var image models.Image
	var userInfo models.UserInformation
	err := db.Table("image").Where("is_active = 1").Where("id = ? ", id).Find(&image).Error
	err = db.Table("user_information").Where("is_active = 1").Where("user_id = ? ", image.UserID).Find(&userInfo).Error
	return image, userInfo, err
}

// GetGalleryImage 获取主题图片信息
func GetGalleryImage(label string, page int64) ([]models.Image, int, error) {
	var image []models.Image
	var total int
	err := db.Table("image").Where("is_active = 1").Where("label= ?", label).Order("score DESC").Limit(20).Offset((page - 1) * 20).Find(&image).Error
	err = db.Table("image").Where("is_active = 1").Where("label= ?", label).Count(&total).Error
	return image, total, err
}
