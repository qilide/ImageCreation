package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strings"
)

const API_KEY = "SmZui0dD3Z5soC04Q0OP6V13"
const SECRET_KEY = "U9yISUIZBzzuP9KmpwlDDDGDRXiw9Gjx"

// 图像无损放大
func main() {
	url := "https://aip.baidubce.com/rest/2.0/image-process/v1/image_quality_enhance?access_token=" + GetAccessToken()
	// image 可以通过 GetFileContentAsBase64("C:\fakepath\_o412MhXDFw.jpg") 方法获取
	imageBase64 := GetFileContentAsBase64("./static/creation/origin/269947940735287296/1E2OXICYBYE.jpg")
	payload := strings.NewReader(fmt.Sprintf("image=%s", imageBase64))
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 解析 JSON 数据
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("JSON 解析失败:", err)
		return
	}
	// 获取 image 字段的值
	imageResultBase64, ok := data["image"].(string)
	if !ok {
		fmt.Println("无法获取 image 字段的值")
		return
	}
	SaveBase64ToImage("output_image.jpg", imageResultBase64)
}

/**
 * 获取文件base64编码
 * param string  path 文件路径
 * return string base64编码信息，不带文件头
 */
func GetFileContentAsBase64(path string) string {
	srcByte, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return url2.QueryEscape(base64.StdEncoding.EncodeToString(srcByte))
}

/**
 * 使用 AK，SK 生成鉴权签名（Access Token）
 * return string 鉴权签名信息（Access Token）
 */
func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
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

func SaveBase64ToImage(outputPath string, base64String string) {
	// 将Base64字符串解码为字节数组
	imageData, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		fmt.Println("Error decoding Base64 string:", err)
		return
	}

	// 保存图像文件
	err = ioutil.WriteFile(outputPath, imageData, 0644)
	if err != nil {
		fmt.Println("Error saving image file:", err)
		return
	}
	fmt.Println("Image saved successfully as output_image.jpg")
}
