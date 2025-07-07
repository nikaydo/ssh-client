package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type TabsCmd struct {
	Tabs *container.DocTabs
}

func InitTabs() (tabs TabsCmd) {
	tabs.Tabs = container.NewDocTabs()
	return
}

func (t *TabsCmd) Add(text string, content fyne.CanvasObject) {
	t.Tabs.Append(container.NewTabItem(text, content))
}
