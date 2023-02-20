package file

import (
	"github.com/gin-gonic/gin"
	"liteserver/global"
	"liteserver/model/file"
	"liteserver/utils/fileutils"
	"liteserver/utils/jwtutils"
	"liteserver/utils/uuidtool"
	"mime/multipart"
	"path/filepath"
)

type FileService struct {
}

// SaveUpload
// @Date 2023-02-20 16:15:20
// @Param fileHeader *multipart.FileHeader
// @Param ctx *gin.Context
// @Return *file.FileInfo
// @Return error
// @Description: 上传文件
func (f FileService) SaveUpload(fileHeaders []*multipart.FileHeader, ctx *gin.Context) ([]file.FileInfo, error) {
	claims, err := jwtutils.ToJwtClaims(ctx)
	if err != nil {
		return nil, err
	}
	// 创建文件信息切片
	fileInfoSlice := make([]file.FileInfo, 0, len(fileHeaders))
	for _, fileHeader := range fileHeaders {
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
			return nil, err
		}
		// 保存文件
		err := ctx.SaveUploadedFile(fileHeader, path)

		if err != nil {
			return nil, err
		} else {
			fileInfoSlice = append(fileInfoSlice, file.FileInfo{
				Name: newFilename,
				Type: suffix,
				Url:  fileutils.ToForwardSlash(userdir),
			})
		}

	}
	return fileInfoSlice, nil
}
