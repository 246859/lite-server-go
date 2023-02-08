package initialize

import (
	"liteserver/config"
	"liteserver/utils/jwtutils"
)

// InitJwtUtils
// @Date 2023-02-07 17:24:49
// @Param cfg *config.JwtConfig
// @Description: 初始化JWT配置
func InitJwtUtils(cfg *config.JwtConfig) {
	jwtutils.SetConfig(cfg)
}
