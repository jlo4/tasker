package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func home(_ fyne.Window) fyne.CanvasObject {
	homeLabel := widget.NewLabel("HOME CONTENT")

	vBox := layout.NewVBoxLayout()
	vBoxContainer := container.New(vBox, homeLabel)
	maxContainer := container.NewMax(vBoxContainer)
	return maxContainer
}
