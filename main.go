package main

import (
	"awesomeProject/broker"
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Print("listen failed, err:", err)
		return
	}
	go broker.Save()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Print("accept failed, err:", err)
			continue
		}
		go broker.Process(conn)

	}
}
