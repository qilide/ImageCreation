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
	imageInfo, err := sc.SetUploadImage(c, userID, image, header)
	fmt.Println(imageInfo)
	fmt.Println(err)
	if err != nil {
		response.Json(c, 200, "上传创作图片失败", err)
		return
	}
	response.Json(c, 200, "上传创作图片成功", imageInfo)
	return
}
