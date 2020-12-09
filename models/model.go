package models

import (
	//"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"pcrm/config"
	"strconv"
)

var DB *gorm.DB

func init() {
	sql := config.Conf.Sql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", sql.User, sql.Pass, sql.Ip, strconv.Itoa(sql.Port), sql.Database)
	//因为DB为全局变量此处必须显式声明error
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	DB.AutoMigrate(&Content{}, &Register{})
	/*
		n := make(map[string]interface{})
		n["1"] = 111
		n["2"] = "222"
		n["responsecode"] = 100
		b, err := json.Marshal(n)
		if err != nil {
			fmt.Println("json.Marshal failed:", err)
			return
		}
		DB.Create(&Content{Mid: 3, Bt: "selfcheck", Sbt: "responserules", Service: "handle", Source: "ms", Op: "add", Data: b})
	*/
}
