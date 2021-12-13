package main

import (
	"fmt"
	"net"
)

func main() {
	var conntype string
	fmt.Printf("TCP or UDP port scan? : ")
	fmt.Scan(&conntype)

	var ipaddr string
	fmt.Printf("Ip address to scan? ip:port : ")
	fmt.Scan(&ipaddr)

	_, err := net.Dial(conntype, ipaddr)

	if err == nil {
		fmt.Println("connection successfull")
	}
}
