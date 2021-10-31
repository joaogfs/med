package main

import "strings"


type Buffer struct {
	data []Line
	cursorChar int
	cursorLine int
}

type Line []rune

func (b *Buffer) loadString(s string) {
	for _, strLine := range strings.Split(s, "\n") {
		b.data = append(b.data, Line(strLine))
	}

}

func (b *Buffer) String() string {
	str := ""

	for _, line := range b.data {
		str += string(line)
		str += "\n"
	}

	return str
}

var MainBuffer = Buffer{
	data:    []Line{},
	cursorChar: 0,
	cursorLine: 0,
}
