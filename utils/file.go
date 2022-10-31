package utils

import (
	"os"
)

// FileExists ...
func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// FileCreate ...
func FileCreate(path string) {
	f, err := os.Create(path)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	if err != nil {
		println(err)
	}
}
