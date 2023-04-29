package dao

import "fmt"

func (p *Personnel) CreatePersonnel(users ...*User) error {

	if err := Db.Create(p).Error; err != nil {
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
		p = nil
		return fmt.Errorf("the user does not exist")
	}
	*p = *Perls[0]
	fmt.Println(p)
	return nil
}

func (p *Personnel) UpdatePersonnel() error {
	err := Db.Updates(p).Error
	return err
}

func DeletePersonnel(u *User) error {
	var pList []*Personnel
	if err := Db.Model(u).Association("Personnels").Find(&pList); err != nil {
		return err
	}
	if len(pList) > 0 {
		if err := Db.Model(u).Association("Personnels").Clear(); err != nil {
			return err
		}
		if err := Db.Delete(pList).Error; err != nil {
			return err
		}
	}
	return nil
}
