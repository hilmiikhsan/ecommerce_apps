package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func WriteError(c *fiber.Ctx, err error) error {
	switch {
	case err == ErrUnAuthorized:
		return write(c, http.StatusUnauthorized, "unauthorized", err.Error(), "40001", nil)
	case err == ErrUnAuthorized:
		return write(c, http.StatusUnauthorized, "unauthorized", err.Error(), "40002", nil)
	default:
		return write(c, http.StatusInternalServerError, "internal server error", "unknown error", "99999", nil)
	}
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
