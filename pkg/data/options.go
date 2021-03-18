package data

import (
	"fmt"
	"os"
	"strings"

	"github.com/DarthCucumber/gofuzz/pkg/utils"
)

type Options struct {
	FuzzUrl       string
	NumRange      string
	CharList      string
	AsciiRange    string
	InputFilePath string
	MethodData    string
	showHelp      bool
}

//parses the -u flag input
func (o Options) ParseUrl() []string {
	//split url based on <@>
	urlSplitList := strings.Split(o.FuzzUrl, "<@>")
	return urlSplitList
}

//parses the -s flag input
func (o Options) ParseCharList() []string {
	//split based on ,
	chars := strings.Split(o.CharList, ",")
	return chars
}

//parses the -n flag input
func (o Options) ParseNumRange() []string {
	var numList []string
	//split based on ,
	list := strings.Split(o.NumRange, ",")
	listLen := len(list)
	//OPTIMISE:if contains empty string, return
	if listLen == 1 && len(list[0]) == 0 {
		return numList
	}

	switch listLen {
	case 1:
		isChar, n1 := utils.ToInt(list[0])
		if isChar {
			fmt.Printf("[x] %s in -n option contains character\n", list)
			fmt.Printf("[!] user -h for usage guide\n")
			os.Exit(0)
		}
		numList = utils.MakeNumList(0, n1)
	case 2:
		isChar, n1 := utils.ToInt(list[0])
		isChar, n2 := utils.ToInt(list[1])
		if isChar {
			fmt.Printf("[x] %s in -n option contains character\n", list)
			fmt.Printf("[!] user -h for usage guide\n")
			os.Exit(0)
		}
		numList = utils.MakeNumList(n1, n2)
	default:
		numList = list
	}
	return numList
}

//parses the -a flag input
func (o Options) ParseAsciiRange() []string {
	var asciiCharList []string
	//split based on ,
	list := strings.Split(o.AsciiRange, ",")
	listLen := len(list)
	if listLen == 1 && len(list[0]) == 0 {
		return asciiCharList
	}
	switch listLen {
	case 1: // -a 65 => "A"
		isChar, a1 := utils.ToInt(list[0])
		if isChar {
			fmt.Printf("[x] %s in -a option contains character\n", list)
			fmt.Printf("[!] user -h for usage guide\n")
			os.Exit(0)
		}
		asciiCharList = append(asciiCharList, string(rune(a1)))
	case 2: // -a 65,90 => 65...90 => "A"..."Z"
		isChar, a1 := utils.ToInt(list[0])
		isChar, a2 := utils.ToInt(list[1])
		if isChar {
			fmt.Printf("[x] %s in -a option contains character\n", list)
			fmt.Printf("[!] user -h for usage guide\n")
			os.Exit(0)
		}
		asciiCharList = utils.AsciiToChar(a1, a2)
	default: // -a 65,70,90 => "A","F","Z"
		for i := range list {
			isChar, j := utils.ToInt(list[i])
			if isChar {
				fmt.Printf("[x] %s in -a option contains character\n", list)
				fmt.Printf("[!] user -h for usage guide\n")
				os.Exit(0)
			}
			asciiCharList = append(asciiCharList, string(rune(j)))
		}
	}
	return asciiCharList
}

func (o Options) DisplayHelp() {
	if !o.showHelp {
		return
	}
	usage := `
░▄▀▒░▄▀▄▒█▀░█▒█░▀█▀░▀█▀
░▀▄█░▀▄▀░█▀░▀▄█░█▄▄░█▄▄	v1.0.0
`
	fmt.Printf("%s", usage)
	os.Exit(0)
}
