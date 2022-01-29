package colors

import (
	"fmt"
)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var blue = "\033[34m"
var purple = "\033[35m"
var cyan = "\033[36m"
var gray = "\033[37m"
var white = "\033[97m"

func AsRed(text ...interface{}) string {
	return fmt.Sprint(red, text, reset)
}

func AsGreen(text ...interface{}) string {
	return fmt.Sprint(green, text, reset)

}

func AsYellow(text ...interface{}) string {
	return fmt.Sprint(yellow, text, reset)

}

func AsBlue(text ...interface{}) string {
	return fmt.Sprint(blue, text, reset)

}

func AsPurple(text ...interface{}) string {
	return fmt.Sprint(purple, text, reset)

}

func AsCyan(text ...interface{}) string {
	return fmt.Sprint(cyan, text, reset)

}

func AsGray(text ...interface{}) string {
	return fmt.Sprint(gray, text, reset)

}

func AsWhite(text ...interface{}) string {
	return fmt.Sprint(white, text, reset)

}
