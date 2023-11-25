// author gmfan
// date 2023/07/01

package router

import (
	"acsupport/common/errs"
	"acsupport/common/result"
	"acsupport/v1/models"
	"acsupport/v1/service/web"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tkgfan/got/core/env"
)

var updatePath = "/update"

func init() {
	env.LoadStr(&updatePath, "UPDATE_PATH", false)
}

func InitRouter(r *gin.Engine) {
	// 获取页面
	r.GET("/:path", get)
	r.GET("", get)

	// 更新配置文件
	r.POST(updatePath, func(c *gin.Context) {
		req := new(models.Config)
		jsonHandle(c, req, func(ctx context.Context) (resp any, err error) {
			return web.UpdateConfig(req)
		})
	})
}

func get(c *gin.Context) {
	path := c.Param("path")
	// 获取文件
	data, err := web.GetFile(path)
	if err != nil {
		result.HttpResult(c, data, err)
		return
	}
	c.Writer.Write(data)
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
