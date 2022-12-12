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

	root := Element{line: Create("JUNGLE", color.Bold, DefaultColor)}
	root.Style(CENTER)

	structs := root.Text("Struct Names:", color.Bold, DefaultColor).Style(Odd{color.FgBlack}, Even{color.FgWhite})

	structs.Text("1")
	structs.Text("2")
	structs.Text("3")

	top.Print()

	lines := make([]*Element, 0)
	root.Read(&lines)

	for _, e := range lines {
		for _, l := range e.line.Split() {
			e.Align(l).Print()
		}
	}

	bottom.Print()
}
