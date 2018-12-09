package color

// Color: A simple color library for Go
// Reference Links:
//    http://ascii-table.com/ansi-escape-sequences-vt-100.php
//

import (
	"fmt"
)

const (
	// Attributes
	ATTR_RESET = iota
	ATTR_BOLD
	ATTR_FAINT
	ATTR_ITALIC
	ATTR_UNDERLINE
	ATTR_BLINKSLOW
	ATTR_BLINKRAPID
	ATTR_REVERSEVIDEO
	ATTR_CONCEALED
	ATTR_CROSSEDOUT
)

const (
	// 8 Colors
	BLACK = iota + 30
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
	RESET
)

const (
	// 16 Colors
	BRIGHT_BLACK = iota + 90
	BRIGHT_RED
	BRIGHT_GREEN
	BRIGHT_YELLOW
	BRIGHT_BLUE
	BRIGHT_MAGENTA
	BRIGHT_CYAN
	BRIGHT_WHITE
)

const (
	// Background Colors
	BACKGROUND_BLACK = iota + 40
	BACKGROUND_RED
	BACKGROUND_GREEN
	BACKGROUND_YELLOW
	BACKGROUND_BLUE
	BACKGROUND_MAGENTA
	BACKGROUND_CYAN
	BACKGROUND_WHITE
	BACKGROUND_RESET
)

const (
	// Background Hi-Intense Colors
	BACKGROUND_BRIGHT_BLACK = iota + 100
	BACKGROUND_BRIGHT_RED
	BACKGROUND_BRIGHT_GREEN
	BACKGROUND_BRIGHT_YELLOW
	BACKGROUND_BRIGHT_BLUE
	BACKGROUND_BRIGHT_MAGENTA
	BACKGROUND_BRIGHT_CYAN
	BACKGROUND_BRIGHT_WHITE
	BACKGROUND_BRIGHT_RESET
)

type colored struct {
	text string
	attr int
}

// Attributes
func Bold(text string, args ...interface{}) string {
	hued := &colored{text, ATTR_BOLD}
	return hued.Proc(text, args...)
}

func Faint(text string, args ...interface{}) string {
	hued := &colored{text, ATTR_FAINT}
	return hued.Proc(text, args...)
}

func Italic(text string, args ...interface{}) string {
	hued := &colored{text, ATTR_ITALIC}
	return hued.Proc(text, args...)
}

func Underline(text string, args ...interface{}) string {
	hued := &colored{text, ATTR_UNDERLINE}
	return hued.Proc(text, args...)
}

func BlinkSlow(text string, args ...interface{}) string {
	hued := &colored{text, ATTR_BLINKSLOW}
	return hued.Proc(text, args...)
}

func BlinkRapid(text string, args ...interface{}) string {
	hued := &colored{text, ATTR_BLINKRAPID}
	return hued.Proc(text, args...)
}

func ReverseVideo(text string, args ...interface{}) string {
	hued := &colored{text, ATTR_REVERSEVIDEO}
	return hued.Proc(text, args...)
}

func Concealed(text string, args ...interface{}) string {
	hued := &colored{text, ATTR_CONCEALED}
	return hued.Proc(text, args...)
}

func Crossedout(text string, args ...interface{}) string {
	hued := &colored{text, ATTR_CROSSEDOUT}
	return hued.Proc(text, args...)
}

// 8 Colors

func Black(text string, args ...interface{}) string {
	hued := &colored{text, BLACK}
	return hued.Proc(text, args...)
}

func Red(text string, args ...interface{}) string {
	hued := &colored{text, RED}
	return hued.Proc(text, args...)
}

func Green(text string, args ...interface{}) string {
	hued := &colored{text, GREEN}
	return hued.Proc(text, args...)
}

func Yellow(text string, args ...interface{}) string {
	hued := &colored{text, YELLOW}
	return hued.Proc(text, args...)
}

func Blue(text string, args ...interface{}) string {
	hued := &colored{text, BLUE}
	return hued.Proc(text, args...)
}

func Magenta(text string, args ...interface{}) string {
	hued := &colored{text, MAGENTA}
	return hued.Proc(text, args...)
}

func Cyan(text string, args ...interface{}) string {
	hued := &colored{text, CYAN}
	return hued.Proc(text, args...)
}

func White(text string, args ...interface{}) string {
	hued := &colored{text, WHITE}
	return hued.Proc(text, args...)
}

