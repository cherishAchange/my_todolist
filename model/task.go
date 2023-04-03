package model

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

var (
	taskDao  *TaskDAO
	taskOnce sync.Once
)

func NewTaskDAO() *TaskDAO {
	taskOnce.Do(func() {
		taskDao = &TaskDAO{}
	})

	return taskDao
}

type Task struct {
	gorm.Model
	UserInfoID  uint
	FollowUsers []*UserInfo `json:"followUsers" gorm:"many2many:user_follow_tasks;"`
	Title       string      `json:"title,omitempty"`
	Describe    string      `json:"describe,omitempty"`
	Status      string      `json:"status,omitempty"`
}

type TaskDAO struct {
}

func (t *TaskDAO) Create(task *Task) error {
	if task == nil {
		return errors.New("创建任务失败")
	}

	return DB.Create(task).Error
}

func (t *TaskDAO) Follow() error {
	return DB.Updates("followUsers").Error
}
