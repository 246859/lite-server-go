package initialize

import (
	"github.com/246859/lite-server-go/config"
	"github.com/246859/lite-server-go/utils/jwtutils"
)

// InitJwtUtils
// @Date 2023-02-07 17:24:49
// @Param cfg *config.JwtConfig
// @Description: 初始化JWT配置
func InitJwtUtils(cfg *config.JwtConfig) {
	jwtutils.SetConfig(cfg)
}
