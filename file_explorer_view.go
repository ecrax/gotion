package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
	"os/exec"
)

var docStyle = lipgloss.NewStyle().Margin(0, 0)

type item struct {
	title, desc string
}

func (i item) Title() string {
	return i.title
}
func (i item) Description() string {
	return i.desc
}
func (i item) FilterValue() string {
	return i.title
}

func fileExplorerView(m Model) string {
	return m.fileList.View()
}

func updateFileExplore(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyCtrlE:
			m.write = !m.write
		case tea.KeyEnter:
			//return m, openEditor(m.rootPath + string(os.PathSeparator) + m.fileList.SelectedItem().FilterValue())
			return m, openEditor(m.fileList.SelectedItem().FilterValue())
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.fileList.SetSize(msg.Width-h, msg.Height-v)
	case errMsg:
		m.err = msg
		return m, nil
	case editorFinishedMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, tea.Quit
		}
		return m, nil
	}

	var cmd tea.Cmd
	m.fileList, cmd = m.fileList.Update(msg)

	return m, cmd
}

func openEditor(file string) tea.Cmd {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	return func() tea.Msg {
		cmd := exec.Command(editor, file)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
		return editorFinishedMsg{}
	}
}

type editorFinishedMsg struct{ err error }
