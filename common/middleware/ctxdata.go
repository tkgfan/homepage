// author gmfan
// date 2023/8/12
package middleware

import (
	"acsupport/common/ctxdata"
	"acsupport/common/errs"
	"acsupport/common/result"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

// SetCtxData 解析 Header 数据并设置到 Request 的上下文中
func SetCtxData() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		// 获取 UID
		uidStr := c.Request.Header.Get(ctxdata.UIDKey)
		if uidStr != "" {
			uid, err := strconv.ParseInt(uidStr, 10, 64)
			if err != nil {
				result.HttpResult(c, nil, errs.NewCodeErr(errs.NotLogin))
				c.Abort()
				return
			}
			ctx = context.WithValue(ctx, ctxdata.UIDKey, uid)
		}

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
