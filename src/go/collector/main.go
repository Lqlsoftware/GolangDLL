package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"log"
)

func main() {
	// Print message
	fmt.Println("#########################################");
	fmt.Println("###          Golang Collector         ###");
	fmt.Println("#########################################");

	// Create new context
	context, _ := zmq.NewContext()

	// New socket for receiving (use tcp to fix on windows platform).
	receiver, _ := context.NewSocket(zmq.PULL)
	err := receiver.Connect("tcp://127.0.0.1:5001")
	if err != nil {
		panic(err)
	}
	defer receiver.Close()

	// Receive and process messages
	for {
		received, err := receiver.Recv(0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[COLLECTOR] Receive a message from queue: \"%s\"\n", string(received))

		/*
			do something...
		 */
	}
}