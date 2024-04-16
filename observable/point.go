package observable

import "syscall/js"

type Point struct {
	jsv js.Value
}

func (v *Point) X() float64       { return v.jsv.Get("x").Float() }
func (v *Point) Y() float64       { return v.jsv.Get("y").Float() }
func (v *Point) SetX(x float64)   { v.jsv.Set("x", x) }
func (v *Point) SetY(y float64)   { v.jsv.Set("y", y) }
func (v *Point) Set(x, y float64) { v.jsv.Call("set", x, y) }

func (v *Point) JSV() js.Value         { return v.jsv }
func (v *Point) SetJSV(value js.Value) { v.jsv = value }
