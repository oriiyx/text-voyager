package main

import "github.com/charmbracelet/bubbles/list"

// Provides the mock data to fill the kanban board

func (b *Board) initLists() {
	b.cols = []column{
		newColumn(searchQueryView),
		newColumn(navigationView),
		newColumn(renderView),
	}
	// Init To Do
	b.cols[searchQueryView].list.Title = "To Do"
	b.cols[searchQueryView].list.SetItems([]list.Item{
		Task{status: searchQueryView, title: "buy milk", description: "strawberry milk"},
		Task{status: searchQueryView, title: "eat sushi", description: "negitoro roll, miso soup, rice"},
		Task{status: searchQueryView, title: "fold laundry", description: "or wear wrinkly t-shirts"},
	})
	// Init in progress
	b.cols[navigationView].list.Title = "In Progress"
	b.cols[navigationView].list.SetItems([]list.Item{
		Task{status: navigationView, title: "write code", description: "don't worry, it's Go"},
	})
	// Init renderView
	b.cols[renderView].list.Title = "Done"
	b.cols[renderView].list.SetItems([]list.Item{
		Task{status: renderView, title: "stay cool", description: "as a cucumber"},
	})
}
