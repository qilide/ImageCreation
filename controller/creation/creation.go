package creation

import (
	"ImageCreation/controller/response"
	"ImageCreation/logic/creation"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
)

// UploadImage 上传创作图片
// @Summary 上传创作图片
// @Description 用于上传创作图片
// @Tags 上传创作图片
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "上传创作图片成功"
// @failure 401 {object}  response.Information "上传创作图片失败"
// @failure 402 {object}  response.Information "请上传文件"
// @Router /creation/upload/image [POST]
func UploadImage(c *gin.Context) {
	userID := c.PostForm("id")
	image, header, _ := c.Request.FormFile("image")
	if image == nil {
		response.Json(c, 200, "请上传文件", 0)
		return
	}
	var sc creation.ShowCreation
	imageInfo, img, err := sc.SetUploadImage(c, userID, image, header)
	if err != nil {
		response.Json(c, 200, "上传创作图片失败", err)
		return
	}
	response.Json(c, 200, "上传创作图片成功", gin.H{"imageInfo": imageInfo, "img": img})
	return
}

// GeneralRecognition 通用物体和场景识别
// @Summary 通用物体和场景识别
// @Description 用于通用物体和场景识别
// @Tags 通用物体和场景识别
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "通用物体和场景识别成功"
// @failure 401 {object}  response.Information "通用物体和场景识别失败"
// @Router /creation/zoom/pro [POST]
func GeneralRecognition(c *gin.Context) {
	userID := c.PostForm("userId")
	imageID := c.PostForm("imageId")
	imagePath := c.PostForm("imagePath")
	prefix := "/static/"
	index := strings.Index(imagePath, prefix)
	// 提取路径部分
	path := "./static/" + imagePath[index+len(prefix):]
	decodedPath, err := url.QueryUnescape(path)
	if err != nil {
		fmt.Println("解码错误:", err)
		response.Json(c, 200, "通用物体和场景识别失败", err)
		return
	}
	var sc creation.ShowCreation
	imageInfo, err := sc.ImageGeneralRecognition(userID, imageID, decodedPath)
	if err != nil || imageInfo.Root == nil {
		fmt.Println(err)
		response.Json(c, 200, "通用物体和场景识别失败", err)
		return
	}
	response.Json(c, 200, "通用物体和场景识别成功", gin.H{"imageInfo": imageInfo})
	return
}

// ZoomPro 图像无损放大
// @Summary 图像无损放大
// @Description 用于图像无损放大
// @Tags 图像无损放大
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "图像无损放大成功"
// @failure 401 {object}  response.Information "图像无损放大失败"
// @Router /creation/zoom/pro [POST]
func ZoomPro(c *gin.Context) {
	userID := c.PostForm("userId")
	imageID := c.PostForm("imageId")
	imagePath := c.PostForm("imagePath")
	prefix := "/static/"
	index := strings.Index(imagePath, prefix)
	// 提取路径部分
	path := "./static/" + imagePath[index+len(prefix):]
	var sc creation.ShowCreation
	apiUrl := "https://aip.baidubce.com/rest/2.0/image-process/v1/image_quality_enhance?access_token="
	imageInfo, err := sc.ImageEnhancement(userID, imageID, path, apiUrl, "图像无损放大")
	if err != nil {
		fmt.Println(err)
		response.Json(c, 200, "图像无损放大失败", err)
		return
	}
	response.Json(c, 200, "图像无损放大成功", gin.H{"imageInfo": imageInfo})
	return
}

