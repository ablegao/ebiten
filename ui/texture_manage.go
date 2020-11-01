package ui

import (
	"bytes"
	"encoding/xml"
	"image"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

var globalTextureManger *TextureManager

func init() {

	globalTextureManger = new(TextureManager)
	globalTextureManger.cache = map[string]*ebiten.Image{}
	globalTextureManger.isGlobal = true
}

func NewTextureManager() (textureManager *TextureManager) {
	textureManager = new(TextureManager)
	textureManager.cache = map[string]*ebiten.Image{}
	textureManager.isGlobal = false
	return
}

type TextureManager struct {
	cache    map[string]*ebiten.Image
	rw       sync.RWMutex
	isGlobal bool
}

func (tm *TextureManager) Add(name string, img *ebiten.Image) {
	tm.rw.Lock()
	tm.cache[name] = img
	tm.rw.Unlock()
	// ("insert into :", name)
}
func (tm *TextureManager) GetSprite(name string) *ebiten.Image {
	tm.rw.RLock()

	v, ok := tm.cache[name]
	if !ok {
		v = globalTextureManger.GetSprite(name)
		if v == nil {
			panic("Texture " + name + " not found")
		}
	}
	tm.rw.RUnlock()
	return v
}

func (tm *TextureManager) Destroy() {
	// do some
	tm.rw.Lock()
	for k, v := range tm.cache {
		delete(tm.cache, k)
		_ = v
	}
	tm.rw.Unlock()
}

func (tm *TextureManager) LoadTexture(fname, fxml []byte) {
	MarshalTexture(tm, fname, fxml)
}

func GetSprite(name string) *ebiten.Image {
	return globalTextureManger.GetSprite(name)
}

func LoadTexture(fname, fxml []byte) (out *TextureAtlas, err error) {
	return MarshalTexture(globalTextureManger, fname, fxml)
}

func LoadResImg(name string, res []byte) (*ebiten.Image, error) {

	img, _, err := image.Decode(bytes.NewBuffer(res))
	if err != nil {
		return nil, err
	}
	eimg := ebiten.NewImageFromImage(img)
	globalTextureManger.Add(name, eimg)
	return eimg, nil
}

func MarshalTexture(manager *TextureManager, fname, fxml []byte) (out *TextureAtlas, err error) {

	out = new(TextureAtlas)
	out.frames = []string{}

	err = xml.Unmarshal(fxml, &out)
	if err != nil {
		return
	}
	var img image.Image
	buf := bytes.NewBuffer(fname)
	img, _, err = image.Decode(buf)
	if err != nil {
		return nil, err
	}

	out.baseFile = img

	err = out.Marshal(manager)
	if err != nil {
		return
	}
	out.baseFile = nil
	img = nil
	return

}
