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

	context, _ := zmq.NewContext()
	defer context.Close()

	receiver, _ := context.NewSocket(zmq.PULL)
	receiver.Connect("ipc://collector.ipc")
	defer receiver.Close()

	for {
		received, _ := receiver.Recv(0)
		fmt.Printf("[COLLECTOR] Receive a message from queue: \"%s\"\n", string(received))
	}
}