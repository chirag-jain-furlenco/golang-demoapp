package api_users

type SCreateUserResponse struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Description string `json:"description" binding:"required"`
}
