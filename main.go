package main

import (
	"flag"
	"os"

	"github.com/DarthCucumber/gofuzz/pkg/data"
	"github.com/DarthCucumber/gofuzz/pkg/utils"
)

func main() {
	var options data.Options
	var session data.SessionData

	var parsedNum data.FuzzData
	var parsedAscii data.FuzzData
	var parsedChar data.FuzzData
	var parsedInput data.FuzzData

	flag.BoolVar(&options.ShowHelp, "h", false, "shows usage details")
	flag.StringVar(&options.TargetUrl, "u", "", "takes in URL for fuzzing")
	flag.StringVar(&options.NumRange, "n", "", "takes in range of numbers for fuzzing")
	flag.StringVar(&options.CharList, "c", "", "takes in range of characters for fuzzing")
	flag.StringVar(&options.AsciiRange, "a", "", "takes in range of ascii values and fuzzes for corresponding charater")
	flag.StringVar(&options.OutputDir, "o", "./output", "set output folder to save the results")
	flag.StringVar(&options.InputFile, "f", "", "file path to list of fuzz data")
	flag.StringVar(&options.ExportType, "e", "txt", "data format in which the result will be stored in the output file")
	flag.StringVar(&options.Method, "m", "HEAD", "Request method [HEAD/GET/POST]")
	flag.IntVar(&options.Timeout, "t", 60000, "takes in timout for each requests in milliseconds. (Default: 2000 ms or 2 s)")
	flag.StringVar(&options.Exclude, "ex", "", "takes in status code separated by commas to be excluded from display result, however everything is included in the result files")
	flag.Parse()

	//detect -h and show help options
	options.DisplayHelp()

	if len(options.TargetUrl) == 0 {
		utils.ShowError("No URL provided for fuzzing")
		utils.ShowWarning("use -h option for usage options")
		os.Exit(0)
	}

	//parse target url
	session.ParsedUrl = options.ParseUrl()
	//set timeout
	session.Timeout = options.Timeout
	//check for valid export type(-e)
	session.ExportType = options.SetExportType()
	//check for valid request method(-m)
	session.Method = options.SetRequestMethod()
	//set status code to be excluded from the results
	session.ExcludeStatus = options.ExcludeStatusCode()

	//parse option data and store 'em
	parsedNum.InputData = options.ParseNumRange()
	parsedAscii.InputData = options.ParseAsciiRange()
	parsedChar.InputData = options.ParseCharList()
	parsedInput.InputData = options.ReadFuzzFile()

	//if no data exists for fuzzing then throw error
	if len(parsedInput.InputData) == 0 && len(parsedNum.InputData) == 0 && len(parsedAscii.InputData) == 0 && len(parsedChar.InputData) == 0 {
		utils.ShowError("No fuzzing data provided")
		utils.ShowInfo("Use -h option to display usage menu")
	}

	//function to create output folder
	session.OutDir = options.SetOutputDir()
	//setting metaData to each entity
	parsedNum.MetaData = session
	parsedAscii.MetaData = session
	parsedChar.MetaData = session
	parsedInput.MetaData = session

	session.DisplayInfo()
	//begin the fuzzing process

	parsedNum.BeginFuzzing("numeric")
	parsedAscii.BeginFuzzing("ascii")
	parsedChar.BeginFuzzing("character")
	parsedInput.BeginFuzzing("file data")

	utils.ShowSuccess("Fuzzing Complete!\n")
	// fmt.Printf("%+v\n", parsedNum.Result)
	// fmt.Printf("%+v\n", parsedAscii.Result)
	// fmt.Printf("%+v\n", parsedChar.Result)
	// fmt.Printf("%+v\n", parsedInput.Result)
}
