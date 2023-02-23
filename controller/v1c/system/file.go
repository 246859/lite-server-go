package system

import (
	"github.com/246859/lite-server-go/controller/v1c"
	"github.com/246859/lite-server-go/utils/response"
	"github.com/gin-gonic/gin"
)

type FileController struct {
}

// Upload
// @Date 2023-02-20 15:11:51
// @Param ctx *gin.Context
// @Method http.MethodPut
// @Description: 上传文件
func (f FileController) Upload(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}
	files := form.File["file"]
	if len(files) == 0 {
		response.FailWithMsg(ctx, "文件列表为空")
		return
	} else if len(files) > 5 {
		response.FailWithMsg(ctx, "单次上传文件数量不能大于5个")
		return
	}
	for _, file := range files {
		// 单个文件大于 10MB
		if file.Size > 10*1024*1024 {
			response.FailWithMsg(ctx, "单个文件最大不能超过10MB")
			return
		}
	}
	fileInfo, err := v1c.FileService.SaveUploadMultiFile(files, ctx)
	if err != nil {
		response.FailWithMsg(ctx, err.Error())
	} else {
		response.OkWithMsgAndData(ctx, fileInfo, "文件上传成功")
	}
}
