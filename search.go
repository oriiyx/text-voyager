package main

import "github.com/charmbracelet/bubbles/textinput"

type Search struct {
	status status
	title  string
	query  textinput.Model
}

// implement the list.Item interface
func (s Search) FilterValue() string { return s.query.Value() }

func (s Search) Title() string { return s.title }

func (s Search) Query() string { return s.query.Value() }
