package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SearchForm struct {
	query textinput.Model
	index int
}

func newDefaultSearchForm() *SearchForm {
	return NewSearchForm()
}

func NewSearchForm() *SearchForm {
	form := SearchForm{
		query: textinput.New(),
	}
	form.query.Placeholder = "..."
	form.query.Focus()
	return &form
}

func (f SearchForm) CreateSearch() Search {
	return Search{searchQueryView, f.query.Value(), f.query}
}

func (f SearchForm) Init() tea.Cmd {
	return nil
}

func (f SearchForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	// case column:
	// 	f.col = msg
	// 	f.col.list.Index()
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return f, tea.Quit

		case key.Matches(msg, keys.Back):
			return board.Update(nil)
		case key.Matches(msg, keys.Enter):
			// Return the completed form as a message.
			return board.Update(f)
		}
	}

	if f.query.Focused() {
		f.query, cmd = f.query.Update(msg)
		return f, cmd
	}

	return f, cmd
}

func (f SearchForm) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,
		"Search the internet",
		f.query.View(),
	)
}
