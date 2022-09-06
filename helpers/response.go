package helpers

import (
	"demoapp/schemas"
	"fmt"
)

type SError struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func BaseError(Error schemas.SErrorResponse) schemas.SErrorResponse {
	_message := ""

	if Error.Message != "" {
		_message = Error.Message
	}

	return schemas.SErrorResponse{
		Code:              Error.Code,
		Name:              Error.Name,
		Message:           _message,
		ResolutionMessage: Error.ResolutionMessage,
		Error:             fmt.Sprintf("Error - %v", Error.Error),
	}
}

func InternalServerError(error SError) schemas.SErrorResponse {
	var __message = "Something Went Wrong"

	if error.Message != "" {
		__message = error.Message
	}

	return BaseError(schemas.SErrorResponse{
		Code:              "DEMO_100",
		Name:              "INTERNAL_SERVER_ERROR",
		Message:           __message,
		ResolutionMessage: "Please try again later",
		Error:             error.Error,
	})
}

func BadRequest(error SError) schemas.SErrorResponse {
	var __message = "Invalid Request"

	if error.Message != "" {
		__message = error.Message
	}

	return BaseError(schemas.SErrorResponse{
		Code:              "DEMO_101",
		Name:              "BAD_REQUEST",
		Message:           __message,
		ResolutionMessage: "Check your request",
		Error:             error.Error,
	})
}
