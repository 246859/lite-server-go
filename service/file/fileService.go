package file

import (
	"github.com/246859/lite-server-go/global"
	"github.com/246859/lite-server-go/model/response"
	"github.com/246859/lite-server-go/utils/fileutils"
	"github.com/246859/lite-server-go/utils/jwtutils"
	"github.com/246859/lite-server-go/utils/uuidtool"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
)

type FileService struct {
}

// SaveUploadMultiFile
// @Date 2023-02-20 16:15:20
// @Param fileHeader *multipart.FileHeader
// @Param ctx *gin.Context
// @Return *file.FileInfo
// @Return error
// @Description: 上传文件
func (f FileService) SaveUploadMultiFile(fileHeaders []*multipart.FileHeader, ctx *gin.Context) ([]response.FileInfo, error) {
	// 创建文件信息切片
	fileInfoSlice := make([]response.FileInfo, 0, len(fileHeaders))
	for _, fileHeader := range fileHeaders {
		if fileInfo, err := SaveUploadSingle(fileHeader, ctx); err != nil {
			return nil, err
		} else {
			fileInfoSlice = append(fileInfoSlice, fileInfo)
		}
	}
	return fileInfoSlice, nil
}

func SaveUploadSingle(fileHeader *multipart.FileHeader, ctx *gin.Context) (response.FileInfo, error) {
	claims, err := jwtutils.ToJwtClaims(ctx)
	if err != nil {
		return response.FileInfo{}, err
	}
	// 获取文件后缀
	suffix := fileutils.FileSuffix(fileHeader.Filename)
	// 新文件名
	newFilename := uuidtool.NewUUIDv5() + "." + suffix
	// 用户目录
	userdir := filepath.Join(claims.UserUUID, newFilename)
	// 目录
	path := filepath.Join(global.WorkDir, filepath.Join("static", userdir))
	// 创建目录
	if err := fileutils.CreateDir(path); err != nil {
		return response.FileInfo{}, err
	}
	// 保存文件
	err = ctx.SaveUploadedFile(fileHeader, path)

	if err == nil {
		return response.FileInfo{
			Name: newFilename,
			Type: suffix,
			Url:  fileutils.ToForwardSlash(userdir),
		}, nil
	} else {
		return response.FileInfo{}, err
	}
}
