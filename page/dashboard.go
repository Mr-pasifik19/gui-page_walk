package pages

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func ShowDashboardPage() {
	var mw *walk.MainWindow

	MainWindow{
		AssignTo: &mw,
		Title:    "Dashboard",
		Size:     Size{Width: 400, Height: 300},
		Layout:   VBox{},
		Children: []Widget{
			Label{Text: "Welcome to the Dashboard!"},
			PushButton{
				Text: "Logout",
				OnClicked: func() {
					mw.Close()
					ShowLoginPage()
				},
			},
		},
	}.Run()
}
