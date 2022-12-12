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

	root := Element{line: Create("test")}
	root.Style(CENTER)

	root.Text("lets\ngo monkey \nthis is \n(also inline)", color.Bold).Style(InferCenter(true))
	root.Text("this is multi\nline ", color.Bold, "max", color.FgHiBlue).Style(InferCenter(true), CENTER)
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
