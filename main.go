// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/jroimartin/gocui"
)

type ViewEditor struct {
	app           *App
	g             *gocui.Gui
	backTabEscape bool
	origEditor    gocui.Editor
}

var defaultEditor ViewEditor

var DefaultConfig = Config{}

const (
	ALL_VIEWS = ""

	SearchPromptView = "prompt"
	NavigationView   = "navigation"
	RENDER_VIEW      = "render"
	HELP_VIEW        = "help"
	StatuslineView   = "status-line"
	POPUP_VIEW       = "popup_view"

	SearchPromptPlaceholder = "search> "
)

var VIEWS = []string{
	SearchPromptView,
	StatuslineView,
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
	POPUP_VIEW: {
		title:    "Info",
		frame:    true,
		editable: false,
		wrap:     false,
		editor:   nil,
	},
}

type App struct {
	viewIndex    int
	historyIndex int
	currentPopup string
	userLocale   string
	userLanguage string
	statusLine   *StatusLine
	config       *Config
}

func quit(g *gocui.Gui, v *gocui.View) error {
	fmt.Println("Quitting")
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
	refreshStatusLine(a, g)

	return nil
}

func refreshStatusLine(a *App, g *gocui.Gui) {
	sv, _ := g.View(StatuslineView)
	sv.BgColor = gocui.ColorDefault | gocui.AttrReverse
	sv.FgColor = gocui.ColorDefault | gocui.AttrReverse
	a.statusLine.Update(sv, a)
}

func main() {
	file, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	log.SetOutput(file)

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

	app.Layout(g)
	g.SetCurrentView(VIEWS[app.viewIndex])

	_ = app.SetKeys(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func initApp(a *App, g *gocui.Gui) {
	lang, loc := getLocale()

	a.userLocale = lang
	a.userLocale = loc

	g.Cursor = true
	g.InputEsc = false
	g.BgColor = gocui.ColorDefault
	g.FgColor = gocui.ColorDefault
	g.SetManagerFunc(a.Layout)
}

func getLocale() (string, string) {
	osHost := runtime.GOOS
	defaultLang := "en"
	defaultLoc := "US"
	switch osHost {
	case "windows":
		// Exec powershell Get-Culture on Windows.
		cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
		output, err := cmd.Output()
		if err == nil {
			langLocRaw := strings.TrimSpace(string(output))
			langLoc := strings.Split(langLocRaw, "-")
			lang := langLoc[0]
			loc := langLoc[1]
			return lang, loc
		}
	case "darwin":
		// Exec shell Get-Culture on MacOS.
		cmd := exec.Command("osascript", "-e", "user locale of (get system info)")
		output, err := cmd.Output()
		if err == nil {
			langLocRaw := strings.TrimSpace(string(output))
			langLoc := strings.Split(langLocRaw, "_")
			lang := langLoc[0]
			loc := langLoc[1]
			return lang, loc
		}
	case "linux":
		envlang, ok := os.LookupEnv("LANG")
		if ok {
			langLocRaw := strings.TrimSpace(envlang)
			langLocRaw = strings.Split(envlang, ".")[0]
			langLoc := strings.Split(langLocRaw, "_")
			lang := langLoc[0]
			loc := langLoc[1]
			return lang, loc
		}
	}
	return defaultLang, defaultLoc
}
