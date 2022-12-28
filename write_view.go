package views

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

func writeView(m Model) string {
	return fmt.Sprintf("Write.\n\n%s\n\n%s", m.textarea.View(), "(ctrl+c to quit | ctrl+s to save | ctrl+e to render)") + "\n\n"
}

func updateWrite(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyCtrlS:
			data := m.textarea.Value()
			err := os.WriteFile("out.md", []byte(data), 0)
			if err != nil {
				log.Fatal(err)
			}
		case tea.KeyCtrlE:
			m.write = !m.write
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
