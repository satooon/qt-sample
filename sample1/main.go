package main

import (
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Hello World Example")
	window.SetMinimumSize2(200, 200)

	layout := widgets.NewQVBoxLayout()

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("1. write something")
	layout.AddWidget(input, 0, 0)

	button := widgets.NewQPushButton2("2. click me", nil)
	button.ConnectClicked(func(checked bool) {
		widgets.QMessageBox_Information(nil, "OK", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
	layout.AddWidget(button, 0, 0)

	window.SetCentralWidget(widget)

	window.Show()

	widgets.QApplication_Exec()
}
