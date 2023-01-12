package initialize

import (
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"liteserver/config"
)

// InitGorm
// @Date: 2023-01-08 22:52:21
// @Description: 初始化GORM连接数据库
// @Param: configGroup config.DataBaseConfigGroup
// @Return: error
func InitGorm(configGroup *config.DataBaseConfigGroup) *config.GormDBGroup {

	gormGroup := config.GormDBGroup{}

	// 遍历db map，根据dp的类型进行初始化
	for name, databaseConfig := range *configGroup {
		// 是否启用
		if !databaseConfig.Enable {
			continue
		}
		sqlType := strings.ToLower(databaseConfig.Type)
		switch sqlType {
		case config.Mysql:
			{
				gormGroup[name] = gormMysql(databaseConfig)
			}
		case config.PostgreSql:
			{
				gormGroup[name] = gormPostgreSql(databaseConfig)
			}
		}
	}

	return &gormGroup
}

func CloseGormGroup(group config.GormDBGroup) error {
	return nil
}

// gormMysql
// @Date: 2023-01-09 10:46:47
// @Description: mysql建立连接
// @Param: generalConfig *config.DataBaseConfig
// @Return: *gorm.DB
func gormMysql(generalConfig *config.DataBaseConfig) *gorm.DB {
	mysqlConfig := config.MysqlConfig{C: generalConfig}
	// 初始化Gorm
	db, err := gorm.Open(mysql.Open(mysqlConfig.Dsn()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	// 进行配置设置
	sqlDB.SetMaxIdleConns(mysqlConfig.C.MaxIldeCons)
	sqlDB.SetMaxOpenConns(mysqlConfig.C.MaxOpenCons)
	return db
}

func gormPostgreSql(config *config.DataBaseConfig) *gorm.DB {
	panic("暂不支持该数据库")
}
