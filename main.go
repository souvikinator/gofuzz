package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DarthCucumber/gofuzz/pkg/data"
)

func main() {
	var menu data.Options
	var meta data.FuzzData //TODO: rename variable
	flag.BoolVar(&menu.ShowHelp, "h", false, "shows usage details")
	flag.StringVar(&menu.FuzzUrl, "u", "", "takes in URL for fuzzing")
	flag.StringVar(&menu.NumRange, "n", "", "takes in range of numbers for fuzzing")
	flag.StringVar(&menu.CharList, "c", "", "takes in range of characters for fuzzing")
	flag.StringVar(&menu.AsciiRange, "a", "", "takes in range of ascii values and fuzzes for corresponding charater")
	flag.StringVar(&menu.OutDir, "o", "./output", "set output folder to save the results")
	flag.StringVar(&menu.InputFile, "f", "", "file path to list of fuzz data")
	flag.StringVar(&menu.ExportType, "e", "txt", "data format in which the result will be stored in the output file")
	flag.Parse()

	//detect -h and show help menu
	menu.DisplayHelp()

	if len(menu.FuzzUrl) == 0 {
		fmt.Println("[x] No URL provided for fuzzing")
		fmt.Println("[*] user -h option for usage menu")
		os.Exit(0)
	}

	//check for valid export type(-e)
	menu.SetExportType()

	meta.NumFuzzData = menu.ParseNumRange()
	meta.AsciiFuzzData = menu.ParseAsciiRange()
	meta.CharFuzzData = menu.ParseCharList()
	meta.InputFuzzData = menu.ReadFuzzFile()

	if meta.IsEmpty() {
		fmt.Println("[x] No fuzz data provided for fuzzing")
		os.Exit(0)
	}

	fmt.Printf("%+v\n", meta)
}
