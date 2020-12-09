package controller

import (
	"github.com/gin-gonic/gin"
	"pcrm/models"
	"strconv"
)

func GetOplog(c *gin.Context) {
	version, _ := strconv.Atoi(c.DefaultQuery("startversion", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if version == 0 || limit == 0 {
		c.JSON(200, gin.H{})
		return
	}
	c.JSON(200, models.GetOplogs(version, limit))
}
