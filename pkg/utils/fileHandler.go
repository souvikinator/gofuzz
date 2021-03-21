package utils

import (
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

//function to write data to file
//only for text files
func WriteFile(filePath string, fileData string) {
	f, err := os.Create(filePath)
	CheckErr(err, "Error occured while creating file: ", filePath)
	_, err = f.WriteString(fileData)
	CheckErr(err, "Error occured while writing data to file: ", filePath)
	ShowSuccess("Exported at: ", filePath)
}
