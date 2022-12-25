package print

import (
	"strings"

	"github.com/fatih/color"
)

var (
	TopLeft     = Create("┌")
	BottomLeft  = Create("└")
	TopRight    = Create("┐")
	BottomRight = Create("┘")
	Horizontal  = Create("│")
	Vertical    = Create("─")
)

var (
	Width  = 60
	Margin = " "
)

func Jungle() {
	top := Create(TopLeft, strings.Repeat(Vertical.String(), Width), TopRight).
		Style(DefaultColor)

	bottom := Create(BottomLeft, strings.Repeat(Vertical.String(), Width), BottomRight).
		Style(DefaultColor)

	root := Element{line: Create("\nJUNGLE\n", color.Bold, DefaultColor)}
	root.Style(CENTER)

	structs := root.Text("Struct Names:", color.Bold, DefaultColor).
		Style(Even{color.FgBlack}, Odd{color.FgWhite})

	structs.Text("1")
	structs.Text("2")
	structs.Text("3")
	structs.Text("4")
	structs.Text("5")
	structs.Text("6")

	root.Text()

	methods := root.Text("Comment Methods:", color.Bold, DefaultColor).
		Style(Even{color.FgBlack}, Odd{color.FgWhite})

	methods.Text("1")
	methods.Text("2")
	methods.Text("3")

	root.Text()

	top.Print()
	root.Print()
	bottom.Print()
}
