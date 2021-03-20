package main

import (
	"flag"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/DarthCucumber/gofuzz/pkg/color"
	"github.com/DarthCucumber/gofuzz/pkg/connection"
	"github.com/DarthCucumber/gofuzz/pkg/data"
)

func main() {
	var menu data.Options
	var wg sync.WaitGroup
	var meta data.FuzzData //TODO: rename variable

	flag.BoolVar(&menu.ShowHelp, "h", false, "shows usage details")
	flag.StringVar(&menu.FuzzUrl, "u", "", "takes in URL for fuzzing")
	flag.StringVar(&menu.NumRange, "n", "", "takes in range of numbers for fuzzing")
	flag.StringVar(&menu.CharList, "c", "", "takes in range of characters for fuzzing")
	flag.StringVar(&menu.AsciiRange, "a", "", "takes in range of ascii values and fuzzes for corresponding charater")
	flag.StringVar(&menu.OutDir, "o", "./output", "set output folder to save the results")
	flag.StringVar(&menu.InputFile, "f", "", "file path to list of fuzz data")
	flag.StringVar(&menu.ExportType, "e", "txt", "data format in which the result will be stored in the output file")
	flag.StringVar(&menu.ReqMethod, "m", "HEAD", "Request method [HEAD/GET/POST]")
	flag.Parse()

	//detect -h and show help menu
	menu.DisplayHelp()

	if len(menu.FuzzUrl) == 0 {
		color.ShowError("No URL provided for fuzzing")
		color.ShowWarning("use -h option for usage menu")
		os.Exit(0)
	}

	//parse target url
	meta.TargetUrl = menu.ParseUrl()

	//check for valid export type(-e)
	menu.SetExportType()
	//check for valid request method(-m)
	menu.SetRequestMethod()

	meta.NumFuzzData = menu.ParseNumRange()
	meta.AsciiFuzzData = menu.ParseAsciiRange()
	meta.CharFuzzData = menu.ParseCharList()
	meta.InputFuzzData = menu.ReadFuzzFile()

	if meta.IsEmpty() {
		color.ShowError("No fuzz data provided for fuzzing")
		os.Exit(0)
	}

	//TODO: make for other fuzzers
	//FIXME: comeup with elegant solution
	if len(meta.NumFuzzData) != 0 {
		color.ShowInfo("Fuzzing for Numeric Values")
		for _, p := range meta.NumFuzzData {
			u := strings.Join(meta.TargetUrl, url.QueryEscape(p))
			wg.Add(1)
			go connection.MakeReq(u, menu.ReqMethod, &wg)
		}
	}
	wg.Wait()
	if len(meta.AsciiFuzzData) != 0 {
		color.ShowInfo("Fuzzing for Ascii Values")
		for _, p := range meta.AsciiFuzzData {
			u := strings.Join(meta.TargetUrl, url.QueryEscape(p))
			wg.Add(1)
			go connection.MakeReq(u, menu.ReqMethod, &wg)
		}
	}
	wg.Wait()
	if len(meta.CharFuzzData) != 0 {
		color.ShowInfo("Fuzzing for Character List")
		for _, p := range meta.CharFuzzData {
			u := strings.Join(meta.TargetUrl, url.QueryEscape(p))
			wg.Add(1)
			go connection.MakeReq(u, menu.ReqMethod, &wg)
		}
	}
	wg.Wait()
	if len(meta.InputFuzzData) != 0 {
		color.ShowInfo("Fuzzing for Input List")
		for _, p := range meta.InputFuzzData {
			u := strings.Join(meta.TargetUrl, url.QueryEscape(p))
			wg.Add(1)
			go connection.MakeReq(u, menu.ReqMethod, &wg)
		}
	}
	wg.Wait()
	color.ShowSuccess("Fuzzing Complete")
}
