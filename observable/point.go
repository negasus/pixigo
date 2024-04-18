package observable

import "syscall/js"

type Point struct {
	jsv js.Value
}

func New(x, y float64) *Point {
	v := &Point{}
	v.jsv = js.Global().Get("PIXI").Get("ObservablePoint").New(x, y)
	return v
}

func (v *Point) Marshal() map[string]any {
	return map[string]any{
		"x": v.X(),
		"y": v.Y(),
	}
}

func (v *Point) X() float64       { return v.jsv.Get("x").Float() }
func (v *Point) Y() float64       { return v.jsv.Get("y").Float() }
func (v *Point) SetX(x float64)   { v.jsv.Set("x", x) }
func (v *Point) SetY(y float64)   { v.jsv.Set("y", y) }
func (v *Point) Set(x, y float64) { v.jsv.Call("set", x, y) }

func (v *Point) JSV() js.Value         { return v.jsv }
func (v *Point) SetJSV(value js.Value) { v.jsv = value }
