package pages

import (
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func ShowLoginPage() {
	var inTE, outTE *walk.LineEdit
	var mw *walk.MainWindow

	MainWindow{
		AssignTo: &mw,
		Title:    "Login",
		Size:     Size{Width: 300, Height: 150},
		Layout:   VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					Label{
						Font: Font{PointSize: 12},
						Text: "Username:"},
					LineEdit{
						AssignTo: &inTE,
						Name:     "Username",
						Font:     Font{PointSize: 12},
						MinSize:  Size{Width: 200, Height: 35},
						MaxSize:  Size{Width: 300, Height: 35},
						Text:     Bind("Username"),
					},
				},
			},
			HSplitter{
				Children: []Widget{
					Label{
						Font: Font{PointSize: 12},
						Text: "Password:"},

					LineEdit{
						AssignTo:     &outTE,
						Name:         "Password",
						Font:         Font{PointSize: 12},
						MinSize:      Size{Width: 200, Height: 35},
						MaxSize:      Size{Width: 300, Height: 35},
						Text:         Bind("Password"),
						PasswordMode: true,
					},
				},
			},
			PushButton{
				Text:       "Login",
				Font:       Font{PointSize: 14, Bold: true},
				Background: SolidColorBrush{Color: walk.RGB(0, 0, 255)},
				OnClicked: func() {
					if strings.ToUpper(inTE.Text()) == "USERNAME" && outTE.Text() == "PASSWORD" {
						mw.Close()
						ShowDashboardPage()
					} else {
						walk.MsgBox(mw, "Login Failed", "Invalid credentials", walk.MsgBoxOK|walk.MsgBoxIconError)
					}
				},
			},
		},
	}.Run()
}
