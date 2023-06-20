package widgets

import (
	"image"
	"image/color"

	"github.com/Asabs18/furexSpike/src/sprites"
	"github.com/Asabs18/furexSpike/src/text"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/ganim8/v2"
)

type TextBox struct {
	Color    color.Color
	FontSize int
	pressed  bool
	runes    []rune
	Text     string
	Counter  int
}

var (
	_ furex.ButtonHandler = (*TextBox)(nil)
	_ furex.Drawer        = (*TextBox)(nil)
	_ furex.Updater       = (*TextBox)(nil)
)

func (b *TextBox) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	x, y := float64(frame.Min.X+frame.Dx()/2), float64(frame.Min.Y+frame.Dy()/2)

	sprite := "textbox.png"

	spriteWidth, spriteHeight := sprites.Get(sprite).Size()

	spriteOpts := ganim8.DrawOpts(x, y, 0, (float64(view.Width) / float64(spriteWidth)), (float64(view.Height) / float64(spriteHeight)), .5, .5)

	ganim8.DrawSpriteWithOpts(screen, sprites.Get(sprite), 0, spriteOpts, nil)
	text.R.SetSizePx(b.FontSize)
	text.R.SetTarget(screen)
	text.R.SetColor(b.Color)
	if b.pressed {
		t := b.Text
		if b.Counter%60 < 30 {
			t += "_"
		}
		text.R.Draw(t, int(x), int(y))
	} else {
		text.R.Draw(b.Text, int(x), int(y))
	}
}

func (b *TextBox) HandlePress(x, y int, t ebiten.TouchID) {
	if !b.pressed {
		b.pressed = true
	} else {
		b.pressed = false
	}
}

func (b *TextBox) HandleRelease(x, y int, isCancel bool) {

}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

func (b *TextBox) Update(v *furex.View) {
	if b.pressed {

		b.runes = ebiten.AppendInputChars(b.runes[:0])
		b.Text += string(b.runes)

		if repeatingKeyPressed(ebiten.KeyBackspace) {
			if len(b.Text) >= 1 {
				b.Text = b.Text[:len(b.Text)-1]
			}
		}

		b.Counter++
	}
}
