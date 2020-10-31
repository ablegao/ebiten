package texture

import (
	"encoding/xml"
	"image"
	"io/ioutil"
	"os"
	"sync"

	"github.com/hajimehoshi/ebiten"
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

func (tm *TextureManager) LoadTexture(fname, fxml string) {
	MarshalTexture(tm, fname, fxml)
}

func GetSprite(name string) *ebiten.Image {
	return globalTextureManger.GetSprite(name)
}

func LoadTexture(fname, fxml string) (out *TextureAtlas, err error) {
	return MarshalTexture(globalTextureManger, fname, fxml)
}

func LoadResImg(res string) (*ebiten.Image, error) {
	var img image.Image

	img_f, err := os.Open(res)
	if err != nil {
		return nil, err
	}
	defer img_f.Close()
	img, _, err = image.Decode(img_f)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img, ebiten.FilterLinear)
}

func MarshalTexture(manager *TextureManager, fname, fxml string) (out *TextureAtlas, err error) {
	buf, err := ioutil.ReadFile(fxml)
	if err != nil {
		return
	}

	out = new(TextureAtlas)
	out.frames = []string{}

	err = xml.Unmarshal(buf, &out)
	if err != nil {
		return
	}
	var img image.Image

	img_f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer img_f.Close()
	img, _, err = image.Decode(img_f)
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
