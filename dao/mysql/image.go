package mysql

import (
	"ImageCreation/models"
)

// GetIndexImage 获取所有图片信息
func GetIndexImage() ([]models.Image, error) {
	var image []models.Image
	err := db.Table("image").Where("is_active = 1").Order("RAND()").Limit(20).Find(&image).Error
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
	subQuery := db.Table("browse").
		Select("MAX(view_time) as max_view_time, user_id, image_id").
		Where("is_active = 1").
		Group("user_id, image_id").
		SubQuery()
	err := db.Table("browse").
		Select("*").
		Joins("JOIN (?) AS b ON browse.user_id = b.user_id AND browse.image_id = b.image_id AND browse.view_time = b.max_view_time", subQuery).
		Where("browse.is_active = 1").
		Order("browse.view_time DESC").
		Limit(20).
		Offset((page - 1) * 20).
		Find(&browses).Error
	for _, browse := range browses {
		var image models.Image
		err = db.Table("image").Where("id = ?", browse.ImageId).Where("is_active = 1").Find(&image).Error
		images = append(images, image)
	}
	err = db.Table("browse").
		Select("*").
		Joins("JOIN (?) AS b ON browse.user_id = b.user_id AND browse.image_id = b.image_id AND browse.view_time = b.max_view_time", subQuery).
		Where("browse.is_active = 1").Count(&total).Error
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

// CheckImageLike 检查是否有图片点赞数据
func CheckImageLike(userID, imageID int64) (models.Like, error) {
	var like models.Like
	err := db.Table("like").Where("user_id = ?", userID).Where("image_id = ?", imageID).Where("is_active = 1").Find(&like).Error
	return like, err
}

// UpdateImageLike 修改图片点赞数据
func UpdateImageLike(like models.Like) error {
	return db.Table("like").Save(&like).Error
}

// CreateImageLike 新建图片点赞数据
func CreateImageLike(like models.Like) error {
	return db.Table("like").Create(&like).Error
}

// UpdateImageInfo 修改图片信息
func UpdateImageInfo(image models.Image) (models.Image, error) {
	err := db.Table("image").Save(&image).Error
	return image, err
}

// GetImageSingleInfo 获取图片个体详细信息
func GetImageSingleInfo(id int64) (models.Image, error) {
	var image models.Image
	err := db.Table("image").Where("is_active = 1").Where("id = ? ", id).Find(&image).Error
	return image, err
}

// CheckImageCollect 检查是否有图片收藏数据
func CheckImageCollect(userID, imageID int64) (models.Collect, error) {
	var collect models.Collect
	err := db.Table("collect").Where("user_id = ?", userID).Where("image_id = ?", imageID).Where("is_active = 1").Find(&collect).Error
	return collect, err
}

// UpdateImageCollect 修改图片收藏数据
func UpdateImageCollect(collect models.Collect) error {
	return db.Table("collect").Save(&collect).Error
}

// CreateImageCollect 新建图片收藏数据
func CreateImageCollect(collect models.Collect) error {
	return db.Table("collect").Create(&collect).Error
}

// CheckImageIsLike 检查是否有图片点赞数据,并点赞
func CheckImageIsLike(userID, imageID int64) (models.Like, error) {
	var like models.Like
	err := db.Table("like").Where("user_id = ?", userID).Where("image_id = ?", imageID).Where("is_active = 1").Where("is_like = 1").Find(&like).Error
	return like, err
}

// CheckImageIsCollect 检查是否有图片收藏数据,并收藏
func CheckImageIsCollect(userID, imageID int64) (models.Collect, error) {
	var collect models.Collect
	err := db.Table("collect").Where("user_id = ?", userID).Where("image_id = ?", imageID).Where("is_active = 1").Where("is_collect = 1").Find(&collect).Error
	return collect, err
}

// CheckImageIsScore 检查是否有图片评分数据,并评分
func CheckImageIsScore(userID, imageID int64) (models.Score, error) {
	var score models.Score
	err := db.Table("score").Where("user_id = ?", userID).Where("image_id = ?", imageID).Where("is_active = 1").Where("score <> ?", 0).Find(&score).Error
	return score, err
}

// CreateImageBrowse 新建图片浏览数据
func CreateImageBrowse(browse models.Browse) error {
	return db.Table("browse").Create(&browse).Error
}

// UpdateImageBrowse 修改图片浏览数据
func UpdateImageBrowse(imageID int64) error {
	var image models.Image
	db.Table("image").Where("id = ?", imageID).Where("is_active = 1").Find(&image)
	image.BrowseCount += 1
	return db.Table("image").Save(&image).Error
}
