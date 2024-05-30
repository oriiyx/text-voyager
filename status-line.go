package main

import (
	"fmt"
	"text/template"

	"github.com/jroimartin/gocui"
)

type StatusLine struct {
	tpl *template.Template
}

type StatusLineFunctions struct {
	app *App
}

func (s *StatusLine) Update(v *gocui.View, a *App) {
	v.Clear()
	err := s.tpl.Execute(v, &StatusLineFunctions{app: a})
	if err != nil {
		fmt.Fprintf(v, "StatusLine update error: %v", err)
	}
}
