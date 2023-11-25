// author gmfan
// date 2023/11/25
package web

import (
	"acsupport/v1/models"
	"bytes"
	"encoding/json"
	"io"
	"os"
	"sync"
)

const (
	dbFile     = "db"
	configFile = "config.js"
)

var mu = sync.Mutex{}

// UpdateConfig 修改配置文件，实际修改 dbFile，在获取配置文件 config.js
// 时会优先使用 dbFile 内的配置。
func UpdateConfig(src *models.Config) (resp any, err error) {
	// 获取目标文件
	dst, err := getDB()
	if err != nil {
		return
	}

	// 更新
	if src.Header.Title != "" {
		dst.Header = src.Header
	}
	if len(src.Links) != 0 {
		dst.Links = src.Links
	}

	// 保存到文件
	err = saveDB(dst)
	if err != nil {
		return
	}
	return
}

// GetConfig 获取配置文件，优先使用 dbFile 内配置
func GetConfig() (res []byte, err error) {
	// 获取 db
	db, err := getDB()
	if err != nil {
		return
	}

	// 获取 config.js 内的配置
	prefix, cfg, err := getConfigJS()
	if err != nil {
		return
	}

	// 优先使用 db 配置
	if db.Header.Title != "" {
		cfg.Header = db.Header
	}
	if len(db.Links) > 0 {
		cfg.Links = db.Links
	}

	// 序列化结果
	res, err = json.Marshal(cfg)
	if err != nil {
		return
	}
	// 拼接前缀
	res = append(prefix, res...)
	return
}

func getConfigJS() (prefix []byte, cfg models.Config, err error) {
	fs, err := os.OpenFile(webRoot+configFile, os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer fs.Close()

	// 解析 config.json，其格式为 prefix={}
	bs, err := io.ReadAll(fs)
	if err != nil {
		return
	}
	prefix = bs[:indexByte(bs, '{')]
	bs = bs[len(prefix):]
	err = json.Unmarshal(bs, &cfg)
	if err != nil {
		return
	}
	return
}

func indexByte(arr []byte, r byte) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == r {
			return i
		}
	}
	return -1
}

func saveDB(db models.Config) (err error) {
	mu.Lock()
	defer mu.Unlock()
	
	fs, err := os.OpenFile(webRoot+dbFile, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer fs.Close()

	// 序列化 db
	bs, err := json.Marshal(db)
	if err != nil {
		return
	}

	// 写入文件
	_, err = io.Copy(fs, bytes.NewReader(bs))
	if err != nil {
		return
	}
	return
}

func getDB() (res models.Config, err error) {
	mu.Lock()
	defer mu.Unlock()

	fs, err := os.OpenFile(webRoot+dbFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer fs.Close()

	// 解析结果
	bs, err := io.ReadAll(fs)
	if err != nil {
		return
	}
	err = json.Unmarshal(bs, &res)
	if err != nil {
		return
	}
	return
}
