package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/DarthCucumber/gofuzz/pkg/data"
	"github.com/DarthCucumber/gofuzz/pkg/utils"
)

func main() {
	var options data.Options
	var wg sync.WaitGroup
	var metaData data.SessionData

	flag.BoolVar(&options.ShowHelp, "h", false, "shows usage details")
	flag.StringVar(&options.TargetUrl, "u", "", "takes in URL for fuzzing")
	flag.StringVar(&options.NumRange, "n", "", "takes in range of numbers for fuzzing")
	flag.StringVar(&options.CharList, "c", "", "takes in range of characters for fuzzing")
	flag.StringVar(&options.AsciiRange, "a", "", "takes in range of ascii values and fuzzes for corresponding charater")
	flag.StringVar(&options.OutputDir, "o", "./output", "set output folder to save the results")
	flag.StringVar(&options.InputFile, "f", "", "file path to list of fuzz data")
	flag.StringVar(&options.ExportType, "e", "txt", "data format in which the result will be stored in the output file")
	flag.StringVar(&options.Method, "m", "HEAD", "Request method [HEAD/GET/POST]")
	flag.StringVar(&options.Exclude, "ex", "", "takes in status code separated by commas to be excluded from display result, however everything is included in the result files")
	flag.Parse()

	//detect -h and show help options
	options.DisplayHelp()

	var banner string = `
░▄▀▒░▄▀▄▒█▀░█▒█░▀█▀░▀█▀
░▀▄█░▀▄▀░█▀░▀▄█░█▄▄░█▄▄	v1.0.0

`
	fmt.Println(banner)

	if len(options.TargetUrl) == 0 {
		utils.ShowError("No URL provided for fuzzing")
		utils.ShowWarning("use -h option for usage options")
		os.Exit(0)
	}

	//parse target url
	metaData.ParsedUrl = options.ParseUrl()
	//check for valid export type(-e)
	metaData.ExportType = options.SetExportType()
	//check for valid request method(-m)
	metaData.Method = options.SetRequestMethod()
	//set of status code to be excluded from the results
	metaData.ExcludeStatus = options.ExcludeStatusCode()

	//parse option data
	metaData.ParsedNum = options.ParseNumRange()
	metaData.ParsedAscii = options.ParseAsciiRange()
	metaData.ParsedChar = options.ParseCharList()
	metaData.ParsedFileInput = options.ReadFuzzFile()

	//if no data exists for fuzzing then throw error
	if metaData.IsEmpty() {
		utils.ShowError("No fuzz data provided for fuzzing")
		os.Exit(0)
	}

	//function to create output folder
	metaData.OutDir = options.SetOutputDir()
	//initializing result map
	metaData.NumRes = make(map[string][]string)
	metaData.AsciiRes = make(map[string][]string)
	metaData.CharRes = make(map[string][]string)
	metaData.InputRes = make(map[string][]string)

	//channel to get result from go routine
	c := make(chan []string)
	//TODO: any improvements to this?
	//fuzzing part
	if len(metaData.ParsedNum) != 0 {
		utils.ShowInfo("Fuzzing Numeric List")
		//iterate over provided data
		for _, u := range metaData.ParsedNum {
			wg.Add(1)
			go utils.Fuzz(metaData.ParsedUrl, u, metaData.Method, c, &wg)
			res := <-c
			//res[0]:statuscode, res[1]:fuzzing data, res[2]:result URL
			//check if status code included in
			//exclude list
			if !metaData.ContainsCode(res[0]) {
				metaData.NumRes[res[0]] = append(metaData.NumRes[res[0]], res[1])
				fmt.Printf("[%s] %s\n", res[0], res[2])
			}
		}
	}
	wg.Wait()
	if len(metaData.ParsedAscii) != 0 {
		utils.ShowInfo("Fuzzing ASCII List")
		//iterate over provided data
		for _, u := range metaData.ParsedAscii {
			wg.Add(1)
			go utils.Fuzz(metaData.ParsedUrl, u, metaData.Method, c, &wg)
			res := <-c
			//res[0]:statuscode, res[1]:fuzzing data, res[2]:result URL
			//check if status code included in
			//exclude list
			if !metaData.ContainsCode(res[0]) {
				metaData.AsciiRes[res[0]] = append(metaData.AsciiRes[res[0]], res[1])
				fmt.Printf("[%s] %s\n", res[0], res[2])
			}
		}
	}
	wg.Wait()
	if len(metaData.ParsedChar) != 0 {
		utils.ShowInfo("Fuzzing Character List")
		//iterate over provided data
		for _, u := range metaData.ParsedChar {
			wg.Add(1)
			go utils.Fuzz(metaData.ParsedUrl, u, metaData.Method, c, &wg)
			res := <-c
			//res[0]:statuscode, res[1]:fuzzing data, res[2]:result URL
			//check if status code included in
			//exclude list
			if !metaData.ContainsCode(res[0]) {
				metaData.CharRes[res[0]] = append(metaData.CharRes[res[0]], res[1])
				fmt.Printf("[%s] %s\n", res[0], res[2])
			}
		}
	}
	wg.Wait()
	if len(metaData.ParsedFileInput) != 0 {
		utils.ShowInfo("Fuzzing User Input")
		//iterate over provided data
		for _, u := range metaData.ParsedFileInput {
			wg.Add(1)
			go utils.Fuzz(metaData.ParsedUrl, u, metaData.Method, c, &wg)
			res := <-c
			//res[0]:statuscode, res[1]:fuzzing data, res[2]:result URL
			//check if status code included in
			//exclude list
			if !metaData.ContainsCode(res[0]) {
				metaData.InputRes[res[0]] = append(metaData.InputRes[res[0]], res[1])
				fmt.Printf("[%s] %s\n", res[0], res[2])
			}
		}
	}
	//wait and close the data
	wg.Wait()
	close(c)

	utils.ShowSuccess("Fuzzing done...")
	//Export
	utils.ShowInfo("Exporting results...")
	metaData.ExportData()
}
