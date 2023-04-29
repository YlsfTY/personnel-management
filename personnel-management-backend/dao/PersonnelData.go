package dao

import "github.com/go-playground/validator/v10"

type Personnel struct {
	Name        string `gorm:"column:name; type:varchar(20); not null" json:"name"`
	Age         uint   `gorm:"column:age;size:4; not null" json:"age"`
	Phone       string `gorm:"column:phone; type:varchar(11); not null" json:"phone"`
	Birthday    uint   `gorm:"column:birthday; not null" json:"birthday"`
	IDNumber    string `gorm:"column:id_number; type:varchar(20);primaryKey" json:"IDNumber"`
	Residence   string `gorm:"column:residence; type:varchar(120); not null" json:"residence"`
	Education   string `gorm:"column:education; type:varchar(20); not null" json:"education"`
	School      string `gorm:"column:school; type:varchar(80); not null" json:"school"`
	Department  string `gorm:"column:department; type:varchar(50); not null" json:"department"`
	EntryTime   uint   `gorm:"column:entry_time; not null" json:"entryTime"`
	RegularTime *uint  `gorm:"column:regular_time" json:"regularTime"`
	Salary      uint   `gorm:"column:salary; not null" json:"salary"`
	Additional  string `gorm:"column:additional; type:varchar(800)" json:"additional"`
}

type ListData struct {
	UserName      string `json:"userName"`
	PersonnelName string `json:"name"`
	Department    string `json:"department"`
}

func (p *Personnel) Validate() error {
	var validate = validator.New()
	if err := validate.Struct(p); err != nil {
		return err
	}
	return nil
}
