package main

import (
	"arterning/go-mq/common"
	"container/list"
	"sync"
)

type Queue struct {
	len  int
	data list.List
}

var lock sync.Mutex

func (queue *Queue) offer(msg common.Msg) {
	queue.data.PushBack(msg)
	queue.len = queue.data.Len()
}

func (queue *Queue) poll() common.Msg {
	if queue.len == 0 {
		return common.Msg{}
	}
	msg := queue.data.Front()
	return msg.Value.(common.Msg)
}

func (queue *Queue) delete(id int64) {
	lock.Lock()
	for msg := queue.data.Front(); msg != nil; msg = msg.Next() {
		if msg.Value.(common.Msg).Id == id {
			queue.data.Remove(msg)
			queue.len = queue.data.Len()
			break
		}
	}
	lock.Unlock()
}
