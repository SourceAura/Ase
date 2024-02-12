package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	epochs   int
	lossData string
	time     string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle messages here
	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("Epochs: %d | Loss Data: %s | Time: %s", m.epochs, m.lossData, m.time)
}

func main() {
	p := tea.NewProgram(model{})

	if err := p.Start(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
