package data

import (
	"fmt"
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
	OutDir          string
	ExportType      string
	Method          string
	// result section
	NumRes   map[string][]string
	AsciiRes map[string][]string
	CharRes  map[string][]string
	InputRes map[string][]string
}

var banner string = `
░▄▀▒░▄▀▄▒█▀░█▒█░▀█▀░▀█▀
░▀▄█░▀▄▀░█▀░▀▄█░█▄▄░█▄▄	v1.0.0

`

//check if we have some data to fuzz?
func (m SessionData) IsEmpty() bool {
	isEmpty := 4
	//display banner
	fmt.Println(banner)
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

var txtTemplate string = `
Export Type: %s
Method: %s
Date: %s
-------------------------------------

`

//TODO: export function
func (sd SessionData) ExportData() {
	switch sd.ExportType {
	case "txt":
		dateNtime := time.Now().Format("2006-01-02 15:04:05")
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
		utils.ShowSuccess("Finished")
	//TODO: add json and csv support
	default:
		utils.ShowError("Invalid export type `", sd.ExportType, "` provided in exportData() method")
	}
}
