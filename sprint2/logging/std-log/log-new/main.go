package main

import (
	"log"
	"os"
)

func main() {
	flog, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer flog.Close()
	// log.New(io.Writer, prefix, )
	myLog := log.New(flog, "server ", log.LstdFlags|log.Lshortfile)
	myLog.Println("Start server")
	// SetFlags(flag int) Sets new flags to Logger
	myLog.SetFlags(log.Llongfile)
	myLog.Println("Finish server")
}
