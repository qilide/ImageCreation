package creation

import (
	"ImageCreation/dao/mysql"
	"ImageCreation/models"
	"ImageCreation/pkg/snowflake"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	image2 "image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	url2 "net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type ShowCreation struct {
}

// SetUploadImage 上传图片
func (sc ShowCreation) SetUploadImage(c *gin.Context, userID string, image multipart.File, header *multipart.FileHeader) (models.Image, image2.Config, error) {
	var sf snowflake.Snowflake
	id := sf.NextVal()
	strInt64 := strconv.FormatInt(id, 10)
	id16, _ := strconv.Atoi(strInt64)
	userId, _ := strconv.ParseInt(userID, 10, 64)
	defer image.Close()
	// 将文件保存到指定的路径，这里保存在当前目录下的 uploads 文件夹中
	savePath := "./static/creation/origin/" + strInt64 + "/"
	savePath1 := savePath + header.Filename
	path := "/static/creation/origin/" + strInt64 + "/" + header.Filename
	// 如果目录不存在，则创建目录
	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		err := os.Mkdir(savePath, 0755)
		if err != nil {
			return models.Image{}, image2.Config{}, err
		}
	}
	err := c.SaveUploadedFile(header, savePath1)
	if err != nil {
		return models.Image{}, image2.Config{}, err
	}
	// 获取图片信息
	img, _, err := image2.DecodeConfig(image)
	if err != nil {
		return models.Image{}, img, err
	}
	imageInfo := models.Image{
		ID:         id16,
		UserID:     int(userId),
		Path:       path,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		IsActive:   1,
		IsCreate:   1,
	}
	info, err := mysql.CreateUploadImage(imageInfo)
	return info, img, err
}

// ImageEnhancement 图像增强接口
// 适用于 图像无损放大 图像去雾 图像对比度增强 图像清晰度增强 图像色彩增强等功能
func (sc ShowCreation) ImageEnhancement(userId, imageId, imagePath, apiUrl, operation string) (models.Creation, error) {
	url := apiUrl + GetAccessToken()
	imageBase64, err := GetFileContentAsBase64(imagePath)
	if err != nil {
		return models.Creation{}, err
	}
	payload := strings.NewReader(fmt.Sprintf("image=%s", imageBase64))
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return models.Creation{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return models.Creation{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.Creation{}, err
	}
	// 解析 JSON 数据
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return models.Creation{}, err
	}
	// 获取 image 字段的值
	imageResultBase64, ok := data["image"].(string)
	if !ok {
		return models.Creation{}, errors.New("无法获取 image 字段的值")
	}
	var sf snowflake.Snowflake
	id := sf.NextVal()
	strInt64 := strconv.FormatInt(id, 10)
	resultImageId, _ := strconv.Atoi(strInt64[:len(strInt64)-2])
	path := "./static/creation/created/" + strconv.Itoa(resultImageId) + ".jpg"
	pathSave := "/static/creation/created/" + strconv.Itoa(resultImageId) + ".jpg"
	err = SaveBase64ToImage(path, imageResultBase64)
	if err != nil {
		return models.Creation{}, err
	}
	userId1, err := strconv.ParseInt(userId, 10, 64)
	imageId1, err := strconv.ParseInt(imageId, 10, 64)
	imageCreation := models.Creation{
		ID:         resultImageId,
		UserId:     int(userId1),
		ImageId:    int(imageId1),
		Path:       pathSave,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		IsActive:   1,
		Operation:  operation,
	}
	imageInfo, err := mysql.CreateCreationImage(imageCreation)
	if err != nil {
		return models.Creation{}, err
	}
	return imageInfo, nil
}

