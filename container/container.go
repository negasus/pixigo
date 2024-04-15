package container

import (
	"syscall/js"

	"github.com/negasus/pixigo"
	"github.com/negasus/pixigo/observable"
)

type Container struct {
	jsv   js.Value
	pivot *observable.ObservablePoint
}

func New() *Container {
	c := &Container{}

	c.jsv = pixigo.PIXI().Get("Container").New()
	c.pivot = &observable.ObservablePoint{}
	c.pivot.SetJSV(c.jsv.Get("pivot"))

	return c
}

func (v *Container) JSV() js.Value {
	return v.jsv
}

func (v *Container) SetJSV(value js.Value) {
	v.jsv = value
}

func (v *Container) AddChild(c ...pixigo.Container) {
	for _, child := range c {
		v.jsv.Call("addChild", child.JSV())
	}
}

func (v *Container) SetX(x float64) {
	v.jsv.Set("x", x)
}

func (v *Container) SetY(y float64) {
	v.jsv.Set("y", y)
}

func (v *Container) Pivot() *observable.ObservablePoint {
	return v.pivot
}

func (v *Container) Width() float64 {
	return v.jsv.Get("width").Float()
}

func (v *Container) Height() float64 {
	return v.jsv.Get("height").Float()
}

func (v *Container) SetRotation(i float64) {
	v.jsv.Set("rotation", i)
}

func (v *Container) Rotation() float64 {
	return v.jsv.Get("rotation").Float()
}
