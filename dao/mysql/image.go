package mysql

import "ImageCreation/models"

// GetIndexImage 获取所有图片信息
func GetIndexImage() ([]models.Image, error) {
	var image []models.Image
	err := db.Table("image").Where("is_active = 1").Limit(16).Find(&image).Error
	return image, err
}

// GetImageInfo 获取图片详细信息
func GetImageInfo(id int64) (models.Image, error) {
	var image models.Image
	err := db.Table("image").Where("is_active = 1").Where("id = ? ", id).Find(&image).Error
	return image, err
}
