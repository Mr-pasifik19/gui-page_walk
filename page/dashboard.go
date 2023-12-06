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
			Label{
				Text: "Welcome to the Dashboard!",
				Font: Font{Family: "Arial", PointSize: 14, Bold: true}},

			PushButton{
				Text:       "Logout",
				Font:       Font{Family: "Arial", PointSize: 14, Bold: true},
				Background: SolidColorBrush{Color: walk.RGB(0, 0, 255)},
				OnClicked: func() {
					mw.Close()
					ShowLoginPage()
				},
			},
		},
	}.Run()
}
