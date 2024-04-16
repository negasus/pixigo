package bounds

import "syscall/js"

type Bounds struct {
	jsv js.Value
}

func (v *Bounds) X() float64 {
	return v.jsv.Get("x").Float()
}

func (v *Bounds) SetX(value float64) {
	v.jsv.Set("x", value)
}
