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

// CONSTANTS
const (
	CHECKBOX_FONT_SCALAR = 10
)

type CheckBox struct {
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

func (c *CheckBox) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	x, y := float64(frame.Min.X+frame.Dx()/2), float64(frame.Min.Y+frame.Dy()/2)

	sprite := "unclickedCheckbox.png"

	spriteWidth, spriteHeight := sprites.Get(sprite).Size()
	//Scales unpressed sprite to the desired width divided by the current sprite width
	spriteOpts := ganim8.DrawOpts(x, y, 0, (float64(view.Width) / float64(spriteWidth)), (float64(view.Height) / float64(spriteHeight)), .5, .5)
	//Scales pressed sprite to the desired width divided by the current sprite width however the desired width is reduced by 5% of the unpressed size
	spritePressedOpts := ganim8.DrawOpts(x, y, 0,
		(float64(view.Width)*PRESSED_BUTTON_SCALAR)/float64(spriteWidth),
		(float64(view.Height)*PRESSED_BUTTON_SCALAR)/float64(spriteHeight), .5, .5)
	text.R.SetSizePx((view.Width + view.Height) / CHECKBOX_FONT_SCALAR)

	if c.mouseover {
		//Scale spriteOpts.ColorM by 10% to make the button brighter using "github.com/hajimehoshi/ebiten/v2/colorm"
		//spriteOpts.ColorM = colorm.ScaleRGB(spriteOpts.ColorM, colorm.Scale(1.1, 1.1, 1.1, 1))
		//TODO: Fix this
		spriteOpts.ColorM.Scale(1.1, 1.1, 1.1, 1)
	}
	if c.pressed && sprite != "" {
		sprite = "clickedCheckbox.png"
		ganim8.DrawSpriteWithOpts(screen, sprites.Get(sprite), 0, spritePressedOpts, nil)
	} else if sprite != "" {
		sprite := "unclickedCheckbox.png"
		ganim8.DrawSpriteWithOpts(screen, sprites.Get(sprite), 0, spriteOpts, nil)
	}

	text.R.SetTarget(screen)
	text.R.SetColor(c.Color)
	text.R.Draw(view.Text, int(x), int(y))
}

func (c *CheckBox) HandlePress(x, y int, t ebiten.TouchID) {
	if !c.pressed {
		c.pressed = true
	} else {
		c.pressed = false
	}
}

func (c *CheckBox) HandleRelease(x, y int, isCancel bool) {
	if !isCancel {
		if c.OnClick != nil {
			c.OnClick()
		}
	}
}

func (c *CheckBox) HandleMouseEnter(x, y int) bool {
	c.mouseover = true
	return true
}

func (c *CheckBox) HandleMouseLeave() {
	c.mouseover = false
}
