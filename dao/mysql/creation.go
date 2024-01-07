package mysql

import "ImageCreation/models"

// CreateUploadImage 新建上传图片
func CreateUploadImage(imageInfo models.Image) (models.Image, error) {
	err := db.Table("image").Create(&imageInfo).Error
	return imageInfo, err
}
