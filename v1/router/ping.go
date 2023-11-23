// author gmfan
// date 2023/07/01

package router

import (
	"acsupport/common/result"
	"acsupport/v1/service"
	"github.com/gin-gonic/gin"
)

// pingApi 注册测试路由
func pingApi(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		resp, err := service.Pong(c.Request.Context())
		result.HttpResult(c, resp, err)
	})
}
