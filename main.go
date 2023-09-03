package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:12345")

	student := Student{
		"name",
		"man",
	}
	fmt.Println(student)

	if err != nil {
		fmt.Print("listen failed, err:", err)
		return
	}
	go Save()
	fmt.Println("go mq started")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Print("accept failed, err:", err)
			continue
		}
		go Process(conn)

	}
}
