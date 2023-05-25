package widgets

import (
	"image"

	"github.com/Asabs18/furexSpike/src/sprites"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/ganim8/v2"
)

const (
	CURSOR_WIDTH_SCALAR  = .00009
	CURSOR_HEIGHT_SCALAR = .000075
)

type Mouse struct {
	isMouseActive bool
	x, y          int
}

var (
	_ furex.DrawHandler            = (*Mouse)(nil)
	_ furex.MouseHandler           = (*Mouse)(nil)
	_ furex.MouseEnterLeaveHandler = (*Mouse)(nil)
)

func (m *Mouse) HandleDraw(screen *ebiten.Image, frame image.Rectangle) {
	if m.isMouseActive {
		spr := sprites.Get("cursor.png")
		// set origin to .1, .1 to make the cursor point to the mouse position
		sprWidth, sprHeight := spr.Size()
		opts := ganim8.DrawOpts(float64(m.x), float64(m.y), 0, float64(sprWidth)*float64(CURSOR_WIDTH_SCALAR), float64(sprHeight)*float64(CURSOR_HEIGHT_SCALAR), .1, .1)
		ganim8.DrawSpriteWithOpts(screen, spr, 0, opts, nil)
	}
}

func (m *Mouse) HandleMouse(x, y int) bool {
	m.x = x
	m.y = y
	return true
}

func (m *Mouse) HandleMouseEnter(x int, y int) bool {
	m.isMouseActive = true
	m.x = x
	m.y = y
	return true
}

func (m *Mouse) HandleMouseLeave() {
	m.isMouseActive = false
}
