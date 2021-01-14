package main

import (
	"fmt"
	"net"
)

func main() {
	ip := "1.1.1.257"
	address := net.ParseIP(ip)
	if address == nil {
		// 没有匹配上
		fmt.Println("error:", ip)
	} else {
		// 匹配上
		fmt.Println(address)
	}
}