// ClarityEnhancement 图像清晰度增强
// @Summary 图像清晰度增强
// @Description 用于图像清晰度增强
// @Tags 图像清晰度增强
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "图像清晰度增强成功"
// @failure 401 {object}  response.Information "图像清晰度增强失败"
// @Router /creation/clarity/enhancement [POST]
func ClarityEnhancement(c *gin.Context) {
	userID := c.PostForm("userId")
	imageID := c.PostForm("imageId")
	imagePath := c.PostForm("imagePath")
	prefix := "/static/"
	index := strings.Index(imagePath, prefix)
	// 提取路径部分
	path := "./static/" + imagePath[index+len(prefix):]
	var sc creation.ShowCreation
	apiUrl := "https://aip.baidubce.com/rest/2.0/image-process/v1/image_definition_enhance?access_token="
	imageInfo, err := sc.ImageEnhancement(userID, imageID, path, apiUrl, "图像清晰度增强")
	if err != nil {
		fmt.Println(err)
		response.Json(c, 200, "图像清晰度增强失败", err)
		return
	}
	response.Json(c, 200, "图像清晰度增强成功", gin.H{"imageInfo": imageInfo})
	return
}

// ColorEnhancement 图像色彩增强
// @Summary 图像色彩增强
// @Description 用于图像色彩增强
// @Tags 图像色彩增强
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "图像色彩增强成功"
// @failure 401 {object}  response.Information "图像色彩增强失败"
// @Router /creation/color/enhancement [POST]
func ColorEnhancement(c *gin.Context) {
	userID := c.PostForm("userId")
	imageID := c.PostForm("imageId")
	imagePath := c.PostForm("imagePath")
	prefix := "/static/"
	index := strings.Index(imagePath, prefix)
	// 提取路径部分
	path := "./static/" + imagePath[index+len(prefix):]
	var sc creation.ShowCreation
	apiUrl := "https://aip.baidubce.com/rest/2.0/image-process/v1/color_enhance?access_token="
	imageInfo, err := sc.ImageEnhancement(userID, imageID, path, apiUrl, "图像色彩增强")
	if err != nil {
		fmt.Println(err)
		response.Json(c, 200, "图像色彩增强失败", err)
		return
	}
	response.Json(c, 200, "图像色彩增强成功", gin.H{"imageInfo": imageInfo})
	return
}

// ContrastEnhancement 图像对比度增强
// @Summary 图像对比度增强
// @Description 用于图像对比度增强
// @Tags 图像对比度增强
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "图像对比度增强成功"
// @failure 401 {object}  response.Information "图像对比度增强失败"
// @Router /creation/contrast/enhancement [POST]
func ContrastEnhancement(c *gin.Context) {
	userID := c.PostForm("userId")
	imageID := c.PostForm("imageId")
	imagePath := c.PostForm("imagePath")
	prefix := "/static/"
	index := strings.Index(imagePath, prefix)
	// 提取路径部分
	path := "./static/" + imagePath[index+len(prefix):]
	var sc creation.ShowCreation
	apiUrl := "https://aip.baidubce.com/rest/2.0/image-process/v1/contrast_enhance?access_token="
	imageInfo, err := sc.ImageEnhancement(userID, imageID, path, apiUrl, "图像对比度增强")
	if err != nil {
		fmt.Println(err)
		response.Json(c, 200, "图像对比度增强失败", err)
		return
	}
	response.Json(c, 200, "图像对比度增强成功", gin.H{"imageInfo": imageInfo})
	return
}

// Defogging 图像去雾
// @Summary 图像去雾
// @Description 用于图像去雾
// @Tags 图像去雾
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "图像去雾成功"
// @failure 401 {object}  response.Information "图像去雾失败"
// @Router /creation/defogging [POST]
func Defogging(c *gin.Context) {
	userID := c.PostForm("userId")
	imageID := c.PostForm("imageId")
	imagePath := c.PostForm("imagePath")
	prefix := "/static/"
	index := strings.Index(imagePath, prefix)
	// 提取路径部分
	path := "./static/" + imagePath[index+len(prefix):]
	var sc creation.ShowCreation
	apiUrl := "https://aip.baidubce.com/rest/2.0/image-process/v1/dehaze?access_token="
	imageInfo, err := sc.ImageEnhancement(userID, imageID, path, apiUrl, "图像去雾")
	if err != nil {
		fmt.Println(err)
		response.Json(c, 200, "图像去雾失败", err)
		return
	}
	response.Json(c, 200, "图像去雾成功", gin.H{"imageInfo": imageInfo})
	return
}
