package data

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/DarthCucumber/gofuzz/pkg/utils"
)

//to store processed/parsed data given by user
// the fuzzing result is stored over here as well
type SessionData struct {
	ParsedUrl       []string
	ParsedNum       []string //from -n
	ParsedChar      []string //from -c
	ParsedAscii     []string //from -a
	ParsedFileInput []string //from file input -f
	ExcludeStatus   []string //list of status codes to be excluded
	OutDir          string
	ExportType      string
	Method          string
	// result section
	NumRes   map[string][]string
	AsciiRes map[string][]string
	CharRes  map[string][]string
	InputRes map[string][]string
}

//function to check if status code
//exists?
func (m SessionData) ContainsCode(code string) bool {
	list := m.ExcludeStatus
	x := sort.SearchStrings(list, code)
	if x < len(list) && list[x] == code {
		return true
	}
	return false
}

//check if we have some data to fuzz?
func (m SessionData) IsEmpty() bool {
	isEmpty := 4
	//display banner
	target := strings.Join(m.ParsedUrl, "__")
	utils.ShowSuccess("Target:", target)
	utils.ShowSuccess("Fuzz for:")

	if len(m.ParsedAscii) != 0 {
		fmt.Println("> Ascii List")
		isEmpty -= 1
	}
	if len(m.ParsedChar) != 0 {
		fmt.Println("> Character List")
		isEmpty -= 1
	}
	if len(m.ParsedNum) != 0 {
		fmt.Println("> Numeric List")
		isEmpty -= 1
	}
	if len(m.ParsedFileInput) != 0 {
		fmt.Println("> Input List")
		isEmpty -= 1
	}
	//=> all are empty
	if isEmpty == 0 {
		return true
	}
	//=> not all are empty
	return false
}

//TODO: find alternative for text template, use struct or something
var txtTemplate string = `
Export Type: %s
Method: %s
Date: %s
-------------------------------------

`

//OPTIMISE: export function, can do better
func (sd SessionData) ExportData() {
	dateNtime := time.Now().Format("2006-01-02 15:04:05")
	switch sd.ExportType {
	case "txt":
		if len(sd.NumRes) != 0 {
			expData := fmt.Sprintf(txtTemplate, "TEXT", sd.Method, dateNtime)
			for statusCode, res := range sd.NumRes {
				expData += fmt.Sprintf("%s: \n%s", statusCode, strings.Join(res, "\n"))
			}
			//save to file outDir/numeric_result.txt
			path := sd.OutDir + "/numeric_result_" + dateNtime + ".txt"
			utils.WriteFile(path, expData)
		}
		if len(sd.AsciiRes) != 0 {
			expData := fmt.Sprintf(txtTemplate, "TEXT", sd.Method, dateNtime)
			for statusCode, res := range sd.AsciiRes {
				expData += fmt.Sprintf("%s: \n%s", statusCode, strings.Join(res, "\n"))
			}
			//save to file outDir/Ascii_result.txt
			path := sd.OutDir + "/ascii_result_" + dateNtime + ".txt"
			utils.WriteFile(path, expData)
		}
		if len(sd.CharRes) != 0 {
			expData := fmt.Sprintf(txtTemplate, "TEXT", sd.Method, dateNtime)
			for statusCode, res := range sd.CharRes {
				expData += fmt.Sprintf("%s: \n%s", statusCode, strings.Join(res, "\n"))
			}
			//save to file outDir/character_result.txt
			path := sd.OutDir + "/char_result_" + dateNtime + ".txt"
			utils.WriteFile(path, expData)
		}
		if len(sd.InputRes) != 0 {
			expData := fmt.Sprintf(txtTemplate, "TEXT", sd.Method, dateNtime)
			for statusCode, res := range sd.InputRes {
				expData += fmt.Sprintf("%s: \n%s", statusCode, strings.Join(res, "\n"))
			}
			//save to file outDir/input_result.txt
			path := sd.OutDir + "/input_file_result_" + dateNtime + ".txt"
			utils.WriteFile(path, expData)
		}
	case "json":
		var jsonexp utils.JsonExportTemplate
		jsonexp.Date = dateNtime
		jsonexp.Export = "JSON"
		jsonexp.Method = sd.Method
		jsonexp.Target = strings.Join(sd.ParsedUrl, "__")
		if len(sd.NumRes) != 0 {
			jsonexp.Result = sd.NumRes
			//save to file outDir/numeric_result.txt
			path := sd.OutDir + "/numeric_result_" + dateNtime + ".json"
			jsonexp.WriteJson(path)
		}
		if len(sd.AsciiRes) != 0 {
			jsonexp.Result = sd.AsciiRes
			//save to file outDir/Ascii_result.txt
			path := sd.OutDir + "/ascii_result_" + dateNtime + ".json"
			jsonexp.WriteJson(path)
		}
		if len(sd.CharRes) != 0 {
			jsonexp.Result = sd.CharRes
			//save to file outDir/character_result.txt
			path := sd.OutDir + "/char_result_" + dateNtime + ".json"
			jsonexp.WriteJson(path)
		}
		if len(sd.InputRes) != 0 {
			jsonexp.Result = sd.InputRes
			//save to file outDir/input_result.txt
			path := sd.OutDir + "/input_file_result_" + dateNtime + ".json"
			jsonexp.WriteJson(path)
		}

	//TODO: add csv support
	default:
		utils.ShowError("Invalid export type `", sd.ExportType, "` provided in exportData() method")
	}
	utils.ShowSuccess("Finished")
}
