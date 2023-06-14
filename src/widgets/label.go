package widgets

import (
	"image"
	"image/color"

	"github.com/Asabs18/furexSpike/src/sprites"
	"github.com/Asabs18/furexSpike/src/text"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/ganim8/v2"
)

type Label struct {
	Color color.Color
}

var (
	_ furex.Drawer = (*Label)(nil)
)

func (l *Label) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	x, y := float64(frame.Min.X+frame.Dx()/2), float64(frame.Min.Y+frame.Dy()/2)

	sprite := "textbox.png"

	spriteWidth, spriteHeight := sprites.Get(sprite).Size()

	spriteOpts := ganim8.DrawOpts(x, y, 0, (float64(view.Width) / float64(spriteWidth)), (float64(view.Height) / float64(spriteHeight)), .5, .5)

	ganim8.DrawSpriteWithOpts(screen, sprites.Get(sprite), 0, spriteOpts, nil)
	text.R.SetSizePx(15)
	text.R.SetTarget(screen)
	text.R.SetColor(l.Color)
	text.R.Draw(view.Text, int(x), int(y))
}
