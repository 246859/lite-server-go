package fileutils

import (
	"go.uber.org/zap"
	"liteserver/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
func MustCreateDirAndFile(path string) {
	MustMkdir(filepath.Dir(path))
	MustCreateFile(path)
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

// JoinPath
// @Date 2023-01-13 21:44:17
// @Param path string
// @Return string
// @Description: 如果是相对路径就拼接，如果是绝对路径就返回
func JoinPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(GetCurrentPath(), path)
}

// FileSuffix
// @Date 2023-02-20 14:55:19
// @Param file string
// @Return string
// @Description: 获取一个文件的后缀
func FileSuffix(file string) string {
	_, fp := filepath.Split(file)
	return fp[strings.Index(fp, ".")+1:]
}

// CreateFile
// @Date 2023-02-20 15:36:23
// @Description: 根据路径创建一个文件，如果文件夹不存在，则先创建文件夹，再创建文件
func CreateFile(path string) error {
	dir, file := filepath.Split(path)
	// 创建目录
	if err := CreateDir(dir); err == nil {
		if _, err := os.Create(file); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return err
	}
}

// CreateDir
// @Date 2023-02-20 15:42:10
// @Param dir string
// @Return error
// @Method
// @Description: 创建一个文件夹，如果文件夹已存在则不会返回错误
func CreateDir(path string) error {
	dir, _ := filepath.Split(path)
	if err := os.Mkdir(dir, os.ModeDir); os.IsExist(err) || err == nil {
		return nil
	} else {
		return err
	}
}

// ToForwardSlash
// @Date 2023-02-20 16:09:25
// @Description: 路径正斜杠
func ToForwardSlash(path string) string {
	return strings.ReplaceAll(path, `\`, `/`)
}
