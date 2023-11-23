// author gmfan
// date 2023/7/24
package middleware

import (
	"acsupport/common/errs"
	"acsupport/common/result"
	"github.com/gin-gonic/gin"
	"github.com/tkgfan/got/core/logx"
)

// TraceLog 处理链路日志
func TraceLog() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 设置链路日志
		ctx, err := logx.SetTraceLog(c.Request.Context(), c.Request.Header.Get(logx.TraceLogKey), c.Request.URL.Path)
		if err != nil {
			logx.Error(err)
			result.HttpResult(c, nil, errs.NewCodeErrMgs(errs.ParamErr, "未设置链路日志"))
			c.Abort()
			return
		}
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
