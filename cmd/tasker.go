package tasker

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jlo4/tasker/views"
)

const HOME = "Home"
const SEARCH = "Search"

func Show() {
	taskerApp := app.New()
	mainWindow := taskerApp.NewWindow("Tasker")
	mainWindow.SetMaster()

	maxContainer := container.NewMax()
	setView := func(v views.View) {
		if views.Cache[strings.ToLower(v.Title)] == nil {
			views.UpdateCache(mainWindow, v)
		}
		maxContainer.Objects = []fyne.CanvasObject{views.Cache[strings.ToLower(v.Title)].CurrentState}
		maxContainer.Refresh()
	}
	setView(views.Views[strings.ToLower(HOME)])
	sideBarTree := createSideBarTree(setView)
	sideBarBorderContainer := container.NewHBox(sideBarTree)

	allContainers := container.NewBorder(nil, nil, sideBarBorderContainer, nil, maxContainer)

	mainWindow.Resize(fyne.NewSize(800, 400))
	mainWindow.SetContent(allContainers)

	mainWindow.ShowAndRun()
}

func createSideBarTree(setView func(v views.View)) *widget.Tree {

	searchItems := map[string][]string{
		"": {HOME, SEARCH},
	}
	tree := widget.NewTreeWithStrings(searchItems)
	tree.Select(HOME)
	tree.OnUnselected = func(id string) {
		if v, ok := views.Views[strings.ToLower(id)]; ok {
			setView(v)
		}
	}
	tree.OnSelected = func(id string) {
		if v, ok := views.Views[strings.ToLower(id)]; ok {
			setView(v)
		}
	}
	return tree
}
