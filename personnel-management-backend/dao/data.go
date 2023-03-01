package dao

import "gorm.io/gorm"

// 用户表
type User struct {
	gorm.Model        //带有ID CreateAt、UpdateAt、DeletedAt的结构体
	Name       string `gorm:"column:user_name; type:varchar(50); unique:true" json:"userName"`
	Password   string `gorm:"column:user_password; type:varchar(500); not null" json:"userPassword"`
	State      bool   `gorm:"column:state; not null; type:tinyint(1); default:1" json:"userState"`
}
