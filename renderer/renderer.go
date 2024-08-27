package renderer

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zeozeozeo/ebitengine-microui-go/icons"
	"github.com/zeozeozeo/microui-go"
)

type Manager struct {
	Ctx                *microui.Context // microui context
	Face               text.Face        // font face
	FaceHeight         int
	ClipMask           bool // whether to use clip mask
	cx                 int
	cy                 int
	isHoldingBackspace bool
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
	target := screen
	ctx.Render(func(cmd *microui.Command) {
		mgr.renderCommand(cmd, &target, screen)
	})
}

func (mgr *Manager) renderCommand(cmd *microui.Command, target **ebiten.Image, origScreen *ebiten.Image) {
	screen := *target
	switch cmd.Type {
	case microui.MU_COMMAND_CLIP:
		if !mgr.ClipMask {
			return
		}
		if cmd.Clip.Rect == microui.UnclippedRect {
			*target = origScreen
			return
		}
		*target = screen.SubImage(image.Rect(
			cmd.Clip.Rect.X,
			cmd.Clip.Rect.Y,
			min(cmd.Clip.Rect.X+cmd.Clip.Rect.W, screen.Bounds().Dx()),
			min(cmd.Clip.Rect.Y+cmd.Clip.Rect.H, screen.Bounds().Dy()),
		)).(*ebiten.Image)
	case microui.MU_COMMAND_RECT:
		mgr.renderRect(cmd.Rect, screen)
	case microui.MU_COMMAND_TEXT:
		mgr.renderText(cmd.Text, screen)
	case microui.MU_COMMAND_ICON:
		mgr.renderIcon(cmd.Icon, screen)
	}
}

func (mgr *Manager) renderRect(cmd microui.RectCommand, screen *ebiten.Image) {
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

func (mgr *Manager) renderText(cmd microui.TextCommand, screen *ebiten.Image) {
	x := cmd.Pos.X
	y := cmd.Pos.Y

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(cmd.Color.ToRGBA())

	text.Draw(screen, cmd.Str, mgr.Face, op)
}

func (mgr *Manager) renderIcon(cmd microui.IconCommand, screen *ebiten.Image) {
	icon := icons.DefaultIcons[cmd.Id]

	x := cmd.Rect.X + (cmd.Rect.W-icon.Bounds().Dx())/2
	y := cmd.Rect.Y + (cmd.Rect.H-icon.Bounds().Dy())/2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(icon, op)

}
