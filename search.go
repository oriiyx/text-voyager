package main

type Search struct {
	status status
	title  string
	query  string
}

// implement the list.Item interface
func (s Search) FilterValue() string { return s.query }

func (s Search) Title() string { return s.title }

func (s Search) Query() string { return s.query }
