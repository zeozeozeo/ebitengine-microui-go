package renderer

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zeozeozeo/microui-go"
	"github.com/zeozeozeo/microui-go-ebiten/atlas"
)

type Manager struct {
	Ctx        *microui.Context // microui context
	Face       text.Face        // font face
	FaceHeight int
	ClipMask   bool // whether to use clip mask
	clip       microui.Rect
	mask       *ebiten.Image
}

func NewManagerWithContext(ctx *microui.Context, face text.Face, faceHeight int) *Manager {
	mgr := &Manager{
		Ctx:        ctx,
		Face:       face,
		FaceHeight: faceHeight,
		ClipMask:   true,
	}
	ctx.TextWidth = mgr.textWidth
	ctx.TextHeight = mgr.textHeight
	return mgr
}

func NewManager(face text.Face, faceHeight int) *Manager {
	ctx := microui.NewContext()
	return NewManagerWithContext(ctx, face, faceHeight)
}

func (mgr *Manager) textWidth(fnt microui.Font, str string) int {
	w, _ := text.Measure(str, mgr.Face, 0.0)
	return int(w)
}

func (mgr *Manager) textHeight(fnt microui.Font) int {
	return mgr.FaceHeight
}

func (mgr *Manager) Draw(screen *ebiten.Image) {
	ctx := mgr.Ctx
	ctx.Render(func(cmd *microui.Command) {
		mgr.renderCommand(cmd, screen)
	})
}

func (mgr *Manager) renderCommand(cmd *microui.Command, screen *ebiten.Image) {
	switch cmd.Type {
	case microui.MU_COMMAND_CLIP:
		if cmd.Clip.Rect == microui.UnclippedRect || !mgr.ClipMask {
			return
		}
		mgr.clip = cmd.Clip.Rect
		mgr.mask = ebiten.NewImage(mgr.clip.W, mgr.clip.H)
	case microui.MU_COMMAND_RECT:
		mgr.renderRect(cmd.Rect, screen)
	case microui.MU_COMMAND_TEXT:
		mgr.renderText(cmd.Text, screen)
	case microui.MU_COMMAND_ICON:
		mgr.renderIcon(cmd.Icon, screen)
	}
}

func (mgr *Manager) renderRect(cmd microui.RectCommand, screen *ebiten.Image) {
	if mgr.mask != nil {
		// draw to mask
		mgr.mask.Clear()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(
			float64(cmd.Rect.X-mgr.clip.X),
			float64(cmd.Rect.Y-mgr.clip.Y),
		)
		vector.DrawFilledRect(
			mgr.mask,
			0, 0,
			float32(cmd.Rect.W),
			float32(cmd.Rect.H),
			cmd.Color.ToRGBA(),
			false,
		)

		// draw mask to screen
		op2 := &ebiten.DrawImageOptions{}
		op2.GeoM.Translate(float64(cmd.Rect.X), float64(cmd.Rect.Y))
		screen.DrawImage(
			mgr.mask,
			op2,
		)
	} else {
		vector.DrawFilledRect(
			screen,
			float32(cmd.Rect.X),
			float32(cmd.Rect.Y),
			float32(cmd.Rect.W),
			float32(cmd.Rect.H),
			cmd.Color.ToRGBA(),
			false,
		)
	}
}

func (mgr *Manager) renderText(cmd microui.TextCommand, screen *ebiten.Image) {
	x := cmd.Pos.X
	y := cmd.Pos.Y

	if mgr.mask != nil {
		// draw to mask
		mgr.mask.Clear()
		op := &text.DrawOptions{}
		op.GeoM.Translate(float64(x-mgr.clip.X), float64(y-mgr.clip.Y))
		op.ColorScale.ScaleWithColor(cmd.Color.ToRGBA())

		text.Draw(mgr.mask, cmd.Str, mgr.Face, op)

		// draw mask to screen
		op2 := &ebiten.DrawImageOptions{}
		op2.GeoM.Translate(float64(mgr.clip.X), float64(mgr.clip.Y))
		screen.DrawImage(mgr.mask, op2)
	} else {
		op := &text.DrawOptions{}
		op.GeoM.Translate(float64(x), float64(y))
		op.ColorScale.ScaleWithColor(cmd.Color.ToRGBA())

		text.Draw(screen, cmd.Str, mgr.Face, op)
	}
}

func (mgr *Manager) renderIcon(cmd microui.IconCommand, screen *ebiten.Image) {
	if mgr.mask != nil {
		// draw to mask
		mgr.mask.Clear()
	}
	rect := atlas.DefaultAtlasRects[cmd.Id]
	x := cmd.Rect.X + (cmd.Rect.W-rect.W)/2
	y := cmd.Rect.Y + (cmd.Rect.H-rect.H)/2

	op := &ebiten.DrawImageOptions{}
	if mgr.mask != nil {
		op.GeoM.Translate(
			float64(x-mgr.clip.X),
			float64(y-mgr.clip.Y),
		)
	} else {
		op.GeoM.Translate(
			float64(x),
			float64(y),
		)
	}

	icon := atlas.DefaultAtlas.SubImage(image.Rect(rect.X, rect.Y, rect.X+rect.W, rect.Y+rect.H))
	if mgr.mask != nil {
		mgr.mask.DrawImage(icon.(*ebiten.Image), op)
		// draw mask to screen
		op2 := &ebiten.DrawImageOptions{}
		op2.GeoM.Translate(float64(mgr.clip.X), float64(mgr.clip.Y))
		screen.DrawImage(mgr.mask, op2)
	} else {
		// draw to screen
		op2 := &ebiten.DrawImageOptions{}
		op2.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(icon.(*ebiten.Image), op2)
	}
}
