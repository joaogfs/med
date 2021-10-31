package main

type EditorState struct {
	shouldQuit bool
}

func InitEditor() EditorState {
	return EditorState{shouldQuit: false}
}
