package main

type Msg struct {
	Id       int64
	TopicLen int64
	Topic    string
	// 1-consumer 2-producer 3-comsumer-ack 4-error
	MsgType int64  // 消息类型
	Len     int64  // 消息长度
	Payload []byte // 消息
}
