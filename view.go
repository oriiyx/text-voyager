package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func setViewProperties(v *gocui.View, name string) {
	v.Title = VIEW_PROPERTIES[name].title
	v.Frame = VIEW_PROPERTIES[name].frame
	v.Editable = VIEW_PROPERTIES[name].editable
	v.Wrap = VIEW_PROPERTIES[name].wrap
	v.Editor = VIEW_PROPERTIES[name].editor
	setViewTextAndCursor(v, VIEW_PROPERTIES[name].text)
	log.Println(VIEW_PROPERTIES[name].text)
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
	p := VIEW_PROPERTIES[view]
	p.text = msg
	VIEW_PROPERTIES[view] = p

	v, err := setView(g, view)
	if err != nil {
		return
	}

	setViewProperties(v, view)
	g.SetViewOnTop(view)
}
