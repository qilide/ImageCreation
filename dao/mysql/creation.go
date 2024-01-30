package mysql

import "ImageCreation/models"

// CreateUploadImage 新建上传图片
func CreateUploadImage(imageInfo models.Image) (models.Image, error) {
	err := db.Table("image").Create(&imageInfo).Error
	return imageInfo, err
}

// CreateCreationImage 新建创作图片
func CreateCreationImage(imageInfo models.Creation) (models.Creation, error) {
	err := db.Table("creation").Create(&imageInfo).Error
	return imageInfo, err
}

// CreateGeneralRecognition 新建通用物体和场景识别
func CreateGeneralRecognition(imageInfo models.Recognition) (models.Recognition, error) {
	err := db.Table("recognition").Create(&imageInfo).Error
	return imageInfo, err
}
