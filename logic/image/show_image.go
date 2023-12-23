package image

import (
	"ImageCreation/dao/mysql"
	"ImageCreation/models"
)

type ShowImage struct {
}

// IndexImage 展示主页图片
func (si ShowImage) IndexImage() ([]models.Image, error) {
	return mysql.GetIndexImage()
}

// ImageInfo 获取图片详细信息
func (si ShowImage) ImageInfo(id int64) (models.Image, models.UserInformation, error) {
	return mysql.GetImageInfo(id)
}

// GalleryImage 展示主题图片
func (si ShowImage) GalleryImage(label string) ([]models.Image, error) {
	return mysql.GetGalleryImage(label)
}
