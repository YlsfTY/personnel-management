package dao

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

const (
	username = "root"         //数据库账号
	password = "admin123"     //密码
	host     = "127.0.0.1"    //地址
	port     = 3306           //端口
	DBname   = "personnel_db" //数据库名称
	timeout  = "10s"          //超时
)

var MysqlDb *gorm.DB
var MysqlDbErr error

// 连接数据库
func init() {
	// 连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&&parseTime=True&loc=Local&timeout=%s",username,password,host,port,DBname,timeout)

	// 连接
	MysqlDb,MysqlDbErr = gorm.Open(mysql.Open(dsn),&gorm.Config{
		// 日志
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if MysqlDbErr != nil {
		panic("数据库连接失败,failed to connect mysql,err:"+MysqlDbErr.Error())
	}
	fmt.Println("连接数据库成功,connect success!!")
	// 进行数据表初始化，如果表已存在，无操作
	AutoMigrate()

}

// 初始化数据表/自动迁移
func AutoMigrate(){
	err := MysqlDb.AutoMigrate(&User{})
	if err != nil {
		panic("初始化Users数据表失败")
	}
}
