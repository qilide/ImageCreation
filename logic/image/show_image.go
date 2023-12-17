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