package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// State represents the state of the model
type State struct {
	Epoch      int
	TotalEpoch int
	// Add other relevant fields as needed
}

// LoadState loads the state from a file
func LoadState() (State, error) {
	// Load the state from the file
	// Implementation depends on how the state is stored
	return State{}, nil // Placeholder for demonstration
}

// SaveState saves the current state to a file
func SaveState(state State) error {
	// Save the state to the file
	// Implementation depends on how the state is stored
	return nil // Placeholder for demonstration
}

// Define the model struct
type model struct {
	State State
}

// Define the Init function for the model
func (m model) Init() tea.Cmd {
	// Load the training state from the file
	state, err := LoadState()
	if err != nil {
		// Use default state if loading fails
		state = State{
			Epoch:      1,
			TotalEpoch: 10, // Default to 10 epochs
		}
	}
	_ = state // Ignore the return value
	m.State = state

	// Start the training loop in the background
	return train(m)
}

// Define the Update function for the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

// Define the View function for the model
func (m model) View() string {
	return fmt.Sprintf("Epoch: %d / %d", m.State.Epoch, m.State.TotalEpoch)
}

// Define the training loop function
func train(m model) tea.Cmd {
	return func() tea.Msg {
		for {
			// Simulate training
			time.Sleep(time.Second)

			// Increment epoch
			m.State.Epoch++

			// Save state
			err := SaveState(m.State)
			if err != nil {
				// Handle error
			}

			// Return a message to trigger an update
			return m
		}
	}
}

func main() {
	p := tea.NewProgram(model{})

	if err := p.Start(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
