package main

import (
	"github.com/jroimartin/gocui"
)

func popup(g *gocui.Gui, msg string) {
	pos := VIEW_POSITIONS[POPUP_VIEW]
	pos.x0.abs = -len(msg)/2 - 1
	pos.x1.abs = len(msg)/2 + 1
	VIEW_POSITIONS[POPUP_VIEW] = pos

	p := VIEW_PROPERTIES[POPUP_VIEW]
	p.text = msg
	VIEW_PROPERTIES[POPUP_VIEW] = p

	if v, err := setView(g, POPUP_VIEW); err != nil {
		if err != gocui.ErrUnknownView {
			return
		}

		setViewProperties(v, POPUP_VIEW)
		g.SetViewOnTop(POPUP_VIEW)
	}
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