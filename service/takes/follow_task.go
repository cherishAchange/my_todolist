package takes

import (
	"my_todolist/model"
)

type FollowFlow struct {
	UserInfoID uint
	TaskID     uint
}

func Follow(userId, taskId uint) (*TaskResponse, error) {
	return (&FollowFlow{UserInfoID: userId, TaskID: taskId}).Do()
}

func (f *FollowFlow) Do() (*TaskResponse, error) {
	followM := model.UserFollowTasks{UserInfoID: f.UserInfoID, TaskID: f.TaskID}
	followF := model.NewTaskDAO()
	err := followF.Follow(&followM)

	if err != nil {
		return &TaskResponse{StatusCode: 1, StatusMsg: "加入任务失败"}, err
	}

	return &TaskResponse{StatusCode: 0, StatusMsg: "操作成功"}, nil
}
