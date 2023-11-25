// author gmfan
// date 2023/8/12
package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/tkgfan/got/core/env"
	"github.com/tkgfan/got/core/logx"
	"net"
	"net/http"
)

// 有权访问列表
var accessList []*net.IPNet

func init() {
	var listStr string
	env.LoadStr(&listStr, "ACCESS_LIST", false)
	if listStr == "" {
		return
	}
	var list []string
	err := json.Unmarshal([]byte(listStr), &list)
	if err != nil {
		logx.Error(`ACCESS_LIST 格式不合法，正确格式为：["192.168.1.0/24"]`)
		panic(err)
	}
	for _, v := range list {
		_, ipNet, err := net.ParseCIDR(v)
		if err != nil {
			logx.Error("无法解析CIDR:", err)
			panic(err)
		}
		accessList = append(accessList, ipNet)
	}
}

// SetCtxData 解析 Header 数据并设置到 Request 的上下文中
func SetCtxData() func(c *gin.Context) {
	return func(c *gin.Context) {
		if len(accessList) == 0 {
			// 直接放行
			c.Next()
			return
		}

		// 判断 IP 是否由访问权限
		access := false
		cHost, _, err := net.SplitHostPort(clientIP(c.Request))
		if err != nil {
			logx.Errorf("解析【%s】失败", clientIP(c.Request))
			c.Abort()
			return
		}
		//本地ip
		if cHost == "::1" {
			cHost = "127.0.0.1"
		}
		cip := net.ParseIP(cHost)
		for i := 0; i < len(accessList); i++ {
			if accessList[i].Contains(cip) {
				access = true
				break
			}
		}
		if access {
			// 有权访问
			logx.Infof("【%s】有权访问", cip)
			c.Next()
		} else {
			// 禁止访问
			logx.Errorf("【%s】禁止访问", cip)
			c.Abort()
		}
	}
}

// 获取客户端 IP
func clientIP(r *http.Request) (remoteIp string) {
	remoteIp = r.RemoteAddr
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		remoteIp = ip
	} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
		remoteIp = ip
	}
	return remoteIp
}
