package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"reflect"
)

type Switch struct {
	Switch string `json:"switch"`
}

type Rcode struct {
	Responsecode int `json:"responsecode"`
}

type Content struct {
	Id      uint   `gorm:"primary_key" json:"vid"`
	Mid     uint   `json:"id"`
	Bt      string `gorm:"type:varchar(40);not null" json:"bt"`
	Sbt     string `gorm:"type:varchar(40);not null" json:"sbt"`
	Source  string `gorm:"type:varchar(4);not null" json:"source"`
	Service string `gorm:"type:varchar(8);not null" json:"service"`
	Op      string `gorm:"type:varchar(8);not null" json:"op"`
	Reason  string `gorm:"type:varchar(80)" json:"reason"`
	//Data    datatypes.JSON `gorm:"type:json" json:"data"`
	Data interface{} `gorm:"type:json" json:"data"`
}

type Contents struct {
	MsRelease string `json:"msrelease"`
	//Contents  []Content `json:"contents"`
	Contents []map[string]interface{} `json:"contents"`
}

func testPost(request Contents) {
	url := "http://127.0.0.1:22222/"

	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(request)

	fmt.Println(requestBody)

	req, err := http.NewRequest("POST", url, requestBody)
	req.Header.Set("Content-Type", "application/json")

	fmt.Println(req)
	client := &http.Client{}
	fmt.Println(client)
	resp, err := client.Do(req)
	fmt.Println(resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func PostConfig(c *gin.Context) {
	var d Contents
	if err := c.BindJSON(&d); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"rcode": 1, "description": "Data format error cannot be handled"})
		return
	} else {
		fmt.Println("recv:", reflect.TypeOf(d), d)
		fmt.Println("recv data:", reflect.TypeOf(d.Contents[0]["data"]), d.Contents[0]["data"])
		fmt.Println("recv bt:", d.Contents[0]["bt"])

		/*
			for i := 0; i < len(d.Contents); i++ {
				fmt.Println("recv:", reflect.TypeOf(d.Contents[i].Data), d.Contents[i].Data)

					if _, ok := utils.CheckMethods[d.Contents[i].Service][d.Contents[i].Bt][d.Contents[i].Sbt]; ok {
						//校验模块
						if res := utils.CheckMethods[d.Contents[i].Service][d.Contents[i].Bt][d.Contents[i].Sbt](&d.Contents[i]); res == "" {
							//校验通过 入队(通道)
							testPost(d)
							fmt.Println("数据校验通过")
						} else {
							//校验失败，直接记录oplog，或者发到通道，在另一个协程中收取写入
							fmt.Println("数据校验失败", res)
						}
					} else {
						//service bt sbt或对应的函数有问题，直接记录oplog失败，或者发到通道，在另一个协程中收取写入
						fmt.Println("数据大格式错误")
					}
			}
		*/
	}
	c.JSON(200, gin.H{
		"status": "received",
		"rcode":  0,
	})
}

func GetConfigs(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "received",
		"rcode":  0,
		"conf":   "all",
	})
}

func AllConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "received",
		"rcode":  0,
		"conf":   "all",
	})
}

func main() {
	//gin日志
	//f, _ := os.Create("/tmp/gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	v1 := r.Group("/api/v1.0/internal")
	{
		//v1.GET("/status", route.Heartbeat)
		v1.GET("/configs", GetConfigs)
		v1.POST("/configs", PostConfig)
		//v1.POST("/all-configs", route.AllConfig)
		//v1.GET("/oplog", route.GetOplog)
		//v1.GET("/tasks", route.GetTask)
		//v1.POST("/tasks", route.PostTask)
		//v1.DELETE("/tasks", route.DeleteTask)
		//v1.GET("/pro", route.GetProduct)
	}

	r.Run(":9999")
}
