package escapes

type background struct{}

// # escape sequence string-returning func()s for setting terminal background colors
//
// Reset is the same as Style.Reset()
var Background = background{}

func (background) Reset() string {
	return reset
}

func (background) Black() string {
	return bgBlack
}

func (background) Red() string {
	return bgRed
}

func (background) Green() string {
	return bgGreen
}

func (background) Yellow() string {
	return bgYellow
}

func (background) Blue() string {
	return bgBlue
}

func (background) Magenta() string {
	return bgMagenta
}

func (background) Cyan() string {
	return bgCyan
}

func (background) White() string {
	return bgWhite
}