// 16 Colors

func BrightBlack(text string, args ...interface{}) string {
	hued := &colored{text, BRIGHT_BLACK}
	return hued.Proc(text, args...)
}

func BrightRed(text string, args ...interface{}) string {
	hued := &colored{text, BRIGHT_RED}
	return hued.Proc(text, args...)
}

func BrightGreen(text string, args ...interface{}) string {
	hued := &colored{text, BRIGHT_GREEN}
	return hued.Proc(text, args...)
}

func BrightYellow(text string, args ...interface{}) string {
	hued := &colored{text, BRIGHT_YELLOW}
	return hued.Proc(text, args...)
}

func BrightBlue(text string, args ...interface{}) string {
	hued := &colored{text, BRIGHT_BLUE}
	return hued.Proc(text, args...)
}

func BrightMagenta(text string, args ...interface{}) string {
	hued := &colored{text, BRIGHT_MAGENTA}
	return hued.Proc(text, args...)
}

func BrightCyan(text string, args ...interface{}) string {
	hued := &colored{text, BRIGHT_CYAN}
	return hued.Proc(text, args...)
}

func BrightWhite(text string, args ...interface{}) string {
	hued := &colored{text, BRIGHT_WHITE}
	return hued.Proc(text, args...)
}

// Background Colors

func BackgroundBlack(text string, args ...interface{}) string {
	hued := &colored{text, BACKGROUND_BLACK}
	return hued.Proc(text, args...)
}

func BackgroundRed(text string, args ...interface{}) string {
	hued := &colored{text, BACKGROUND_RED}
	return hued.Proc(text, args...)
}

func BackgroundGreen(text string, args ...interface{}) string {
	hued := &colored{text, BACKGROUND_GREEN}
	return hued.Proc(text, args...)
}

func BackgroundYellow(text string, args ...interface{}) string {
	hued := &colored{text, BACKGROUND_YELLOW}
	return hued.Proc(text, args...)
}

func BackgroundBlue(text string, args ...interface{}) string {
	hued := &colored{text, BACKGROUND_BLUE}
	return hued.Proc(text, args...)
}

func BackgroundMagenta(text string, args ...interface{}) string {
	hued := &colored{text, BACKGROUND_MAGENTA}
	return hued.Proc(text, args...)
}

func BackgroundCyan(text string, args ...interface{}) string {
	hued := &colored{text, BACKGROUND_CYAN}
	return hued.Proc(text, args...)
}

func BackgroundWhite(text string, args ...interface{}) string {
	hued := &colored{text, BACKGROUND_WHITE}
	return hued.Proc(text, args...)
}

// Bright Background Colors

func BackgroundBrightBlack(text string, args ...interface{}) string {
	hued := &colored{text, BACKGROUND_BRIGHT_BLACK}
	return hued.Proc(text, args...)
}

func BackgroundBrightRed(text string, args ...interface{}) string {
	hued := &colored{text, BACKGROUND_BRIGHT_RED}
	return hued.Proc(text, args...)
}

func BackgroundBrightGreen(text string, args ...interface{}) string {
	hued := &colored{text, BACKGROUND_BRIGHT_GREEN}
	return hued.Proc(text, args...)
}

func BackgroundBrightYellow(text string, args ...interface{}) string {
	hued := &colored{text, BACKGROUND_BRIGHT_YELLOW}
	return hued.Proc(text, args...)
}

func BackgroundBrightBlue(text string, args ...interface{}) string {
	hued := &colored{text, BACKGROUND_BRIGHT_BLUE}
	return hued.Proc(text, args...)
}

func BackgroundBrightMagenta(text string, args ...interface{}) string {
	hued := &colored{text, BACKGROUND_BRIGHT_MAGENTA}
	return hued.Proc(text, args...)
}

func BackgroundBrightCyan(text string, args ...interface{}) string {
	hued := &colored{text, BACKGROUND_BRIGHT_CYAN}
	return hued.Proc(text, args...)
}

func BackgroundBrightWhite(text string, args ...interface{}) string {
	hued := &colored{text, BACKGROUND_BRIGHT_WHITE}
	return hued.Proc(text, args...)
}

// Color it!
func (hue *colored) Proc(text string, args ...interface{}) string {
	hue.text = fmt.Sprintf(text, args...)
	return fmt.Sprintf("\u001b[%dm%s\u001b[0m", hue.attr, hue.text)
}
