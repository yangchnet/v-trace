package middleware

import (
	"context"
	"fmt"
	"net/http/httputil"
	"time"

	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Log(ctx context.Context, l *logger.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path

		logStr := fmt.Sprintf("%3d | %v | %s | %s %s %s",
			statusCode,
			latency,
			clientIP, method,
			path,
			c.Errors.String(),
		)

		switch {
		case statusCode >= 400 && statusCode <= 499:
			l.Warn(ctx, logStr)
		case statusCode >= 500:
			l.Error(ctx, logStr)
		default:
			l.Info(ctx, logStr)
		}
	}
}

func Recover(ctx context.Context, l *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				l.Errorf("Panic recovered: %+v\n%s", r, string(httpRequest))
				c.JSON(500, gin.H{
					"title": "error",
					"err":   r,
				})
			}
		}()
		c.Next()
	}
}

func CORSMiddleware(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
