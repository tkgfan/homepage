// author gmfan
// date 2023/11/24

package web

import (
	"acsupport/common/errs"
	"fmt"
	"github.com/tkgfan/got/core/env"
	"io"
	"os"
	"strings"
	"sync"
)

const webRoot = "web/"

var pathMap = map[string]string{
	"":            "index.html",
	"index.html":  "index.html",
	"index":       "index.html",
	"config.json": "config.json",
	"favicon.ico": "favicon.ico",
}

var cache = sync.Map{}

// GetFile 获取文件信息
func GetFile(_path string) (res []byte, err error) {
	_path = strings.ReplaceAll(_path, "/", "")
	// 路径映射
	path := pathMap[_path]
	if path == "" {
		return nil, errs.NewCodeErrMgs(errs.ParamErr, fmt.Sprintf("路径【%s】不存在", _path))
	}

	if env.CurModel == env.DevModel {
		// 开发环境实时加载
		return readFile(path)
	}

	// 缓存加载数据
	if val, ok := cache.Load(path); ok {
		return val.([]byte), nil
	}

	// 加载数据
	res, err = readFile(path)
	if err != nil {
		return
	}

	// 缓存结果
	cache.Store(path, res)

	return
}

func readFile(path string) (res []byte, err error) {
	if path == configFile {
		// configFile 需要特殊处理，其会优先使用 dbFile 内的配置
		return GetConfig()
	}

	fs, err := os.Open(webRoot + path)
	if err != nil {
		return
	}
	res, err = io.ReadAll(fs)
	return
}
