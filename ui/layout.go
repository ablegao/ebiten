package ui

import "encoding/xml"

/*
	Align 0,1,2
	Valign 1,2,3
	 -----------------
	|0,0  | 1,0 | 2,0 |
	|0,1  | 1,1 | 2,1 |
	|0,2  | 1,2 | 2,2 |
    -------------------
*/

type AlignLayoutSprite struct {
	Name   string    `xml:"name,attr"`
	Align  AlignType `xml:"align,attr"`
	VAlign AlignType `xml:"valign,attr"`
}
type AlignLayout struct {
	XMLName xml.Name            `xml:"AlignLayout"`
	Sprites []AlignLayoutSprite `xml:"sprites"`
}
