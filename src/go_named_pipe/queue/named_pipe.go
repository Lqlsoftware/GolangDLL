package main

import (
	"bufio"
	"log"
	"os"
	"syscall"
)

type NamedPipe struct {
	name 	string
	file 	*os.File
}

// implement NamedPipe
type PipeWriter struct {
	name 	string
	file 	*os.File
	opened	bool
}

func NewWriter(filename string) *PipeWriter {
	var err error
	// Delete previous pipe file
	os.Remove(filename)

	// Create fifo named pipe
	err = syscall.Mkfifo(filename, 0666)
	if err != nil {
		log.Fatal("[Error] create named pipe:", err)
	}

	// Open named pipe.
	pipe, err := os.OpenFile(filename, os.O_RDWR, 0777)
	if err != nil {
		log.Fatalf("[Error] opening pipe file: %v", err)
	}

	return &PipeWriter{
		name:  	filename,
		file:   pipe,
		opened: true,
	}
}

func (p *PipeWriter) Write(content []byte) {
	// write to pip
	_, err := p.file.Write(append(content, '\n'))
	if err != nil {
		log.Fatalf("[Error] write to pipe file: %v", err)
	}
}

func (p *PipeWriter) Close() {
	err := p.file.Close()
	if err != nil {
		log.Fatalf("[Error] close pipe file: %v", err)
	}
}


// implement NamedPipe
type PipeReader struct {
	name 	string
	file 	*os.File
	reader	*bufio.Reader
	opened	bool
}

func NewReader(filename string) *PipeReader {
	var err error
	// Open named pipe.
	pipe, err := os.OpenFile(filename, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		log.Fatalf("[Error] opening pipe file: %v", err)
	}

	return &PipeReader{
		name:  	filename,
		file:   pipe,
		reader: bufio.NewReader(pipe),
		opened: true,
	}
}

func (p *PipeReader) Read() []byte {
	content, err := p.reader.ReadBytes('\n')
	// trim '\n'
	content = content[:len(content) - 1]
	if err != nil {
		log.Fatalf("[Error] read from pipe file: %v", err)
	}
	return content
}

func (p *PipeReader) Close() {
	err := p.file.Close()
	if err != nil {
		log.Fatalf("[Error] close pipe file: %v", err)
	}
}