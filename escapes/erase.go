package escapes

type erase struct{}

/*
# escape sequence string-returning func()s for erasing text in the terminal

	Erase.ToScreenEnd() // erases from the cursor's position until the end of the terminal screen
	Erase.ToScreenBegin() // erases from the cursor's position until the start of the terminal screen
	Erase.Screen() // erases all the text on the screen
	Erase.Line() // erases the text on the cursor's current line
	Erase.ToLineEnd() // erases the text on the cursor's current line from the cursor to the end
	Erase.ToLineBegin() // erases the text on the cursor's current line from the cursor back to the beginning of the line
*/
var Erase = erase{}

func (erase) ToScreenEnd() string {
	return eraseScreenToEnd
}

func (erase) ToScreenBegin() string {
	return eraseScreenToBegin
}

func (erase) Screen() string {
	return eraseAll
}

func (erase) Line() string {
	return eraseLine
}

func (erase) ToLineEnd() string {
	return eraseLineToEnd
}

func (erase) ToLineBegin() string {
	return eraseLineToBegin
}
