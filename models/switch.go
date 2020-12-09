package models

import (
	//"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Switch struct {
	gorm.Model `json:"-"`
	Service    string `gorm:"type:varchar(8);not null" json:"-"`
	Bt         string `gorm:"type:varchar(40);not null" json:"-"`
	Sbt        string `gorm:"type:varchar(40);not null" json:"-"`
	Switch     string `gorm:"type:varchar(8);not null" json:"switch"`
}

func GetSwitchs() []Switch {
	switchs := []Switch{}
	DB.Find(&switchs)
	return switchs
}

//func (s *Switch) UpdateSwitch(bt, sbt, service, status, string) {
//}
