/*
# escapes provides Ansi escape sequence utility functions.

	// all escape codes provided will be prepended to the text string
	func Sprint(text string, escapeCodes ...string) string // escapeCodes should be provided functionally from the structs provided
	func Print(text string, escapeCodes ...string)
	func Println(text string, escapeCodes ...string)

	Color // struct containing a bunch of text color func()s
	Background // contains terminal background color func()s
	Style // contains styling (like bold, italics) func()s
	Erase // contains func()s for erasing text from the terminal
	Screen // contains func()s for Screen manipulation
	Cursor // contains func()s for moving the cursor around

	// Example:
	Print("Hello, world!", )
*/
package escapes

import (
	"fmt"
	"strings"
)

const csi string = "\033["

/*
Sprint returns a string with all the escape codes added.

# Params:

	text string
	escapeCodes ...string // any amount of escape sequence func()s, like Style.Reset() or Color.Blue()

# Example:

	var stringWithCodes string = Sprint("test", Color.Green(), Style.Bold())
	// stringWithCodes is a string that will appear green and bold
*/
func Sprint(text string, escapeCodes ...string) string {
	var b strings.Builder
	for _, code := range escapeCodes {
		b.WriteString(code)
	}
	b.WriteString(text)
	return b.String()
}

/*
# Print is fmt.Print but escape codes provided
*/
func Print(text string, escapeCodes ...string) {
	for _, code := range escapeCodes {
		fmt.Print(code)
	}
	fmt.Print(text)
}
func Println(text string, escapeCodes ...string) {
	for _, code := range escapeCodes {
		fmt.Print(code)
	}
	fmt.Println(text)
}

const (
	reset              = csi + "0m"
	bold               = csi + "1m"
	boldOff            = csi + "22m"
	light              = csi + "2m"
	italic             = csi + "3m"
	italicOff          = csi + "23m"
	blink              = csi + "5m"
	blinkOff           = csi + "25m"
	reversed           = csi + "7m"
	reversedOff        = csi + "27m"
	hidden             = csi + "8m"
	hiddenOff          = csi + "28m"
	defaultFont        = csi + "10m"
	black              = csi + "30m"
	red                = csi + "31m"
	green              = csi + "32m"
	yellow             = csi + "33m"
	blue               = csi + "34m"
	magenta            = csi + "35m"
	cyan               = csi + "36m"
	white              = csi + "37m"
	bgBlack            = csi + "40m"
	bgRed              = csi + "41m"
	bgGreen            = csi + "42m"
	bgYellow           = csi + "43m"
	bgBlue             = csi + "44m"
	bgMagenta          = csi + "45m"
	bgCyan             = csi + "46m"
	bgWhite            = csi + "47m"
	eraseScreenToEnd   = csi + "0J"
	eraseScreenToBegin = csi + "1J"
	eraseAll           = csi + "2J"
	eraseLineToEnd     = csi + "0K"
	eraseLineToBegin   = csi + "1K"
	eraseLine          = csi + "2K"
	scrollUp           = csi + "%dS"
	scrollDown         = csi + "%dT"
	alternateScreenOn  = csi + "?1049h"
	alternateScreenOff = csi + "?1049l"
	cursorRight        = csi + "%s%dC"
	cursorLeft         = csi + "%s%dD"
	cursorUp           = csi + "%s%dA"
	cursorDown         = csi + "%s%dB"
	cursorLineNext     = csi + "%s%dE"
	cursorLinePrevious = csi + "%s%dF"
	cursorColumn       = csi + "%s%dG"
	cursorMove         = csi + "%s%d;%dH"
	font               = csi + "%s1%dm"
)
