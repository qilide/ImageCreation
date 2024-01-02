package image

import (
	"ImageCreation/dao/mysql"
	"ImageCreation/dao/redis"
	"ImageCreation/models"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type ShowImage struct {
}

type PhotoInfo struct {
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Results    []struct {
		ID          string `json:"id"`
		Description string `json:"description"`
		Urls        struct {
			Regular string `json:"regular"`
		} `json:"urls"`
	} `json:"results"`
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
func (si ShowImage) GalleryImage(label string, page string) ([]models.Image, int, error) {
	page1, _ := strconv.ParseInt(page, 10, 64)
	return mysql.GetGalleryImage(label, page1)
}

// GetSearchImage 搜索图片
func (si ShowImage) GetSearchImage(label string) (PhotoInfo, error) {
	page := 1
	pageCount := "28"
	var responseData PhotoInfo
	// 设置每种类型获取的图片数量
	var sai redis.SearchApiRedis
	apiKeys, err := sai.SetApiInfo()
	if err != nil {
		return responseData, err
	}
	if len(apiKeys) == 0 {
		return responseData, errors.New("搜索功能已上限，请下一小时重试！")
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(apiKeys))
	key := apiKeys[randomIndex]
	// 构建请求
	url := fmt.Sprintf("https://api.unsplash.com/search/photos?page=%d&query=%s&per_page=%s", page, label, pageCount)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return responseData, err
	}
	// 设置请求头部，添加访问密钥
	req.Header.Set("Authorization", "Client-ID "+key)
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return responseData, err
	}
	defer resp.Body.Close()

	// 解析响应并获取照片信息
	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
			fmt.Println("Error decoding response:", err)
			return responseData, err
		}
		if err = sai.UpdateApiUseCount(key); err != nil {
			return responseData, err
		}
		return responseData, nil
	} else {
		fmt.Println("Failed to fetch photos. Status code:", resp.StatusCode)
	}
	return responseData, err
}

// ImageToLike 图片点赞操作
func (si ShowImage) ImageToLike(userID, imageID, isLike string) (models.Image, error) {
	userID1, _ := strconv.ParseInt(userID, 10, 64)
	imageID1, _ := strconv.ParseInt(imageID, 10, 64)
	isLike1, _ := strconv.ParseInt(isLike, 10, 64)
	image, err := mysql.GetImageSingleInfo(imageID1)
	if err != nil {
		return image, err
	}
	like, err := mysql.CheckImageLike(userID1, imageID1)
	if err != nil {
		newLike := models.Like{
			UserId:     int(userID1),
			ImageId:    int(imageID1),
			IsLike:     int(isLike1),
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			IsActive:   1,
		}
		if err := mysql.CreateImageLike(newLike); err != nil {
			return image, err
		}
		if isLike1 != 0 {
			image.LikeCount += int(isLike1)
		}
		return mysql.UpdateImageInfo(image)
	} else {
		originIsLike := like.IsLike
		if isLike1 != 0 && originIsLike == 0 {
			like.IsLike = int(isLike1)
			like.UpdateTime = time.Now()
			if err := mysql.UpdateImageLike(like); err != nil {
				return image, err
			}
			image.LikeCount += 1
		} else if isLike1 == 0 && originIsLike == 1 {
			like.IsLike = int(isLike1)
			like.UpdateTime = time.Now()
			if err := mysql.UpdateImageLike(like); err != nil {
				return image, err
			}
			image.LikeCount -= 1
		}
		return mysql.UpdateImageInfo(image)
	}
}

