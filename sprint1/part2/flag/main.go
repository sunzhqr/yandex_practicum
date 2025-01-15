package main

import (
	"flag"
	"fmt"
)

var options struct {
	width int
	thumb bool
}
var (
	destDir = flag.String("dest", "./output", "destination folder")
)

func init() {
	flag.IntVar(&options.width, "w", 1024, "width of the image")
	flag.BoolVar(&options.thumb, "thumb", false, "create thumb")
}

func main() {
	somePtr := flag.String("first", "wtf", "watatata")
	flag.StringVar(somePtr, "second", "life", "kfdls;kfs;ld")
	flag.Parse()
	fmt.Println("Destination Folder:", *destDir)
	fmt.Println("Width", options.width)
	fmt.Println("Thumbs:", options.thumb)
	fmt.Println("SomPtr", *somePtr)
	for i, v := range flag.Args() {
		fmt.Printf("Image file (%d): %s\r\n", i, v)
	}
}
