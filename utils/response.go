package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)
type APIResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	ErrorCode  string      `json:"error_code,omitempty"`
	Timestamp  string      `json:"timestamp"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Total      int         `json:"total"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	Pages      int         `json:"total_pages"`
}

func ResponseSuccess(cRes *fiber.Ctx, statusCode int, message string, data interface{}, pagination *Pagination) error {
	return cRes.Status(statusCode).JSON(APIResponse{
		Success:    true,
		Message:    message,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Data:       data,
		Pagination: pagination,
	})
}


func ResponseError(cRes *fiber.Ctx, statusCode int, message string, data interface{}, code ...string) error {
	errorCode := "ERR_UNKNOWN"
	if len(code) > 0 {
		errorCode = code[0]
	}
	return cRes.Status(statusCode).JSON(APIResponse{
		Success:   false,
		Message:   message,
		ErrorCode: errorCode,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Data:      data,
	})
}
