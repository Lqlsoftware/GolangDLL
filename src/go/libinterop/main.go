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
	context, _	= zmq.NewContext()
	socket, _ 	= context.NewSocket(zmq.PUSH)
	socket.Connect("ipc://queue.ipc")
}

//export Enq
func Enq(parameter string) {
	socket.Send([]byte(parameter), 0)
	fmt.Printf("[C] Send a message: \"%s\"\n", parameter)
}