package renderer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/zeozeozeo/microui-go"
)

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 40
		interval = 4
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

// updates the input state of the context
func (mgr *Manager) Update() {
	ctx := mgr.Ctx

	// mouse movement
	cx, cy := ebiten.CursorPosition()
	if cx != mgr.cx || cy != mgr.cy {
		ctx.InputMouseMove(cx, cy)
		mgr.cx, mgr.cy = cx, cy
	}

	// scrollwheel
	wx, wy := ebiten.Wheel()
	if wx != 0 || wy != 0 {
		ctx.InputScroll(int(wx*-30), int(wy*-30))
	}

	// keyboard input
	chars := ebiten.AppendInputChars(nil)
	if len(chars) > 0 {
		ctx.InputText(chars)
	}

	// mouse down
	var buttonsDown int
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		buttonsDown |= microui.MU_MOUSE_LEFT
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonMiddle) {
		buttonsDown |= microui.MU_MOUSE_MIDDLE
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		buttonsDown |= microui.MU_MOUSE_RIGHT
	}
	if buttonsDown != 0 {
		ctx.InputMouseDown(cx, cy, buttonsDown)
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
		ctx.InputMouseUp(cx, cy, buttonsUp)
	}

	// modifiers
	if inpututil.IsKeyJustPressed(ebiten.KeyAlt) {
		ctx.InputKeyDown(microui.MU_KEY_ALT)
	} else if inpututil.IsKeyJustReleased(ebiten.KeyAlt) {
		ctx.InputKeyUp(microui.MU_KEY_ALT)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyControl) {
		ctx.InputKeyDown(microui.MU_KEY_CTRL)
	} else if inpututil.IsKeyJustReleased(ebiten.KeyControl) {
		ctx.InputKeyUp(microui.MU_KEY_CTRL)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		ctx.InputKeyDown(microui.MU_KEY_RETURN)
	} else if inpututil.IsKeyJustReleased(ebiten.KeyEnter) {
		ctx.InputKeyUp(microui.MU_KEY_RETURN)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyShift) {
		ctx.InputKeyDown(microui.MU_KEY_SHIFT)
	} else if inpututil.IsKeyJustReleased(ebiten.KeyShift) {
		ctx.InputKeyUp(microui.MU_KEY_SHIFT)
	}

	// repeating keys
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		ctx.InputKeyDown(microui.MU_KEY_BACKSPACE)
		mgr.isHoldingBackspace = true
	} else if mgr.isHoldingBackspace {
		ctx.InputKeyUp(microui.MU_KEY_BACKSPACE)
		mgr.isHoldingBackspace = false
	}
}
