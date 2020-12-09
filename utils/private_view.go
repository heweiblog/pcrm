package utils

import (
	"encoding/json"
	"pcrm/models"
)

type Pview struct {
	View string   `json:"view"`
	Src  []string `json:"srciplist"`
	Dst  []string `json:"srciplist"`
}

type PviewLimit struct {
	View      string `json:"view" binding:"required"`
	Threshold uint   `json:"threshold" binding:"required"`
}

/*
func PviewCheck(c *models.Content) string {
	var p Pview
	if err := json.Unmarshal(c.Data, &p); err != nil {
		return err.Error()
	}
	fmt.Println(p)
	return ""
}
*/

func PviewLimitCheck(c *models.Content) string {
	p := new(PviewLimit)
	if err := json.Unmarshal(c.Jdata, p); err != nil {
		return err.Error()
	}
	if p.Threshold > 0xffffffff {
		return "view threshold value range error"
	}
	c.Data = p
	return ""
}
