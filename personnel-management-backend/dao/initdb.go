package dao

import (
	"fmt"
	"personnel-management-backend/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
var DbErr error

// 连接数据库
func DBinit() {
	cfg := config.Conf.Database
	// 连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&&parseTime=True&loc=Local&timeout=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Timeout)
	// 连接
	Db, DbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 日志
		Logger: logger.Default.LogMode(logger.Info),
	})
	if DbErr != nil {
		panic("数据库连接失败,failed to connect mysql,err:" + DbErr.Error())
	}
	sqlDb, err := Db.DB()
	if err != nil {
		panic("无法获取DB实例,err:" + err.Error())
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(20)
	fmt.Println("连接数据库成功,connect success!!")
	// 进行数据表初始化，如果表已存在，无操作
	AutoMigrate()
}

// 初始化数据表/自动迁移
func AutoMigrate() {
	err := Db.AutoMigrate(&User{})
	if err != nil {
		panic("初始化Users数据表失败")
	}
}

func CloseDB() {
	db, err := Db.DB()
	if err != nil {
		panic(err)
	}
	db.Close()
}
