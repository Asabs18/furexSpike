package widgets

import (
	"image"
	"image/color"
	"strings"

	"github.com/Asabs18/furexSpike/src/sprites"
	"github.com/Asabs18/furexSpike/src/text"
	"github.com/eiannone/keyboard"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/ganim8/v2"
)

type TextBox struct {
	Color    color.Color
	FontSize int
	Text     string
	pressed  bool
}

var (
	_ furex.ButtonHandler = (*TextBox)(nil)
	_ furex.Drawer        = (*TextBox)(nil)
	_ furex.Updater       = (*TextBox)(nil)
)

func (t *TextBox) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	x, y := float64(frame.Min.X+frame.Dx()/2), float64(frame.Min.Y+frame.Dy()/2)

	sprite := "textbox.png"

	spriteWidth, spriteHeight := sprites.Get(sprite).Size()

	spriteOpts := ganim8.DrawOpts(x, y, 0, (float64(view.Width) / float64(spriteWidth)), (float64(view.Height) / float64(spriteHeight)), .5, .5)

	ganim8.DrawSpriteWithOpts(screen, sprites.Get(sprite), 0, spriteOpts, nil)
	text.R.SetSizePx(t.FontSize)
	text.R.SetTarget(screen)
	text.R.SetColor(t.Color)
	text.R.Draw(t.Text, int(x), int(y))
}

func (b *TextBox) HandlePress(x, y int, t ebiten.TouchID) {
	if !b.pressed {
		b.pressed = true
		b.Text += "|"
	} else {
		b.pressed = false
		b.Text = strings.TrimRight(b.Text, "|")
	}
}

func (b *TextBox) HandleRelease(x, y int, isCancel bool) {

}

func (b *TextBox) Update(v *furex.View) {
	if b.pressed {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		b.Text += string(char)
	}
}
