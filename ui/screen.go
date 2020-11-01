package ui

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ErrWaitScreenEndForWork = errors.New("wait screend end of work")
)

type Screen interface {
	Update() error
	DrawScreen(screen *ebiten.Image)
	Start() error //
	Destroy()     //
}

type WaitScreen interface {
	Start() error
	Wait(now, waitTime int)
	Update() error
	DrawScreen(screen *ebiten.Image)
}
