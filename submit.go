package main

import (
	"log"
	"strings"

	"github.com/jroimartin/gocui"
)

func submit(g *gocui.Gui, v *gocui.View) error {
	log.Println("submitting")
	log.Printf("%#v", g)
	log.Printf("%#v", v)
	test := getViewValue(g, SearchPromptView)
	log.Println(test)
	return nil
}

func getViewValue(g *gocui.Gui, name string) string {
	v, err := g.View(name)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(v.Buffer())
}
