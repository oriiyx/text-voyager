package main

import "github.com/jroimartin/gocui"

const (
	MIN_WIDTH  = 60
	MIN_HEIGHT = 20
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
	position := VIEW_POSITIONS[viewName]
	return g.SetView(viewName,
		position.x0.getCoordinate(maxX+1),
		position.y0.getCoordinate(maxY+1),
		position.x1.getCoordinate(maxX+1),
		position.y1.getCoordinate(maxY+1))
}

var VIEW_POSITIONS = map[string]viewPosition{
	SearchPromptView: {
		position{0.0, 0},
		position{0.0, 0},
		position{1.0, -2},
		position{0.0, 3}},
	NavigationView: {
		position{0.0, 0},
		position{0.0, 3},
		position{0.3, 0},
		position{0.25, 0}},
	StatuslineView: {
		position{0.0, -1},
		position{1.0, -4},
		position{1.0, 0},
		position{1.0, -1}},
}
