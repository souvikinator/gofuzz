package data

import (
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

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
	var pw sync.WaitGroup //request worker
	var ow sync.WaitGroup //output worker
	f.Result = make(map[string][]string)
	count := len(f.InputData)
	maxReq := count
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
	out := make(chan []string, maxReq)
	go func() {
		pw.Wait()
		close(out)
	}()
	for _, d := range f.InputData {
		pw.Add(1)
		//form payload URL
		url := strings.Join(f.MetaData.ParsedUrl, url.PathEscape(d))
		//make concurrent requests
		go utils.MakeRequest(f.MetaData.Method, url, f.MetaData.Timeout, out, &pw)
	}
	//get the results from channel
	ow.Add(1)
	go func() {
		for r := range out {
			bar.Increment()
			if !f.ContainsStatusCode(r[0]) {
				f.Result[r[0]] = append(f.Result[r[0]], r[1])
			}
		}
		bar.Finish()
		ow.Done()
	}()
	ow.Wait()
	// export data
	f.ExportData(displayText)
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
var txtTemplate string = `
Export Type: %s
Method: %s
Date: %s
-------------------------------------

`

func (f FuzzData) ExportData(filename string) {
	if len(f.Result) == 0 {
		return
	}
	utils.ShowInfo("exporting results")
	dateNtime := time.Now().Format("2006-01-02 15:04:05")
	switch f.MetaData.ExportType {
	case "txt":
		expData := fmt.Sprintf(txtTemplate, "TEXT", f.MetaData.Method, dateNtime)
		for statusCode, res := range f.Result {
			expData += fmt.Sprintf("%s: \n%s", statusCode, strings.Join(res, "\n"))
		}
		path := f.MetaData.OutDir + filename + "/" + "_" + dateNtime + ".txt"
		utils.WriteFile(path, expData)
	case "json":
		var jsonexp utils.JsonExportTemplate
		jsonexp.Date = dateNtime
		jsonexp.Export = "JSON"
		jsonexp.Method = f.MetaData.Method
		jsonexp.Target = strings.Join(f.MetaData.ParsedUrl, "__")
		jsonexp.Result = f.Result
		path := f.MetaData.OutDir + "/" + filename + "_" + dateNtime + ".json"
		jsonexp.WriteJson(path)
	//TODO: add csv support
	default:
		utils.ShowError("Invalid export type `", f.MetaData.ExportType, "` provided in exportData() method")
	}
}
