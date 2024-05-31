package main

import (
	"errors"
	"fmt"
	"strings"
	"text/template"

	"github.com/jroimartin/gocui"
)

type Config struct {
	Keys map[string]map[string]string
}

var DefaultKeys = map[string]map[string]string{
	"global": {
		"CtrlC": "quit",
		"Tab":   "nextView",
	},
	"prompt": {
		"Enter": "submit",
	},
}

func (a *App) SetKeys(g *gocui.Gui) error {
	for viewName, keys := range a.config.Keys {
		if viewName == "global" {
			viewName = ALL_VIEWS
		}
		for keyStr, commandStr := range keys {
			if err := a.setKey(g, keyStr, commandStr, viewName); err != nil {
				return err
			}
		}
	}
	return nil
}

func LoadConfig() (*Config, error) {
	conf := DefaultConfig
	conf.Keys = DefaultKeys
	return &conf, nil
}

func (a *App) LoadConfig() error {
	conf, err := LoadConfig()
	if err != nil {
		return err
	}
	a.config = conf
	sl, err := NewStatusLine("Press enter to start browsing")
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

func (a *App) setKey(g *gocui.Gui, keyStr, commandStr, viewName string) error {
	if commandStr == "" {
		return nil
	}
	key, mod, err := parseKey(keyStr)
	if err != nil {
		return err
	}
	commandParts := strings.SplitN(commandStr, " ", 2)
	command := commandParts[0]
	var commandArgs string
	if len(commandParts) == 2 {
		commandArgs = commandParts[1]
	}
	keyFnGen, found := COMMANDS[command]
	if !found {
		return fmt.Errorf("Unknown command: %v", command)
	}
	keyFn := keyFnGen(commandArgs, a)

	if err := g.SetKeybinding(viewName, key, mod, keyFn); err != nil {
		return fmt.Errorf("Failed to set key '%v': %v", keyStr, err)
	}
	return nil
}

func parseKey(k string) (interface{}, gocui.Modifier, error) {
	mod := gocui.ModNone
	if strings.Index(k, "Alt") == 0 {
		mod = gocui.ModAlt
		k = k[3:]
	}
	switch len(k) {
	case 0:
		return 0, 0, errors.New("Empty key string")
	case 1:
		if mod != gocui.ModNone {
			k = strings.ToLower(k)
		}
		return rune(k[0]), mod, nil
	}

	key, found := KEYS[k]
	if !found {
		return 0, 0, fmt.Errorf("Unknown key: %v", k)
	}
	return key, mod, nil
}

var KEYS = map[string]gocui.Key{
	"F1":             gocui.KeyF1,
	"F2":             gocui.KeyF2,
	"F3":             gocui.KeyF3,
	"F4":             gocui.KeyF4,
	"F5":             gocui.KeyF5,
	"F6":             gocui.KeyF6,
	"F7":             gocui.KeyF7,
	"F8":             gocui.KeyF8,
	"F9":             gocui.KeyF9,
	"F10":            gocui.KeyF10,
	"F11":            gocui.KeyF11,
	"F12":            gocui.KeyF12,
	"Insert":         gocui.KeyInsert,
	"Delete":         gocui.KeyDelete,
	"Home":           gocui.KeyHome,
	"End":            gocui.KeyEnd,
	"PageUp":         gocui.KeyPgup,
	"PageDown":       gocui.KeyPgdn,
	"ArrowUp":        gocui.KeyArrowUp,
	"ArrowDown":      gocui.KeyArrowDown,
	"ArrowLeft":      gocui.KeyArrowLeft,
	"ArrowRight":     gocui.KeyArrowRight,
	"CtrlTilde":      gocui.KeyCtrlTilde,
	"Ctrl2":          gocui.KeyCtrl2,
	"CtrlSpace":      gocui.KeyCtrlSpace,
	"CtrlA":          gocui.KeyCtrlA,
	"CtrlB":          gocui.KeyCtrlB,
	"CtrlC":          gocui.KeyCtrlC,
	"CtrlD":          gocui.KeyCtrlD,
	"CtrlE":          gocui.KeyCtrlE,
	"CtrlF":          gocui.KeyCtrlF,
	"CtrlG":          gocui.KeyCtrlG,
	"Backspace":      gocui.KeyBackspace,
	"CtrlH":          gocui.KeyCtrlH,
	"Tab":            gocui.KeyTab,
	"CtrlI":          gocui.KeyCtrlI,
	"CtrlJ":          gocui.KeyCtrlJ,
	"CtrlK":          gocui.KeyCtrlK,
	"CtrlL":          gocui.KeyCtrlL,
	"Enter":          gocui.KeyEnter,
	"CtrlM":          gocui.KeyCtrlM,
	"CtrlN":          gocui.KeyCtrlN,
	"CtrlO":          gocui.KeyCtrlO,
	"CtrlP":          gocui.KeyCtrlP,
	"CtrlQ":          gocui.KeyCtrlQ,
	"CtrlR":          gocui.KeyCtrlR,
	"CtrlS":          gocui.KeyCtrlS,
	"CtrlT":          gocui.KeyCtrlT,
	"CtrlU":          gocui.KeyCtrlU,
	"CtrlV":          gocui.KeyCtrlV,
	"CtrlW":          gocui.KeyCtrlW,
	"CtrlX":          gocui.KeyCtrlX,
	"CtrlY":          gocui.KeyCtrlY,
	"CtrlZ":          gocui.KeyCtrlZ,
	"Esc":            gocui.KeyEsc,
	"CtrlLsqBracket": gocui.KeyCtrlLsqBracket,
	"Ctrl3":          gocui.KeyCtrl3,
	"Ctrl4":          gocui.KeyCtrl4,
	"CtrlBackslash":  gocui.KeyCtrlBackslash,
	"Ctrl5":          gocui.KeyCtrl5,
	"CtrlRsqBracket": gocui.KeyCtrlRsqBracket,
	"Ctrl6":          gocui.KeyCtrl6,
	"Ctrl7":          gocui.KeyCtrl7,
	"CtrlSlash":      gocui.KeyCtrlSlash,
	"CtrlUnderscore": gocui.KeyCtrlUnderscore,
	"Space":          gocui.KeySpace,
	"Backspace2":     gocui.KeyBackspace2,
	"Ctrl8":          gocui.KeyCtrl8,
}
