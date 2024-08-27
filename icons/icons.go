package icons

import (
	"bytes"
	_ "embed"
	"image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed mu_icon_close.png
var muIconClose []byte

//go:embed mu_icon_check.png
var muIconCheck []byte

//go:embed mu_icon_collapsed.png
var muIconCollapsed []byte

//go:embed mu_icon_expanded.png
var muIconExpanded []byte

//go:embed mu_icon_max.png
var muIconMax []byte

var (
	DefaultIcons []*ebiten.Image
)

func loadPng(data []byte) *ebiten.Image {
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

func init() {
	// load icons
	iconClose := loadPng(muIconClose)
	iconCheck := loadPng(muIconCheck)
	iconExpanded := loadPng(muIconExpanded)
	iconCollapsed := loadPng(muIconCollapsed)
	atlasWhite := loadPng(muIconMax)

	// set default icon order
	DefaultIcons = []*ebiten.Image{
		{},
		iconClose,
		iconCheck,
		iconCollapsed,
		iconExpanded,
		atlasWhite,
	}
}
