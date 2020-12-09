//package config
package main

import (
	"encoding/json"
	"fmt"
	"os"
	//"io/ioutil"
)

type Listen struct {
	Ip   string
	Port int
}

type Mysql struct {
	Listen
	User     string
	Pass     string
	Database string
}

type Bind struct {
	Listen
	Conf string
}

type Ems struct {
	Ip string
}

type Out struct {
	Conf int
	Task int
}

type Config struct {
	Net     Listen
	Db      Mysql
	Ybind   Bind
	Crm     Listen
	Ms      Ems
	Timeout Out
}

//var Conf *Config

//func init() {
//}

func readFile() {

	filePtr, err := os.Open("test.json")
	//filePtr, err := os.Open("conf.json")
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()

	var person Config

	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println("Decoder failed", err.Error())

	} else {
		fmt.Println("Decoder success")
		fmt.Println(person)
	}
}

func write() {
	/*
		bytes, err := ioutil.ReadFile("./conf.json")
		if err != nil {
			fmt.Println("读取json文件失败", err)
			return
		}
		fmt.Printf("%+v\n", bytes)
	*/
	//Conf = new(Config)
	l := Listen{Ip: ":", Port: 9999}
	m := Mysql{Listen: l, User: "root", Pass: "123456", Database: "test"}
	fmt.Printf("%+v\n", m)
	//y := Bind{Listen: l, Conf: "/etc/named.conf"}
	y := Bind{l, "/etc/named.conf"}
	fmt.Printf("%+v\n", y)
	ms := Ems{Ip: "127.0.0.1"}
	fmt.Printf("%+v\n", ms)
	t := Out{Conf: 5, Task: 3600}
	fmt.Printf("%+v\n", t)
	c := Config{l, m, y, l, ms, t}
	fmt.Printf("%+v\n", c)

	filePtr, err := os.Create("test.json")
	if err != nil {
		fmt.Println("Create file failed", err.Error())
		return
	}
	defer filePtr.Close()

	/*
		encoder := json.NewEncoder(filePtr)

		err = encoder.Encode(c)
		if err != nil {
			fmt.Println("Encoder failed", err.Error())

		} else {
			fmt.Println("Encoder success")
		}
	*/

	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		fmt.Println("Encoder failed", err.Error())

	} else {
		fmt.Println("Encoder success")
	}

	filePtr.Write(data)

	//Conf := Config{Net: l, Db: {Listen: l, User: "root", Pass: "123456", Database: "test"}, Ybind: {Listen: l, Conf: "/etc/named.conf"}, Crm: l, Ms: {Ip: "127.0.0.1"}, Timeout: {Conf: 5, Task: 3600}}
	//fmt.Printf("%+v\n", Conf)
	/*
		fmt.Printf("%+v\n", Conf)
		err = json.Unmarshal(bytes, Conf)
		if err != nil {
			fmt.Println("解析数据失败", err)
			return
		}
		fmt.Printf("%+v\n", Conf)
	*/
}

func main() {
	readFile()
	//write()
}
