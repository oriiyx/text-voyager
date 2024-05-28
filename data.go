package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
)

// Provides the mock data to fill the kanban board

func (boardModel *Board) initLists() {
	textInput := textinput.New()
	textInput.Placeholder = "Enter your new todo"
	textInput.Focus()

	boardModel.cols = []column{
		newColumn(searchQueryView),
		newColumn(navigationView),
		newColumn(renderView),
	}

	// Init Search Query View
	boardModel.cols[searchQueryView].list.Title = "Search"
	boardModel.cols[searchQueryView].list.SetItems([]list.Item{
		Search{status: searchQueryView, title: "Search", query: textInput},
	})

	// Init Navigation View
	boardModel.cols[navigationView].list.Title = "Navigation"
	boardModel.cols[navigationView].list.SetItems([]list.Item{
		Task{status: navigationView, title: "write code", description: "don't worry, it's Go"},
	})

	// Init Review View
	boardModel.cols[renderView].list.Title = "Render"
	boardModel.cols[renderView].list.SetItems([]list.Item{
		Task{status: renderView, title: "stay cool", description: "as a cucumber"},
	})
}
