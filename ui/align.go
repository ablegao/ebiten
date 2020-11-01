package ui

import "image"

func AlignOption(screenRet image.Rectangle, imgRet image.Rectangle, align, valign AlignType) (x, y float64) {

	switch align {
	case AlignLeft:
		x = float64(imgRet.Min.X)
	case AlignCenter:
		x = float64(screenRet.Max.X/2 - imgRet.Max.X/2)
	case AlignRight:
		x = float64(screenRet.Max.X - imgRet.Max.X)
	}

	switch valign {

	case VerticalAilgnTop:
		y = float64(imgRet.Min.Y)
	case VerticalAlignMiddle:
		y = float64(screenRet.Max.Y/2 - imgRet.Max.Y/2)
	case VerticalAlignBottom:
		y = float64(screenRet.Max.Y - imgRet.Max.Y)
	}
	return
}
