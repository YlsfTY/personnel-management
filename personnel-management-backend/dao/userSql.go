package dao

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

// ---------- User表操作
// 寻找用户
// true用户名存在
func (u *User) SearchUserName() (bool, error) {
	// result := Db.Limit(1).Find(user, "user_name = ?", user.Name)
	err := Db.Take(u, "user_name = ?", u.Name).Error
	if err != nil {
		// 判断err是否是未找到数据错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			log.Printf("Error while checking if username exists: %v", err)
			return false, err
		}
	} else {
		return true, nil
	}
}

// 添加用户
// true操作成功;false err=nil用户存在
func (u *User) CreateUser() (bool, error) {
	exist, err := u.SearchUserName()
	if err != nil {
		return false, err
	} else if exist {
		return false, nil
	}

	err = Db.Select("user_name", "user_password").Create(&u).Error
	if err != nil {
		log.Printf("Error while creating user: %v", err)
		return false, err
	}
	return true, nil
}

// 删除用户(state赋0，不删除数据库记录)
// true操作成功;false err=nil用户不存在
func (u *User) DeleteUser() (bool, error) {
	exist, err := u.SearchUserName()
	if err != nil {
		return false, err
	} else if !exist {
		return false, nil
	}

	u.State = false
	// u.DeletedAt = gorm.DeletedAt{
	// Time:time.Now(),
	// Valid: true,
	// }
	err = Db.Save(u).Error
	if err != nil {
		log.Printf("Error while updating user state: %v", err)
		return false, err
	}
	return true, nil
}

// 检验用户状态
func (u *User) CheckUserState() (bool, error) {
	exist, err := u.SearchUserName()
	if err != nil {
		// 查询错误
		return true, err
	} else if !exist {
		// 用户不存在
		return false, fmt.Errorf("The user does not exist.")
	}
	if u.State {
		// 用户正常
		return true, nil
	}
	// state=false
	return false, nil
}

func (user *User) IsUserWithPer() (bool, error) {
	// 检查 `User` 对象和 `Personnel` 对象之间是否存在关联关系
	// var count int64
	result := Db.Model(user).Association("Personnels")

	// 如果存在关联关系，则返回 true；否则返回 false
	return result.Count() > 0, result.Error
}
