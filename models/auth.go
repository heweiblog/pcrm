package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PrivateZone struct {
	gorm.Model     `json:"-"`
	View           string          `gorm:"type:varchar(80)" json:"view_name"`
	Zone           string          `gorm:"type:varchar(256)" json:"zone_name"`
	Email          string          `gorm:"type:varchar(256)" json:"email"`
	Recursion      bool            `gorm:"not null" json:"proxy_external"`
	PrivateRecords []PrivateRecord `json:"-"`
}

type PrivateRecord struct {
	gorm.Model    `json:"-"`
	View          string         `gorm:"type:varchar(80)" json:"view_name"`
	Zone          string         `gorm:"type:varchar(256)" json:"zone_name"`
	Name          string         `gorm:"type:varchar(63)" json:"name"`
	Type          string         `gorm:"type:varchar(10)" json:"type"`
	Records       datatypes.JSON `gorm:"type:json" json:"records"`
	Ttl           uint           `gorm:"json:"ttl"`
	PrivateZoneID uint           `json:"-"`
}
