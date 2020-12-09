package main

import (
	"fmt"
	"net"
)

func main() {
	ip := "1.1.1.0/24"
	address := net.ParseIP(ip)
	if address == nil {
		// 没有匹配上
		fmt.Println("NOT IP:", ip)
	} else {
		// 匹配上
		fmt.Println("IS IP:", address)
	}
}
