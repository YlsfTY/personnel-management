package dao

// ---------- User表操作
// 查询用户是否存在(检验用户名)
func SearchUserName(name string) bool {
	result := MysqlDb.Limit(1).Find(&User{}, "user_name = ?", name)
	// 结果条数不为0返回真
	return result.RowsAffected != 0
}

// 用户校验（检验用户名与密码）,也可用于获取目标用户
func CheckUser(user *User) (bool, error) {
	// var targetUser User
	result := MysqlDb.Limit(1).
		Find(user, "user_name = ? AND user_password = ?", user.Name, user.Password)
	return result.RowsAffected != 0, result.Error
}

// 添加用户
func CreateUser(user *User) (bool, error) {
	// 添加用户不进行用户校验，需要自行调用
	result := MysqlDb.Select("user_name", "user_password").Create(&user)
	return result.Error == nil, result.Error
}

// 删除用户(state赋0，不删除数据库记录)
func DeleteUser(user *User) (bool, error) {
	CheckUser(user)
	user.State = false
	result := MysqlDb.Save(user)
	return result.Error == nil, result.Error
}
