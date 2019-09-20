package screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
)

func WaveformScreen() fyne.CanvasObject {
	return fyne.NewContainerWithLayout(layout.NewGridLayout(1), LayoutPanel())
}
