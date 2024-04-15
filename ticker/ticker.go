package ticker

import "syscall/js"

type Ticker struct {
	jsv js.Value
}

func (v *Ticker) Add(fn func(t *Ticker)) {
	v.jsv.Call("add", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		fn(v)
		return nil
	}))
}

func (v *Ticker) DeltaTime() float64 {
	return v.jsv.Get("deltaTime").Float()
}

func (v *Ticker) JSV() js.Value {
	return v.jsv
}

func (v *Ticker) SetJSV(value js.Value) {
	v.jsv = value
}
