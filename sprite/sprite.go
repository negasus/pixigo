package sprite

import (
	"syscall/js"

	"github.com/negasus/pixigo"
	"github.com/negasus/pixigo/texture"
)

type Sprite struct {
	jsv js.Value
}

func FromTexture(txt *texture.Texture) *Sprite {
	s := pixigo.PIXI().Get("Sprite").New(txt.JSV())

	v := &Sprite{
		jsv: s,
	}

	return v
}

func (v *Sprite) SetJSV(value js.Value) {
	v.jsv = value
}

func (v *Sprite) JSV() js.Value {
	return v.jsv
}

func (v *Sprite) SetX(x float64) {
	v.jsv.Set("x", x)
}

func (v *Sprite) SetY(y float64) {
	v.jsv.Set("y", y)
}
