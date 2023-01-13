package fileutils

import (
	"go.uber.org/zap"
	"liteserver/utils"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

// IsExist
// @Date 2023-01-12 18:04:48
// @Param path string
// @Return bool
// @Description: 判断一个文件或者文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}

// MustMkdir
// @Date 2023-01-12 18:11:35
// @Param path string
// @Description: 必须创建文件夹，否则就panic
func MustMkdir(path string) {
	if !IsExist(path) {
		utils.MustOrLogPanic(func() error {
			return os.Mkdir(path, os.ModeDir)
		}, "文件目录创建失败", zap.String("文件目录", path))
	}
}

// MustCreateFile
// @Date 2023-01-12 18:21:26
// @Param filepath string
// @Description: 必须创建文件，否则就panic
func MustCreateFile(filepath string) {
	if !IsExist(filepath) {
		utils.MustOrLogPanic(func() error {
			_, err := os.Create(filepath)
			return err
		}, "文件创建失败", zap.String("文件路径", filepath))
	}
}

// MustCreateDirAndFile
// @Date 2023-01-12 18:19:36
// @Param filepath string
// @Description: 必须创建文件，否则就panic
func MustCreateDirAndFile(filepath string) {
	MustMkdir(path.Dir(filepath))
	MustCreateFile(filepath)
}

// GetCurrentPath
// @Date 2023-01-13 19:42:52
// @Return string
// @Description: 获取当前执行文件的系统绝对路径
func GetCurrentPath() string {
	lookPath, _ := exec.LookPath(os.Args[0])
	abs, err := filepath.Abs(lookPath)
	if err != nil {
		return ""
	}
	return filepath.Dir(abs)
}
