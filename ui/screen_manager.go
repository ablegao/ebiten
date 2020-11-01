package ui

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

var DefaultGame = new(ScreenManager)

func SetScreenSize(x, y int) {
	DefaultGame.SetScreenSize(x, y)
}

func GetScreenSize() (int, int) {
	return DefaultGame.w, DefaultGame.h
}

type ScreenManager struct {
	ScreenWidth  int
	ScreenHeight int
	w            int
	h            int
	runingScreen Screen
	swScreen     WaitScreen

	locker sync.RWMutex
}

func (g *ScreenManager) SetScreenSize(x, y int) {
	g.locker.Lock()

	g.ScreenWidth = x
	g.ScreenHeight = y

	g.locker.Unlock()
}

func (g *ScreenManager) RunScreen(screen Screen, wait WaitScreen) {
	g.locker.Lock()

	if g.runingScreen != nil {
		g.runingScreen.Destroy()
		g.runingScreen = nil
	}

	if wait != nil {
		wait.Start()
		g.swScreen = wait
		// go screen.Start()
		go g.asyncChangeScreen(screen)

	} else {
		screen.Start()
		g.runingScreen = screen
	}
	g.locker.Unlock()
}

func (g *ScreenManager) asyncChangeScreen(screen Screen) {
	screen.Start()
	g.locker.Lock()
	g.runingScreen = screen
	g.swScreen = nil
	g.locker.Unlock()
}

// func (g *ScreenManager) RunScreenAsync(screen Screen) {
// 	screen.Start()
// 	g.locker.Lock()
// 	g.runingScreen = screen
// 	g.swScreen = nil
// 	g.locker.Unlock()
// }

func RunScreen(screen Screen, wait WaitScreen) {
	DefaultGame.RunScreen(screen, wait)
}
func RunScreenAsync(screen Screen) {
	go DefaultGame.asyncChangeScreen(screen)
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *ScreenManager) Update() error {
	g.locker.RLock()
	switch true {
	case g.swScreen != nil:
		err := g.swScreen.Update()
		if err == ErrWaitScreenEndForWork {
			g.locker.RUnlock()
			{
				g.locker.Lock()
				g.swScreen = nil
				g.locker.Unlock()
			}
			g.locker.RLock()
		}
	case g.runingScreen != nil:
		g.runingScreen.Update()
	}
	g.locker.RUnlock()
	return nil
}

// Draw draws the game screen.
// Draw is called ever5y frame (typically 1/60[s] for 60Hz display).
func (g *ScreenManager) Draw(screen *ebiten.Image) {
	g.locker.RLock()
	switch true {
	case g.swScreen != nil:
		g.swScreen.DrawScreen(screen)
	case g.runingScreen != nil:
		g.runingScreen.DrawScreen(screen)
	}
	g.locker.RUnlock()

}

func (g *ScreenManager) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	if g.ScreenWidth == 0 || g.ScreenHeight == 0 {
		g.w = outsideWidth
		g.h = outsideHeight
		return g.w, g.h
	} else {
		s := float64(outsideWidth) / float64(outsideHeight)
		g.w = int(float64(g.ScreenHeight) * s)
		g.h = g.ScreenHeight
		return g.w, g.h
	}
}
