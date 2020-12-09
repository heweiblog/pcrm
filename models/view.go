package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PrivateView struct {
	gorm.Model `json:"-"`
	View       string         `gorm:"type:varchar(80)" json:"view"`
	SrcIpList  datatypes.JSON `gorm:"type:json" json:"srciplist"`
	DstIpList  datatypes.JSON `gorm:"type:json" json:"dstiplist"`
	Threshold  uint           `json:"threshold"`
}
