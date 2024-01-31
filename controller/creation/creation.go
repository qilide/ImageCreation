package creation

import (
	"ImageCreation/controller/response"
	"ImageCreation/logic/creation"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// PageCreation 显示创作页面
// @Summary 显示创作页面
// @Description 用于显示创作页面
// @Tags 显示创作页面
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "显示创作页面成功"
// @failure 401 {object}  response.Information "显示创作页面失败"
// @Router /creation [GET]
func PageCreation(c *gin.Context) {
	page := c.Query("page")
	if page == "" || len(page) == 0 {
		page = "1"
	}
	var sc creation.ShowCreation
	if imageInfos, err := sc.GetCreationImage(); err != nil {
		c.HTML(http.StatusOK, "errors.html", err)
	} else {
		var pages []int
		for i := 1; i < len(imageInfos)/8+1; i++ {
			pages = append(pages, i)
		}
		page1, _ := strconv.Atoi(page)
		creationImages := imageInfos[(page1-1)*8 : page1*8]
		c.HTML(http.StatusOK, "creation.html", gin.H{"ImageInfo": creationImages, "Pages": pages})
	}
	return
}

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

// Revoke 回撤图像
// @Summary 回撤图像
// @Description 用于回撤图像
// @Tags 回撤图像
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "回撤图像成功"
// @failure 401 {object}  response.Information "回撤图像失败"
// @Router /creation/revoke [POST]
func Revoke(c *gin.Context) {
	imageId := c.PostForm("origin_image_id")
	fmt.Println(imageId)
	var sc creation.ShowCreation
	imageInfo, err := sc.RevokeImage(imageId)
	if err != nil {
		fmt.Println(err)
		response.Json(c, 200, "回撤图像失败", err)
		return
	}
	response.Json(c, 200, "回撤图像成功", gin.H{"imageInfo": imageInfo})
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

// AnimatedPortraits 人像动漫化
// @Summary 人像动漫化
// @Description 用于人像动漫化
// @Tags 人像动漫化
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "人像动漫化成功"
// @failure 401 {object}  response.Information "人像动漫化失败"
// @Router /creation/animated/portraits [POST]
func AnimatedPortraits(c *gin.Context) {
	userID := c.PostForm("userId")
	imageID := c.PostForm("imageId")
	imagePath := c.PostForm("imagePath")
	type1 := c.PostForm("type")
	prefix := "/static/"
	index := strings.Index(imagePath, prefix)
	// 提取路径部分
	path := "./static/" + imagePath[index+len(prefix):]
	var sc creation.ShowCreation
	apiUrl := "https://aip.baidubce.com/rest/2.0/image-process/v1/selfie_anime?access_token=" + creation.GetAccessToken()
	imageBase64, err := creation.GetFileContentAsBase64(path)
	if err != nil {
		fmt.Println(err)
		response.Json(c, 200, "人像动漫化失败", err)
		return
	}
	var payload *strings.Reader
	if type1 == "anime" {
		payload = strings.NewReader(fmt.Sprintf("type=%s&image=%s", type1, imageBase64))
	} else {
		payload = strings.NewReader(fmt.Sprintf("type=%s&mask_id=1&image=%s", type1, imageBase64))
	}
	imageInfo, err := sc.ImageEffects(userID, imageID, apiUrl, "人像动漫化", payload)
	if err != nil {
		fmt.Println(err)
		response.Json(c, 200, "人像动漫化失败", err)
		return
	}
	response.Json(c, 200, "人像动漫化成功", gin.H{"imageInfo": imageInfo})
	return
}

// StyleConversion 图像风格转换
// @Summary 图像风格转换
// @Description 用于图像风格转换
// @Tags 图像风格转换
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "图像风格转换成功"
// @failure 401 {object}  response.Information "图像风格转换失败"
// @Router /creation/style/conversion [POST]
func StyleConversion(c *gin.Context) {
	userID := c.PostForm("userId")
	imageID := c.PostForm("imageId")
	imagePath := c.PostForm("imagePath")
	style := c.PostForm("style")
	prefix := "/static/"
	index := strings.Index(imagePath, prefix)
	// 提取路径部分
	path := "./static/" + imagePath[index+len(prefix):]
	var sc creation.ShowCreation
	apiUrl := "https://aip.baidubce.com/rest/2.0/image-process/v1/style_trans?access_token=" + creation.GetAccessToken()
	imageBase64, err := creation.GetFileContentAsBase64(path)
	if err != nil {
		fmt.Println(err)
		response.Json(c, 200, "图像风格转换失败", err)
		return
	}
	payload := strings.NewReader(fmt.Sprintf("option=%s&image=%s", style, imageBase64))
	imageInfo, err := sc.ImageEffects(userID, imageID, apiUrl, "图像风格转换", payload)
	if err != nil {
		fmt.Println(err)
		response.Json(c, 200, "图像风格转换失败", err)
		return
	}
	response.Json(c, 200, "图像风格转换成功", gin.H{"imageInfo": imageInfo})
	return
}

// CustomStyle 自定义图像风格
// @Summary 自定义图像风格
// @Description 用于自定义图像风格
// @Tags 自定义图像风格
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "自定义图像风格成功"
// @failure 401 {object}  response.Information "自定义图像风格失败"
// @Router /creation/custom/style [POST]
func CustomStyle(c *gin.Context) {
	userID := c.PostForm("userId")
	imageID := c.PostForm("imageId")
	imagePath := c.PostForm("imagePath")
	image, header, _ := c.Request.FormFile("image")
	defer image.Close()
	// 将文件保存到指定的路径，这里保存在当前目录下的 uploads 文件夹中
	savePath := "./static/creation/custom_style/"
	savePath1 := savePath + header.Filename
	// 如果目录不存在，则创建目录
	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		err := os.Mkdir(savePath, 0755)
		if err != nil {
			fmt.Println(err)
			response.Json(c, 200, "自定义图像风格失败", err)
			return
		}
	}
	err := c.SaveUploadedFile(header, savePath1)
	imageBase641, err := creation.GetFileContentAsBase64(savePath1)
	if err != nil {
		fmt.Println(err)
		response.Json(c, 200, "自定义图像风格失败", err)
		return
	}
	prefix := "/static/"
	index := strings.Index(imagePath, prefix)
	// 提取路径部分
	path := "./static/" + imagePath[index+len(prefix):]
	var sc creation.ShowCreation
	apiUrl := "https://aip.baidubce.com/rest/2.0/image-process/v1/customize_stylization?access_token=" + creation.GetAccessToken()
	imageBase64, err := creation.GetFileContentAsBase64(path)
	if err != nil {
		fmt.Println(err)
		response.Json(c, 200, "自定义图像风格失败", err)
		return
	}
	payload := strings.NewReader(fmt.Sprintf("style=%s&image=%s", imageBase641, imageBase64))
	imageInfo, err := sc.ImageEffects(userID, imageID, apiUrl, "自定义图像风格", payload)
	if err != nil {
		fmt.Println(err)
		response.Json(c, 200, "自定义图像风格失败", err)
		return
	}
	response.Json(c, 200, "自定义图像风格成功", gin.H{"imageInfo": imageInfo})
	return
}
