package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
)

type Action struct {
	c    Command
	args ArgList
}

type Keypress struct {
	ctrl bool
	alt  bool
	char rune // 0 if key != tcell.KeyRune
	key  tcell.Key // tcell.KeyRune if char != 0
}

// a more visually clean way of representing Keypress
type KeyLiteral string

var GlobalKeymap = map[Keypress]Command{
	keyLiteral("ctrl-q"): ComQuit,
	keyLiteral("ctrl-alt-p"): ComLogMessage,
}

func processKeypress(kp Keypress) {
	var args ArgList
	if comVal, ok := GlobalKeymap[kp]; ok {
		switch comVal {
		case ComInsertRune:
			args = ArgList{kp.char}
		case ComLogMessage:
			args = ArgList{"firmão"}
		}

		comVal.Exec(args)
	}
}

// Converts tcell.EventKey to Keypress
func eventKeyToKeypress(ev tcell.EventKey) Keypress {
	var char rune
	var key tcell.Key
	if ev.Key() == tcell.KeyRune {
		char = ev.Rune()
		key = tcell.KeyRune
	} else {
		nameParts := strings.Split(ev.Name(), "+")
		lastItem := nameParts[len(nameParts)-1]
		if len(lastItem) == 1 {
			lastItem = strings.ToLower(lastItem)
			char = []rune(lastItem)[0]
			key = tcell.KeyRune
		} else {
			char = 0
			key = ev.Key()
		}
	}
	return Keypress{
		ctrl: tcell.ModCtrl & ev.Modifiers() == tcell.ModCtrl,
		alt:  tcell.ModAlt  & ev.Modifiers() == tcell.ModAlt,
		char: char,
		key:  key,
	}
}

// Converts KeyLiteral to Keypress
func keyLiteral(key string) Keypress {
	rv := Keypress{
		false,
		false,
		0,
		tcell.KeyNUL,
	}

	for _, elmt := range strings.Split(key, "-") {
		if elmt == "ctrl" {
			rv.ctrl = true
		} else if elmt == "alt" {
			rv.alt = true
		} else if len(elmt) == 1 {
			rv.char = rune([]rune(elmt)[0])
			rv.key = tcell.KeyRune
		} else {
			for key, value := range tcell.KeyNames {
				if elmt == value {
					rv.key = key
				}
			}
		}
	}

	return rv
}
