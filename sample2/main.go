package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"os"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	quickcontrols2.QQuickStyle_SetStyle("material")

	view := qml.NewQQmlApplicationEngine(nil)

	view.Load(core.NewQUrl3("qrc:///qml/main.qml", 0))

	gui.QGuiApplication_Exec()
}
