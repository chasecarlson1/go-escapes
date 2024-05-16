package escapes

import "fmt"

type style struct{}

/*
# escape sequences for styling.

...note that Reset resets all of these AND COLORING

	Style.Font(0) // resets the font to the terminal default, 1-9 are alternative fonts (if any)
*/
var Style = style{}

func (style) Reset() string {
	return reset
}

func (style) Bold(on bool) string {
	if on {
		return bold
	} else {
		return boldOff
	}
}

func (style) Light(on bool) string {
	if on {
		return light
	} else {
		return boldOff
	}
}

func (style) Italic(on bool) string {
	if on {
		return italic
	} else {
		return italicOff
	}
}

func (style) Blink(on bool) string {
	if on {
		return blink
	} else {
		return blinkOff
	}
}

func (style) Reverse(on bool) string {
	if on {
		return reversed
	} else {
		return reversedOff
	}
}

func (style) Hide(on bool) string {
	if on {
		return hidden
	} else {
		return hiddenOff
	}
}

func (style) Font(index int) string {
	if index > 9 || index < 0 {
		index = 0
	}
	return fmt.Sprintf(font, csi, index)
}
