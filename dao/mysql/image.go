package mysql

import "ImageCreation/models"

// GetIndexImage 获取所有图片信息
func GetIndexImage() ([]models.Image, error) {
	var image []models.Image
	err := db.Table("image").Where("is_active = 1").Limit(16).Find(&image).Error
	return image, err
}
