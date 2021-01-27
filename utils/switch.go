package utils

import (
	"encoding/json"
	"pcrm/models"
)

type Switch struct {
	Switch string `json:"switch" binding:"required"`
}

func SwitchCheck(c *models.Content) string {
	p := new(Switch)
	if err := json.Unmarshal(c.Jdata, p); err != nil {
		return err.Error()
	}
	if p.Switch != "enable" && p.Switch != "diaable" {
		return "switch " + p.Switch + " format error"
	}
	c.Data = p
	return ""
}

/*
func SwitchCheck(c *models.Content) string {
	m := make(map[string]interface{})
	if err := json.Unmarshal(c.Jdata, &m); err != nil {
		return err.Error()
	}
	if res, ok := m["switch"]; ok {
		if res, ok = res.(string); ok {
			if res == "enable" || res == "disable" {
				return ""
			}
		}
	}
	return "switch format error"
}
*/
