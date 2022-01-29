package colors

import (
	"fmt"
)

// var (
// 	reset       = "0m"
// 	black       = "30m"
// 	dark_red    = "31m"
// 	dark_green  = "32m"
// 	dark_yellow = "33m"
// 	dark_blue   = "34m"
// 	dark_magenta = "35m"
// 	dark_cyan   = "36m"
// 	dark_gray   = "37m"
// 	dark_white  = "97m"
// )


type Format string

const (
	Black        Format = "30m"
	DarkRed      Format = "31m"
	DarkGreen    Format = "32m"
	DarkYellow   Format = "33m"
	DarkBlue     Format = "34m"
	DarkMagenta  Format = "35m"
	DarkCyan     Format = "36m"
	LightGray    Format = "37m"
	DarkGray     Format = "90m"
	LightRed     Format = "91m"
	LightGreen   Format = "92m"
	LightYellow  Format = "93m"
	LightBlue    Format = "94m"
	LightMagenta Format = "95m"
	LightCyan    Format = "96m"
	White        Format = "97m"
	// Format
	Bold         Format = "1m"
	Underline    Format = "4m"
	NoUnderline  Format = "24m"
	ReverseText  Format = "7m"
	PositiveText Format = "27m"
	reset        Format = "\033[0m"
)

func As(text interface{}, options ...interface{}) string {
	formats := make([]Format, 0)
	for _, option := range options {
		switch d := option.(type) {
		case Format:
			formats = append(formats, d)
		}
	}
	t := "\033["
	for _, f := range formats {
		t += fmt.Sprint(f, ";")
	}
	t += "m"
	return fmt.Sprint(t, text, reset)
}

// func AsRed(text ...interface{}) string {
// 	return fmt.Sprint(dark_red, text, reset)
// }

// func AsGreen(text ...interface{}) string {
// 	return fmt.Sprint(dark_green, text, reset)

// }

// func AsBlack(text ...interface{}) string {
// 	return fmt.Sprint(black, text, reset)
// }

// func AsYellow(text ...interface{}) string {
// 	return fmt.Sprint(dark_yellow, text, reset)

// }

// func AsBlue(text ...interface{}) string {
// 	return fmt.Sprint(dark_blue, text, reset)

// }

// func AsPurple(text ...interface{}) string {
// 	return fmt.Sprint(dark_magenta, text, reset)

// }

// func AsMagenta(text ...interface{}) string {
// 	return fmt.Sprint(dark_magenta, text, reset)

// }

// func AsCyan(text ...interface{}) string {
// 	return fmt.Sprint(dark_cyan, text, reset)

// }

// func AsGray(text ...interface{}) string {
// 	return fmt.Sprint(dark_gray, text, reset)

// }

// func AsWhite(text ...interface{}) string {
// 	return fmt.Sprint(dark_white, text, reset)

// }
