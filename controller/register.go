package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pcrm/models"
)

var RegisterSource map[string]bool

func init() {
	RegisterSource = make(map[string]bool)
	RegisterSource["proxy"] = true
	RegisterSource["xforward"] = true
	RegisterSource["recursion"] = true
}

//注册接口，更新注册表，返回全量配置
func RegisterConf(c *gin.Context) {
	var m models.Register
	if err := c.BindJSON(&m); err != nil {
		c.JSON(200, gin.H{"rcode": 1, "description": "Data format error cannot be handled"})
		return
	} else {
		source := c.DefaultQuery("source", "")
		//获取请求ip只有ip
		ip := c.ClientIP()
		//ip := c.Request.RemoteAddr 获取请求ip含有端口
		if _, ok := RegisterSource[source]; ok && ip != "" && m.Port > 0 && m.Port < 65535 && m.ConfUrl != "" && m.TaskUrl != "" {
			m.Module = source
			m.Ip = ip
			m.ConfUrl = fmt.Sprintf("http://%s:%d%s", ip, m.Port, m.ConfUrl)
			m.TaskUrl = fmt.Sprintf("http://%s:%d%s", ip, m.Port, m.TaskUrl)
			fmt.Println(m)
			models.UpdateRegister(m)
		}
	}
	//返回全量配置信息，导出注册模块所有配置，models.GetOplogs(1, 2)暂模拟一下
	c.JSON(200, gin.H{"contents": models.GetOplogs(1, 2)})
}
