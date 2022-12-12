package print

import (
	"strings"
	"unicode/utf8"
)

type Element struct {
	level    int
	line     Line
	next     []*Element
	previous *Element
	style    *Style
}

func (e *Element) Text(line Line) *Element {
	new := &Element{
		previous: e,
		line:     line,
		level:    e.level + 1,
		next:     make([]*Element, 0),
	}
	e.next = append(e.next, new)
	return new
}

func (e *Element) Read(lines *[]*Element) {
	for _, child := range e.next {
		*lines = append(*lines, child)
		child.Read(lines)
	}
}

func (e *Element) Center(line Line, width int) Line {
	previous := line.String()

	if e.previous != nil && e.style.InferCenter {
		previous = e.previous.line.String()
	}

	size := (width - utf8.RuneCountInString(previous)) / 2

	if size < 1 {
		return line
	}

	line = Create(strings.Repeat(" ", size), line)
	size = (width - utf8.RuneCountInString(line.String()))
	line = Create(line, strings.Repeat(" ", size))

	return line
}

func (e *Element) Align() Line {
	line := Create(strings.Repeat(" ", e.level*2), e.line)

	style := DefaultStyle()

	if e.style != nil {
		style = e.style
	}

	padding := style.Padding
	align := style.Alignment

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
