// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

var (
	viewArr = []string{"Search", "Navigation", "Render"}
	active  = 0
)

var defaultEditor ViewEditor

var DefaultConfig = Config{}

const (
	ALL_VIEWS = ""

	SearchPromptView = "prompt"
	NavigationView   = "navigation"
	RENDER_VIEW      = "render"
	HELP_VIEW        = "help"
	StatuslineView   = "status-line"

	SearchPromptPlaceholder = "search> "
)

var VIEWS = []string{
	SearchPromptView,
	NavigationView,
}

type viewProperties struct {
	title    string
	frame    bool
	editable bool
	wrap     bool
	editor   gocui.Editor
	text     string
}

var VIEW_PROPERTIES = map[string]viewProperties{
	SearchPromptView: {
		title:    "Search",
		frame:    true,
		editable: true,
		wrap:     false,
		editor:   &SingleLineEditor{&defaultEditor},
	},
	StatuslineView: {
		title:    "",
		frame:    false,
		editable: false,
		wrap:     false,
		editor:   nil,
		text:     "",
	},
}

type App struct {
	viewIndex    int
	historyIndex int
	currentPopup string
	statusLine   *StatusLine
	config       *Config
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (a *App) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if maxX < MIN_WIDTH || maxY < MIN_HEIGHT {
		fmt.Println("Terminal is too small")
		return nil
	}

	for _, name := range []string{
		SearchPromptView,
		StatuslineView,
	} {
		if v, err := setView(g, name); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			setViewProperties(v, name)
		}
	}
	// refreshStatusLine(a, g)

	return nil
}

func refreshStatusLine(a *App, g *gocui.Gui) {
	sv, _ := g.View(StatuslineView)
	sv.BgColor = gocui.ColorDefault | gocui.AttrReverse
	sv.FgColor = gocui.ColorDefault | gocui.AttrReverse
	a.statusLine.Update(sv, a)
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	app := &App{}
	defaultEditor = ViewEditor{app, g, false, gocui.DefaultEditor}
	initApp(app, g)
	err = app.LoadConfig()
	if err != nil {
		g.Close()
		log.Fatalf("Error configuring: %v", err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func initApp(a *App, g *gocui.Gui) {
	g.Cursor = true
	g.InputEsc = false
	g.BgColor = gocui.ColorDefault
	g.FgColor = gocui.ColorDefault
	g.SetManagerFunc(a.Layout)
}