// ImageToCollect 图片收藏操作
func (si ShowImage) ImageToCollect(userID, imageID, isCollect string) (models.Image, error) {
	userID1, _ := strconv.ParseInt(userID, 10, 64)
	imageID1, _ := strconv.ParseInt(imageID, 10, 64)
	isCollect1, _ := strconv.ParseInt(isCollect, 10, 64)
	image, err := mysql.GetImageSingleInfo(imageID1)
	if err != nil {
		return image, err
	}
	collect, err := mysql.CheckImageCollect(userID1, imageID1)
	if err != nil {
		newCollect := models.Collect{
			UserId:     int(userID1),
			ImageId:    int(imageID1),
			IsCollect:  int(isCollect1),
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			IsActive:   1,
		}
		if err := mysql.CreateImageCollect(newCollect); err != nil {
			return image, err
		}
		if isCollect1 != 0 {
			image.CollectCount += int(isCollect1)
		}
		return mysql.UpdateImageInfo(image)
	} else {
		originIsCollect := collect.IsCollect
		if isCollect1 != 0 && originIsCollect == 0 {
			collect.IsCollect = int(isCollect1)
			collect.UpdateTime = time.Now()
			if err := mysql.UpdateImageCollect(collect); err != nil {
				return image, err
			}
			image.CollectCount += 1
		} else if isCollect1 == 0 && originIsCollect == 1 {
			collect.IsCollect = int(isCollect1)
			collect.UpdateTime = time.Now()
			if err := mysql.UpdateImageCollect(collect); err != nil {
				return image, err
			}
			image.CollectCount -= 1
		}
		return mysql.UpdateImageInfo(image)
	}
}

//// ImageToScore 图片评分操作
//func (si ShowImage) ImageToScore(userID, imageID, Score string) (models.Image, error) {
//	userID1, _ := strconv.ParseInt(userID, 10, 64)
//	imageID1, _ := strconv.ParseInt(imageID, 10, 64)
//	Score1, _ := strconv.ParseFloat(Score, 10)
//	image, err := mysql.GetImageSingleInfo(imageID1)
//	if err != nil {
//		return image, err
//	}
//	score, err := mysql.CheckImageScore(userID1, imageID1)
//	if err != nil {
//		newScore := models.Score{
//			UserId:     int(userID1),
//			ImageId:    int(imageID1),
//			Score:      int(Score1),
//			CreateTime: time.Now(),
//			UpdateTime: time.Now(),
//			IsActive:   1,
//		}
//		if err := mysql.CreateImageScore(newScore); err != nil {
//			return image, err
//		}
//		if Score1 != 0 {
//			image.Score += Score1
//		}
//		return mysql.UpdateImageInfo(image)
//	} else {
//		originIsScore := score.Score
//		score.Score = int(Score1)
//		if err := mysql.UpdateImageScore(score); err != nil {
//			return image, err
//		}
//		if Score1 != 0 && originIsScore == 0 {
//			image.Score += 1
//		} else if Score1 == 0 && originIsScore == 1 {
//			image.Score -= 1
//		}
//		return mysql.UpdateImageInfo(image)
//	}
//}

// ImageToOperation 查询当前用户对图片的操作
func (si ShowImage) ImageToOperation(userID, imageID string) (string, string, string) {
	userID1, _ := strconv.ParseInt(userID, 10, 64)
	imageID1, _ := strconv.ParseInt(imageID, 10, 64)
	var isLike, isCollect, isScore string
	like, err := mysql.CheckImageIsLike(userID1, imageID1)
	if err == nil || like.IsLike != 0 {
		isLike = "like"
	}
	collect, err := mysql.CheckImageIsCollect(userID1, imageID1)
	if err == nil || collect.IsCollect != 0 {
		isCollect = "collect"
	}
	score, err := mysql.CheckImageIsScore(userID1, imageID1)
	if err == nil || score.Score != 0 {
		isScore = "score"
	}
	return isLike, isCollect, isScore
}

// ImageToBrowse 查询当前用户对图片的操作
func (si ShowImage) ImageToBrowse(userID, imageID string) error {
	userID1, _ := strconv.ParseInt(userID, 10, 64)
	imageID1, _ := strconv.ParseInt(imageID, 10, 64)
	browse := models.Browse{
		UserId:   int(userID1),
		ImageId:  int(imageID1),
		ViewTime: time.Now(),
		IsActive: 1,
	}
	err := mysql.CreateImageBrowse(browse)
	if err != nil {
		return err
	}
	err = mysql.UpdateImageBrowse(imageID1)
	return err
}
