// author gmfan
// date 2023/07/01

package router

import (
	"acsupport/common/errs"
	"acsupport/common/result"
	"context"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	pingApi(r)
}

// 处理 JSON 请求
func jsonHandle(c *gin.Context, req any, handle func(ctx context.Context) (resp any, err error)) {
	// 解析请求参数
	if err := c.ShouldBindJSON(req); err != nil {
		result.HttpResult(c, nil, errs.NewCodeErr(errs.ParamErr))
		return
	}

	// 调用处理钩子函数
	resp, err := handle(c.Request.Context())
	// 返回处理结果
	result.HttpResult(c, resp, err)
}
