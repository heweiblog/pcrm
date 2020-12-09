package utils

import (
	"encoding/json"
	"pcrm/models"
)

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
