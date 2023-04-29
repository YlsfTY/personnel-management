package dao

import "fmt"

func GetUserList(limit, offset int) ([]*User, error) {
	var users []*User
	if err := Db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetPerList(limit, offset int) ([]*Personnel, error) {
	var perArr []*Personnel
	if err := Db.Offset(offset).Limit(limit).Find(&perArr).Error; err != nil {
		return nil, err
	}
	return perArr, nil
}

func GetDataList() ([]*ListData, error) {
	var users []*User
	if err := Db.Select("id,user_name").Find(&users).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	listData := make([]*ListData, len(users))
	for i, u := range users {
		var per []*Personnel
		if err := Db.Model(u).Association("Personnels").Find(&per); err != nil {
			fmt.Println(err)
			return nil, err
		}
		if len(per) == 0 {
			listData[i] = &ListData{
				UserName:      u.Name,
				PersonnelName: "未填写",
				Department:    "未填写",
			}
		} else {
			listData[i] = &ListData{
				UserName:      u.Name,
				PersonnelName: per[0].Name,
				Department:    per[0].Department,
			}
		}
	}
	return listData, nil
}