// ImageGeneralRecognition 通用物体和场景识别
func (sc ShowCreation) ImageGeneralRecognition(userId, imageId, imagePath string) (models.Recognition, error) {
	url := "https://aip.baidubce.com/rest/2.0/image-classify/v2/advanced_general?access_token=" + GetAccessToken()
	imageBase64, err := GetFileContentAsBase64(imagePath)
	if err != nil {
		return models.Recognition{}, err
	}
	payload := strings.NewReader(fmt.Sprintf("image=%s&baike_num=1", imageBase64))
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return models.Recognition{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return models.Recognition{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.Recognition{}, err
	}
	// 解析 JSON 数据
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return models.Recognition{}, err
	}
	results, ok := data["result"].([]interface{})
	if !ok {
		return models.Recognition{}, errors.New("无法获取 result 数组")
	}
	var baikeURL, description, imageURL, score, root, keyword interface{}
	// 遍历结果数组
	for _, result := range results {
		// 将每个元素断言为 map[string]interface{}
		resultMap, ok := result.(map[string]interface{})
		if !ok {
			continue
		}
		// 获取 "baike_info" map
		baikeInfo, ok := resultMap["baike_info"].(map[string]interface{})
		if ok {
			// 在 baikeInfo 中获取需要的信息
			baikeURL = baikeInfo["baike_url"]
			description = baikeInfo["description"]
			imageURL = baikeInfo["image_url"]
			score = resultMap["score"]
			root = resultMap["root"]
			keyword = resultMap["keyword"]
		} else {
			continue
		}
	}
	if root == nil || keyword == nil || score == nil {
		return models.Recognition{}, err
	}
	var sf snowflake.Snowflake
	id := sf.NextVal()
	strInt64 := strconv.FormatInt(id, 10)
	resultImageId, _ := strconv.Atoi(strInt64[:len(strInt64)-2])
	userId1, err := strconv.ParseInt(userId, 10, 64)
	imageId1, err := strconv.ParseInt(imageId, 10, 64)
	imageRecognition := models.Recognition{
		ID:          resultImageId,
		UserId:      int(userId1),
		ImageId:     int(imageId1),
		Score:       score,
		Root:        root,
		BaikeUrl:    baikeURL,
		ImageUrl:    imageURL,
		Description: description,
		Keyword:     keyword,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		IsActive:    1,
	}
	imageInfo, err := mysql.CreateGeneralRecognition(imageRecognition)
	if err != nil {
		return models.Recognition{}, err
	}
	return imageInfo, nil
}

// ImageEffects 图像特效接口
// 适用于 图像风格转换 人像动漫化 自定义图像风格等功能
func (sc ShowCreation) ImageEffects(userId, imageId, url, operation string, payload *strings.Reader) (models.Creation, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return models.Creation{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return models.Creation{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.Creation{}, err
	}
	// 解析 JSON 数据
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return models.Creation{}, err
	}
	// 获取 image 字段的值
	imageResultBase64, ok := data["image"].(string)
	if !ok {
		return models.Creation{}, errors.New("无法获取 image 字段的值")
	}
	var sf snowflake.Snowflake
	id := sf.NextVal()
	strInt64 := strconv.FormatInt(id, 10)
	resultImageId, _ := strconv.Atoi(strInt64[:len(strInt64)-2])
	path := "./static/creation/created/" + strconv.Itoa(resultImageId) + ".jpg"
	pathSave := "/static/creation/created/" + strconv.Itoa(resultImageId) + ".jpg"
	err = SaveBase64ToImage(path, imageResultBase64)
	if err != nil {
		return models.Creation{}, err
	}
	userId1, err := strconv.ParseInt(userId, 10, 64)
	imageId1, err := strconv.ParseInt(imageId, 10, 64)
	imageCreation := models.Creation{
		ID:         resultImageId,
		UserId:     int(userId1),
		ImageId:    int(imageId1),
		Path:       pathSave,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		IsActive:   1,
		Operation:  operation,
	}
	imageInfo, err := mysql.CreateCreationImage(imageCreation)
	if err != nil {
		return models.Creation{}, err
	}
	return imageInfo, nil
}

// GetFileContentAsBase64 获取文件base64编码
// param string  path 文件路径
// return string base64编码信息，不带文件头
func GetFileContentAsBase64(path string) (string, error) {
	srcByte, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return url2.QueryEscape(base64.StdEncoding.EncodeToString(srcByte)), nil
}

// GetAccessToken 使用 AK，SK 生成鉴权签名（Access Token）
// return string 鉴权签名信息（Access Token）
func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	API_KEY := viper.GetString("creationKeys.api_key")
	SECRET_KEY := viper.GetString("creationKeys.secret_key")
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", API_KEY, SECRET_KEY)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	accessTokenObj := map[string]string{}
	json.Unmarshal(body, &accessTokenObj)
	return accessTokenObj["access_token"]
}

// SaveBase64ToImage 根据Base64字符串保存图片到指定位置
func SaveBase64ToImage(outputPath string, base64String string) error {
	// 将Base64字符串解码为字节数组
	imageData, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return err
	}
	// 保存图像文件
	err = ioutil.WriteFile(outputPath, imageData, 0644)
	if err != nil {
		return err
	}
	return nil
}
