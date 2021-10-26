package data

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/DarthCucumber/gofuzz/pkg/utils"
)

//TODO: add timeout option
type Options struct {
	TargetUrl  string
	NumRange   string
	CharList   string
	AsciiRange string
	Method     string
	ShowHelp   bool
	OutputDir  string
	InputFile  string
	ExportType string
	Host       string
	Exclude    string
	Timeout    int
	Retries    int
}

//exclude status code functionality
func (o Options) ExcludeStatusCode() []string {
	//split based on , (comma)
	list := strings.Split(o.Exclude, ",")
	listLen := len(list)
	//if contains empty string, return
	if listLen == 1 && len(list[0]) == 0 {
		return list
	}
	//sort statuscode list for easy lookup
	sort.Strings(list)
	return list
}

//check for valid request method
func (o Options) SetRequestMethod() string {
	switch o.Method {
	case "HEAD", "GET":
		return o.Method
	case "POST":
		utils.ShowInfo("Sorry! POST method is currently under development")
		os.Exit(0)
	default:
		utils.ShowError("Invalid request method in -m")
		utils.ShowInfo("Only HEAD,GET,POST allowed")
		os.Exit(0)
	}
	return "HEAD"
}

//to check valid export type
func (o Options) SetExportType() string {
	switch o.ExportType {
	//TODO: add check for CSV
	case "txt", "json":
		return o.ExportType
	default:
		utils.ShowError("Invalid Export type `", o.ExportType, "` in -e option")
		os.Exit(0)
	}
	return "txt"
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
		utils.ShowError("Input file `%s` either doesn't exist or is a directory. Unable to access!\n", o.InputFile)
		os.Exit(0)
	}
	//open file
	f, err := os.Open(o.InputFile)
	utils.CheckErr(err, "Error occurred while reading input file\n", o.InputFile)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	//read file line by line
	for scanner.Scan() {
		line := scanner.Text()
		tmp = append(tmp, line)
	}
	return tmp
}

//parses the -u flag input
func (o *Options) ParseUrl() []string {
	//split url based on <@>
	urlSplitList := strings.Split(o.TargetUrl, "<@>")
	//get host from url
	u, err := url.Parse(strings.Join(urlSplitList, ""))
	utils.CheckErr(err, err)
	host := fmt.Sprintf("%s://%s", u.Scheme, u.Host)
	//check if host is reachable
	_, err = http.Head(host)
	utils.CheckErr(err, err)
	//store host and use it to name output folder
	o.Host = u.Host
	return urlSplitList
}

//sets output folder
//and create one if doesn't exists
func (o Options) SetOutputDir() string {
	//check if exits?
	//output dir: ./output/<target_host>/
	out := filepath.Join(o.OutputDir, o.Host)
	if !utils.DirExists(out) {
		//if not, create one
		err := os.MkdirAll(out, os.ModePerm)
		utils.CheckErr(err, "Error occurred while creating output file", out, err)
	}
	// utils.ShowSuccess("Output Folder: ", out)
	return out
}

//parses the -c flag input
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
	//split based on , (comma)
	list := strings.Split(o.NumRange, ",")
	listLen := len(list)
	//if contains empty string, return
	if listLen == 1 && len(list[0]) == 0 {
		return numList
	}

	switch listLen {
	case 1:
		isChar, n1 := utils.ToInt(list[0])
		if isChar {
			utils.ShowError("fuzzing data provided in -n option contains character")
			utils.ShowWarning("user -h for usage guide")
			os.Exit(0)
		}
		numList = utils.MakeNumList(0, n1)
	case 2:
		isChar, n1 := utils.ToInt(list[0])
		isChar, n2 := utils.ToInt(list[1])
		if isChar {
			utils.ShowError("fuzzing data provided in -n option contains character")
			utils.ShowWarning("user -h for usage guide")
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
			utils.ShowError("fuzzing data provided in -a option contains character")
			utils.ShowWarning("user -h for usage guide")
			os.Exit(0)
		}
		asciiCharList = append(asciiCharList, string(rune(a1)))
	case 2: // -a 65,90 => 65...90 => "A"..."Z"
		isChar, a1 := utils.ToInt(list[0])
		isChar, a2 := utils.ToInt(list[1])
		if isChar {
			utils.ShowError("fuzzing data provided in -a option contains character")
			utils.ShowWarning("user -h for usage guide")
			os.Exit(0)
		}
		asciiCharList = utils.AsciiToChar(a1, a2)
	default: // -a 65,70,90 => "A","F","Z"
		for i := range list {
			isChar, j := utils.ToInt(list[i])
			if isChar {
				utils.ShowError("fuzzing data provided in -a option contains character")
				utils.ShowWarning("user -h for usage guide")
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

-h

display help/usage menu

-u

takes in target URL for fuzzing. User placeholder <@>
Ex: -u "http://target.tld/q1=<@>&q2=<@>"
NOTE: Try to enclose URL in quotes as the placeholder may cause issue in terminal

-n

takes in comma separated number.
Ex: -n 12  implies gofuzz will test for numbers from 0 to 12
-n 12,100  implies gofuzz will test for numbers from 12 to 100
-n 12,13,14,11  implies gofuzz will test for numbers 12,13,14,11

-a

takes in comma separated ASCII values and extended ASCII values and test for the corresponding character of those values.
Ex: -a 65  implies gofuzz will test for "A"
-a 65,90  implies gofuzz will test for "A" to "Z"
-a 65,70,66  implies gofuzz will test for "A","F" and "B"

-c

takes in characters as input, mainly used for passing symbols.
NOTE: try to enclose the string in quotes or use forward slash to escape shell characters
Ex: -c "\&,@,#" implies gofuzz will test for "&","@","#"

-m

takes in GET/POST/HEAD request methods as input (default: HEAD)
NOTE: POST doesn't work for now

-t

Time to wait for each request. Takes in time in milliseconds (default: 30000 ms or 30 s)

-r

How many times attempts a request can have in case of error (default: 3)

-o

Output directory where results will be stored. (Default: ./output)

-export

takes txt/json export type as input (default and preferred: json)

-exclude

takes in response status code separated by commas(,) and excludes them from the results. (blacklisting status codes)
Ex: -exclude 404,500 : implies any result corresponding to these status code will not be included in results.

`
	fmt.Printf("%s", usage)
	os.Exit(0)
}
