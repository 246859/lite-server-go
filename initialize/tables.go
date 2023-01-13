package initialize

import (
	"go.uber.org/zap"
	"liteserver/config"
	"liteserver/model"
	"liteserver/utils"
)

// IniTables
// @Date 2023-01-13 15:52:16
// @Param list model.TableList
// @Param db *gorm.DB
// @Description: 初始化数据库表
func IniTables(model *model.TableGroup, gormDB *config.GormDBGroup) {

	modelGroup := *model
	gormGroup := *gormDB
	for name, modelList := range modelGroup {
		if db, exist := gormGroup[name]; exist {
			utils.MustOrLogPanic(func() error {
				return db.AutoMigrate(modelList...)
			}, "数据库表自动迁移失败", zap.String("数据库名称", name))
		} else {
			zap.L().Error("不存在的对应名称的GORM数据库", zap.String("名称", name))
		}
	}
}
