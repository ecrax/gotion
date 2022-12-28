package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"log"
)

func updateRender(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlE:
			m.write = !m.write
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, nil
}

func renderView(m Model) string {
	render, err := glamour.Render(m.textarea.Value(), "dark")
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("Read.\n\n%s\n\n%s", render, "(ctrl+c to quit | ctrl+e to write)") + "\n\n"

}
