package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/alcb1310/kanban/cmd/types"
)

const (
	LIST = iota
)

type Position struct {
	Row, Col uint
}

type Model struct {
	Lists    map[int][]types.List
	Position Position
}

func App() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initialModel() Model {
	return Model{
		Lists: initList(),
		Position: Position{
			Row: 0,
			Col: 0,
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return listUpdate(m, msg)
}

func (m Model) View() string {
	return displayItems(m)
}
