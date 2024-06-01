package main

import (
	"github.com/jroimartin/gocui"
)

func popup(g *gocui.Gui, msg string) {
	pos := ViewPositions[PopupView]
	pos.x0.abs = -len(msg)/2 - 1
	pos.x1.abs = len(msg)/2 + 1
	ViewPositions[PopupView] = pos

	p := ViewProperties[PopupView]
	p.text = msg
	ViewProperties[PopupView] = p

	v, err := setView(g, PopupView)
	if err != nil {
		return
	}

	setViewProperties(v, PopupView)
	g.SetViewOnTop(PopupView)
	v.Clear()
}

func closeRealPopup(g *gocui.Gui) {
	pos := ViewPositions[PopupView]
	pos.x0.abs = -9999
	pos.x1.abs = -9999
	ViewPositions[PopupView] = pos

	v, err := setView(g, PopupView)
	if err != nil {
		return
	}

	setViewProperties(v, PopupView)
	g.SetViewOnTop(PopupView)
}

func (a *App) closePopup(g *gocui.Gui, viewname string) {
	_, err := g.View(viewname)
	if err == nil {
		a.currentPopup = ""
		g.DeleteView(viewname)
		g.SetCurrentView(VIEWS[a.viewIndex%len(VIEWS)])
		g.Cursor = true
	}
}
