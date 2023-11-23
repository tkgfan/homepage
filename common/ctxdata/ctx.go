// author gmfan
// date 2023/07/01

package ctxdata

import "context"

const (
	NoUID  int64 = 0
	UIDKey       = "UID"
)

// GetUID 获取用户 uid
func GetUID(ctx context.Context) int64 {
	uid := ctx.Value(UIDKey)
	if uid == nil {
		return NoUID
	}
	return uid.(int64)
}
