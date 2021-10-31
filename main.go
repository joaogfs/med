package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

var debugString string

var editor = InitEditor()

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	defer s.Fini()


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
