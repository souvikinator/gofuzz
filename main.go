package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DarthCucumber/gofuzz/pkg/data"
)

func main() {
	var menu data.Options
	var tmpStorage data.FuzzData //TODO: rename variable
	flag.StringVar(&menu.FuzzUrl, "u", "", "takes in URL for fuzzing")
	flag.StringVar(&menu.NumRange, "n", "", "takes in range of numbers for fuzzing")
	flag.StringVar(&menu.CharList, "s", "", "takes in range of characters for fuzzing")
	flag.StringVar(&menu.AsciiRange, "a", "", "takes in range of ascii values and fuzzes for corresponding charater")
	flag.BoolVar(&menu.Usage, "h", false, "shows usage details")
	flag.Parse()

	// if len(menu.FuzzUrl) == 0 {
	// 	fmt.Println("[x] No URL provided for fuzzing")
	// 	fmt.Println("[*] user -h option for usage menu")
	// 	os.Exit(0)
	// }

	numlist := menu.ParseNumRange()
	tmpStorage.FuzzList = append(tmpStorage.FuzzList, numlist...)
	asciilist := menu.ParseAsciiRange()
	tmpStorage.FuzzList = append(tmpStorage.FuzzList, asciilist...)
	charlist := menu.ParseCharList()
	tmpStorage.FuzzList = append(tmpStorage.FuzzList, charlist...)

	//TODO:add file input feature

	if len(tmpStorage.FuzzList) == 0 {
		fmt.Println("[x] no options/inputs provided for fuzzing pattern")
		fmt.Println("[!] use -h for usage guide")
		os.Exit(0)
	}
}
