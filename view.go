package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func setViewProperties(v *gocui.View, name string) {
	v.Title = ViewProperties[name].title
	v.Frame = ViewProperties[name].frame
	v.Editable = ViewProperties[name].editable
	v.Wrap = ViewProperties[name].wrap
	v.Editor = ViewProperties[name].editor
	setViewTextAndCursor(v, ViewProperties[name].text)
	log.Println(ViewProperties[name].text)
}

func setViewTextAndCursor(v *gocui.View, s string) {
	v.Clear()
	log.Printf("Printing the following: %s", s)
	fmt.Fprint(v, s)
	v.SetCursor(len(s), 0)
}

func (a *App) NextView(g *gocui.Gui, v *gocui.View) error {
	a.viewIndex = (a.viewIndex + 1) % len(VIEWS)
	return a.setView(g)
}

func (a *App) PrevView(g *gocui.Gui, v *gocui.View) error {
	a.viewIndex = (a.viewIndex - 1 + len(VIEWS)) % len(VIEWS)
	return a.setView(g)
}

func (a *App) setView(g *gocui.Gui) error {
	a.closePopup(g, a.currentPopup)
	_, err := g.SetCurrentView(VIEWS[a.viewIndex])
	return err
}

func (a *App) setViewByName(g *gocui.Gui, name string) error {
	for i, v := range VIEWS {
		if v == name {
			a.viewIndex = i
			return a.setView(g)
		}
	}
	return fmt.Errorf("View not found")
}

func ChangeViewText(g *gocui.Gui, view string, msg string) {
	p := ViewProperties[view]
	p.text = msg
	ViewProperties[view] = p

	v, err := setView(g, view)
	if err != nil {
		return
	}

	setViewProperties(v, view)
	g.SetViewOnTop(view)
}
