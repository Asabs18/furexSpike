package widgets

import (
	"image"
	"image/color"
	"strings"

	"github.com/Asabs18/furexSpike/src/sprites"
	"github.com/Asabs18/furexSpike/src/text"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/ganim8/v2"
)

// CONSTANTS
const (
	BUTTON_FONT_SCALAR    = 10
	PRESSED_BUTTON_SCALAR = .95
	FONT_OFFSET_SCALAR    = .01
)

type Button struct {
	Color   color.Color
	OnClick func()

	mouseover bool
	pressed   bool
}

var (
	_ furex.ButtonHandler          = (*Button)(nil)
	_ furex.Drawer                 = (*Button)(nil)
	_ furex.MouseEnterLeaveHandler = (*Button)(nil)
)

func (b *Button) getButtonSprite(color string, shape string) string {
	sprite := "Cannot find sprite for that color."
	if color == "purple" {
		sprite = shape + "ButtonPurple.png"
	} else if color == "red" {
		sprite = shape + "ButtonRed.png"
	} else if color == "green" {
		sprite = shape + "ButtonGreen.png"
	} else if color == "blue" {
		sprite = shape + "ButtonBlue.png"
	} else if color == "yellow" {
		sprite = shape + "ButtonYellow.png"
	} else if color == "orange" {
		sprite = shape + "ButtonOrange.png"
	}
	return sprite
}

func (b *Button) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	x, y := float64(frame.Min.X+frame.Dx()/2), float64(frame.Min.Y+frame.Dy()/2)

	color := strings.ToLower(view.Attrs["color"])
	shape := strings.ToLower(view.Attrs["shape"])
	sprite := b.getButtonSprite(color, shape)

	spriteWidth, spriteHeight := sprites.Get(sprite).Size()
	//Scales unpressed sprite to the desired width divided by the current sprite width
	spriteOpts := ganim8.DrawOpts(x, y, 0, (float64(view.Width) / float64(spriteWidth)), (float64(view.Height) / float64(spriteHeight)), .5, .5)
	//Scales pressed sprite to the desired width divided by the current sprite width however the desired width is reduced by 5% of the unpressed size
	spritePressedOpts := ganim8.DrawOpts(x, y, 0,
		(float64(view.Width)*PRESSED_BUTTON_SCALAR)/float64(spriteWidth),
		(float64(view.Height)*PRESSED_BUTTON_SCALAR)/float64(spriteHeight), .5, .5)

	if b.mouseover {
		//Scale spriteOpts.ColorM by 10% to make the button brighter using "github.com/hajimehoshi/ebiten/v2/colorm"
		//spriteOpts.ColorM = colorm.ScaleRGB(spriteOpts.ColorM, colorm.Scale(1.1, 1.1, 1.1, 1))
		//TODO: Fix this
		spriteOpts.ColorM.Scale(1.1, 1.1, 1.1, 1)
	}
	if b.pressed && sprite != "" {
		ganim8.DrawSpriteWithOpts(screen, sprites.Get(sprite), 0, spritePressedOpts, nil)
		text.R.SetSizePx(int(float64(view.Width)*PRESSED_BUTTON_SCALAR+float64(view.Height)*PRESSED_BUTTON_SCALAR) / BUTTON_FONT_SCALAR)
		y += float64(view.Width) * .01
	} else if sprite != "" {
		ganim8.DrawSpriteWithOpts(screen, sprites.Get(sprite), 0, spriteOpts, nil)
		text.R.SetSizePx((view.Width + view.Height) / BUTTON_FONT_SCALAR)
	}

	text.R.SetTarget(screen)
	text.R.SetColor(b.Color)
	text.R.Draw(view.Text, int(x), int(y))
}

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

func (b *Button) HandleMouseEnter(x, y int) bool {
	b.mouseover = true
	return true
}

func (b *Button) HandleMouseLeave() {
	b.mouseover = false
}
