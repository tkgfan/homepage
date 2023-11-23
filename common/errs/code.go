// author gmfan
// date 2023/7/8
package errs

// 统一状态码不可修改
const (
	UnknownErr uint32 = 100
	ParamErr   uint32 = 200
	NotLogin   uint32 = 300
	// ServerErr 服务器错误
	ServerErr uint32 = 400
	// TokenExpire token 过期
	TokenExpire uint32 = 500
	// NoPermission 没有权限
	NoPermission uint32 = 600
)

const (
	UnknownErrMsg   = "未知错误"
	ParamErrMsg     = "参数错误"
	NotLoginMsg     = "未登入"
	TokenExpireMsg  = "Token 已过期"
	NoPermissionMsg = "没有权限"
)

var codeMsg = map[uint32]string{
	UnknownErr:   UnknownErrMsg,
	ParamErr:     ParamErrMsg,
	NotLogin:     NotLoginMsg,
	TokenExpire:  TokenExpireMsg,
	NoPermission: NoPermissionMsg,
}
