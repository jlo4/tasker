package views

import (
	"strings"

	"fyne.io/fyne/v2"
)

type View struct {
	Title        string
	View         func(w fyne.Window) fyne.CanvasObject
	CurrentState fyne.CanvasObject
	FocusStart   fyne.CanvasObject
}

var (
	Views = map[string]View{
		"home":   {"Home", home, nil, nil},
		"search": {"Search", search, nil, nil},
	}

	Cache = map[string]*View{}
)

func UpdateCache(w fyne.Window, v View) {
	Cache[strings.ToLower(v.Title)] = &View{
		Title:        v.Title,
		View:         v.View,
		CurrentState: v.View(w),
		FocusStart:   v.FocusStart,
	}
}
