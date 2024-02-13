package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const useHighPerformanceRenderer = false

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4F5D75")).
			Background(lipgloss.Color("#A8D0E6")).
			Padding(0, 1).
			Border(lipgloss.RoundedBorder())

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#A8D0E6")).
			Padding(0, 1).
			Border(lipgloss.RoundedBorder())
)

type model struct {
	content string
	ready   bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		if !m.ready {
			m.ready = true
		}
	}

	return m, nil
}

func (m model) View() string {
	if !m.ready {
		return "\n  Fetching training details..."
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.content, m.footerView())
}

func (m model) headerView() string {
	title := titleStyle.Render("ASE Training Details")
	line := strings.Repeat("─", len(title))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	info := infoStyle.Render("Scroll to view more")
	line := strings.Repeat("─", len(info))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func main() {
	p := tea.NewProgram(
		model{},
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
