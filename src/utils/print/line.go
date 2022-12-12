package print

import (
	"fmt"

	"github.com/fatih/color"
)

type Line []Text

type Text struct {
	value string
	faith []color.Attribute
}

func Create(input ...interface{}) Line {
	return Convert(input)
}

func Convert(input []interface{}) Line {
	lenght := len(input)
	line := make(Line, 0)

	var creating bool
	var text Text

	reset := func() {
		if !creating {
			return
		}

		line = append(line, text)
		creating = false
		text = Text{}
	}

	for i := 0; i < lenght; i++ {
		array, ok := input[i].(Line)
		if ok {
			reset()
			line = append(line, array...)
			continue
		}

		value, ok := input[i].(string)
		if ok {
			reset()
			creating = true
			text = Text{value: value}
			continue
		}

		attribute, ok := input[i].(color.Attribute)
		if ok && creating {
			text.faith = append(text.faith, attribute)
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

func (line Line) Print() {
	for i, text := range line {
		color.New(text.faith...).Print(text.value)

		if i == len(line)-1 {
			fmt.Print("\n")
		}
	}

}

func (line Line) Cut(x int) (Line, int) {
	// start := len(line.String())
	size := len(line) - 1

	for i := size; i > 0; i-- {
		if x < 1 {
			return line, 0
		}
		// value := line[i].value

	}

	return line, 0
}
