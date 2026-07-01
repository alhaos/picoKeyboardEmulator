package main

import (
	"machine"
	"machine/usb/hid/keyboard"
	"time"
)

func main() {
	kb := keyboard.Port()

	led := machine.GP1
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led.High() // поменять на led.Low() если нужно выключить, и перепрошить

	pressKey := func(k keyboard.Keycode, withShift bool) {
		if withShift {
			kb.Down(keyboard.KeyLeftShift)
			time.Sleep(10 * time.Millisecond)
		}
		kb.Down(k)
		time.Sleep(10 * time.Millisecond)
		kb.Up(k)
		if withShift {
			time.Sleep(10 * time.Millisecond)
			kb.Up(keyboard.KeyLeftShift)
		}
		time.Sleep(30 * time.Millisecond)
	}

	type keyDef struct {
		key   keyboard.Keycode
		shift bool
	}

	charMap := map[rune]keyDef{
		' ':  {keyboard.KeySpace, false},
		'!':  {keyboard.Key1, true},
		'"':  {keyboard.KeyQuote, true},
		'#':  {keyboard.Key3, true},
		'$':  {keyboard.Key4, true},
		'%':  {keyboard.Key5, true},
		'&':  {keyboard.Key7, true},
		'\'': {keyboard.KeyQuote, false},
		'(':  {keyboard.Key9, true},
		')':  {keyboard.Key0, true},
		'*':  {keyboard.Key8, true},
		'+':  {keyboard.KeyEqual, true},
		',':  {keyboard.KeyComma, false},
		'-':  {keyboard.KeyMinus, false},
		'.':  {keyboard.KeyPeriod, false},
		'/':  {keyboard.KeySlash, false},
		'0':  {keyboard.Key0, false},
		'1':  {keyboard.Key1, false},
		'2':  {keyboard.Key2, false},
		'3':  {keyboard.Key3, false},
		'4':  {keyboard.Key4, false},
		'5':  {keyboard.Key5, false},
		'6':  {keyboard.Key6, false},
		'7':  {keyboard.Key7, false},
		'8':  {keyboard.Key8, false},
		'9':  {keyboard.Key9, false},
		':':  {keyboard.KeySemicolon, true},
		';':  {keyboard.KeySemicolon, false},
		'<':  {keyboard.KeyComma, true},
		'=':  {keyboard.KeyEqual, false},
		'>':  {keyboard.KeyPeriod, true},
		'?':  {keyboard.KeySlash, true},
		'@':  {keyboard.Key2, true},
		'A':  {keyboard.KeyA, true}, 'B': {keyboard.KeyB, true}, 'C': {keyboard.KeyC, true},
		'D': {keyboard.KeyD, true}, 'E': {keyboard.KeyE, true}, 'F': {keyboard.KeyF, true},
		'G': {keyboard.KeyG, true}, 'H': {keyboard.KeyH, true}, 'I': {keyboard.KeyI, true},
		'J': {keyboard.KeyJ, true}, 'K': {keyboard.KeyK, true}, 'L': {keyboard.KeyL, true},
		'M': {keyboard.KeyM, true}, 'N': {keyboard.KeyN, true}, 'O': {keyboard.KeyO, true},
		'P': {keyboard.KeyP, true}, 'Q': {keyboard.KeyQ, true}, 'R': {keyboard.KeyR, true},
		'S': {keyboard.KeyS, true}, 'T': {keyboard.KeyT, true}, 'U': {keyboard.KeyU, true},
		'V': {keyboard.KeyV, true}, 'W': {keyboard.KeyW, true}, 'X': {keyboard.KeyX, true},
		'Y': {keyboard.KeyY, true}, 'Z': {keyboard.KeyZ, true},
		'[':  {keyboard.KeyLeftBrace, false},
		'\\': {keyboard.KeyBackslash, false},
		']':  {keyboard.KeyRightBrace, false},
		'^':  {keyboard.Key6, true},
		'_':  {keyboard.KeyMinus, true},
		'`':  {keyboard.KeyTilde, false},
		'a':  {keyboard.KeyA, false}, 'b': {keyboard.KeyB, false}, 'c': {keyboard.KeyC, false},
		'd': {keyboard.KeyD, false}, 'e': {keyboard.KeyE, false}, 'f': {keyboard.KeyF, false},
		'g': {keyboard.KeyG, false}, 'h': {keyboard.KeyH, false}, 'i': {keyboard.KeyI, false},
		'j': {keyboard.KeyJ, false}, 'k': {keyboard.KeyK, false}, 'l': {keyboard.KeyL, false},
		'm': {keyboard.KeyM, false}, 'n': {keyboard.KeyN, false}, 'o': {keyboard.KeyO, false},
		'p': {keyboard.KeyP, false}, 'q': {keyboard.KeyQ, false}, 'r': {keyboard.KeyR, false},
		's': {keyboard.KeyS, false}, 't': {keyboard.KeyT, false}, 'u': {keyboard.KeyU, false},
		'v': {keyboard.KeyV, false}, 'w': {keyboard.KeyW, false}, 'x': {keyboard.KeyX, false},
		'y': {keyboard.KeyY, false}, 'z': {keyboard.KeyZ, false},
		'{': {keyboard.KeyLeftBrace, true},
		'|': {keyboard.KeyBackslash, true},
		'}': {keyboard.KeyRightBrace, true},
		'~': {keyboard.KeyTilde, true},
	}

	typeString := func(s string) {
		for _, c := range s {
			if entry, ok := charMap[c]; ok {
				pressKey(entry.key, entry.shift)
			}
		}
	}

	btn := machine.GP0
	btn.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	stable := true
	lastRaw := true
	var lastChange time.Time
	wasPressed := false

	for {
		raw := btn.Get()

		if raw != lastRaw {
			lastChange = time.Now()
			lastRaw = raw
		}

		if time.Since(lastChange) > 20*time.Millisecond {
			stable = raw
		}

		pressed := !stable
		if pressed && !wasPressed {
			typeString("your_pass_hera")
			kb.Down(keyboard.KeyEnter)
			time.Sleep(10 * time.Millisecond)
			kb.Up(keyboard.KeyEnter)
		}
		wasPressed = pressed

		time.Sleep(2 * time.Millisecond)
	}
}
