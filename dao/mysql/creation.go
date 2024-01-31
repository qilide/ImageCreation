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

// GetCreateImage 查询创作图片
func GetCreateImage(imageId string) (models.Creation, error) {
	var imageInfo models.Creation
	err := db.Table("creation").Where("id=?", imageId).Where("is_active = 1").Find(&imageInfo).Error
	return imageInfo, err
}

// GetCreations 获取所有创作图片
func GetCreations() ([]models.Creation, error) {
	var creations []models.Creation
	err := db.Table("creation").Where("is_active = 1").Where("operation IN (?, ?, ?, ?, ?)", "自定义图像风格", "图像风格转换", "人像动漫化", "图像清晰度增强", "图像对比度增强").Order("update_time DESC").Find(&creations).Error
	return creations, err
}
