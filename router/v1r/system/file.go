package system

import (
	"github.com/246859/lite-server-go/router/route"
	"github.com/246859/lite-server-go/router/v1r"
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
