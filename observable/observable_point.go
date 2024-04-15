package observable

import "syscall/js"

type ObservablePoint struct {
	jsv js.Value
}

func (v *ObservablePoint) SetX(x float64) {
	v.jsv.Set("x", x)
}

func (v *ObservablePoint) SetY(y float64) {
	v.jsv.Set("y", y)
}

func (v *ObservablePoint) Set(x, y float64) {
	v.jsv.Call("set", x, y)
}

func (v *ObservablePoint) JSV() js.Value {
	return v.jsv
}

func (v *ObservablePoint) SetJSV(value js.Value) {
	v.jsv = value
}
