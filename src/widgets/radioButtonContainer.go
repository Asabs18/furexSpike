package widgets

import (
	"github.com/yohamta/furex/v2"
)

type RadioButtonContainer struct {
	Buttons     []string
	CurrState   []bool
	CurrPressed int
}

var (
	_ furex.Updater = (*RadioButtonContainer)(nil)
)

func (r *RadioButtonContainer) Update(v *furex.View) {
	for i, button := range r.Buttons {
		buttonView, _ := v.GetByID(button)
		if buttonView.Handler.(*CheckBox).pressed {
			if !r.CurrState[i] && i != r.CurrPressed {
				r.CurrState[i] = true
				r.CurrPressed = i
				buttonView.Handler.(*CheckBox).pressed = true
				for j := 0; j < len(r.Buttons); j++ {
					if j != r.CurrPressed {
						b, _ := v.GetByID(r.Buttons[j])
						r.CurrState[j] = false
						b.Handler.(*CheckBox).pressed = false
					}
				}
			} else if r.CurrState[i] && i != r.CurrPressed {
				r.CurrState[i] = false
				r.CurrPressed = -1
				buttonView.Handler.(*CheckBox).pressed = false
			}
		}
	}
}
