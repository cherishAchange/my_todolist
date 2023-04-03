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
	UserInfoID  uint        `json:"userInfoId"`
	FollowUsers []*UserInfo `json:"followUsers" gorm:"many2many:user_follow_tasks;"`
	Title       string      `json:"title,omitempty"`
	Describe    string      `json:"describe,omitempty"`
	Status      string      `json:"status,omitempty"`
}

type UserFollowTasks struct {
	UserInfoID uint
	TaskID     uint
}

type TaskDAO struct {
}

func (t *TaskDAO) Create(task *Task) error {
	if task == nil {
		return errors.New("创建任务失败")
	}

	return DB.Create(task).Error
}

func (t *TaskDAO) Follow(follow *UserFollowTasks) error {
	if follow == nil {
		return errors.New("加入任务失败")
	}

	return DB.Create(follow).Error
}

func (t *TaskDAO) GetAllTasks() ([]Task, error) {
	tasks := []Task{}
	err := DB.Find(&tasks).Error

	return tasks, err
}

func (t *TaskDAO) GetOwnCreatedTasks(userId uint) ([]Task, error) {
	tasks := []Task{}
	err := DB.Where("user_info_id = ?", userId).Find(&tasks).Error

	return tasks, err
}

func (t *TaskDAO) GetOwnFollowTasks(userId uint) ([]Task, error) {
	tasks := []Task{}
	err := DB.Where("id IN (?)", DB.Table("user_follow_tasks").Select("task_id").Where("user_info_id = ?", userId)).Find(&tasks).Error

	return tasks, err
}
