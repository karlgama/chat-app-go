package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	TraceIDHeader = "request-trace-id"
	// TraceIDKey é a chave para armazenar o trace ID no contexto
	TraceIDKey = "trace_id"
)

func TraceIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader(TraceIDHeader)

		if traceID == "" {
			traceID = uuid.New().String()
		}

		c.Set(TraceIDKey, traceID)

		c.Header(TraceIDHeader, traceID)

		c.Next()
	}
}

func GetTraceID(c *gin.Context) string {
	if traceID, exists := c.Get(TraceIDKey); exists {
		if id, ok := traceID.(string); ok {
			return id
		}
	}
	return "unknown"
}
