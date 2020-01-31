package main

import (
	"fmt"
	"github.com/sheerun/queue"
)

func main() {
	// Print message
	fmt.Println("#########################################");
	fmt.Println("###            Golang Queue           ###");
	fmt.Println("#########################################");

	// New Queue (thread-safe)
	q := queue.New()

	// Starting receive
	go Receive(q)

	// Deq from queue and send to collector
	go Deq(q)

	// Block main thread
	var b chan int
	<-b
}

func Receive(q *queue.Queue) {
	pipeReader := NewReader("dll_queue_pipe.ipc")
	for {
		content := string(pipeReader.Read())
		q.Append(content)
		fmt.Printf("[QUEUE] Receive a enqueue message: \"%s\"\n", content)
	}
}

func Deq(q *queue.Queue) {
	pipeWriter := NewWriter("queue_collector_pipe.ipc")
	for {
		<-q.NotEmpty

		deq := q.Pop().(string)
		pipeWriter.Write([]byte(deq))
		fmt.Printf("[QUEUE] Dequeue and send to collector: \"%s\"\n", deq)
	}
}