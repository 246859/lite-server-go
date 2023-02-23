package system

import (
	"github.com/246859/lite-server-go/controller/v1c"
	"github.com/246859/lite-server-go/global"
	"github.com/246859/lite-server-go/utils/jwtutils"
	"github.com/246859/lite-server-go/utils/responseuils"
	"github.com/gin-gonic/gin"
)

type FileController struct {
}

// UploadMultipart
// @Date 2023-02-20 15:11:51
// @Param ctx *gin.Context
// @Method http.MethodPut
// @Description: 多文件上传接口
func (f FileController) UploadMultipart(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
		return
	}
	files := form.File["file"]
	if len(files) == 0 {
		responseuils.FailWithMsg(ctx, global.I18nRawCN("file.count.empty"))
		return
	} else if len(files) > 5 {
		responseuils.FailWithMsg(ctx, global.I18nDataCN("file.count.exceed", 5))
		return
	}
	for _, file := range files {
		// 单个文件大于 10MB
		if file.Size > 10*1024*1024 {
			responseuils.FailWithMsg(ctx, global.I18nDataCN("file.size.exceed", 10))
			return
		}
	}
	claims, err := jwtutils.ToJwtClaims(ctx)
	if err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
		return
	}
	fileInfo, err := v1c.FileService.SaveUploadMultiFile(files, ctx, claims)
	if err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
	} else {
		responseuils.OkWithMsgAndData(ctx, fileInfo, global.I18nRawCN("file.upload.ok"))
	}
}

// Upload
// @Date 2023-02-23 23:26:31
// @Param ctx *gin.Context
// @Description: 单文件上传接口
func (f FileController) Upload(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
		return
	}
	claims, err := jwtutils.ToJwtClaims(ctx)
	if err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
		return
	}

	// 保存上传的文件
	fileInfo, err := v1c.FileService.SaveUploadSingle(fileHeader, ctx, claims)
	if err != nil {
		responseuils.FailWithMsg(ctx, err.Error())
		return
	} else {
		responseuils.OkWithMsgAndData(ctx, fileInfo, global.I18nRawCN("file.upload.ok"))
	}
}
