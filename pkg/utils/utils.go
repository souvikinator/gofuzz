package utils

import (
	"fmt"
	"os"
	"strconv"
)

//Takes a number as input in string format. Ex: "1","2"
//throws error if arg is not a number in string form
//returns a bool value (true if input is a char, false if everything ok) and input in integer
func ToInt(s string) (bool, int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return true, 0
	}
	return false, i
}

//Takes two integer as input and returns
//list of numbers from first arg -> second arg
func MakeNumList(start int, end int) []string {
	var tmp []string
	for j := start; j <= end; j++ {
		tmp = append(tmp, strconv.Itoa(j))
	}
	return tmp
}

//function to convert ascii to character
func AsciiToChar(start int, end int) []string {
	var tmp []string
	for j := start; j <= end; j++ {
		tmpRune := rune(j)
		tmp = append(tmp, string(tmpRune))
	}
	return tmp
}

//function to check for error
func CheckErr(e error, errMsg ...interface{}) {
	if e != nil {
		fmt.Println(errMsg...)
		os.Exit(0)
	}
}

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
