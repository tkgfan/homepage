// author gmfan
// date 2023/7/8
package result

import (
	"acsupport/common/errs"
	"github.com/gin-gonic/gin"
	"github.com/tkgfan/got/core/errors"
	"github.com/tkgfan/got/core/logx"
	"github.com/tkgfan/got/core/model"
	"github.com/tkgfan/got/core/structs"
	"net/http"
)

const (
	OK uint32 = 0
)

const (
	OkMsg = "OK"
)

// HttpResult 统一处理返回结果
func HttpResult(c *gin.Context, data any, err error) {
	// 合并 Trace Log
	logx.TraceLogMergeLog(c.Request.Context(), err)

	// 返回链路日志
	tlVal := logx.GetTraceLogStr(c.Request.Context())
	c.Writer.Header().Set(logx.TraceLogKey, tlVal)

	if structs.IsNil(err) {
		c.JSON(http.StatusOK, model.NewSuccessResp(OK, OkMsg, data))
		return
	}

	logx.Error(errors.Json(err))

	// 处理错误
	cause := errors.Cause(err)
	if e, ok := errs.IsCodeErr(cause); ok {
		c.JSON(http.StatusOK, model.NewFailResp(e.Code, e.Msg))
	} else {
		c.JSON(http.StatusOK, model.NewFailResp(errs.UnknownErr, errs.UnknownErrMsg))
	}
}
