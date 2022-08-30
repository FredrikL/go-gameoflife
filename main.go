package main

import (
	"coderetreat/game"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	board      map[game.Pos]bool
	generation int
}

func main() {
	board := make(map[game.Pos]bool)
	// Beacon
	board[game.Pos{X: 1, Y: 1}] = true
	board[game.Pos{X: 1, Y: 2}] = true
	board[game.Pos{X: 2, Y: 1}] = true
	board[game.Pos{X: 2, Y: 2}] = true
	board[game.Pos{X: 3, Y: 3}] = true
	board[game.Pos{X: 3, Y: 4}] = true
	board[game.Pos{X: 4, Y: 3}] = true
	board[game.Pos{X: 4, Y: 4}] = true

	// Blinker
	board[game.Pos{X: 6, Y: 6}] = true
	board[game.Pos{X: 7, Y: 6}] = true
	board[game.Pos{X: 8, Y: 6}] = true

	p := tea.NewProgram(model{board: board}, tea.WithAltScreen(), tea.WithMouseAllMotion())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch()
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit

		case "enter":
			b := game.Advance(m.board)
			return model{board: b, generation: (m.generation + 1)}, nil
		}

	case tea.MouseMsg:
		e := tea.MouseEvent(msg)
		if e.Type == tea.MouseLeft {
			m.board[game.Pos{X: e.X, Y: e.Y}] = true
		}
	}

	return m, nil
}

func (m model) View() string {
	str := ""
	for i := 1; i <= 40; i++ {
		for j := 1; j <= 80; j++ {
			if (m.board[game.Pos{X: j, Y: i}]) {
				str += "X"
			} else {
				str += " "
			}
		}
		str += "\n"
	}
	str += "\n"
	str += "[enter] to advance, [q] to quit.\n"
	str += "Use mouse to add live cells.\n"
	str += fmt.Sprintf("Generation: %d, Number of cells: %d\n", m.generation, len(m.board))
	return str
}
