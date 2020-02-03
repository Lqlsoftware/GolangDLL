package main

import (
	"fmt"
	zmq "github.com/alecthomas/gozmq"
)

func main() {
	// Print message
	fmt.Println("#########################################");
	fmt.Println("###          Golang Collector         ###");
	fmt.Println("#########################################");

	// Create new context
	context, _ := zmq.NewContext()
	defer context.Close()

	// New socket for receiving (use tcp to fix on windows platform).
	receiver, _ := context.NewSocket(zmq.PULL)
	receiver.Connect("tcp://127.0.0.1:5001")
	defer receiver.Close()

	// Receive and process messages
	for {
		received, _ := receiver.Recv(0)
		fmt.Printf("[COLLECTOR] Receive a message from queue: \"%s\"\n", string(received))

		/*
			do something...
		 */
	}
}