package data

import (
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/DarthCucumber/gofuzz/pkg/utils"
	"github.com/cheggaaa/pb/v3"
)

//to store processed/parsed data given by user
// the fuzzing result is stored over here as well
type SessionData struct {
	ParsedUrl     []string
	ExcludeStatus []string //list of status codes to be excluded
	OutDir        string
	ExportType    string
	Method        string
	Body          string //request body input by user
	Header        string //request header input by user
	Timeout       int    //timeout for each request
}

type FuzzData struct {
	MetaData  SessionData
	InputData []string
	Result    map[string][]string
}

//mind that pointer xD
func (f *FuzzData) BeginFuzzing(displayText string) {
	//initialize result map
	// f.Result = make(map[int][]string)
	// worker for this specific process is
	var pw sync.WaitGroup
	var ow sync.WaitGroup
	f.Result = make(map[string][]string)
	maxReq := 5
	count := len(f.InputData)
	// fmt.Println(f.InputData, count)
	//create progress bar
	if count == 0 {
		return
	}
	if count > 100 {
		maxReq = 100
	}

	//progress bar
	tmpl := `{{ magenta "` + displayText + `" }} {{ counters . }} {{ bar . "[" "#" (cycle . "" "" "" "" ) "." "]"}} {{speed . | blue }} {{percent .}}`
	// start bar
	bar := pb.ProgressBarTemplate(tmpl).Start(count)

	//make concurrent requests
	out := make(chan []string, maxReq)
	go func() {
		pw.Wait()
		// bar.Finish()
		close(out)
	}()

	for _, d := range f.InputData {
		pw.Add(1)
		//form payload URL
		url := strings.Join(f.MetaData.ParsedUrl, url.PathEscape(d))
		go utils.MakeRequest(f.MetaData.Method, url, f.MetaData.Timeout, out, &pw)
		//storing in result map
	}
	ow.Add(1)
	go func() {
		for r := range out {
			// fmt.Printf("[%s] %s\n", r[0], r[1])
			bar.Increment()
			f.Result[r[0]] = append(f.Result[r[0]], r[1])
		}
		bar.Finish()
		ow.Done()
	}()
	ow.Wait()
}

//function to check if status code
//exists?
func (f FuzzData) ContainsStatusCode(code string) bool {
	list := f.MetaData.ExcludeStatus
	//search for the status code
	x := sort.SearchStrings(list, code)
	//exists? return true
	if x < len(list) && list[x] == code {
		return true
	}
	return false
}

func (sd SessionData) DisplayInfo() {
	var banner string = `
░▄▀▒░▄▀▄▒█▀░█▒█░▀█▀░▀█▀
░▀▄█░▀▄▀░█▀░▀▄█░█▄▄░█▄▄	v1.0.0

`
	fmt.Println(banner)

	utils.ShowSuccess("Target: ", strings.Join(sd.ParsedUrl, "__"))
	utils.ShowSuccess("Method: ", sd.Method)
	utils.ShowSuccess("Exclude: ", sd.ExcludeStatus)
	utils.ShowSuccess("Timeout: ", strconv.Itoa(sd.Timeout), "ms")
	utils.ShowSuccess("Export Type: ", sd.ExportType)
	utils.ShowSuccess("Output: ", sd.OutDir)
}

// export functions
//TODO: find alternative for text template, use struct or something
// var txtTemplate string = `
// Export Type: %s
// Method: %s
// Date: %s
// -------------------------------------

// `

//OPTIMISE: export function, can do better
// func (sd SessionData) ExportData() {
// 	dateNtime := time.Now().Format("2006-01-02 15:04:05")
// 	switch sd.ExportType {
// 	case "txt":
// 		if len(sd.NumRes) != 0 {
// 			expData := fmt.Sprintf(txtTemplate, "TEXT", sd.Method, dateNtime)
// 			for statusCode, res := range sd.NumRes {
// 				expData += fmt.Sprintf("%s: \n%s", statusCode, strings.Join(res, "\n"))
// 			}
// 			//save to file outDir/numeric_result.txt
// 			path := sd.OutDir + "/numeric_result_" + dateNtime + ".txt"
// 			utils.WriteFile(path, expData)
// 		}
// 		if len(sd.AsciiRes) != 0 {
// 			expData := fmt.Sprintf(txtTemplate, "TEXT", sd.Method, dateNtime)
// 			for statusCode, res := range sd.AsciiRes {
// 				expData += fmt.Sprintf("%s: \n%s", statusCode, strings.Join(res, "\n"))
// 			}
// 			//save to file outDir/Ascii_result.txt
// 			path := sd.OutDir + "/ascii_result_" + dateNtime + ".txt"
// 			utils.WriteFile(path, expData)
// 		}
// 		if len(sd.CharRes) != 0 {
// 			expData := fmt.Sprintf(txtTemplate, "TEXT", sd.Method, dateNtime)
// 			for statusCode, res := range sd.CharRes {
// 				expData += fmt.Sprintf("%s: \n%s", statusCode, strings.Join(res, "\n"))
// 			}
// 			//save to file outDir/character_result.txt
// 			path := sd.OutDir + "/char_result_" + dateNtime + ".txt"
// 			utils.WriteFile(path, expData)
// 		}
// 		if len(sd.InputRes) != 0 {
// 			expData := fmt.Sprintf(txtTemplate, "TEXT", sd.Method, dateNtime)
// 			for statusCode, res := range sd.InputRes {
// 				expData += fmt.Sprintf("%s: \n%s", statusCode, strings.Join(res, "\n"))
// 			}
// 			//save to file outDir/input_result.txt
// 			path := sd.OutDir + "/input_file_result_" + dateNtime + ".txt"
// 			utils.WriteFile(path, expData)
// 		}
// 	case "json":
// 		var jsonexp utils.JsonExportTemplate
// 		jsonexp.Date = dateNtime
// 		jsonexp.Export = "JSON"
// 		jsonexp.Method = sd.Method
// 		jsonexp.Target = strings.Join(sd.ParsedUrl, "__")
// 		if len(sd.NumRes) != 0 {
// 			jsonexp.Result = sd.NumRes
// 			//save to file outDir/numeric_result.txt
// 			path := sd.OutDir + "/numeric_result_" + dateNtime + ".json"
// 			jsonexp.WriteJson(path)
// 		}
// 		if len(sd.AsciiRes) != 0 {
// 			jsonexp.Result = sd.AsciiRes
// 			//save to file outDir/Ascii_result.txt
// 			path := sd.OutDir + "/ascii_result_" + dateNtime + ".json"
// 			jsonexp.WriteJson(path)
// 		}
// 		if len(sd.CharRes) != 0 {
// 			jsonexp.Result = sd.CharRes
// 			//save to file outDir/character_result.txt
// 			path := sd.OutDir + "/char_result_" + dateNtime + ".json"
// 			jsonexp.WriteJson(path)
// 		}
// 		if len(sd.InputRes) != 0 {
// 			jsonexp.Result = sd.InputRes
// 			//save to file outDir/input_result.txt
// 			path := sd.OutDir + "/input_file_result_" + dateNtime + ".json"
// 			jsonexp.WriteJson(path)
// 		}

// 	//TODO: add csv support
// 	default:
// 		utils.ShowError("Invalid export type `", sd.ExportType, "` provided in exportData() method")
// 	}
// 	utils.ShowSuccess("Finished")
// }
