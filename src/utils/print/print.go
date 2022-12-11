package print

import (
	"strings"

	"github.com/fatih/color"
)

var (
	UP_LEFT    = Create("┌")
	LOW_LEFT   = Create("└")
	UP_RIGHT   = Create("┐")
	LOW_RIGHT  = Create("┘")
	HORIZONTAL = Create("│")
	VERTICAL   = Create("─")
)

var (
	Width  = 30
	Margin = " "
)

func Wrap(wrapper, line Line) Line {
	return Create(wrapper, line, wrapper)
}

func Jungle() {
	top := Create(UP_LEFT, strings.Repeat(VERTICAL.String(), Width), UP_RIGHT).
		Style(DefaultColor)

	bottom := Create(LOW_LEFT, strings.Repeat(VERTICAL.String(), Width), LOW_RIGHT).
		Style(DefaultColor)

	root := Element{line: Create("test"), level: 0, style: &Style{Padding: 0, Alignment: CENTER}}
	root.Text(Create("ik ben ", "een aap", color.FgHiYellow, color.Bold))

	lines := []*Element{&root}
	root.Read(&lines)

	top.Print()

	for _, e := range lines {
		e.Align().Print()
	}

	bottom.Print()
}
