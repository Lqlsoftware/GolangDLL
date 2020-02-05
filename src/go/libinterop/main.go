package main

import (
	"C"
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"log"
)

func main() {}

var context *zmq.Context
var socket 	*zmq.Socket

//export Init
func Init() {
	// Method to Init ZeroMQ
	var err error
	// New context
	context, err	= zmq.NewContext()
	if err != nil {
		panic(err)
	}
	// New socket for sending
	socket, err 	= context.NewSocket(zmq.PUSH)
	if err != nil {
		panic(err)
	}
	// Connect localhost
	err = socket.Connect("tcp://127.0.0.1:5000")
	if err != nil {
		panic(err)
	}
}

//export Send
func Send(parameter string) {
	// Method to send message to queue
	var err error
	// Send message
	_, err = socket.Send(parameter, 0)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("[C] Send a message: \"%s\"\n", parameter)
	}
}