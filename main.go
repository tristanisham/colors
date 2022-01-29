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
	Black        Format = "30"
	DarkRed      Format = "31"
	DarkGreen    Format = "32"
	DarkYellow   Format = "33"
	DarkBlue     Format = "34"
	DarkMagenta  Format = "35"
	DarkCyan     Format = "36"
	LightGray    Format = "37"
	DarkGray     Format = "90"
	LightRed     Format = "91"
	LightGreen   Format = "92"
	LightYellow  Format = "93"
	LightBlue    Format = "94"
	LightMagenta Format = "95"
	LightCyan    Format = "96"
	White        Format = "97"
	// Background
	BgBlack        Format = "40"
	BgDarkRed      Format = "41"
	BgDarkGreen    Format = "42"
	BgDarkYellow   Format = "43"
	BgDarkBlue     Format = "44"
	BgDarkCyan     Format = "46"
	BgLightGray    Format = "100"
	BgLightRed     Format = "101"
	BgLightGreen   Format = "102"
	BgLightYellow  Format = "103"
	BgLightBlue    Format = "104"
	BgLightMagenta Format = "105"
	BgLightCyan    Format = "106"
	BgWhite        Format = "107"
	// Format
	Bold         Format = "1"
	Underline    Format = "4"
	NoUnderline  Format = "24"
	// ReverseText  Format = "7"
	PositiveText Format = "27"
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
	t := ""
	for i, f := range formats {
		t += fmt.Sprintf("%v", f)
		if i != len(formats)-1 {
			t += ";"
		}
	}
	return fmt.Sprintf("\033[%sm%v%v", t, text, reset)
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
