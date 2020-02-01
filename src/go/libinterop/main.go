package main

import (
	"C"
	"fmt"
	zmq "github.com/alecthomas/gozmq"
)

func main() {}

var context *zmq.Context
var socket 	*zmq.Socket

//export Init
func Init() {
	var err error
	context, err	= zmq.NewContext()
	if err != nil {
		panic(err)
	}
	socket, err 	= context.NewSocket(zmq.PUSH)
	if err != nil {
		panic(err)
	}
	socket.Connect("tcp://127.0.0.1.5000")
}

//export Enq
func Enq(parameter string) {
	err := socket.Send([]byte(parameter), 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[C] Send a message: \"%s\"\n", parameter)
}