package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// 初始化一个 Zap Logger
var logger, _ = zap.NewProduction()

// ZapLoggerMiddleware 是一个 Gin 中间件，它使用 Zap 记录每个请求的日志
func ZapLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求的开始时间
		start := time.Now()

		// 继续处理请求
		c.Next()

		// 获取请求的处理状态
		status := c.Writer.Status()

		// 记录日志
		logger.Info("Request received",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", status),
			zap.Duration("duration", time.Since(start)),
		)
	}
}
