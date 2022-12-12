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
	Width  = 45
	Margin = " "
)

func Wrap(line Line, input ...interface{}) Line {
	wrapper := Convert(input)
	return Create(wrapper, line, wrapper)
}

func Jungle() {
	top := Create(TopLeft, strings.Repeat(Vertical.String(), Width), TopRight).
		Style(DefaultColor)

	bottom := Create(BottomLeft, strings.Repeat(Vertical.String(), Width), BottomRight).
		Style(DefaultColor)

	root := Element{line: Create("TEST", color.Bold), level: 0, style: &Style{Padding: 0, Alignment: CENTER}}
	// test.GetStyle().Alignment = CENTER
	test := root.Text(Create("ik ben een ", "aap", color.FgHiYellow, color.Bold))
	vis := root.Text(Create("ik ben een ", "vis", color.FgHiBlue, color.Bold))
	vis.GetStyle().Alignment = RIGHT

	test.Text(Create("hallo ", color.FgWhite, "thom", color.FgRed, color.Bold))
	// thom.GetStyle().Alignment = CENTER
	test.Text(Create("hallo ", color.FgWhite, "max", color.FgHiMagenta, color.Bold))

	wow := root.Text(Create("wow", color.FgBlack))
	wow.GetStyle().Padding = 0

	t1 := root.Text(Create("test", color.Bold, color.FgBlack))
	t1.GetStyle().Alignment = LEFT
	t1.GetStyle().Padding = 0

	// thom.GetStyle().InferCenter = true
	lines := []*Element{&root}
	root.Read(&lines)

	top.Print()

	for _, e := range lines {
		e.Align().Print()
	}

	bottom.Print()
}
