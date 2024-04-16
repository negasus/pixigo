package sprite

import (
	"syscall/js"

	"github.com/negasus/pixigo"
	"github.com/negasus/pixigo/observable"
	"github.com/negasus/pixigo/texture"
)

//go:generate go run ../cmd/gen --src_dir ../container --src_type Container --dest_type Sprite --dest_package sprite

// Sprite
// @inherits container.Container
type Sprite struct {
	jsv    js.Value
	anchor *observable.Point
	scale  *observable.Point
	pivot  *observable.Point // Container
}

func FromTexture(txt *texture.Texture) *Sprite {
	s := pixigo.PIXI().Get("Sprite").New(txt.JSV())

	v := &Sprite{
		jsv:    s,
		anchor: &observable.Point{},
		scale:  &observable.Point{},
	}

	v.anchor.SetJSV(s.Get("anchor"))
	v.scale.SetJSV(s.Get("scale"))

	return v
}

func (v *Sprite) Anchor() *observable.Point {
	return v.anchor
}

func (v *Sprite) Scale() *observable.Point {
	return v.scale
}

func (v *Sprite) SetTint(value float64) {
	v.jsv.Set("tint", value)
}

func (v *Sprite) JSV() js.Value         { return v.jsv }
func (v *Sprite) SetJSV(value js.Value) { v.jsv = value }
