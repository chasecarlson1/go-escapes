package escapes

import (
	"fmt"
	"strings"
)

type cursor struct{}

/*
# Cursor provides functions to move the cursor around the terminal.

	Up(chars int), Down(chars int), Left(chars int), Right(chars int) // move by character in 4 directions.
	LineDown(lines int), LineUp(lines int) // move to the beginning of the line, either down or up a number of lines.
	MoveTo(x, y int) // moves the cursor to row y, column x
	Move(x, y int) // moves the cursor horizontally by x cells, vertically by y cells
	Column(n int) // moves the cursor to the column n of in the screen without changing the line number
	Alterate(enter bool) // enter (if true) alternate screen mode, or exit if false
*/
var Cursor cursor = cursor{}

func (cursor) Up(by int) string {
	return fmt.Sprintf(cursorUp, csi, by)
}
func (cursor) Down(by int) string {
	return fmt.Sprintf(cursorDown, csi, by)
}
func (cursor) Left(by int) string {
	return fmt.Sprintf(cursorLeft, csi, by)
}
func (cursor) Right(by int) string {
	return fmt.Sprintf(cursorRight, csi, by)
}
func (cursor) LinesUp(by int) string {
	return fmt.Sprintf(cursorLinePrevious, csi, by)
}
func (cursor) LinesDown(by int) string {
	return fmt.Sprintf(cursorLineNext, csi, by)
}
func (cursor) Column(by int) string {
	return fmt.Sprintf(cursorColumn, csi, by)
}
func (cursor) MoveTo(x, y int) string {
	return fmt.Sprintf(cursorMove, csi, y, x)
}
func (cursor) Move(x, y int) string {
	var sb strings.Builder
	if x < 0 {
		x = -x
		sb.WriteString(fmt.Sprintf(cursorLeft, csi, x))
	} else if x > 0 {
		sb.WriteString(fmt.Sprintf(cursorRight, csi, x))
	}
	if y < 0 {
		y = -y
		sb.WriteString(fmt.Sprintf(cursorUp, csi, y))
	} else if y > 0 {
		sb.WriteString(fmt.Sprintf(cursorDown, csi, y))
	}
	return sb.String()
}
