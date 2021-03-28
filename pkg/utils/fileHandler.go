package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//function to check if directory exists?
func DirExists(path string) bool {
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		// path is a directory
		return true
	}
	return false
}

//function to check if file exists?
func FileExists(path string) bool {
	if stat, err := os.Stat(path); err == nil && !stat.IsDir() {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

//TODO: find alternative for text template, use struct or something
//function to write data to file
//only for text files
func WriteFile(filePath string, fileData string) {
	f, err := os.Create(filePath)
	CheckErr(err, "Error occured while creating file: ", filePath)
	_, err = f.WriteString(fileData)
	CheckErr(err, "Error occured while writing data to file: ", filePath)
	ShowSuccess("Exported at: ", filePath)
}

//for json export
//rename the members
type JsonExportTemplate struct {
	Export string              `json:"export_type"`
	Method string              `json:"method"`
	Target string              `json:"target"`
	Date   string              `json:"date_time"`
	Result map[string][]string `json:"result"`
}

func (j JsonExportTemplate) WriteJson(filePath string) {
	f, err := json.MarshalIndent(j, "", " ")
	CheckErr(err, err)
	err = ioutil.WriteFile(filePath, f, 0644)
	CheckErr(err, err)
	ShowSuccess("Exported at: ", filePath)
}
