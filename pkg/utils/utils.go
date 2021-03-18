package utils

import (
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

func AsciiToChar(start int, end int) []string {
	var tmp []string
	for j := start; j <= end; j++ {
		tmpRune := rune(j)
		tmp = append(tmp, string(tmpRune))
	}
	return tmp
}
