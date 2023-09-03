package main

import (
	"arterning/go-mq/common"
	"bytes"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Print("connect failed, err:", err)
	}
	defer conn.Close()

	msg := common.Msg{Topic: "topic-test", MsgType: 1}

	n, err := conn.Write(common.MsgToBytes(msg))
	if err != nil {
		fmt.Println("write failed, err:", err)
	}
	fmt.Println("n", n)

	var res [128]byte
	conn.Read(res[:])
	buf := bytes.NewBuffer(res[:])
	receMsg := common.BytesToMsg(buf)
	fmt.Print(receMsg)

	// ack
	conn, _ = net.Dial("tcp", "127.0.0.1:12345")
	l, e := conn.Write(common.MsgToBytes(common.Msg{Id: receMsg.Id, Topic: receMsg.Topic, MsgType: 3}))
	if e != nil {
		fmt.Println("write failed, err:", err)
	}
	fmt.Println("l:", l)
}
