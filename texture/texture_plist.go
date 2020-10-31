package texture

import (
	"bytes"
	"encoding/xml"
	"image"
	"image/png"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten"
)

type TextureAtlas struct {
	XMLName xml.Name `xml:"TextureAtlas"`
	Version string   `xml:"version,attr"`

	Sprite   []*TextureSprite `xml:"sprite"`
	Width    int              `xml:"width,attr"`
	Height   int              `xml:"height,attr"`
	frames   []string
	baseFile image.Image
}

func (tp *TextureAtlas) Marshal(manager *TextureManager) error {
	for _, v := range tp.Sprite {
		nImg := v.Copy(tp.baseFile)
		// tp.SaveToFiles(v.Name+".png", nImg, "./out/")
		if img, err := ebiten.NewImageFromImage(nImg, ebiten.FilterLinear); err == nil {
			tp.frames = append(tp.frames, v.Name)
			manager.Add("#"+v.Name, img)
		}
	}

	return nil
}

func (tp *TextureAtlas) SaveToFiles(k string, v image.Image, out string) error {
	// debuf file
	buf := bytes.NewBuffer(nil)
	err := png.Encode(buf, v)
	if err != nil {
		return err
	}
	fo, _ := os.OpenFile(out+k, os.O_CREATE|os.O_WRONLY, 0666)
	fo.Write(buf.Bytes())
	fo.Close()
	return nil
}
