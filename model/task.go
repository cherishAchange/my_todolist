package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	CreateAt    time.Time `json:"-"`
	UserInfoID  uint
	FollowUsers []*UserInfo `json:"followUsers" gorm:"many2many:user_follow_tasks;"`
	Title       string      `json:"title,omitempty"`
	Describe    string      `json:"describe,omitempty"`
	Status      string      `json:"status,omitempty"`
}
