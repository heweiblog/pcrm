package controller

import (
	"github.com/gin-gonic/gin"
	"pcrm/models"
	"pcrm/utils"
)

type Contents struct {
	MsRelease string           `json:"msrelease" binding:"required"`
	Contents  []models.Content `json:"contents" binding:"required"`
	//Contents []utils.Content `json:"contents" binding:"required"`
}

var MsRelease string

func PostConfig(c *gin.Context) {
	var cts Contents
	//BindJSON()出错会在header写入400 ShouldBindJSON不会
	if err := c.ShouldBindJSON(&cts); err != nil {
		c.JSON(200, gin.H{"rcode": 1, "description": "Data format error cannot be handled"})
	}
	MsRelease = cts.MsRelease
	go utils.HandleData(cts.Contents)
	c.JSON(200, gin.H{"status": "received", "rcode": 0})
}

func GetConfigs(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "received",
		"rcode":  0,
		"conf":   "all",
	})
}
