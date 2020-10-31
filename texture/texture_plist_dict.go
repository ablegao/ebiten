package texture

import (
	"image"
	"image/draw"
)

/*
 <TextureAtlas imagePath="/Users/ablegao/Downloads/craftpix-897854-cooking-match-3-game-kit/xxx-1.png" width="2024" height="2048">
    <sprite n="button-play-main" x="1449" y="1643" w="420" h="200"/>
    <sprite n="game-icon7" x="1" y="1923" w="180" h="124"/>
    <sprite n="map-bg" x="1" y="1" w="1080" h="1920"/>
    <sprite n="map-point-blue" x="463" y="1923" w="137" h="114"/>
    <sprite n="map-point-grey" x="183" y="1923" w="138" h="115"/>
    <sprite n="map-point-pink" x="323" y="1923" w="138" h="114"/>
    <sprite n="map-prize-container" x="1083" y="1643" w="364" h="383"/>
    <sprite n="map-stars1" x="602" y="1923" w="142" h="73"/>
    <sprite n="map-stars2" x="746" y="1923" w="142" h="73"/>
    <sprite n="map-stars3" x="890" y="1923" w="142" h="72"/>
    <sprite n="shop-coins-popup" x="1083" y="1" w="940" h="1640"/>
    <sprite n="wheel-button-quit" x="1449" y="1845" w="394" h="148"/>
    <sprite n="wheel-button-spin" x="1871" y="1643" w="149" h="394" r="y"/>
</TextureAtlas>

*/
type TextureSprite struct {
	Name string `xml:"n,attr"`
	X    int    `xml:"x,attr"`
	Y    int    `xml:"y,attr"`
	W    int    `xml:"w,attr"`
	H    int    `xml:"h,attr"`
	R    string `xml:"r,attr"`
}

var POINT_0 = image.Point{0, 0}

func (tp *TextureSprite) Copy(base image.Image) image.Image {
	r := image.Rect(0, 0, tp.W, tp.H)

	pic := image.NewRGBA(r)

	draw.Draw(pic, r, base, image.Point{tp.X, tp.Y}, draw.Src)

	if tp.IsRetated() {
		rotate270 := image.NewRGBA(image.Rect(0, 0, pic.Bounds().Dy(), pic.Bounds().Dx()))
		// 矩阵旋转
		for x := pic.Bounds().Min.Y; x < pic.Bounds().Max.Y; x++ {
			for y := pic.Bounds().Max.X - 1; y >= pic.Bounds().Min.X; y-- {
				// 设置像素点
				rotate270.Set(x, pic.Bounds().Max.X-y, pic.At(y, x))
			}
		}

		return rotate270
	}
	return pic
}

func (tp *TextureSprite) IsRetated() bool {
	return tp.R == "y"
}
