package mine

import (
	"ImageCreation/dao/mysql"
	"ImageCreation/models"
	"strconv"
)

type MeImageInfo struct {
}

// MeLikeImage 获取我的点赞图片
func (mii MeImageInfo) MeLikeImage(page string) ([]models.Image, int, error) {
	page1, _ := strconv.ParseInt(page, 10, 64)
	return mysql.GetLikeImage(page1)
}

// MeCollectImage 获取我的收藏图片
func (mii MeImageInfo) MeCollectImage(page string) ([]models.Image, int, error) {
	page1, _ := strconv.ParseInt(page, 10, 64)
	return mysql.GetCollectImage(page1)
}

// MeBrowseImage 获取我的浏览图片
func (mii MeImageInfo) MeBrowseImage(page string) ([]models.Image, int, error) {
	page1, _ := strconv.ParseInt(page, 10, 64)
	return mysql.GetBrowseImage(page1)
}

// MeScoreImage 获取我的评分图片
func (mii MeImageInfo) MeScoreImage(page string) ([]models.Image, int, error) {
	page1, _ := strconv.ParseInt(page, 10, 64)
	return mysql.GetScoreImage(page1)
}
