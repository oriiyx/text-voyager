package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type status int

func (s status) getNext() status {
	if s == renderView {
		return searchQueryView
	}
	return s + 1
}

func (s status) getPrev() status {
	if s == searchQueryView {
		return renderView
	}
	return s - 1
}

const margin = 4

var board *Board

const (
	searchQueryView status = iota
	navigationView
	renderView
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	board = NewBoard()
	board.initLists()
	p := tea.NewProgram(board)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
