package models

import "time"

type Browse struct {
	ID       int       `json:"id" gorm:"column:id"`
	UserId   int       `json:"user_id" gorm:"column:user_id"`
	ImageId  int       `json:"image_id" gorm:"column:image_id"`
	ViewTime time.Time `json:"view_time" gorm:"column:view_time"`
	IsActive int       `json:"is_active" gorm:"column:is_active"`
}
