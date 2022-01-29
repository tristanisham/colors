package colors

import (
	"fmt"
)

var (
	reset       = "\033[0m"
	black       = "\033[30m"
	dark_red    = "\033[31m"
	dark_green  = "\033[32m"
	dark_yellow = "\033[33m"
	dark_blue   = "\033[34m"
	dark_magenta = "\033[35m"
	dark_cyan   = "\033[36m"
	dark_gray   = "\033[37m"
	dark_white  = "\033[97m"
)



func AsRed(text ...interface{}) string {
	return fmt.Sprint(dark_red, text, reset)
}

func AsGreen(text ...interface{}) string {
	return fmt.Sprint(dark_green, text, reset)

}

func AsBlack(text ...interface{}) string {
	return fmt.Sprint(black, text, reset)
}

func AsYellow(text ...interface{}) string {
	return fmt.Sprint(dark_yellow, text, reset)

}

func AsBlue(text ...interface{}) string {
	return fmt.Sprint(dark_blue, text, reset)

}

func AsPurple(text ...interface{}) string {
	return fmt.Sprint(dark_magenta, text, reset)

}

func AsMagenta(text ...interface{}) string {
	return fmt.Sprint(dark_magenta, text, reset)

}

func AsCyan(text ...interface{}) string {
	return fmt.Sprint(dark_cyan, text, reset)

}

func AsGray(text ...interface{}) string {
	return fmt.Sprint(dark_gray, text, reset)

}

func AsWhite(text ...interface{}) string {
	return fmt.Sprint(dark_white, text, reset)

}
