package texture

import "syscall/js"

type Texture struct {
	jsv js.Value
}

func (v *Texture) SetJSV(value js.Value) {
	v.jsv = value
}

func (v *Texture) JSV() js.Value {
	return v.jsv
}
