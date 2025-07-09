package app

import (
	"log"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	myssh "github.com/nikaydo/ssh-client/internal/ssh"
)

type Tab struct {
	Tab       *container.TabItem
	Session   myssh.Session
	CmdButton *widget.Entry
	Grid      *widget.TextGrid
	Scroll    *container.Scroll
}

type TabsCmd struct {
	DocTabs *container.DocTabs
	Items   map[int]Tab
}

func InitTabs() (tabs TabsCmd) {
	tabs.DocTabs = container.NewDocTabs()
	tabs.Items = make(map[int]Tab)
	return
}

func (t *TabsCmd) Add(text string, session *myssh.Session) (id int, grid *widget.TextGrid, scroll *container.Scroll) {
	id = len(t.Items) + 1
	grid = widget.NewTextGrid()
	grid.Hide()
	scroll = container.NewScroll(container.NewStack(grid))
	b := widget.NewEntry()
	b.Hide()
	b.OnSubmitted = func(s string) {
		if session.StdinPipe != nil {
			_, err := session.StdinPipe.Write([]byte(s + "\n"))
			if err != nil {
				log.Println(err)
			}
		}
		b.SetText("")
		scroll.ScrollToBottom()
	}
	session_cmd := container.NewBorder(nil, b, nil, nil, scroll)
	t.Items[id] = Tab{
		Tab:       container.NewTabItem(text, session_cmd),
		Session:   *session,
		CmdButton: b,
		Grid:      grid,
		Scroll:    scroll,
	}
	return
}

func (t *TabsCmd) ShowAll() {
	for _, i := range t.Items {
		i.Grid.Show()
		i.Scroll.Show()
		i.CmdButton.Show()
	}
}

func (t *TabsCmd) HideAll() {
	for _, i := range t.Items {
		i.Grid.Hide()
		i.Scroll.Hide()
		i.CmdButton.Hide()
	}
}

func (t *TabsCmd) AppendDocTab() {
	for _, item := range t.Items {
		t.DocTabs.Append(item.Tab)
	}
}
