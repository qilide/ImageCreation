package models

import "time"

type UserInformation struct {
	ID               int       `json:"id" gorm:"column:id"`
	UserID           int       `json:"user_id" gorm:"column:user_id"`
	Nickname         string    `json:"nickname" gorm:"column:nickname"`
	Age              int       `json:"age" gorm:"column:age"`
	Sex              string    `json:"sex" gorm:"column:sex"`
	BrithDate        time.Time `json:"brith_date" gorm:"column:brith_date"`
	Avatar           string    `json:"avatar" gorm:"column:avatar"`
	Biography        string    `json:"biography" gorm:"column:biography"`
	Address          string    `json:"address" gorm:"column:address"`
	IsAuthor         int       `json:"is_author" gorm:"column:is_author"`
	IsActive         int       `json:"is_active" gorm:"column:is_active"`
	Description      string    `json:"description" gorm:"column:description"`
	Style            string    `json:"style" gorm:"column:style"`
	Posts            string    `json:"posts" gorm:"column:posts"`
	TotalLikes       int       `json:"total_likes" gorm:"column:total_likes"`
	TotalCollections int       `json:"total_collections" gorm:"column:total_collections"`
	TotalBrowse      int       `json:"total_browse" gorm:"column:total_browse"`
	TotalPhotos      int       `json:"total_photos" gorm:"column:total_photos"`
}
