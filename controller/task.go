package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

type Task struct {
	Source   string                 `json:"source"`
	Service  string                 `json:"service"`
	TaskType string                 `json:"tasktype"`
	Data     map[string]interface{} `json:"data"`
}

type Tasks struct {
	Contents []Task `json:"contents"`
}

func GetTask(c *gin.Context) {
	id := c.Query("taskid")
	c.JSON(200, gin.H{
		"description": "success",
		"taskid":      id,
	})
}

func DeleteTask(c *gin.Context) {
	id := c.Query("taskid")
	c.JSON(200, gin.H{
		"description": "success",
		"taskid":      id,
	})
}

func PostTask(c *gin.Context) {
	var d Tasks
	if err := c.BindJSON(&d); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"rcode": 1, "description": "Data format error cannot be handled"})
		return
	} else {
		fmt.Println(d)
		for _, i := range d.Contents {
			fmt.Println(i)
			for k, v := range i.Data {
				fmt.Println(k, v, reflect.TypeOf(v))
			}
		}
	}
	c.JSON(200, gin.H{
		"description": "success",
		"tasktype":    "test",
		"taskid":      110,
	})
}
