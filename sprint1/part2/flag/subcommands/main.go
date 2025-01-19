package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Options struct {
	// width   int
	// thumb   bool
	effects []string
}

// String должен уметь сериализовать переменную типа в строку.
func (o *Options) String() string {
	return fmt.Sprint(strings.Join(o.effects, ","))
}

// Set связывает переменную типа со значением флага
// и устанавливает правила парсинга для пользовательского типа.
func (o *Options) Set(flagValue string) error {
	o.effects = strings.Split(flagValue, ",")
	return nil
}

func main() {
	// Declaring flag sets for subcommands
	cnvFlags := flag.NewFlagSet("cnv", flag.ExitOnError)
	filterFlags := flag.NewFlagSet("filter", flag.ExitOnError)
	options := new(Options) // &Options{}

	// Flags of the cnvFlags set
	destDir := cnvFlags.String("dest", "./output", "destination folder")
	width := cnvFlags.Int("w", 1024, "width of the image")
	isThumb := cnvFlags.Bool("thumb", false, "create thumb")

	// Flags of the filterFlags set
	isGray := filterFlags.Bool("gray", false, "convert to grayscale")
	isSepia := filterFlags.Bool("sepia", false, "convert to sepia")
	var effects []string

	if len(os.Args) < 2 {
		fmt.Println("set or get subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "cnv":
		fmt.Println("cnv subcommand")
		cnvFlags.Var(options, "effects", "Rotation and mirroring")
		cnvFlags.Parse(os.Args[2:])
	case "filter":
		fmt.Println("filter subcommand")
		filterFlags.Func("effects", "Rotation and mirroring", func(effectValue string) error {
			effects = strings.Split(effectValue, ",")
			return nil
			// return errors.New("not implemented")
		})
		filterFlags.Parse(os.Args[2:])
	default:
		fmt.Println("cnv usage:")
		cnvFlags.PrintDefaults()
		fmt.Println("filter usage:")
		filterFlags.PrintDefaults()
		os.Exit(1)
	}
	// Функция FlagSet.Parsed() возвращает false, если
	// парсинг флагов набора не проводился
	if cnvFlags.Parsed() {
		fmt.Println("destDir:", *destDir)
		fmt.Println("width:", *width)
		fmt.Println("isThumb:", *isThumb)
		fmt.Println("effects:", options)
	}
	if filterFlags.Parsed() {
		fmt.Println("isGray:", *isGray)
		fmt.Println("isSepia:", *isSepia)
		fmt.Println("effects:", effects)
	}
}
