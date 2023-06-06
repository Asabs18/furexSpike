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
	MECHECKBOX_FONT_SCALAR = 10
)

type MEcheckBox struct {
	Color      color.Color
	CheckBoxes []CheckBox

	mouseover bool
	pressed   bool
}

var (
	_ furex.ButtonHandler = (*Button)(nil)
	_ furex.Drawer        = (*Button)(nil)
)

func (m *MEcheckBox) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	x, y := float64(frame.Min.X+frame.Dx()/2), float64(frame.Min.Y+frame.Dy()/2)

	sprite := "unclickedCheckbox.png"

	spriteWidth, spriteHeight := sprites.Get(sprite).Size()
	//Scales unpressed sprite to the desired width divided by the current sprite width
	spriteOpts := ganim8.DrawOpts(x, y, 0, (float64(view.Width) / float64(spriteWidth)), (float64(view.Height) / float64(spriteHeight)), .5, .5)
	//Scales pressed sprite to the desired width divided by the current sprite width however the desired width is reduced by 5% of the unpressed size
	spritePressedOpts := ganim8.DrawOpts(x, y, 0,
		(float64(view.Width)*PRESSED_BUTTON_SCALAR)/float64(spriteWidth),
		(float64(view.Height)*PRESSED_BUTTON_SCALAR)/float64(spriteHeight), .5, .5)
	text.R.SetSizePx((view.Width + view.Height) / MECHECKBOX_FONT_SCALAR)

	if m.mouseover {
		//Scale spriteOpts.ColorM by 10% to make the button brighter using "github.com/hajimehoshi/ebiten/v2/colorm"
		//spriteOpts.ColorM = colorm.ScaleRGB(spriteOpts.ColorM, colorm.Scale(1.1, 1.1, 1.1, 1))
		//TODO: Fix this
		spriteOpts.ColorM.Scale(1.1, 1.1, 1.1, 1)
	}
	if m.pressed && sprite != "" {
		sprite = "clickedCheckbox.png"
		ganim8.DrawSpriteWithOpts(screen, sprites.Get(sprite), 0, spritePressedOpts, nil)
	} else if sprite != "" {
		sprite := "unclickedCheckbox.png"
		ganim8.DrawSpriteWithOpts(screen, sprites.Get(sprite), 0, spriteOpts, nil)
	}

	text.R.SetTarget(screen)
	text.R.SetColor(m.Color)
	text.R.Draw(view.Text, int(x), int(y))
}

func (m *MEcheckBox) HandlePress(x, y int, t ebiten.TouchID) {
	if !m.pressed {
		m.pressed = true
	} else {
		m.pressed = false
	}
}

func (m *MEcheckBox) HandleRelease(x, y int, isCancel bool) {

}
