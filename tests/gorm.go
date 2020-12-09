package main

import (
	"fmt"
	//"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
)

type Product struct {
	gorm.Model
	Code  string `gorm:"unique"`
	Price uint
}

type Switch struct {
	gorm.Model
	Id     int
	Bt     string
	Sbt    string
	Switch string
}

type Tabler interface {
	TableName() string
}

// TableName 会将 User 的表名重写为 `profiles`
func (Switch) TableName() string {
	return "switch"
}

func main() {
	dsn := "root:123456@tcp(192.168.5.41:3306)/test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	//db.AutoMigrate(&Product{})
	//db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product, "code = ?", "D42")
	fmt.Println(product)

	//var s Switch
	s := make([]Switch, 0, 100)
	//result := db.First(&s)
	result := db.Find(&s)
	fmt.Println(result, reflect.TypeOf(result), result.RowsAffected, result.Error)
	for _, v := range s {
		fmt.Println(v)
	}
}
