package models

import "time"

type Score struct {
	ID         int       `json:"id" gorm:"column:id"`
	UserId     int       `json:"user_id" gorm:"column:user_id"`
	ImageId    int       `json:"image_id" gorm:"column:image_id"`
	Score      int       `json:"score" gorm:"column:score"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
	IsActive   int       `json:"is_active" gorm:"column:is_active"`
}
