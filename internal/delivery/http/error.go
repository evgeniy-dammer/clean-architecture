package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ErrorResponse struct {
	Info   interface{} `json:"info,omitempty"`
	Error  string      `json:"message,omitempty"`
	Errors []string    `json:"errors,omitempty"`
	ID     uuid.UUID   `json:"id"`
}

func SetError(ctx *gin.Context, statusCode int, errs ...error) {
	response := ErrorResponse{
		ID: uuid.New(),
	}

	if len(errs) == 0 {
		return
	}

	if len(errs) > 0 {
		response.Error = errs[0].Error()

		if len(errs) > 1 {
			for _, err := range errs {
				response.Errors = append(response.Errors, err.Error())
			}
		}
	}

	ctx.JSON(statusCode, response)
}
