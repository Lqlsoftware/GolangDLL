package main

import (
	"C"
	"fmt"
	"github.com/zeromq/goczmq"
)

func main() {}

var pipeWriter *PipeWriter

//export Init
func Init() {
	pipeWriter = NewWriter("dll_queue_pipe.ipc")
}

//export Enq
func Enq(parameter string) {
	pipeWriter.Write([]byte(parameter))
	fmt.Printf("[C] Send a message: \"%s\"\n", parameter)
}