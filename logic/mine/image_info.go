package mine

import (
	"ImageCreation/dao/mysql"
	"ImageCreation/models"
	"strconv"
)

type MeImageInfo struct {
}

// MeLikeImage 获取我的点赞图片
func (mii MeImageInfo) MeLikeImage(id string, page string) ([]models.Image, int, error) {
	page1, _ := strconv.ParseInt(page, 10, 64)
	return mysql.GetLikeImage(id, page1)
}

// MeCollectImage 获取我的收藏图片
func (mii MeImageInfo) MeCollectImage(id string, page string) ([]models.Image, int, error) {
	page1, _ := strconv.ParseInt(page, 10, 64)
	return mysql.GetCollectImage(id, page1)
}

// MeBrowseImage 获取我的浏览图片
func (mii MeImageInfo) MeBrowseImage(id string, page string) ([]models.Image, int, error) {
	page1, _ := strconv.ParseInt(page, 10, 64)
	return mysql.GetBrowseImage(id, page1)
}

// MeScoreImage 获取我的评分图片
func (mii MeImageInfo) MeScoreImage(id string, page string) ([]models.Image, int, error) {
	page1, _ := strconv.ParseInt(page, 10, 64)
	return mysql.GetScoreImage(id, page1)
}
