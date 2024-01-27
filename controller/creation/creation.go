package creation

import (
	"ImageCreation/controller/response"
	"ImageCreation/logic/creation"
	"fmt"
	"github.com/gin-gonic/gin"
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
	fmt.Println(imageInfo)
	fmt.Println(err)
	// 打印图片信息
	fmt.Println("宽度:", img.Width, "像素")
	fmt.Println("高度:", img.Height, "像素")
	fmt.Println("颜色模式:", img.ColorModel)
	if err != nil {
		response.Json(c, 200, "上传创作图片失败", err)
		return
	}
	response.Json(c, 200, "上传创作图片成功", gin.H{"imageInfo": imageInfo, "img": img})
	return
}
