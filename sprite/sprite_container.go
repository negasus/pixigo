// Code generated by cmd/gen --src_dir ../container --src_type Container --dest_type Sprite --dest_package sprite. DO NOT EDIT.
package sprite

import (
	"github.com/negasus/pixigo"
	"github.com/negasus/pixigo/observable"
)

func (v *Sprite) AddChild(c ...pixigo.Container) {
	for _, child := range c {
		v.jsv.Call("addChild", child.JSV())
	}
}
func (v *Sprite) X() float64 {
	return v.jsv.Get("x").Float()
}
func (v *Sprite) Y() float64 {
	return v.jsv.Get("y").Float()
}
func (v *Sprite) SetX(x float64) {
	v.jsv.Set("x", x)
}
func (v *Sprite) SetY(y float64) {
	v.jsv.Set("y", y)
}
func (v *Sprite) Pivot() *observable.Point {
	return v.pivot
}
func (v *Sprite) Width() float64 {
	return v.jsv.Get("width").Float()
}
func (v *Sprite) Height() float64 {
	return v.jsv.Get("height").Float()
}
func (v *Sprite) SetRotation(i float64) {
	v.jsv.Set("rotation", i)
}
func (v *Sprite) Rotation() float64 {
	return v.jsv.Get("rotation").Float()
}
