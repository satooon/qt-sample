package main

import (
	"github.com/therecipe/qt/widgets"
	"os"
	"github.com/therecipe/qt/webengine"
	"github.com/therecipe/qt/core"
	"github.com/ashwanthkumar/slack-go-webhook"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("sample3")
	window.SetMinimumSize2(520, 480)

	widget := widgets.NewQWidget(nil, 0)

	vLayout := widgets.NewQVBoxLayout()
	widget.SetLayout(vLayout)

	var webView = webengine.NewQWebEngineView(nil)
	webView.SetUrl(core.NewQUrl3("https://slack.com/", 0))
	vLayout.AddWidget(webView, 0, 0)

	inputIcon := widgets.NewQLineEdit(nil)
	inputIcon.SetPlaceholderText("icon url")
	vLayout.AddWidget(inputIcon, 0, 0)

	inputTxt := widgets.NewQLineEdit(nil)
	inputTxt.SetPlaceholderText("message")
	vLayout.AddWidget(inputTxt, 0, 0)

	inputUser := widgets.NewQLineEdit(nil)
	inputUser.SetPlaceholderText("user")
	vLayout.AddWidget(inputUser, 0, 0)

	inputCh := widgets.NewQLineEdit(nil)
	inputCh.SetPlaceholderText("channel")
	vLayout.AddWidget(inputCh, 0, 0)

	inputHook := widgets.NewQLineEdit(nil)
	inputHook.SetPlaceholderText("hook")
	vLayout.AddWidget(inputHook, 0, 0)

	var button = widgets.NewQPushButton2("send", nil)
	button.ConnectClicked(func(flag bool) {
		payload := slack.Payload{
			IconUrl: inputIcon.Text(),
			Text:     inputTxt.Text(),
			Username: inputUser.Text(),
			Channel:  inputCh.Text(),
		}
		slack.Send(inputHook.Text(), "", payload)
	})
	vLayout.AddWidget(button, 0, 0)

	window.SetCentralWidget(widget)
	window.Show()
	widgets.QApplication_Exec()
}
