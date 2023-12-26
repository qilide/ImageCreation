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
func (si ShowImage) GalleryImage(label string) ([]models.Image, error) {
	return mysql.GetGalleryImage(label)
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
