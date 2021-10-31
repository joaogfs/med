package main

import "github.com/gdamore/tcell/v2"

func InitScreen() tcell.Screen {
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}

	return s
}

func Render() {
	renderString(MainBuffer.String(), 0, 0)
	s.Show()
	s.ShowCursor(MainBuffer.cursorChar, MainBuffer.cursorLine)
}

func renderString(str string, origX int, origY int) {
	var charCount = 0
	for _, char := range str {
		if char == '\n' {
			origY++
			origX = 0
			charCount = 0
			continue
		}

		s.SetContent(origX+charCount, origY, char, []rune{}, tcell.StyleDefault)
		charCount++
	}
}
