package main

import (
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/yohamta/furex/examples/game/sprites"
	//"github.com/yohamta/furex/examples/game/text"

	

	"github.com/yohamta/furex/examples/game/widgets"
	"github.com/yohamta/furex/v2"

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
	screen.Fill(color.RGBA{63, 124, 182, 255})
	g.gameUI.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.screen.Width = outsideWidth
	g.screen.Height = outsideHeight
	return g.screen.Width, g.screen.Height
}

func NewGame() (*Game, error) {
	game := &Game{}
	return game, nil
}

//go:embed assets/html/main.html
var mainHTML string

func (g *Game) setupUI() {

	// Setup the UI parsed from HTML.
	g.gameUI = furex.Parse(mainHTML, &furex.ParseOptions{

		Width:  g.screen.Width,
		Height: g.screen.Height,

		Components: furex.ComponentsMap{
			"bottom-button": func() *furex.View {
				return &furex.View{
					Width:  45,
					Height: 49,
					Handler: &widgets.Button{
						Color:   color.RGBA{210, 178, 144, 255},
						OnClick: func() { println("button clicked") },
					}}
			},
		},
	})

	// // panels that draws mouse cursor
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

func main() {
	ebiten.SetWindowSize(480, 640)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
