package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/alcb1310/kanban/cmd/types"
)

const (
	width    = 96
	colWidth = 50
)

var (
	subtle     = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight  = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#383838"}
	special    = lipgloss.AdaptiveColor{Light: "$43BF6D", Dark: "#73F59F"}
	greenTitle = lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}

	divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(subtle).
		String()

	titleStyle = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Foreground(highlight).
			Bold(true).
			MarginTop(1).
			MarginBottom(1).
			Render

	list = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, true).
		BorderForeground(special).
		MarginRight(2).
		MarginLeft(2).
		Padding(1, 1).
		Height(8).
		Width(colWidth + 1)

	listHeader = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			Foreground(greenTitle).
			BorderBottom(true).
			BorderForeground(greenTitle).
			MarginRight(2).
			MarginBottom(1).
			Render

	listText = lipgloss.NewStyle().
			Foreground(special).
			Render

	listItem = lipgloss.NewStyle().SetString("○ ").
			Foreground(lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}).
			PaddingRight(1).
			String()

	checkMark = lipgloss.NewStyle().SetString("● ").
			Foreground(lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}).
			PaddingRight(1).
			String()

	listDone = lipgloss.NewStyle().
			Foreground(subtle).
			Strikethrough(true).
			Render

	historyStyle = lipgloss.NewStyle().
			Align(lipgloss.Left).
			Border(lipgloss.DoubleBorder(), false, true, false, true).
			BorderForeground(lipgloss.Color("69")).
			Margin(1, 3, 0, 0).
			Padding(1, 2).
			Width(colWidth)
)

func listUpdate(m Model, msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.Position.Row == 0 {
				m.Position.Row = uint(len(m.Lists[int(m.Position.Col)]) - 1)
			} else {
				m.Position.Row--
			}

		case "down", "j":
			if m.Position.Row == uint(len(m.Lists[int(m.Position.Col)])-1) {
				m.Position.Row = 0
			} else {
				m.Position.Row++
			}

		case "left", "h":
			m.Position.Row = 0
			if m.Position.Col == types.TODO {
				m.Position.Col = types.DONE
			} else {
				m.Position.Col--
			}

		case "right", "l":
			m.Position.Row = 0
			if m.Position.Col == types.DONE {
				m.Position.Col = types.TODO
			} else {
				m.Position.Col++
			}

		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func displayItems(m Model) string {
	t := listHeader("TODO Items")
	t += "\n"

	for i, list := range m.Lists[types.TODO] {
		c := listItem
		if m.Position.Col == types.TODO && m.Position.Row == uint(i) {
			c = checkMark
		}

		t += c + listText(list.Title) + "\n"
	}

	i := listHeader("INPROGRESS Items")
	i += "\n"

	for k, list := range m.Lists[types.INPROGRESS] {
		c := listItem
		if m.Position.Col == types.INPROGRESS && m.Position.Row == uint(k) {
			c = checkMark
		}

		i += c + listText(list.Title) + "\n"
	}

	d := listHeader("DONE Items")
	d += "\n"

	for i, list := range m.Lists[types.DONE] {
		c := listItem
		if m.Position.Col == types.DONE && m.Position.Row == uint(i) {
			c = checkMark
		}

		d += c + listDone(list.Title) + "\n"
	}

	title := titleStyle(strings.ToUpper("Kanban Board"))

	s := lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			list.Copy().Render(t),
			list.Copy().Render(i),
			list.Copy().Render(d),
		),
	)
	return s
}
