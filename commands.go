package main

import "fmt"

type ArgList []interface{}

type Command int

const (
	ComQuit Command = iota
	ComSaveToFile
	ComLoadFile
	ComInsertRune
	ComLogMessage
	ComMoveCursor
)

func (c Command) Exec(args ...interface{}) {
	CommandAction[c](args)
}

var CommandAction = map[Command]func(ArgList) {
	ComQuit:       quit,
	ComSaveToFile: saveToFile,
	ComLoadFile:   loadFile,
	ComInsertRune: insertRune,
	ComLogMessage: logMessage,
	ComMoveCursor: moveCursor,
}

// no args
func quit(args ArgList) {
	editor.shouldQuit = true
}
// args[0] - filename string
func saveToFile(args ArgList) {}
// args[0] - filename string
func loadFile(args ArgList) {}
// args[0] - char rune
func insertRune(args ArgList) {

}
// args[0] - dir rune {'u', 'd', 'l', 'r'}
func moveCursor(args ArgList) {

}

func logMessage(args ArgList) {
	arg :=  fmt.Sprintf("%s\n", args[0])
	log(arg)
}
