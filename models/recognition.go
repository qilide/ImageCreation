package models

import "time"

type Recognition struct {
	ID          int         `json:"id" gorm:"column:id"`
	UserId      int         `json:"user_id" gorm:"column:user_id"`
	ImageId     int         `json:"image_id" gorm:"column:image_id"`
	Score       interface{} `json:"score" gorm:"column:score"`
	Root        interface{} `json:"root" gorm:"column:root"`
	BaikeUrl    interface{} `json:"baike_url" gorm:"column:baike_url"`
	ImageUrl    interface{} `json:"image_url" gorm:"column:image_url"`
	Description interface{} `json:"description" gorm:"column:description"`
	Keyword     interface{} `json:"keyword" gorm:"column:keyword"`
	CreateTime  time.Time   `json:"create_time" gorm:"column:create_time"`
	UpdateTime  time.Time   `json:"update_time" gorm:"column:update_time"`
	IsActive    int         `json:"is_active" gorm:"column:is_active"`
}
