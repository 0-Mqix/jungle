package print

import (
	"github.com/fatih/color"
)

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
	InferLevel  bool
	Alignment   Alignment
}

func DefaultStyle() Style {
	return Style{Padding: DefaultPadding}
}

func (e *Element) Style(input ...interface{}) *Element {
	for _, v := range input {
		switch v := v.(type) {
		case Padding:
			e.style.Padding = int(v)
		case InferCenter:
			e.style.InferCenter = bool(v)
		case InferLevel:
			e.style.InferLevel = bool(v)
		case Alignment:
			e.style.Alignment = v
		}
	}

	return e
}

func InStyle(input []interface{}) Style {
	for _, v := range input {
		style, ok := v.(Style)

		if ok {
			return style
		}
	}

	return DefaultStyle()
}
