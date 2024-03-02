package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/alcb1310/kanban/cmd/types"
)

func addUpdate(m Model, msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.input.Value() == "" {
				return m, nil
			}

			l := types.List{
				Title: m.input.Value(),
			}
			m.database.AddList(l)

			m.Position.Col = types.TODO
			m.Position.Row = 0
			m.Lists = initList(m.database)

			m.mode = LIST
		}
	}

	return m, nil
}

func addView(m Model) string {
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Left,
		lipgloss.Top,
		lipgloss.JoinVertical(
			lipgloss.Left,
			listHeader("Add Item"),
			m.input.View(),
		),
	)
}
