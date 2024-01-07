package creation

import (
	"ImageCreation/dao/mysql"
	"ImageCreation/models"
	"ImageCreation/pkg/snowflake"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

type ShowCreation struct {
}

// SetUploadImage 上传图片
func (sc ShowCreation) SetUploadImage(c *gin.Context, userID string, image multipart.File, header *multipart.FileHeader) (models.Image, error) {
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
			return models.Image{}, err
		}
	}
	err := c.SaveUploadedFile(header, savePath1)
	if err != nil {
		return models.Image{}, err
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
	return mysql.CreateUploadImage(imageInfo)
}
