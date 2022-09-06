package schemas

type SErrorResponse struct {
	Name              string      `json:"name"`
	Code              string      `json:"code"`
	Message           string      `json:"message"`
	ResolutionMessage string      `json:"resolutionMessage"`
	Error             interface{} `json:"error"`
}

type SInternalServerError struct {
	Name              string `json:"name" default:"Internal Server Error"`
	Code              string `json:"code" default:"INTERNAL_SERVER_ERROR"`
	Message           string `json:"message" default:"Something Went Wrong"`
	ResolutionMessage string `json:"resolutionMessage"`
}
