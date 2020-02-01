package main

import (
	"fmt"
)

func main() {
	// Print message
	fmt.Println("#########################################");
	fmt.Println("###          Golang Collector         ###");
	fmt.Println("#########################################");

	pipeReader := NewReader("queue_collector_pipe.ipc")
	for {
		content := string(pipeReader.Read())
		fmt.Printf("[COLLECTOR] Receive a message from queue: \"%s\"\n", content)
	}
}