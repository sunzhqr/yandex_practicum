package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Declaring flag sets for subcommands
	cnvFlags := flag.NewFlagSet("cnv", flag.ExitOnError)
	filterFlags := flag.NewFlagSet("filter", flag.ExitOnError)
	// Flags of the cnvFlags set
	destDir := cnvFlags.String("dest", "./output", "destination folder")
	width := cnvFlags.Int("w", 1024, "width of the image")
	isThumb := cnvFlags.Bool("thumb", false, "create thumb")
	_ = destDir
	_ = width
	_ = isThumb

	// Flags of the filterFlags set
	isGray := filterFlags.Bool("gray", false, "convert to grayscale")
	isSepia := filterFlags.Bool("sepia", false, "convert to sepia")

	if len(os.Args) < 2 {
		fmt.Println("set or get subcommand required")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "cnv":
		cnvFlags.Parse(os.Args[2:])
	case "filter":
		filterFlags.Parse(os.Args[2:])
	default:
		cnvFlags.PrintDefaults()
		filterFlags.PrintDefaults()
		os.Exit(1)
	}

	// Функция FlagSet.Parsed() возвращает false, если
	// парсинг флагов набора не проводился
	if cnvFlags.Parsed() {
		fmt.Println("cnv command is using...")
	}
	if filterFlags.Parsed() {
		if *isGray {
			fmt.Println("Converted to gray")
		} else if *isSepia {
			fmt.Println("Converted to sepia")
		} else {
			fmt.Println("Default color, nothing changed")
		}
	}
}
