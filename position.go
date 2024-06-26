package main

import "github.com/jroimartin/gocui"

const (
	MinWidth  = 60
	MinHeight = 20
)

type position struct {
	// value = prc * MAX + abs
	pct float32
	abs int
}

type viewPosition struct {
	x0, y0, x1, y1 position
}

func (p position) getCoordinate(max int) int {
	return int(p.pct*float32(max)) + p.abs
}

func setView(g *gocui.Gui, viewName string) (*gocui.View, error) {
	maxX, maxY := g.Size()
	position := ViewPositions[viewName]
	return g.SetView(viewName,
		position.x0.getCoordinate(maxX+1),
		position.y0.getCoordinate(maxY+1),
		position.x1.getCoordinate(maxX+1),
		position.y1.getCoordinate(maxY+1))
}

var ViewPositions = map[string]viewPosition{
	SearchPromptView: {
		position{0.0, 0},
		position{0.0, 0},
		position{1.0, -2},
		position{0.0, 2},
	},
	NavigationView: {
		position{0.0, 0},
		position{0.0, 3},
		position{0.2, 0},
		position{1, -5},
	},
	RenderView: {
		position{0.2, 3},
		position{0.0, 3},
		position{1, -2},
		position{1, -5},
	},
	StatuslineView: {
		position{0.0, -1},
		position{1.0, -4},
		position{1.0, 0},
		position{1.0, -1},
	},
	PopupView: {
		position{0.5, -9999}, // set before usage using len(msg)
		position{0.5, -1},
		position{0.5, -9999}, // set before usage using len(msg)
		position{0.5, 1},
	},
}
