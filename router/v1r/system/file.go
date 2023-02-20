package system

import (
	"liteserver/router/route"
	"liteserver/router/v1r"
	"net/http"
)

const File = "file"

const Upload = "upload"

type FileRouter struct {
}

func (f FileRouter) InitRouter() route.RouterMap {
	return route.RouterMap{
		Upload: route.Api{
			Path:    Upload,
			Method:  http.MethodPut,
			Handler: v1r.FileController.Upload,
		},
	}
}
