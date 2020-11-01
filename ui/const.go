package ui

import "image"

var Point0 = image.Point{0, 0}

type ScreenStatus int

type AppStatus int

type AlignType int

const (
	AlignLeft AlignType = iota
	AlignCenter
	AlignRight
)

const (
	VerticalAilgnTop = iota
	VerticalAlignMiddle
	VALIGN_BOTTOM
	VerticalAlignBottom
)
