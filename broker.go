package main

import (
	"bufio"
	"net"
	"os"
	"sync"
	"time"
)

var topics = sync.Map{}

func handleErr(conn net.Conn) {
	if err := recover(); err != nil {
		println(err.(string))
		conn.Write(MsgToBytes(Msg{MsgType: 4}))
	}
}

func Process(conn net.Conn) {
	defer handleErr(conn)
	reader := bufio.NewReader(conn)
	msg := BytesToMsg(reader)
	queue, ok := topics.Load(msg.Topic)
	var res Msg
	if msg.MsgType == 1 {
		// comsumer
		if queue == nil || queue.(*Queue).len == 0 {
			return
		}
		msg = queue.(*Queue).poll()
		msg.MsgType = 1
		res = msg
	} else if msg.MsgType == 2 {
		// producer
		if !ok {
			queue = &Queue{}
			queue.(*Queue).data.Init()
			topics.Store(msg.Topic, queue)
		}
		queue.(*Queue).offer(msg)
		res = Msg{Id: msg.Id, MsgType: 2}
	} else if msg.MsgType == 3 {
		// consumer ack
		if queue == nil {
			return
		}
		queue.(*Queue).delete(msg.Id)

	}
	conn.Write(MsgToBytes(res))

}

func Save() {
	ticker := time.NewTicker(60)
	for {
		select {
		case <-ticker.C:
			topics.Range(func(key, value interface{}) bool {
				if value == nil {
					return false
				}
				file, _ := os.Open(key.(string))
				if file == nil {
					file, _ = os.Create(key.(string))
				}
				for msg := value.(*Queue).data.Front(); msg != nil; msg = msg.Next() {
					file.Write(MsgToBytes(msg.Value.(Msg)))
				}
				file.Close()
				return false
			})
		default:
			time.Sleep(1)
		}
	}
}
