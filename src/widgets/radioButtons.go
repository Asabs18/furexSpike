package widgets

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
)

type RadioButtons struct {
	Color      color.Color
	CheckBoxes []MEcheckBox

	mouseover bool
	pressed   bool
}

var (
	_ furex.ButtonHandler = (*Button)(nil)
	_ furex.Drawer        = (*Button)(nil)
)

func (r *RadioButtons) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {

}

func (r *RadioButtons) HandlePress(x, y int, t ebiten.TouchID) {

}

func (r *RadioButtons) HandleRelease(x, y int, isCancel bool) {

}
