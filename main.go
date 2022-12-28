package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.fileExplorer {
		return updateFileExplore(msg, m)
	}

	if m.write {
		return updateWrite(msg, m)
	} else {
		return updateRender(msg, m)
	}
}

func (m Model) View() string {
	if m.fileExplorer {
		return fileExplorerView(m)
	}

	if m.write {
		return writeView(m)
	} else {
		return renderView(m)
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There has been an error: %s", err)
		os.Exit(1)

	}
}
