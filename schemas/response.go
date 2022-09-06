package schemas

type SSuccessResponse[T interface{}] struct {
	Success bool   `json:"success" binding:"required"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
