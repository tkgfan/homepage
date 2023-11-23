// author gmfan
// date 2023/7/8
package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

// Timeout 设置请求上下文超时时间
func Timeout(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, _ := context.WithTimeout(c.Request.Context(), timeout)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
