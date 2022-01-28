package colors

import (
	"fmt"
	"runtime"
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

func init() {
	if runtime.GOOS == "windows" {
		reset = ""
		red = ""
		green = ""
		yellow = ""
		blue = ""
		purple = ""
		cyan = ""
		gray = ""
		white = ""
	}
}

func AsRed(text string) string {
	return fmt.Sprintf("%s%s%s", red, text, reset)
}

func AsGreen(text string) string {
	return fmt.Sprintf("%s%s%s", green, text, reset)

}

func AsYellow(text string) string {
	return fmt.Sprintf("%s%s%s", yellow, text, reset)

}

func AsBlue(text string) string {
	return fmt.Sprintf("%s%s%s", blue, text, reset)

}

func AsPurple(text string) string {
	return fmt.Sprintf("%s%s%s", purple, text, reset)

}

func AsCyan(text string) string {
	return fmt.Sprintf("%s%s%s", cyan, text, reset)

}

func AsGray(text string) string {
	return fmt.Sprintf("%s%s%s", gray, text, reset)

}

func AsWhite(text string) string {
	return fmt.Sprintf("%s%s%s", white, text, reset)

}