package data

import (
	"fmt"
	"strings"

	"github.com/DarthCucumber/gofuzz/pkg/color"
)

//TODO: rename the file, sounds weird
type FuzzData struct {
	TargetUrl     []string
	AsciiFuzzData []string //from -a
	CharFuzzData  []string //from -c
	NumFuzzData   []string //from -n options
	InputFuzzData []string //from file input
}

type FuzzResult struct {
	AsciiFuzzRes map[string]string
	CharFuzzRes  map[string]string
	NumFuzzRes   map[string]string
	InputFuzzRes map[string]string
}

var banner string = `
░▄▀▒░▄▀▄▒█▀░█▒█░▀█▀░▀█▀
░▀▄█░▀▄▀░█▀░▀▄█░█▄▄░█▄▄	v1.0.0

`

//check if we have some data to fuzz?
func (m FuzzData) IsEmpty() bool {
	isEmpty := 4
	//display banner
	fmt.Println(banner)
	target := strings.Join(m.TargetUrl, "__")
	color.ShowSuccess("Target:", target)
	color.ShowSuccess("Fuzz for:")

	if len(m.AsciiFuzzData) != 0 {
		fmt.Println("> Ascii List")
		isEmpty -= 1
	}
	if len(m.CharFuzzData) != 0 {
		fmt.Println("> Character List")
		isEmpty -= 1
	}
	if len(m.NumFuzzData) != 0 {
		fmt.Println("> Numeric List")
		isEmpty -= 1
	}
	if len(m.InputFuzzData) != 0 {
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
