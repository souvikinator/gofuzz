package data

import (
	"fmt"
	"strings"

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

// var txtTemplate string = `
// Export Type: %s
// Method: %s
// Date: %s
// -------------------------------------

// %s
// `

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

//TODO: export function
// func (fr SessionData) ExportData(exportType string, outDir string) {
// 	switch exportType {
// 	case "TXT":
// 		// expData := fmt.Sprintf(txtTemplate, "TEXT", fr.Method, fr.Date)
// 		//find a way to extract data from ft.Result
// 		fmt.Printf("%+v\n", fr.Data)
// 	case "JSON":
// 		fmt.Printf("%+v\n", fr.Data)

// 	case "CSV":
// 		fmt.Printf("%+v\n", fr.Data)

// 	default:
// 		color.ShowError("Invalid export type `", exportType, "` provided in exportData() method")
// 	}
// }
