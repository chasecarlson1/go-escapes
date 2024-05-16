package escapes

type color struct{}

// # escape sequence func()s for setting text/foreground colors
//
// Reset is the same as Style.Reset()
var Color = color{}

func (color) Reset() string {
	return reset
}

func (color) Black() string {
	return black
}

func (color) Red() string {
	return red
}

func (color) Green() string {
	return green
}

func (color) Yellow() string {
	return yellow
}

func (color) Blue() string {
	return blue
}

func (color) Magenta() string {
	return magenta
}

func (color) Cyan() string {
	return cyan
}

func (color) White() string {
	return white
}
