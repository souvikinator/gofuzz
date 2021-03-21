package utils

import (
	"fmt"
)

var (
	infoColor = "\033[1;34m[i]%s\033[0m\n"
	// noticeColor  = "\033[1;36m%s\033[0m"
	warningColor = "\033[1;33m[!]%s\033[0m\n"
	errorColor   = "\033[1;31m[x]%s\033[0m\n"
	// debugColor   = "\033[0;36m[]%s\033[0m"
	successColor = "\033[1;32m[*]%s\033[0m\n"
)

func ShowInfo(msg ...string) {
	var tmp string
	for _, s := range msg {
		tmp += fmt.Sprintf(" %s", s)
	}
	fmt.Printf(infoColor, tmp)
}
func ShowWarning(msg ...interface{}) {
	var tmp string
	for _, s := range msg {
		tmp += fmt.Sprintf(" %s", s)
	}
	fmt.Printf(warningColor, tmp)
}
func ShowError(msg ...interface{}) {
	var tmp string
	for _, s := range msg {
		tmp += fmt.Sprintf(" %s", s)
	}
	fmt.Printf(errorColor, tmp)
}
func ShowSuccess(msg ...interface{}) {
	var tmp string
	for _, s := range msg {
		tmp += fmt.Sprintf(" %s", s)
	}
	fmt.Printf(successColor, tmp)
}
