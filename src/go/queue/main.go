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

	context, _ := zmq.NewContext()
	defer context.Close()

	// Socket facing clients
	frontend, _ := context.NewSocket(zmq.PULL)
	defer frontend.Close()
	frontend.Bind("tcp://127.0.0.1:5000")

	// Socket facing services
	backend, _ := context.NewSocket(zmq.PUSH)
	defer backend.Close()
	backend.Bind("tcp://127.0.0.1:5001")

	// Start built-in device
	zmq.Device(zmq.QUEUE, frontend, backend)
}