package main

import (
	"container/list"
	"sync"
)

type Queue struct {
	len  int
	data list.List
}

var lock sync.Mutex

func (queue *Queue) offer(msg Msg) {
	queue.data.PushBack(msg)
	queue.len = queue.data.Len()
}

func (queue *Queue) poll() Msg {
	if queue.len == 0 {
		return Msg{}
	}
	msg := queue.data.Front()
	return msg.Value.(Msg)
}

func (queue *Queue) delete(id int64) {
	lock.Lock()
	for msg := queue.data.Front(); msg != nil; msg = msg.Next() {
		if msg.Value.(Msg).Id == id {
			queue.data.Remove(msg)
			queue.len = queue.data.Len()
			break
		}
	}
	lock.Unlock()
}
