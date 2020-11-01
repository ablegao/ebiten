package ui

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewButton(img *ebiten.Image) *Button {
	return &Button{
		img: img}
}

type ButtonCallback func()

type Button struct {
	img                 *ebiten.Image
	onPressUpCallback   ButtonCallback
	onPressDownCallback ButtonCallback
	press               bool
	Pos                 image.Rectangle
}

func (bc *Button) SetOnPressUpListener(fun ButtonCallback) {
	bc.onPressUpCallback = fun
}

func (bc *Button) SetOnPressDownListener(fun ButtonCallback) {
	bc.onPressDownCallback = fun
}

func (bc *Button) Size() (x, y int) {

	return bc.img.Size()
}
func (bc *Button) Update() {

	var p image.Point
	touchids := ebiten.TouchIDs()
	if len(touchids) > 0 {
		touched := false
		for _, touchid := range touchids {
			p.X, p.Y = ebiten.TouchPosition(touchid)
			if p.In(bc.Pos) {
				touched = true
				break
			}
		}
		if touched {
			bc.changeStatus(true)
		} else {
			bc.changeStatus(false)
		}
	} else {
		p.X, p.Y = ebiten.CursorPosition()
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && p.In(bc.Pos) {
			bc.changeStatus(true)
		} else if bc.press {
			bc.changeStatus(false)
		}
	}
}

func (bc *Button) changeStatus(press bool) {
	if bc.press && !press {
		if bc.onPressUpCallback != nil {
			bc.onPressUpCallback()
		}
	} else if !bc.press && press {
		if bc.onPressDownCallback != nil {
			bc.onPressDownCallback()
		}
	}

	bc.press = press
}

func (bc *Button) DrawTo(screen *ebiten.Image, x, y float64) {
	opt := &ebiten.DrawImageOptions{}
	// opt.GeoM.Translate()
	bx, by := bc.Size()

	opt.GeoM.Translate(-float64(bx/2), -float64(by/2))
	if bc.press {
		opt.GeoM.Scale(1.1, 1.1)
	}
	// x, _ := famtree.AlignOption(screen.Bounds(), bc.img.Bounds(), bc.Align, bc.VAlign)
	bc.Pos = image.Rectangle{
		Min: image.Point{int(x) - bx/2, int(y) - by/2},
		Max: image.Point{int(x) + bx/2, int(y) + by/2},
	}
	opt.GeoM.Translate(x, y)
	screen.DrawImage(bc.img, opt)
}
