package rectangle

import "syscall/js"

type Rectangle struct {
	jsv js.Value
}

func (v *Rectangle) JSV() js.Value {
	return v.jsv
}

func (v *Rectangle) SetJSV(value js.Value) {
	v.jsv = value
}

func (v *Rectangle) Width() float64 {
	return v.jsv.Get("width").Float()
}

func (v *Rectangle) Height() float64 {
	return v.jsv.Get("height").Float()
}
