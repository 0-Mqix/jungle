package print

import "github.com/fatih/color"

type Alignment int

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
