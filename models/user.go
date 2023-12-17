package models

import "time"

type User struct {
	ID          int       `json:"id" gorm:"column:id"`
	Username    string    `json:"username" gorm:"column:username;unique;not null"`
	Password    string    `json:"password" gorm:"column:password"`
	IsSuperuser int       `json:"is_superuser" gorm:"column:is_superuser"`
	IsActive    int       `json:"is_active" gorm:"column:is_active"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`
	LastLogin   time.Time `json:"last_login" gorm:"column:last_login"`
	Email       string    `json:"email" gorm:"column:email"`
	Phone       string    `json:"phone" gorm:"column:phone"`
}
