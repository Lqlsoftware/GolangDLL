package main

import (
	"fmt"
	"github.com/sheerun/queue"
	zmq "github.com/pebbe/zmq4"
	"log"
)

func main() {
	// Print message
	fmt.Println("#########################################");
	fmt.Println("###            Golang Queue           ###");
	fmt.Println("#########################################");

	// New Queue (thread-safe)
	q := queue.New()

	// zmq
	var err error
	context, err := zmq.NewContext()
	if err != nil {
		panic(err)
	}

	// Starting receive
	go Receive(context, q)

	// Deq from queue and send to collector
	go Deq(context, q)

	// Block main thread
	var b chan int
	<-b
}

func Receive(context *zmq.Context, q *queue.Queue) {
	var err error
	// Socket to receive
	receiver, err := context.NewSocket(zmq.PULL)
	if err != nil {
		panic(err)
	}
	defer receiver.Close()
	err = receiver.Bind("tcp://127.0.0.1:5000")
	if err != nil {
		panic(err)
	}

	for {
		received, err := receiver.Recv(0)
		if err != nil {
			log.Fatal(err)
		}
		q.Append(received)
		fmt.Printf("[QUEUE] Receive a enqueue message: \"%s\"\n", received)
	}
}

func Deq(context *zmq.Context, q *queue.Queue) {
	var err error
	// Socket to send
	sender, err := context.NewSocket(zmq.PUSH)
	if err != nil {
		panic(err)
	}
	defer sender.Close()
	err = sender.Bind("tcp://127.0.0.1:5001")
	if err != nil {
		panic(err)
	}

	for {
		<-q.NotEmpty

		deq := q.Pop().(string)
		_, err := sender.Send(deq, 0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[QUEUE] Dequeue and send to collector: \"%s\"\n", deq)
	}
}