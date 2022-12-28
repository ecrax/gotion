package models

import "github.com/charmbracelet/bubbles/textarea"

type errMsg error
type Model struct {
	textarea textarea.Model
	err      error
	write    bool
}
