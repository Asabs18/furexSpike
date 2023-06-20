package main

import (
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"

	"github.com/Asabs18/furexSpike/src/sprites"
	"github.com/Asabs18/furexSpike/src/text"
	"github.com/Asabs18/furexSpike/src/widgets"

	_ "embed"
)

type Game struct {
	initOnce sync.Once
	screen   screen
	gameUI   *furex.View
}

type screen struct {
	Width  int
	Height int
}

func (g *Game) Update() error {
	g.initOnce.Do(func() { g.setupUI() })
	g.gameUI.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 120, 120, 255})
	g.gameUI.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.screen.Width = outsideWidth
	g.screen.Height = outsideHeight
	return g.screen.Width, g.screen.Height
}

func NewGame() (*Game, error) {
	text.LoadFonts()
	sprites.LoadSprites(
		"assets/images/cursor.xml",
		"assets/images/cursor.png",
		sprites.LoadOpts{
			PanelOpts: map[string]sprites.PanelOpts{},
		})
	sprites.LoadSprites(
		"assets/images/blankButtonColors.xml",
		"assets/images/blankButtonColors.png",
		sprites.LoadOpts{
			PanelOpts: map[string]sprites.PanelOpts{},
		})
	sprites.LoadSprites(
		"assets/images/textbox.xml",
		"assets/images/textbox.png",
		sprites.LoadOpts{
			PanelOpts: map[string]sprites.PanelOpts{},
		})
	sprites.LoadSprites(
		"assets/images/checkbox.xml",
		"assets/images/checkbox.png",
		sprites.LoadOpts{
			PanelOpts: map[string]sprites.PanelOpts{},
		})
	game := &Game{}
	return game, nil
}

//go:embed assets/html/main.html
var mainHTML string

func (g *Game) setupUI() {

	g.gameUI = furex.Parse(mainHTML, &furex.ParseOptions{

		Width:  g.screen.Width,
		Height: g.screen.Height,

		Components: furex.ComponentsMap{
			"circle-button": func() *furex.View {
				return &furex.View{
					Handler: &widgets.Button{
						Color:   color.RGBA{0, 0, 0, 255},
						OnClick: func() { println("circle button clicked") },
					}}
			},
			"start-button": func() *furex.View {
				return &furex.View{
					Handler: &widgets.Button{
						Color:   color.RGBA{0, 0, 0, 255},
						OnClick: func() { println("rect button clicked") },
					}}
			},
			"label": func() *furex.View {
				return &furex.View{
					Handler: &widgets.Label{
						Color:    color.RGBA{0, 0, 0, 255},
						FontSize: 18,
					}}
			},
			"test-checkbox": func() *furex.View {
				return &furex.View{
					Handler: &widgets.CheckBox{
						Color:   color.RGBA{0, 0, 0, 255},
						OnClick: func() { println("test checkbox toggled") },
					}}
			},
			"test-container": func() *furex.View {
				return &furex.View{
					Handler: &widgets.RadioButtonContainer{
						Buttons:     []string{"testRB1", "testRB2", "testRB3", "testRB4"},
						CurrState:   []bool{false, false, false, false},
						CurrPressed: -1,
					}}
			},
			"test-rb": func() *furex.View {
				return &furex.View{
					Handler: &widgets.CheckBox{
						Color:   color.RGBA{0, 0, 0, 255},
						OnClick: func() { println("test rb toggled") },
					}}
			},
			"textbox": func() *furex.View {
				return &furex.View{
					Handler: &widgets.TextBox{
						Color:    color.RGBA{0, 0, 0, 255},
						FontSize: 20,
						Text:     "test",
						Counter:  0,
					}}
			},
		},
	})

	g.gameUI.AddChild(
		&furex.View{
			Width:    g.screen.Width,
			Height:   g.screen.Height,
			Position: furex.PositionAbsolute,
			Left:     0,
			Top:      0,
			Handler:  &widgets.Mouse{},
		},
	)
}

// "radio": func() *furex.View {
// 	return furex.Parse(mainHTML, &furex.ParseOptions{
// 		Width:  g.screen.Width,
// 		Height: g.screen.Height,

// 		Components: furex.ComponentsMap{
// 			"mecheckbox": func() *furex.View {
// 				return &furex.View{
// 					Handler: &widgets.MEcheckBox{
// 						Color:   color.RGBA{0, 0, 0, 255},
// 						OnClick: func() { println("checkbox toggled") },
// 					}}
// 			},
// 		},
// 	})
// },

// <radio class="radioButtons" orientation="vertical" id="testButtons">
// 	   <mecheckbox class="MEcheckbox" for="testButtons"></mecheckbox>
//     <mecheckbox class="MEcheckbox" for="testButtons"></mecheckbox>
//     <mecheckbox class="MEcheckbox" for="testButtons"></mecheckbox>
// </radio>

func main() {
	ebiten.SetWindowSize(1000, 750)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
