package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
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

func getContextFields(ctx *gin.Context) []zap.Field {
	fields := []zap.Field{
		zap.Int("status", ctx.Writer.Status()),
		zap.String("method", ctx.Request.Method),
		zap.String("path", ctx.Request.URL.Path),
		zap.String("query", ctx.Request.URL.RawQuery),
		zap.String("ip", ctx.ClientIP()),
		zap.String("user-agent", ctx.Request.UserAgent()),
	}

	if span := opentracing.SpanFromContext(ctx.Request.Context()); span != nil {
		if jaegerSpan, ok := span.Context().(jaeger.SpanContext); ok {
			fields = append(fields, zap.Stringer("traceID", jaegerSpan.TraceID()))
		}
	}

	return fields
}
