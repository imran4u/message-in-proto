package main

import (
	"fmt"
	"log"
	"my-protobuf/basic"
	"time"
)

type logWriter struct{}

func (logWriter) Write(p []byte) (n int, err error) {
	// return fmt.Print(time.Now().Format("15:04:05.000 ") + string(p))
	return fmt.Println(time.Now().String() + string(p))
}

func main() {
	log.SetFlags(0)
	logWriter := logWriter{}
	b := []byte("hello")
	log.SetOutput(logWriter)
	logWriter.Write(b)
	basic.BasicHello()
}
