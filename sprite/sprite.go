package sprite

import (
	"syscall/js"

	"github.com/negasus/pixigo"
	"github.com/negasus/pixigo/observable"
	"github.com/negasus/pixigo/texture"
)

//go:generate go run ../cmd/gen --src_dir ../container --src_type Container --dest_type Sprite --dest_package sprite

type Sprite struct {
	jsv      js.Value
	anchor   *observable.Point
	scale    *observable.Point
	pivot    *observable.Point // Container
	position *observable.Point // Container
}

type Options struct {
	Texture *texture.Texture
	Width   float64
	Height  float64
	//Anchor    *observable.Point
	Anchor    float64
	BlendMode string
	Position  *observable.Point
}

func (o Options) marshal() map[string]any {
	args := map[string]any{}

	if o.Texture != nil {
		args["texture"] = o.Texture.JSV()
	}

	if o.Width != 0 {
		args["width"] = o.Width
	}

	if o.Height != 0 {
		args["height"] = o.Height
	}

	if o.Anchor != 0 {
		args["anchor"] = o.Anchor
	}

	if o.BlendMode != "" {
		args["blendMode"] = o.BlendMode
	}

	if o.Position != nil {
		args["position"] = o.Position.Marshal()
	}

	return args

}

func New(opts Options) *Sprite {
	s := pixigo.PIXI().Get("Sprite").New(opts.marshal())

	v := &Sprite{
		jsv:    s,
		anchor: &observable.Point{},
		scale:  &observable.Point{},
	}

	v.anchor.SetJSV(s.Get("anchor"))
	v.scale.SetJSV(s.Get("scale"))

	return v
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
