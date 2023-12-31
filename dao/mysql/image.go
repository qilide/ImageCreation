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

// GetLikeImage 获取我的点赞图片
func GetLikeImage(id string, page int64) ([]models.Image, int, error) {
	var likes []models.Like
	var images []models.Image
	var total int
	err := db.Table("like").Where("user_id = ?", id).Where("is_active = 1").Where("is_like = 1").Order("update_time DESC").Limit(20).Offset((page - 1) * 20).Find(&likes).Error
	for _, like := range likes {
		var image models.Image
		err = db.Table("image").Where("id = ?", like.ImageId).Where("is_active = 1").Find(&image).Error
		images = append(images, image)
	}
	err = db.Table("like").Where("user_id = ?", id).Where("is_active = 1").Where("is_like = 1").Count(&total).Error
	return images, total, err
}

// GetCollectImage 获取我的收藏图片
func GetCollectImage(id string, page int64) ([]models.Image, int, error) {
	var collects []models.Collect
	var images []models.Image
	var total int
	err := db.Table("collect").Where("user_id = ?", id).Where("is_active = 1").Where("is_collect = 1").Order("update_time DESC").Limit(20).Offset((page - 1) * 20).Find(&collects).Error
	for _, collect := range collects {
		var image models.Image
		err = db.Table("image").Where("id = ?", collect.ImageId).Where("is_active = 1").Find(&image).Error
		images = append(images, image)
	}
	err = db.Table("collect").Where("user_id = ?", id).Where("is_active = 1").Where("is_collect = 1").Count(&total).Error
	return images, total, err
}

// GetBrowseImage 获取我的浏览图片
func GetBrowseImage(id string, page int64) ([]models.Image, int, error) {
	var browses []models.Browse
	var images []models.Image
	var total int
	err := db.Table("browse").Where("user_id = ?", id).Where("is_active = 1").Order("view_time DESC").Limit(20).Offset((page - 1) * 20).Find(&browses).Error
	for _, browse := range browses {
		var image models.Image
		err = db.Table("image").Where("id = ?", browse.ImageId).Where("is_active = 1").Find(&image).Error
		images = append(images, image)
	}
	err = db.Table("browse").Where("user_id = ?", id).Where("is_active = 1").Count(&total).Error
	return images, total, err
}

// GetScoreImage 获取我的评分图片
func GetScoreImage(id string, page int64) ([]models.Image, int, error) {
	var scores []models.Score
	var images []models.Image
	var total int
	err := db.Table("score").Where("user_id = ?", id).Where("is_active = 1").Order("update_time DESC").Limit(20).Offset((page - 1) * 20).Find(&scores).Error
	for _, score := range scores {
		var image models.Image
		err = db.Table("image").Where("id = ?", score.ImageId).Where("is_active = 1").Find(&image).Error
		images = append(images, image)
	}
	err = db.Table("score").Where("user_id = ?", id).Where("is_active = 1").Count(&total).Error
	return images, total, err
}
