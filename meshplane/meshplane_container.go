// Code generated by cmd/gen --src_dir ../container --src_type Container --dest_type MeshPlane --dest_package meshplane. DO NOT EDIT.
package meshplane

import (
	"github.com/negasus/pixigo"
	"github.com/negasus/pixigo/observable"
)

func (v *MeshPlane) AddChild(c ...pixigo.Container) {
	for _, child := range c {
		v.jsv.Call("addChild", child.JSV())
	}
}
func (v *MeshPlane) X() float64 {
	return v.jsv.Get("x").Float()
}
func (v *MeshPlane) Y() float64 {
	return v.jsv.Get("y").Float()
}
func (v *MeshPlane) SetX(x float64) {
	v.jsv.Set("x", x)
}
func (v *MeshPlane) SetY(y float64) {
	v.jsv.Set("y", y)
}
func (v *MeshPlane) Pivot() *observable.Point {
	return v.pivot
}
func (v *MeshPlane) Width() float64 {
	return v.jsv.Get("width").Float()
}
func (v *MeshPlane) Height() float64 {
	return v.jsv.Get("height").Float()
}
func (v *MeshPlane) SetRotation(i float64) {
	v.jsv.Set("rotation", i)
}
func (v *MeshPlane) Rotation() float64 {
	return v.jsv.Get("rotation").Float()
}
func (v *MeshPlane) Position() *observable.Point {
	return v.position
}
