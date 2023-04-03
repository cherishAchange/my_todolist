package takes

import (
	"my_todolist/model"
)

func QueryAllTasks() ([]model.Task, error) {
	taskDao := model.NewTaskDAO()
	tasks, err := taskDao.GetAllTasks()

	return tasks, err
}

func QueryCreated(userId uint) ([]model.Task, error) {
	taskDao := model.NewTaskDAO()
	tasks, err := taskDao.GetOwnCreatedTasks(userId)

	return tasks, err
}

func QueryFollowedTasks(userId uint) ([]model.Task, error) {
	taskDao := model.NewTaskDAO()
	tasks, err := taskDao.GetOwnFollowTasks(userId)

	return tasks, err
}

func QueryFollower(taskId uint) ([]model.UserInfo, error) {
	userDao := model.NewUserInfoDAO()
	users, err := userDao.GetUsersByTaskId(taskId)

	return users, err
}
