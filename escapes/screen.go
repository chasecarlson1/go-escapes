package escapes

import "fmt"

type screen struct{}

/*
# Screen provides functions to scroll the terminal and enter/exit the alternate screen buffer.

	ScrollUp(lines int), ScrollDown(lines int) // scroll the terminal up and down `n` lines.

	AlternateBuffer(enter bool) // enter alternate screen mode if true, or exit the mode if false is passed
*/
var Screen screen = screen{}

func (screen) ScrollUp(lines int) string {
	return fmt.Sprintf(scrollUp, lines)
}
func (screen) ScrollDown(lines int) string {
	return fmt.Sprintf(scrollDown, lines)
}
func (screen) AlternateBuffer(enter bool) string {
	if enter {
		return alternateScreenOn
	} else {
		return alternateScreenOff
	}
}
