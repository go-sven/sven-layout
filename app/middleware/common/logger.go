package common

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github/go-sven/sven-layout/logger"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set(logger.TraceKey,uuid.New().String())
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		logger.InfoWithCtx(c,"status:",c.Writer.Status(), "path:",path, "method:",c.Request.Method,"query:",query,"ip:",c.ClientIP(),"errors:",c.Errors.ByType(gin.ErrorTypePrivate).String(),"cost:",cost)
	}
}