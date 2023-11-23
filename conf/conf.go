// author gmfan
// date 2023/07/01

package conf

import (
	"github.com/tkgfan/got/core/env"
	"github.com/tkgfan/got/core/logx"
)

var (
	Port     = "8888"
	Timeout  = 60 * 1000
	LogLevel = logx.InfoLevel
)

func init() {
	ginTemplateInit()
}

func ginTemplateInit() {
	must := false

	env.LoadStr(&Port, "PORT", must)
	env.LoadInt(&Timeout, "TIMEOUT", must)
	env.LoadStr(&LogLevel, "LOG_LEVEL", must)
	logx.SetLevel(LogLevel)
}
