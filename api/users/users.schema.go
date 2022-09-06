package api_users

type SCreateUser struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type SGetUserParams struct {
	Id uint `uri:"userId" binding:"required"`
}

type SCreateTaskBody struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type SCreateTaskParams struct {
	UserId uint `uri:"userId" binding:"required"`
}

type SGetUserTasks struct {
	UserId uint `uri:"userId" binding:"required"`
}
