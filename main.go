package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"log"
	"os"
)

func initialModel() model {
	ti := textarea.New()
	ti.Placeholder = ""
	ti.SetHeight(20)
	ti.Focus()
	return model{
		textarea: ti,
		err:      nil,
		write:    true,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.write {
		return updateWrite(msg, m)
	} else {
		return updateRender(msg, m)
	}
}

func (m model) View() string {
	if m.write {
		return writeView(m)
	} else {
		return renderView(m)
	}
}

func updateRender(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
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

func renderView(m model) string {
	render, err := glamour.Render(m.textarea.Value(), "dark")
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("Read.\n\n%s\n\n%s", render, "(ctrl+c to quit | ctrl+e to write)") + "\n\n"

}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There has been an error: %s", err)
		os.Exit(1)
	}
}
