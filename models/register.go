package models

import (
	"gorm.io/gorm"
)

type Register struct {
	gorm.Model `json:"-"`
	Ip         string `gorm:"type:varchar(40);not null"`
	Port       uint   `gorm:"not null" json:"port"`
	Module     string `gorm:"type:varchar(40);not null"`
	ConfUrl    string `gorm:"type:varchar(80);not null" json:"conf_url"`
	TaskUrl    string `gorm:"type:varchar(80);not null" json:"task_url"`
}

func (Register) TableName() string {
	return "register"
}

func GetRegisters() []Register {
	registers := []Register{}
	DB.Find(&registers)
	return registers
}

func UpdateRegister(m Register) {
	var r Register
	result := DB.Where("module = ?", m.Module).First(&r)
	if result.Error == nil {
		r.Ip = m.Ip
		r.Port = m.Port
		r.Module = m.Module
		r.ConfUrl = m.ConfUrl
		r.TaskUrl = m.TaskUrl
		DB.Save(&r)
	} else {
		DB.Create(&m)
	}
}
