package controller

import (
	"github.com/gin-gonic/gin"
	"pcrm/config"
	"pcrm/models"
)

//心跳接口get请求，返回设备id，ms id等信息
func Heartbeat(c *gin.Context) {
	m := make(map[string]interface{})
	m["status"] = "running"
	m["msrelease"] = MsRelease
	m["devicerelease"] = "1.0"
	m["msversion"] = models.GetMsId()
	m["deviceversion"] = models.GetDevId()
	m["softwareversion"] = config.Version
	m["licenseinfo"] = "abcABC=="

	c.JSON(200, m)
}

//心跳接口post请求，接收一个ms id，更新到数据库
func ResetMsId(c *gin.Context) {
	m := make(map[string]interface{})
	if err := c.BindJSON(&m); err != nil {
		c.JSON(200, gin.H{"rcode": 1, "description": "Data format error cannot be handled"})
		return
	} else {
		if mid, ok := m["msversion"]; ok {
			if id, ok := mid.(float64); ok {
				models.SetMsId(uint(id))
				c.JSON(200, gin.H{"status": "success"})
				return
			}
		}
	}
	c.JSON(200, gin.H{"status": "failed"})
}
