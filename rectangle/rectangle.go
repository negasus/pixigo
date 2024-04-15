package rectangle

import "syscall/js"

type Rectangle struct {
	jsv js.Value
}

func New(x float64, y float64, w float64, h float64) *Rectangle {
	r := &Rectangle{}

	r.jsv = js.Global().Get("PIXI").Get("Rectangle").New(x, y, w, h)

	return r
}

func (v *Rectangle) X() float64 {
	return v.jsv.Get("x").Float()
}

func (v *Rectangle) Y() float64 {
	return v.jsv.Get("y").Float()
}

func (v *Rectangle) Width() float64 {
	return v.jsv.Get("width").Float()
}

func (v *Rectangle) Height() float64 {
	return v.jsv.Get("height").Float()
}

func (v *Rectangle) JSV() js.Value {
	return v.jsv
}

func (v *Rectangle) SetJSV(value js.Value) {
	v.jsv = value
}
