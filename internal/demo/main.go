package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/zeozeozeo/ebitengine-microui-go/renderer"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	src  *text.GoTextFaceSource
	face text.Face
)

type Game struct {
	mgr *renderer.Manager
}

func (g *Game) Update() error {
	g.mgr.Update() // updates microui input
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ui
	DrawUI(g.mgr.Ctx)
	g.mgr.Draw(screen)

	// tps/fps counter
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf(
			"tps: %.2f\nfps: %.2f",
			ebiten.ActualTPS(),
			ebiten.ActualFPS(),
		),
		5,
		5,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	sf := ebiten.Monitor().DeviceScaleFactor()
	return int(float64(outsideWidth) * sf), int(float64(outsideHeight) * sf)
}

func init() {
	// load font
	var err error

	src, err = text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal("err: ", err)
	}

	face = &text.GoTextFace{
		Source: src,
		Size:   14,
	}
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetTPS(ebiten.SyncWithFPS)

	g := &Game{}
	g.mgr = renderer.NewManager(face, 14)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
