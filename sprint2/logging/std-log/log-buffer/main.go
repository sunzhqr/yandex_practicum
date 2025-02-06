package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var buf bytes.Buffer
	// 0 means no flag
	bufLog := log.New(&buf, `Buffer:	`, 0)
	bufLog.Println("Hello, from Buffer!")
	bufLog.Println("Buffer executed successfully!")
	// if we pass just buf it prints bytes of line
	fmt.Println(&buf)
}
