package sprite

import (
	"syscall/js"

	"github.com/negasus/pixigo"
	"github.com/negasus/pixigo/observable"
	"github.com/negasus/pixigo/texture"
)

type Sprite struct {
	jsv    js.Value
	anchor *observable.ObservablePoint
	scale  *observable.ObservablePoint
}

func FromTexture(txt *texture.Texture) *Sprite {
	s := pixigo.PIXI().Get("Sprite").New(txt.JSV())

	v := &Sprite{
		jsv:    s,
		anchor: &observable.ObservablePoint{},
		scale:  &observable.ObservablePoint{},
	}

	v.anchor.SetJSV(s.Get("anchor"))
	v.scale.SetJSV(s.Get("scale"))

	return v
}

func (v *Sprite) X() float64 {
	return v.jsv.Get("x").Float()
}

func (v *Sprite) Y() float64 {
	return v.jsv.Get("y").Float()
}

func (v *Sprite) SetX(x float64) {
	v.jsv.Set("x", x)
}

func (v *Sprite) SetY(y float64) {
	v.jsv.Set("y", y)
}

func (v *Sprite) Anchor() *observable.ObservablePoint {
	return v.anchor
}

func (v *Sprite) Scale() *observable.ObservablePoint {
	return v.scale
}

func (v *Sprite) SetJSV(value js.Value) {
	v.jsv = value
}

func (v *Sprite) JSV() js.Value {
	return v.jsv
}

func (v *Sprite) SetTint(value float64) {
	v.jsv.Set("tint", value)
}

func (v *Sprite) SetRotation(f float64) {
	v.jsv.Set("rotation", f)
}
