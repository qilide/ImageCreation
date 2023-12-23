package author

import (
	"ImageCreation/dao/mysql"
	"ImageCreation/models"
)

type ShowAuthor struct {
}

// AllAuthors 展示所有摄影者信息
func (sa ShowAuthor) AllAuthors() ([]models.UserInformation, error) {
	return mysql.GetAuthors()
}

// AuthorInfo 获取摄影者详细信息
func (sa ShowAuthor) AuthorInfo(id int64) (models.UserInformation, error) {
	return mysql.GetAuthorInfo(id)
}
