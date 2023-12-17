package models

import "time"

type Image struct {
	ID               int       `json:"id" gorm:"column:id"`
	UserID           int       `json:"user_id" gorm:"column:user_id"`
	ImageName        string    `json:"image_name" gorm:"column:image_name"`
	Path             string    `json:"path" gorm:"column:path"`
	Theme            string    `json:"theme" gorm:"column:theme"`
	Label            string    `json:"label" gorm:"column:label"`
	Description      string    `json:"description" gorm:"column:description"`
	CreateTime       time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime       time.Time `json:"update_time" gorm:"column:update_time"`
	IsActive         int       `json:"is_active" gorm:"column:is_active"`
	IsCreate         int       `json:"is_create" gorm:"column:is_create"`
	Score            float64   `json:"score" gorm:"column:score"`
	CollectCount     int       `json:"collect_count" gorm:"column:collect_count"`
	LikeCount        int       `json:"like_count" gorm:"column:like_count"`
	BrowseCount      int       `json:"browse_count" gorm:"column:browse_count"`
	Heat             int       `json:"heat" gorm:"column:heat"`
	Location         string    `json:"location" gorm:"column:location"`
	AlgorithmTheme   string    `json:"algorithm_theme" gorm:"column:algorithm_theme"`
	AlgorithmContent string    `json:"algorithm_content" gorm:"column:algorithm_content"`
	AlgorithmScore   string    `json:"algorithm_score" gorm:"column:algorithm_score"`
}
