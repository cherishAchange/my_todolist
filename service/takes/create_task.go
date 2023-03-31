package takes

import "my_todolist/model"

type CreateResponse struct {
	StatusCode int16
	StatusMsg  string
}

func NewCreateTaskFlow(title, describe string) (*CreateResponse, error) {
	return (&CreateTaskFlow{Title: title, Describe: describe}).Do()
}

type CreateTaskFlow struct {
	UserInfoID  uint
	FollowUsers []*model.UserInfo `json:"followUsers" gorm:"many2many:user_follow_tasks;"`
	Title       string            `json:"title,omitempty"`
	Describe    string            `json:"describe,omitempty"`
	Status      string            `json:"status,omitempty"`
}

func (c *CreateTaskFlow) Do() (*CreateResponse, error) {

	// 填充 UserInfoID
	c.UserInfoID = 2

	task := model.Task{UserInfoID: c.UserInfoID, Title: c.Title, Describe: c.Describe, Status: c.Status}

	taskDao := model.NewTaskDAO()

	err := taskDao.Create(&task)

	if err != nil {
		return &CreateResponse{StatusCode: 1, StatusMsg: "创建失败"}, err
	}

	return &CreateResponse{StatusCode: 0, StatusMsg: "操作成功"}, nil
}
