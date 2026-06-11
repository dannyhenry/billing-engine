package api_response

import (
	"fmt"
	"strings"
)

type DataResponse struct {
	Status interface{} `json:"status"`
	Data   interface{} `json:"data"`
}

type ItemsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Items interface{} `json:"items,omitempty"`
	} `json:"data"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
	Page    interface{} `json:"page,omitempty"`
	Data    interface{} `json:"data"`
}

func BuildSuccessResponse(message string, status int, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, status int, err error, data interface{}) Response {
	errorMessage := err.Error()

	splitError := strings.Split(errorMessage, "\n")
	res := Response{
		Status:  status,
		Message: message,
		Errors:  splitError,
		Data:    data,
	}
	return res
}

func BuildSuccessPageResponse(status int, message string, data interface{}, args ...interface{}) Response {
	res := Response{}
	res.Status = status
	res.Message = message
	if data != "" && data != nil {
		res.Data = data
	}
	if len(args) > 0 {
		res.Page = args[0]
	}
	return res
}

type Page struct {
	Offset    int   `json:"offset"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
}

func ApplicationError(message string, status int, err error, data interface{}) Response {
	if status == 0 {
		status = 500
	}

	if message == "" {
		message = "Internal Server Error."
	}

	return BuildErrorResponse(message, status, err, data)
}

func ValidationError(message string, status int, err error, data interface{}) Response {
	if status == 0 {
		status = 102
	}

	if message == "" {
		message = "Validation Error."
	}
	fmt.Println("status", status)

	return BuildErrorResponse(message, status, err, data)
}

func DatabaseError(message string, status int, err error, data interface{}) Response {
	if status == 0 {
		status = 200
	}

	if message == "" {
		message = "Database Error."
	}

	return BuildErrorResponse(message, status, err, data)
}

func AuthorizationError(message string, status int, err error, data interface{}) Response {
	if status == 0 {
		status = 107
	}

	if message == "" {
		message = "Unauthorized."
	}

	return BuildErrorResponse(message, status, err, data)
}
