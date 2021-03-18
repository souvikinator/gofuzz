package data

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DarthCucumber/gofuzz/pkg/utils"
)

type Options struct {
	FuzzUrl    string
	NumRange   string
	CharList   string
	AsciiRange string
	MethodData string
	ShowHelp   bool
	OutDir     string
	InputFile  string
	ExportType string
}

//sets output folder
//and create one if doesn't exists
func (o Options) SetOutputDir() {
	//check if exits?
	if !utils.DirExists(o.OutDir) {
		//if not, create one
		err := os.Mkdir(o.OutDir, 0755)
		utils.CheckErr(err, "[x] Error occured while creating output file", o.OutDir)
	}
}

//to check valid export type
func (o Options) SetExportType() {
	switch o.ExportType {
	case "json", "txt", "csv":
		return
	default:
		fmt.Printf("[x] Invalid Export type `%s` in -e option\n", o.ExportType)
		os.Exit(0)
	}
}

//parse input from the list -f
func (o Options) ReadFuzzFile() []string {
	var tmp []string
	if len(o.InputFile) == 0 {
		return tmp
	}
	//check if exists? if not throw error.
	flExist := utils.FileExists(o.InputFile)
	if !flExist {
		fmt.Printf("[x] Input file `%s` either doesn't exist or is a directory. Unable to access!\n", o.InputFile)
		os.Exit(0)
	}
	//open file
	f, err := os.Open(o.InputFile)
	utils.CheckErr(err, "[x] Error occured while reading input file\n", o.InputFile)
	scanner := bufio.NewScanner(f)
	//read file line by line
	for scanner.Scan() {
		line := scanner.Text()
		tmp = append(tmp, line)
	}
	return tmp
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
	list := strings.Split(o.CharList, ",")
	listLen := len(list)
	//for empty charlist
	if listLen == 1 && len(list[0]) == 0 {
		return []string{}
	}
	return list
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
	if !o.ShowHelp {
		return
	}
	usage := `
░▄▀▒░▄▀▄▒█▀░█▒█░▀█▀░▀█▀
░▀▄█░▀▄▀░█▀░▀▄█░█▄▄░█▄▄	v1.0.0

Usage: gofuzz [options...]

Options:

-u	takes in target URL for fuzzing. User placeholder <@>
	Ex: -u http://target.com/q1=<@>&q2=<@>

-n	takes in comma separated number. 
	Ex: -n 12  implies gofuzz will test for numbers from 0 to 12
	    -n 12,100  implies gofuzz will test for numbers from 12 to 100
	    -n 12,13,14,11  implies gofuzz will test for numbers 12,13,14,11

-a  takes in comma sparated ASCII values and extended ASCII values and test for the corresponding
	character of those values.
	Ex: -a 65  implies gofuzz will test for "A"
	    -a 65,90  implies gofuzz will test for "A" to "Z"
	    -a 65,70,66  implies gofuzz will test for "A","F" and "B"

-c	takes in characters as input, mainly used for passing symbols.
	NOTE: try to enclose the string in quotes and use forward slash to escape shell characters
	Ex: -c "\&,@,#" implies gofuzz will test for "&","@","#"
`
	fmt.Printf("%s", usage)
	os.Exit(0)
}
