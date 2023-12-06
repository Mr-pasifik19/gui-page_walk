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
					Label{Text: "Username:"},
					LineEdit{
						AssignTo: &inTE,
						Name:     "Username",
						MinSize:  Size{Width: 200},
						Text:     Bind("Username"),
					},
				},
			},
			HSplitter{
				Children: []Widget{
					Label{Text: "Password:"},
					LineEdit{
						AssignTo:     &outTE,
						Name:         "Password",
						MinSize:      Size{Width: 200},
						Text:         Bind("Password"),
						PasswordMode: true,
					},
				},
			},
			PushButton{
				Text: "Login",
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
