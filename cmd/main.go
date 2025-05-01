// cmd/gbb/main.go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var appVersion = "v0.1"

var statusColors = map[string]string{
	"red":    lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Render("Not ready"),
	"yellow": lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Render("Can build with changes"),
	"green":  lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Render("Ready"),
	"blue":   lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Render("Ready (Offline)"),
}

type menuItem string

func (i menuItem) Title() string       { return string(i) }
func (i menuItem) Description() string { return "" }
func (i menuItem) FilterValue() string { return string(i) }

type model struct {
	list       list.Model
	readyState string
	selected   string
}

func initialModel() model {
	items := []list.Item{
		menuItem("Check online for DosBox build plugins"),
		menuItem("Change install location"),
		menuItem("Change/specify build folder"),
		menuItem("Change builder (currently: CMake)"),
		menuItem("Options"),
		menuItem("\U0001F680 Build DosBox now"),
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "GBB Menu"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

	return model{
		list:       l,
		readyState: "green",
		selected:   "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.selected = m.list.SelectedItem().(menuItem).Title()
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height-4)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	now := time.Now().Format("Jan 2, 2006 15:04:05")
	status := statusColors[m.readyState]

	header := fmt.Sprintf("GBB %s — [%s] — %s\nBuild: DosBox | Method: CMake | Dir: ./dosbox-staging\n", appVersion, status, now)
	return header + m.list.View()
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	fmt.Println("You selected:", p.Model().(model).selected)
}
