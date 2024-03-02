package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/alcb1310/kanban/cmd/database"
	"github.com/alcb1310/kanban/cmd/types"
)

const (
	LIST = iota
	ADD
)

type Position struct {
	Row, Col uint
}

type Model struct {
	Lists         map[int][]types.List
	Position      Position
	mode          int
	database      database.Database
	input         textinput.Model
	width, height int
}

func App() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initialModel() Model {
	p := Position{
		Row: 0,
		Col: 0,
	}

	db := database.Connect()

	textInput := textinput.New()
	textInput.Placeholder = "Enter a new item"
	textInput.Prompt = "> "
	textInput.Focus()

	return Model{
		Lists:    initList(db),
		Position: p,
		mode:     LIST,
		database: db,
		input:    textInput,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.mode {
	case LIST:
		return listUpdate(m, msg)

	case ADD:
		m.input, _ = m.input.Update(msg)
		return addUpdate(m, msg)

	default:
		return m, nil
	}

}

func (m Model) View() string {
	m.Lists = initList(m.database)

	switch m.mode {
	case LIST:
		return displayItems(m)

	case ADD:
		return addView(m)

	default:
		return ""
	}
}
