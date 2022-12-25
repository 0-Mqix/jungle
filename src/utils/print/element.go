package print

import (
	"strings"
	"unicode/utf8"
)

type Element struct {
	level     int
	line      Line
	childeren []*Element
	parent    *Element
	style     Style
}

func (e *Element) Text(input ...interface{}) *Element {
	style := InStyle(input)
	line := Convert(input)

	new := &Element{
		parent:    e,
		line:      line,
		childeren: make([]*Element, 0),
		style:     style,
		level:     e.level + 1,
	}
	e.childeren = append(e.childeren, new)
	return new

}

func (e *Element) Read(lines *[]*Element) {
	if len(*lines) == 0 {
		*lines = append(*lines, e)
	}

	for _, child := range e.childeren {
		*lines = append(*lines, child)
		child.Read(lines)
	}
}

func (e *Element) Print() {
	if e.parent == nil {
		for _, l := range e.line.Split() {
			e.Align(l).Print()
		}
	}

	style := e.style

	for i, child := range e.childeren {
		if len(style.Even) > 0 && i%2 == 0 {
			child.line.Style(style.Even...)
		}

		if len(style.Odd) > 0 && i%2 != 0 {
			child.line.Style(style.Odd...)
		}

		for _, l := range child.line.Split() {
			child.Align(l).Print()
		}

		child.Print()
	}
}

func (e *Element) Center(line Line, width int) Line {
	previous := Size(line)

	if e.parent != nil && e.style.InferCenter {
		previous = Size(e.parent.line.Split()...)
	}

	size := (width - previous) / 2

	if size < 1 {
		return line
	}

	line = Create(strings.Repeat(" ", size), line)
	size = (width - utf8.RuneCountInString(line.String()))
	line = Create(line, strings.Repeat(" ", size))

	return line
}

func (e *Element) Align(line Line) Line {

	if e.style.InferLevel {
		line = Create(strings.Repeat(" ", e.level*2), line)
	}

	padding := e.style.Padding
	align := e.style.Alignment

	space := Width - len(line.String()) - padding*2

	switch align {
	case CENTER:
		line = e.Center(line, Width)
	case LEFT:
		line = Create(line, strings.Repeat(" ", space))
		line = Wrap(line, strings.Repeat(" ", padding))

	case RIGHT:
		line = Create(strings.Repeat(" ", space), line)
		line = Wrap(line, strings.Repeat(" ", padding))
	}

	return Wrap(line, Horizontal.Style(DefaultColor))
}
