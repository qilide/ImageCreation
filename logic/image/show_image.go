package image

import (
	"ImageCreation/dao/mysql"
	"ImageCreation/models"
	"encoding/json"
	"errors"
	"fmt"
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

type ApiKey struct {
	Key            string
	RemainingCalls int
	LastUsed       time.Time
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
	apiKeys := []ApiKey{
		{"6ZUUwUDCk_woNQeTa08uok7E67XdTMuJXzNRNkNC2gQ", 50, time.Now()},
		{"soSfySaTOOCDbT4c6x2uAtHT9vAv_vcREfaHMUm5TG8", 50, time.Now()},
		{"bdGWF-_MVxtzgiHv2V6PSSD_mDu9Ga88NRWYDABHU74", 50, time.Now()},
		{"VhKv_iprwfuhUru9sL7MrE8iAOqo9iCzU6OyYQp1xAM", 50, time.Now()},
		{"js1UEyCaFesf3B66Ie7WPxQpWIStlhLRPh2dgriP1uo", 50, time.Now()},
		{"TtVo04sWt-bL2CDFagV46b7ZtXP-_LfGrkSd4q8a2RM", 50, time.Now()},
		{"8dws2jf0LMNW4pd573pC582cvuN0QH6h2hTk8z41SYc", 50, time.Now()},
		{"OHc2ydJXgBs1klbPC8kspt3PrswLpEs0n-qBflwsZzw", 50, time.Now()},
		{"xflPjF7W10dlu3ajhPXK4_IeW5flUyfxb3yfRmM1WfQ", 50, time.Now()},
		{"B4dTWWYOFcnwTlYYWDJb8ZhnF8hAqsl5RXa_KVhENrM", 50, time.Now()},
	}

	// 遍历每种类型进行请求
	for _, key := range apiKeys {
		if key.RemainingCalls > 0 {
			// 构建请求
			url := fmt.Sprintf("https://api.unsplash.com/search/photos?page=%d&query=%s&per_page=%s", page, label, pageCount)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Println("Error creating request:", err)
				return responseData, err
			}
			// 设置请求头部，添加访问密钥
			req.Header.Set("Authorization", "Client-ID "+key.Key)
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
				return responseData, nil
			} else {
				fmt.Println("Failed to fetch photos. Status code:", resp.StatusCode)
			}
			key.RemainingCalls--
		} else {
			return responseData, errors.New("API key 已达到调用限制")
		}
	}
	return responseData, nil
}
