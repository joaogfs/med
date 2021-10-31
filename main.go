package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

var debugString string
var editor = InitEditor()
var s tcell.Screen = InitScreen()

func main() {
	defer s.Fini()

	content, _ := os.ReadFile("./dic-master/pt_BR.dic")

	MainBuffer.loadString(string(content))

	Render()

	for !editor.shouldQuit {
		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			processKeypress(eventKeyToKeypress(*ev))
			log(fmt.Sprintf("%+v\n", eventKeyToKeypress(*ev)))
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventError:
			panic(ev.Error())
		}

		Render()
	}
}

func log(s string) {
	debugString += s
	writeLogFile()
}

func writeLogFile() {
	file, _ := os.Create("./log.txt")
	file.WriteString(debugString)
	file.Close()
}
