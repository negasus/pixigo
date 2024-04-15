package pixigo

import (
	"syscall/js"
)

type Container interface {
	SetJSV(js.Value)
	JSV() js.Value
}

func PIXI() js.Value {
	return js.Global().Get("PIXI")
}

func Await(v js.Value, name string, args ...any) (then, catch []js.Value) {
	thenChan := make(chan []js.Value)
	catchChan := make(chan []js.Value)
	v.Call(name, args...).
		Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
			thenChan <- args
			return nil
		})).
		Call("catch", js.FuncOf(func(this js.Value, args []js.Value) any {
			catchChan <- args
			return nil
		}))

	select {
	case then = <-thenChan:
		return then, nil
	case catch = <-catchChan:
		return nil, catch
	}
}
