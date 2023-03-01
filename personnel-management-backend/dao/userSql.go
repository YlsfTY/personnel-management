package dao

// import "golang.org/x/crypto/bcrypt"

// ---------- User表操作
// 查询用户是否存在(检验用户名)
func SearchUserName(user *User) bool {
	result := MysqlDb.Limit(1).Find(user, "user_name = ?", user.Name)
	// 结果条数不为0返回真
	return result.RowsAffected != 0
}

// 添加用户
func CreateUser(user *User) (bool, error) {
	// 添加用户不进行用户校验，需要自行调用
	result := MysqlDb.Select("user_name", "user_password").Create(&user)
	return result.Error == nil, result.Error
}

// 删除用户(state赋0，不删除数据库记录)
func DeleteUser(user *User) (bool, error) {
	SearchUserName(user)
	user.State = false
	result := MysqlDb.Save(user)
	return result.Error == nil, result.Error
}

func CheckState(user *User) (bool, error) {
	result := MysqlDb.Limit(1).
		Find(user, "state = ?", true)
	return result.RowsAffected != 0, result.Error
}
