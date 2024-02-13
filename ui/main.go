package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Model holds the application state
type Model struct {
	TrainingData string
	Error        error
}

// Init initializes the application
func (m Model) Init() tea.Cmd {
	return fetchTrainingData()
}

// Update updates the application state in response to messages
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

// View renders the application UI
func (m Model) View() string {
	var b strings.Builder

	// Clear the screen
	b.WriteString("\x1b[2J")

	if m.Error != nil {
		b.WriteString(fmt.Sprintf("Error: %v", m.Error))
	} else {
		b.WriteString(fmt.Sprintf("Training Data: %s", m.TrainingData))
	}

	return b.String()
}

func main() {
	p := tea.NewProgram(Model{})
	if err := p.Start(); err != nil {
		fmt.Printf("Error starting program: %v", err)
		os.Exit(1)
	}
}

// fetchTrainingData simulates fetching training data from a server
func fetchTrainingData() tea.Cmd {
	return func() tea.Msg {
		// Simulate a forbidden error
		return Model{Error: fmt.Errorf("403 Client Error: Forbidden for url: http://localhost:5000/")}
	}
}
