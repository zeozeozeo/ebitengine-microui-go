package renderer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/zeozeozeo/microui-go"
)

// updates the input state of the context
func (mgr *Manager) Update() {
	ctx := mgr.Ctx
	mx, my := ebiten.CursorPosition()
	ctx.InputMouseMove(mx, my)

	// mouse down
	var buttonsDown int
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		buttonsDown |= microui.MU_MOUSE_LEFT
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
		buttonsDown |= microui.MU_MOUSE_MIDDLE
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		buttonsDown |= microui.MU_MOUSE_RIGHT
	}
	if buttonsDown != 0 {
		ctx.InputMouseDown(mx, my, buttonsDown)
	}

	// mouse up
	var buttonsUp int
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		buttonsUp |= microui.MU_MOUSE_LEFT
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonMiddle) {
		buttonsUp |= microui.MU_MOUSE_MIDDLE
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		buttonsUp |= microui.MU_MOUSE_RIGHT
	}
	if buttonsUp != 0 {
		ctx.InputMouseUp(mx, my, buttonsUp)
	}

	// text input
	ctx.InputText(ebiten.AppendInputChars(nil))

	// key pressed
	if ebiten.IsKeyPressed(ebiten.KeyControl) {
		ctx.InputKeyDown(microui.MU_KEY_CTRL)
	}
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		ctx.InputKeyDown(microui.MU_KEY_SHIFT)
	}
	if ebiten.IsKeyPressed(ebiten.KeyAlt) {
		ctx.InputKeyDown(microui.MU_KEY_ALT)
	}
	if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		ctx.InputKeyDown(microui.MU_KEY_BACKSPACE)
	}
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		ctx.InputKeyDown(microui.MU_KEY_RETURN)
	}
	// key released
	if inpututil.IsKeyJustReleased(ebiten.KeyControl) {
		ctx.InputKeyUp(microui.MU_KEY_CTRL)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyShift) {
		ctx.InputKeyUp(microui.MU_KEY_SHIFT)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyAlt) {
		ctx.InputKeyUp(microui.MU_KEY_ALT)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyBackspace) {
		ctx.InputKeyUp(microui.MU_KEY_BACKSPACE)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyEnter) {
		ctx.InputKeyUp(microui.MU_KEY_RETURN)
	}
}
