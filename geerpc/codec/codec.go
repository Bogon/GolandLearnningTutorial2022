package codec

import "io"

// 一个典型的 RPC 调用
// 远程过程调用
// err = client.Call("Arith.Multiply", args, &reply)

// Header 定义消息头
type Header struct {
	ServiceMethod string // format "Service.Method"
	Seq           uint64 // sequence number chosen by client
	Error         string
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}
