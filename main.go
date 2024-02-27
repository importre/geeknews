package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"importre.com/geeknews/programs/board"
)

func main() {
	p := tea.NewProgram(
		Model{board: board.New()},
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type Model struct {
	board board.Model
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.board.Init("", 1),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.board, cmd = m.board.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.board.View()
}
