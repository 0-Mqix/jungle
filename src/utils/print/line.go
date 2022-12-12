package print

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Line []Text

type Text struct {
	value string
	faith []color.Attribute
}

func CreateText(s string, faith []color.Attribute) Text { return Text{value: s, faith: faith} }

func Create(input ...interface{}) Line {
	return Convert(input)
}

func Convert(input []interface{}) Line {
	var creating bool
	var text Text
	var line Line

	reset := func() {
		if !creating {
			return
		}

		line = append(line, text)
		creating = false
		text = Text{}
	}

	for _, v := range input {
		switch v := v.(type) {

		case Line:
			reset()
			line = append(line, v...)

		case string:
			reset()
			creating = true
			text = Text{value: v}

		case color.Attribute:
			if creating {
				text.faith = append(text.faith, v)
			}
		}
	}

	if text.value != "" {
		line = append(line, text)
	}

	return line
}

func (line Line) Style(atributes ...color.Attribute) Line {
	for i, text := range line {
		line[i].faith = append(text.faith, atributes...)
	}
	return line
}

func (line Line) String() (s string) {
	for _, text := range line {
		s += text.value
	}

	return
}

func (line Line) Split() []Line {
	result := make([]Line, 0)
	new := make(Line, 0)

	for _, text := range line {
		parts := strings.Split(text.value, "\n")
		size := len(parts)

		if size > 1 {
			new = append(new, CreateText(parts[0], text.faith))
			result = append(result, new)

			for i := 1; i < size-1; i++ {
				result = append(result, Line{CreateText(parts[i], text.faith)})
			}

			new = Line{CreateText(parts[size-1], text.faith)}
			continue
		}

		new = append(new, text)
	}

	result = append(result, new)

	return result
}

func (line Line) Print() {
	for i, text := range line {
		color.New(text.faith...).Print(text.value)

		if i == len(line)-1 {
			fmt.Print("\n")
		}
	}
}

func Wrap(line Line, input ...interface{}) Line {
	wrapper := Convert(input)
	return Create(wrapper, line, wrapper)
}
