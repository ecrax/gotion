package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type errMsg error
type Model struct {
	textarea     textarea.Model
	err          error
	write        bool
	fileExplorer bool
	fileList     list.Model
	rootPath     string
}

func initialModel() Model {
	ti := textarea.New()
	ti.Placeholder = ""
	ti.SetHeight(20)
	ti.Focus()

	path, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	fl := list.New(findFilesInPath(path), list.NewDefaultDelegate(), 0, 0)
	fl.Title = "Find."

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)
	//log.Fatal(exPath)

	return Model{
		textarea:     ti,
		fileList:     fl,
		err:          nil,
		write:        true,
		fileExplorer: true,
		rootPath:     exPath,
	}
}

func findFilesInPath(path string) []list.Item {
	items := make([]list.Item, 0)
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		isMD := strings.HasSuffix(file.Name(), "md")
		if !file.IsDir() && isMD {
			items = append(items, item{title: file.Name()})
		}
	}
	return items
}
