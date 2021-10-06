package utils

import (
	"fmt"

	"github.com/mattn/go-colorable"
)

var (
	infoColor = "\033[1;34m[i]%s\033[0m\n"
	// noticeColor  = "\033[1;36m%s\033[0m"
	warningColor = "\033[1;33m[!]%s\033[0m\n"
	errorColor   = "\033[1;31m[x]%s\033[0m\n"
	debugColor   = "\033[0;36m[>]%s\033[0m\n"
	successColor = "\033[1;32m[*]%s\033[0m\n"

	out = colorable.NewColorableStdout()
)

func ShowDebug(msg ...string) {
	var tmp string
	for _, s := range msg {
		tmp += fmt.Sprintf(" %s", s)
	}
	fmt.Fprintf(out, debugColor, tmp)
}
func ShowInfo(msg ...string) {
	var tmp string
	for _, s := range msg {
		tmp += fmt.Sprintf(" %s", s)
	}
	fmt.Fprintf(out, infoColor, tmp)
}
func ShowWarning(msg ...interface{}) {
	var tmp string
	for _, s := range msg {
		tmp += fmt.Sprintf(" %s", s)
	}
	fmt.Fprintf(out, warningColor, tmp)
}
func ShowError(msg ...interface{}) {
	var tmp string
	for _, s := range msg {
		tmp += fmt.Sprintf(" %s", s)
	}
	fmt.Fprintf(out, errorColor, tmp)
}
func ShowSuccess(msg ...interface{}) {
	var tmp string
	for _, s := range msg {
		tmp += fmt.Sprintf(" %s", s)
	}
	fmt.Fprintf(out, successColor, tmp)
}
