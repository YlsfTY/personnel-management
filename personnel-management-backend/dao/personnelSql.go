package dao

import "fmt"

// func (p *Personnel) CreatePersonnel() error {
// 	admin := User{}
// 	if err := Db.Where("name = ?", "admin").First(&admin).Error; err != nil {
// 		return err
// 	}
// 	if err := Db.Create(p).Error; err != nil {
// 		return err
// 	}
// 	if err := Db.Model(&admin).Association("Personnels").Append(p); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (p *Personnel) CreatePersonnel(u *User) error {
// 	if err := p.CreatePersonnel();err != nil {
// 		return err
// 	}
// 	if err := Db.Model(u).Association("Personnels").Append(p); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (p *Personnel) CreatePersonnel(users ...*User) error {
	admin := &User{
		Name: "admin",
	}

	if _, err := admin.SearchUserName(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	if err := Db.Create(p).Error; err != nil {
		fmt.Println(err.Error())
		return err
	}
	if err := Db.Model(&admin).Association("Personnels").Append(p); err != nil {
		fmt.Println(err.Error())
		return err
	}
	var Perls []*Personnel
	Perls = append(Perls, p)
	for _, u := range users {
		if u == nil {
			continue
		}
		if err := Db.Model(u).Association("Personnels").Replace(Perls); err != nil {
			return err
		}
	}
	return nil
}

func (p *Personnel) GetPersonnel(u *User) error {
	var Perls []*Personnel
	if err := Db.Model(u).Association("Personnels").Find(&Perls); err != nil {
		return err
	}
	if len(Perls) == 0 {
		return fmt.Errorf("The user does not exist.")
	}
	*p = *Perls[0]
	fmt.Println(p)
	return nil
}

func (p *Personnel) UpdatePersonnel() error {
	err := Db.Updates(p).Error
	return err
}
