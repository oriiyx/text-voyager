package main

import (
	"text/template"

	"github.com/jroimartin/gocui"
)

type Config struct {
	Keys map[string]map[string]string
}

var DefaultKeys = map[string]map[string]string{
	"global": {
		"CtrlR": "submit",
		"CtrlC": "quit",
		"F2":    "focus url",
	},
	"url": {
		"Enter": "submit",
	},
	"help": {
		"ArrowUp":   "scrollUp",
		"ArrowDown": "scrollDown",
		"PageUp":    "pageUp",
		"PageDown":  "pageDown",
	},
}

func (a *App) SetKeys(g *gocui.Gui) error {
	return nil
}

func LoadConfig() (*Config, error) {
	conf := DefaultConfig

	if conf.Keys == nil {
		conf.Keys = DefaultKeys
	} else {
		// copy default keys
		for keyCategory, keys := range DefaultKeys {
			confKeys, found := conf.Keys[keyCategory]
			if found {
				for key, action := range keys {
					if _, found := confKeys[key]; !found {
						conf.Keys[keyCategory][key] = action
					}
				}
			} else {
				conf.Keys[keyCategory] = keys
			}
		}
	}

	return &conf, nil
}

func (a *App) LoadConfig() error {
	conf, err := LoadConfig()
	if err != nil {
		return err
	}
	a.config = conf
	sl, err := NewStatusLine("example")
	if err != nil {
		a.config.Keys = DefaultKeys
		return err
	}
	a.statusLine = sl
	return nil
}

func NewStatusLine(format string) (*StatusLine, error) {
	tpl, err := template.New("status line").Parse(format)
	if err != nil {
		return nil, err
	}
	return &StatusLine{
		tpl: tpl,
	}, nil
}
