package models

import (
	"gorm.io/datatypes"
)

type Content struct {
	Id      uint           `gorm:"primary_key" json:"vid"`
	Mid     uint           `json:"id" binding:"required"`
	Bt      string         `gorm:"type:varchar(40)" json:"bt" binding:"required"`
	Sbt     string         `gorm:"type:varchar(40)" json:"sbt" binding:"required"`
	Source  string         `gorm:"type:varchar(4)" json:"source" binding:"required"`
	Service string         `gorm:"type:varchar(8)" json:"service" binding:"required"`
	Op      string         `gorm:"type:varchar(8)" json:"op" binding:"required"`
	Status  string         `gorm:"type:varchar(8)" json:"status"`
	Reason  string         `gorm:"type:varchar(80)" json:"reason"`
	Data    interface{}    `gorm:"-"`
	Jdata   datatypes.JSON `gorm:"type:json;column:data" json:"data" binding:"required"`
}

//自定义表名
func (Content) TableName() string {
	return "oplog"
}

func GetOplogs(start, limit int) []Content {
	oplogs := []Content{}
	DB.Where("id >= ? AND id < ?", start, start+limit).Find(&oplogs)
	return oplogs
}

func GetDevId() uint {
	var c Content
	r := DB.Last(&c)
	if r.Error != nil {
		return 0
	}
	return c.Id
}

func GetMsId() uint {
	var c Content
	r := DB.Order("mid desc, mid").First(&c)
	if r.Error != nil {
		return 0
	}
	return c.Mid
}

func SetMsId(mid uint) {
	var c Content
	r := DB.Order("mid desc, mid").First(&c)
	if r.Error != nil {
		return
	}
	DB.Model(&c).Update("mid", mid)
}
