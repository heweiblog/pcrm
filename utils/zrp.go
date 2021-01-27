package utils

import (
	"encoding/json"
	"pcrm/models"
)

type Zid struct {
	Zid string `json:"zid" binding:"required"`
}

func isZid(zid string) bool {
	if len(zid) > 0 && len(zid) < 128 {
		return true
	}
	return false
}

func ZidCheck(c *models.Content) string {
	p := new(Zid)
	if err := json.Unmarshal(c.Jdata, p); err != nil {
		return err.Error()
	}
	if isZid(p.Zid) == false {
		return "zid " + p.Zid + " format error"
	}
	c.Data = p
	return ""
}

type BackendForward struct {
	Group string `json:"group" binding:"required"`
}

func BackendForwardCheck(c *models.Content) string {
	p := new(Zid)
	if err := json.Unmarshal(c.Jdata, p); err != nil {
		return err.Error()
	}
	if isZid(p.Zid) == false {
		return "zid " + p.Zid + " format error"
	}
	c.Data = p
	return ""
}
