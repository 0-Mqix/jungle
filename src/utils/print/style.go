package print

import "github.com/fatih/color"

type Alignment int
type Padding int

type InferCenter bool
type InferLevel bool

const (
	LEFT Alignment = iota + 0
	RIGHT
	CENTER
)

var (
	DefaultPadding = 2
	DefaultColor   = color.FgHiGreen
)

type Style struct {
	Padding     int
	InferCenter bool
	Alignment   Alignment
}

func DefaultStyle() *Style {
	return &Style{Padding: DefaultPadding}
}

func (e *Element) GetStyle() *Style {
	if e.style != nil {
		return e.style
	}

	style := DefaultStyle()
	e.style = style
	return style
}

func (e *Element) Style(input ...interface{}) *Element {
	style = e.GetStyle()

	padding := 

	return e
}
