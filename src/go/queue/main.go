package main

import (
	"fmt"
	zmq "github.com/alecthomas/gozmq"
)

func main() {
	// Print message
	fmt.Println("#########################################");
	fmt.Println("###            Golang Queue           ###");
	fmt.Println("#########################################");

	var err error
	context, err := zmq.NewContext()
	if err != nil {
		panic(err)
	}
	defer context.Close()

	// Socket facing libinterop
	frontend, err := context.NewSocket(zmq.PULL)
	if err != nil {
		panic(err)
	}
	defer frontend.Close()
	err = frontend.Bind("tcp://*:5000")
	if err != nil {
		panic(err)
	}

	// Socket facing collector
	backend, err := context.NewSocket(zmq.PUSH)
	if err != nil {
		panic(err)
	}
	defer backend.Close()
	err = backend.Bind("tcp://*:5001")
	if err != nil {
		panic(err)
	}

	// Start built-in device
	err = zmq.Device(zmq.QUEUE, frontend, backend)
	if err != nil {
		panic(err)
	}
}