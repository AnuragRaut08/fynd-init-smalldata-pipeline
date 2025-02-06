// internal/project/tui.go
package project

import (
    "fmt"
    "strings"

    tea "github.com/charmbracelet/bubbletea"
)

// Model defines the data structure for the TUI.
type Model struct {
    ProjectName string
    CreatedItems []string
}

// Init initializes the TUI model.
func (m Model) Init() tea.Cmd {
    return nil
}

// Update handles messages and updates the TUI model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "esc", "ctrl+c":
            return m, tea.Quit
        }
    }
    return m, nil
}

// View renders the TUI.
func (m Model) View() string {
    var b strings.Builder
    b.WriteString(fmt.Sprintf("Project '%s' has been initialized successfully.\n\n", m.ProjectName))
    b.WriteString("The following items were created:\n")
    for _, item := range m.CreatedItems {
        b.WriteString(fmt.Sprintf("- %s\n", item))
    }
    b.WriteString("\nPress 'q' or 'esc' to exit.")
    return b.String()
}

// StartTUI starts the Bubble Tea program.
func StartTUI(projectName string, createdItems []string) {
    p := tea.NewProgram(Model{ProjectName: projectName, CreatedItems: createdItems})
    if err := p.Start(); err != nil {
        fmt.Printf("Error starting TUI: %v\n", err)
    }
}
