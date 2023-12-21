package file

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrInvalidFileType = errors.New("invalid file type")
	ErrInvalidFileSize = errors.New("invalid file size")
)

func WriteError(c *fiber.Ctx, err error) error {
	switch {
	case err == ErrInvalidFileType:
		return write(c, http.StatusBadRequest, "bad request", err.Error(), "40001", nil)
	case err == ErrInvalidFileSize:
		return write(c, http.StatusBadRequest, "bad request", err.Error(), "40002", nil)
	default:
		return write(c, http.StatusInternalServerError, "internal server error", "unknown error", "99999", nil)
	}
}

func WriteSuccess(c *fiber.Ctx, message string, payload interface{}, statusCode int) error {
	resp := response{
		Success: true,
		Message: message,
		Payload: payload,
	}
	c = c.Status(statusCode)
	return c.JSON(resp)
}

type response struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Payload   interface{} `json:"payload,omitempty"`
	Error     *string     `json:"error,omitempty"`
	ErrorCode *string     `json:"error_code,omitempty"`
}

func write(c *fiber.Ctx, statusCode int, message, errorMessage, errorCode string, payload interface{}) error {
	c = c.Status(statusCode)
	isSuccess := statusCode >= 200 && statusCode < 300

	if isSuccess {
		return c.JSON(response{
			Success: true,
			Message: message,
			Payload: payload,
		})
	}

	return c.JSON(response{
		Success:   false,
		Message:   message,
		Error:     &errorMessage,
		ErrorCode: &errorCode,
	})
}
