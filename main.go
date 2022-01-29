package colors

import (
	"fmt"
)


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
// As takes an output to return and any modifiers of the Format type
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
