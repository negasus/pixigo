package graphics

import (
	"syscall/js"

	"github.com/negasus/pixigo"
)

type FillStyleInputs struct {
	Width int
	Color int
}

func (i *FillStyleInputs) marshal() map[string]any {
	res := map[string]any{}

	if i.Width != 0 {
		res["width"] = i.Width
	}

	if i.Color != 0 {
		res["color"] = i.Color
	}

	return res
}

type Graphics struct {
	jsv js.Value
}

func New() *Graphics {
	v := &Graphics{}

	v.jsv = pixigo.PIXI().Get("Graphics").New()

	return v
}

func (v *Graphics) Rect(x, y, w, h float64) {
	v.jsv.Call("rect", x, y, w, h)
}

func (v *Graphics) Circle(x, y, r float64) {
	v.jsv.Call("circle", x, y, r)
}

func (v *Graphics) Fill(color int) {
	v.jsv.Call("fill", color)
}

func (v *Graphics) Stroke(i *FillStyleInputs) {
	v.jsv.Call("stroke", i.marshal())
}

func (v *Graphics) JSV() js.Value {
	return v.jsv
}

func (v *Graphics) SetJSV(jsv js.Value) {
	v.jsv = jsv
}
