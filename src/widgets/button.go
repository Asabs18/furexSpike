package widgets

import (
	"image"
	"image/color"

	"github.com/Asabs18/furexSpike/src/sprites"
	"github.com/Asabs18/furexSpike/src/text"
	"github.com/Asabs18/furexSpike/src/functions"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/ganim8/v2"
)

type Button struct {
	Color   color.Color
	OnClick func()

	mouseover bool
	pressed   bool

	width  int
	height int
}

var (
	_ furex.ButtonHandler          = (*Button)(nil)
	_ furex.Drawer                 = (*Button)(nil)
	_ furex.MouseEnterLeaveHandler = (*Button)(nil)
)

func (b *Button) HandlePress(x, y int, t ebiten.TouchID) {
	b.pressed = true
}

func (b *Button) HandleRelease(x, y int, isCancel bool) {
	b.pressed = false
	if !isCancel {
		if b.OnClick != nil {
			b.OnClick()
		}
	}
}

func (b *Button) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	x, y := float64(frame.Min.X+frame.Dx()/2), float64(frame.Min.Y+frame.Dy()/2)

	sprite := view.Attrs["sprite"]
	spritePressed := view.Attrs["sprite_pressed"]

	if First(sprites.Get(sprite).Size()) != First(sprites.Get(spritePressed).Size()) || Second(sprites.Get(sprite).Size()) != Second(sprites.Get(spritePressed).Size()) {
		panic("sprite size mismatch")
	}

	butWidth, butHeight := sprites.Get(sprite).Size()

	opts := ganim8.DrawOpts(x, y, 0, 1, 1, .5, .5)
	if b.mouseover {
		opts.ColorM.Scale(1.1, 1.1, 1.1, 1)
	}
	if b.pressed && spritePressed != "" {
		ganim8.DrawSpriteWithOpts(screen, sprites.Get(spritePressed), 0, opts, nil)
	} else if sprite != "" {
		ganim8.DrawSpriteWithOpts(screen, sprites.Get(sprite), 0, opts, nil)
	}

	text.R.SetAlign(etxt.YCenter, etxt.XCenter)
	text.R.SetTarget(screen)
	if b.Color != nil {
		text.R.SetColor(b.Color)
	} else {
		text.R.SetColor(color.White)
	}
	text.R.Draw(view.Text, int(x), int(y))
}

func (b *Button) HandleMouseEnter(x, y int) bool {
	b.mouseover = true
	return true
}

func (b *Button) HandleMouseLeave() {
	b.mouseover = false
}
