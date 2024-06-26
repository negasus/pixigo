// Code generated by cmd/gen --src_dir ../container --src_type Container --dest_type Text --dest_package text. DO NOT EDIT.
package text

import (
	"github.com/negasus/pixigo"
	"github.com/negasus/pixigo/observable"
)

func (v *Text) AddChild(c ...pixigo.Container) {
	for _, child := range c {
		v.jsv.Call("addChild", child.JSV())
	}
}
func (v *Text) X() float64 {
	return v.jsv.Get("x").Float()
}
func (v *Text) Y() float64 {
	return v.jsv.Get("y").Float()
}
func (v *Text) SetX(x float64) {
	v.jsv.Set("x", x)
}
func (v *Text) SetY(y float64) {
	v.jsv.Set("y", y)
}
func (v *Text) Pivot() *observable.Point {
	return v.pivot
}
func (v *Text) Width() float64 {
	return v.jsv.Get("width").Float()
}
func (v *Text) Height() float64 {
	return v.jsv.Get("height").Float()
}
func (v *Text) SetRotation(i float64) {
	v.jsv.Set("rotation", i)
}
func (v *Text) Rotation() float64 {
	return v.jsv.Get("rotation").Float()
}
func (v *Text) Position() *observable.Point {
	return v.position
}
