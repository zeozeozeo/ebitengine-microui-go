package main

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/zeozeozeo/ebitengine-microui-go/renderer"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	width, height = 1280, 720
	src           *text.GoTextFaceSource
	face          text.Face
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
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
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
	ebiten.SetTPS(ebiten.SyncWithFPS)
	ebiten.SetWindowSize(width, height)
	g := &Game{}
	g.mgr = renderer.NewManager(face, 14)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
