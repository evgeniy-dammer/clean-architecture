package http

import (
	"net/http"

	log "github.com/evgeniy-dammer/clean-architecture/pkg/type/logger"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

func Tracer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		span := opentracing.SpanFromContext(ctx.Request.Context())

		if span == nil {
			span = StartSpanWithHeader(&ctx.Request.Header, "rest-request-"+ctx.Request.Method, ctx.Request.Method, ctx.Request.URL.Path) //nolint:lll
		}

		defer span.Finish()

		ctx.Request = ctx.Request.WithContext(opentracing.ContextWithSpan(ctx.Request.Context(), span))

		if traceID, ok := span.Context().(jaeger.SpanContext); ok {
			ctx.Header("uber-trace-id", traceID.TraceID().String())
		}

		ctx.Next()

		ext.HTTPStatusCode.Set(span, uint16(ctx.Writer.Status()))

		if len(ctx.Errors) == 0 {
			log.Info("", getContextFields(ctx)...)
		}
	}
}

func StartSpanWithHeader(header *http.Header, operationName, method, path string) opentracing.Span {
	var wireContext opentracing.SpanContext

	if header != nil {
		wireContext, _ = opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(*header))
	}

	return StartSpanWithParent(wireContext, operationName, method, path)
}

// StartSpanWithParent will start a new span with a parent span.
// example:
//
//	span:= StartSpanWithParent(c.Get("tracing-context"),
func StartSpanWithParent(parent opentracing.SpanContext, operationName, method, path string) opentracing.Span {
	options := []opentracing.StartSpanOption{
		opentracing.Tag{Key: ext.SpanKindRPCServer.Key, Value: ext.SpanKindRPCServer.Value},
		opentracing.Tag{Key: string(ext.HTTPMethod), Value: method},
		opentracing.Tag{Key: string(ext.HTTPUrl), Value: path},
	}
	if parent != nil {
		options = append(options, opentracing.ChildOf(parent))
	}

	return opentracing.StartSpan(operationName, options...)
}
