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
	frontend.Bind("ipc://queue.ipc")

	// Socket facing services
	backend, _ := context.NewSocket(zmq.PUSH)
	defer backend.Close()
	backend.Bind("ipc://collector.ipc")

	// Start built-in device
	zmq.Device(zmq.QUEUE, frontend, backend)
}