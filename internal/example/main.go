package main

import (
	_ "embed"
	"fmt"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zeozeozeo/microui-go"
	"github.com/zeozeozeo/microui-go-ebiten/renderer"
	"golang.org/x/image/font"
)

//go:embed proggy.ttf
var proggyTtf []byte
var proggySize = 16
var proggyFace font.Face

var (
	width, height = 1280, 720
)

type Game struct {
	mgr *renderer.Manager
}

func (g *Game) Update() error {
	g.mgr.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ctx := g.mgr.Ctx
	ctx.Begin()
	if ctx.BeginWindow("Hello, world!", microui.NewRect(100, 100, 200, 400)) {
		if ctx.Button("some button") {
			fmt.Println("button pressed")
		}
		ctx.EndWindow()
	}
	ctx.End()

	g.mgr.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func init() {
	// load font
	f, err := truetype.Parse(proggyTtf)
	if err != nil {
		panic(err)
	}
	proggyFace = truetype.NewFace(f, &truetype.Options{
		DPI:     72,
		Size:    float64(proggySize),
		Hinting: font.HintingVertical,
	})
}

func main() {
	ebiten.SetTPS(ebiten.SyncWithFPS)
	ebiten.SetWindowSize(width, height)
	g := &Game{}
	g.mgr = renderer.NewManager(proggyFace, proggySize)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
