package DCGoWebFramework

import (
	"os"
	"path/filepath"
)

// golang新版本的应该
func IsPathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// 拿到静态文件的本地路径
func GetFilePath(upath string, staticPath string) (string, bool) {
	staticFilePath := ""
	basePath, _ := os.Getwd()

	if upath == "favicon.ico" {
		staticFilePath = filepath.Join(basePath, staticPath, "favicon.ico")
	} else if string([]rune(upath)[:8]) == "/"+staticPath+"/" {
		staticFilePath = filepath.Join(basePath, upath)
	}

	if IsPathExist(staticFilePath) {
		return staticFilePath, false
	}

	return "", true
}