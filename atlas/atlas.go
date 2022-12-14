package atlas

import (
	"bytes"
	_ "embed"
	"image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zeozeozeo/microui-go"
)

//go:embed default_atlas.png
var defaultAtlasPng []byte
var DefaultAtlas *ebiten.Image

var (
	iconClose     = microui.NewRect(0, 0, 16, 16)
	iconResize    = microui.NewRect(24, 24, 6, 6)
	iconCheck     = microui.NewRect(16, 0, 16, 16)
	iconCollapsed = microui.NewRect(32, 0, 16, 16)
	iconExpanded  = microui.NewRect(48, 0, 16, 16)
	atlasWhite    = microui.NewRect(2, 18, 3, 3)

	DefaultAtlasRects = []microui.Rect{
		{},
		iconClose,
		iconResize,
		iconCheck,
		iconCollapsed,
		iconExpanded,
		atlasWhite,
	}
)

func init() {
	// load atlas
	img, err := png.Decode(bytes.NewReader(defaultAtlasPng))
	if err != nil {
		panic(err)
	}
	DefaultAtlas = ebiten.NewImageFromImage(img)
}