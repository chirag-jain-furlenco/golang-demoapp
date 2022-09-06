package api_models

import "github.com/gin-gonic/gin"

type SRoute struct {
	Route      string
	Method     string
	Controller func(req *gin.Context)
}

type SRouteGroupDef struct {
	Path   string
	Routes []SRoute
}

type SBaseSuccessResponse[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}
