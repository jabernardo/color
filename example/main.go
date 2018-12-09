package main

import (
	"fmt"

	".."
)

func main() {
	// Attributes
	fmt.Println("This is", color.Bold("Bold"))
	fmt.Println("This is", color.Faint("Faint"))
	fmt.Println("This is", color.Italic("Italic"))
	fmt.Println("This is", color.Underline("Underline"))
	fmt.Println("This is", color.BlinkSlow("Blink Slow"))
	fmt.Println("This is", color.BlinkRapid("Blink Rapid"))
	fmt.Println("This is", color.ReverseVideo("Reverse Video"))
	fmt.Println("This is", color.Concealed("Concealed"))
	fmt.Println("This is", color.Crossedout("Crossedout"))

	// 8 Colors
	fmt.Println("This is", color.Black("Black"))
	fmt.Println("This is", color.Red("Red"))
	fmt.Println("This is", color.Green("Green"))
	fmt.Println("This is", color.Yellow("Yellow"))
	fmt.Println("This is", color.Blue("Blue"))
	fmt.Println("This is", color.Magenta("Magenta"))
	fmt.Println("This is", color.Cyan("Cyan"))
	fmt.Println("This is", color.White("White"))

	// 16 Colors
	fmt.Println("This is", color.BrightBlack("Bright Black"))
	fmt.Println("This is", color.BrightRed("Bright Red"))
	fmt.Println("This is", color.BrightGreen("Bright Green"))
	fmt.Println("This is", color.BrightYellow("Bright Yellow"))
	fmt.Println("This is", color.BrightBlue("Bright Blue"))
	fmt.Println("This is", color.BrightMagenta("Bright Magenta"))
	fmt.Println("This is", color.BrightCyan("Bright Cyan"))
	fmt.Println("This is", color.BrightWhite("Bright White"))

	// Background Colors
	fmt.Println("This is", color.BackgroundBlack("Black Background"))
	fmt.Println("This is", color.BackgroundRed("Black Red"))
	fmt.Println("This is", color.BackgroundGreen("Black Green"))
	fmt.Println("This is", color.BackgroundYellow("Black Yellow"))
	fmt.Println("This is", color.BackgroundBlue("Black Blue"))
	fmt.Println("This is", color.BackgroundMagenta("Black Magenta"))
	fmt.Println("This is", color.BackgroundCyan("Black Cyan"))
	fmt.Println("This is", color.BackgroundWhite("Black White"))

	fmt.Println("This is", color.BackgroundBrightBlack("Black Bright Background"))
	fmt.Println("This is", color.BackgroundBrightRed("Black Bright Red"))
	fmt.Println("This is", color.BackgroundBrightGreen("Black Bright Green"))
	fmt.Println("This is", color.BackgroundBrightYellow("Black Bright Yellow"))
	fmt.Println("This is", color.BackgroundBrightBlue("Black Bright Blue"))
	fmt.Println("This is", color.BackgroundBrightMagenta("Black Bright Magenta"))
	fmt.Println("This is", color.BackgroundBrightCyan("Black Bright Cyan"))
	fmt.Println("This is", color.BackgroundBrightWhite("Black Bright White"))
}
