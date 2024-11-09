package rest

import (
	"fernandezlautt/feedback-service/lib"
	"fernandezlautt/feedback-service/lib/log"

	"github.com/gin-gonic/gin"
)

func newGinLogger(c *gin.Context, ctx ...interface{}) log.LogRusEntry {
	return log.Get(ctx...).
		WithField(log.LOG_FIELD_CORRELATION_ID, lib.GetCorrelationId(c)).
		WithField(log.LOG_FIELD_CONTROLLER, "Rest").
		WithField(log.LOG_FIELD_HTTP_METHOD, c.Request.Method).
		WithField(log.LOG_FIELD_HTTP_PATH, c.Request.URL.Path)
}

func GinLoggerMiddleware(ctx ...interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := newGinLogger(c, ctx...)
		c.Set("logger", logger)

		c.Next()

		if c.Request.Method != "OPTIONS" {
			ctx := GinCtx(c)
			log.Get(ctx...).WithField(log.LOG_FIELD_HTTP_STATUS, c.Writer.Status()).Info("Completed")
		}
	}
}

func ginLogger(c *gin.Context) log.LogRusEntry {
	logger, exist := c.Get("logger")
	if !exist {
		return newGinLogger(c)
	}
	return logger.(log.LogRusEntry)
}
