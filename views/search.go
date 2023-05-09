package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func focusStart(focusable fyne.CanvasObject) *fyne.CanvasObject {
	return &focusable
}

func search(_ fyne.Window) fyne.CanvasObject {
	addButton := widget.NewButtonWithIcon("Add a task", theme.ContentAddIcon(), func() {})
	searchButton := widget.NewButtonWithIcon("", theme.SearchIcon(), func() {})

	inputText := binding.NewString()
	taskInput := widget.NewEntryWithData(inputText)
	taskInput.SetPlaceHolder("Search for a task")

	mLContainer := container.New(&myLayout{}, taskInput, searchButton, addButton)
	containerWrapper := container.NewVBox(mLContainer)
	vBox := layout.NewVBoxLayout()
	vBoxContainer := container.New(vBox, containerWrapper)
	taskInput.FocusGained()
	focusStart(taskInput)
	return vBoxContainer
}

type myLayout struct{}

func (mL *myLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	minHeight := objects[0].MinSize().Height
	for _, o := range objects {
		w += o.MinSize().Width
		if minHeight > h {
			h = o.MinSize().Height
		}
	}
	return fyne.NewSize(w, h)
}

func (mL *myLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	const PADDING = 4
	var containerWidth float32 = containerSize.Width * .70
	for i, o := range objects {
		if i == 0 {
			o.Resize(fyne.NewSize(containerWidth, o.MinSize().Height))
			continue
		}
		size := o.MinSize()
		o.Resize(size)
		o.Move(fyne.NewPos(containerWidth+PADDING, 0))
		containerWidth += size.Width + PADDING
	}
}
