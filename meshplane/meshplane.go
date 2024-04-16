package meshplane

import (
	"syscall/js"

	"github.com/negasus/pixigo/texture"
)

//go:generate go run ../cmd/gen --src_dir ../container --src_type Container --dest_type MeshPlane --dest_package meshplane

type MeshPlane struct {
	jsv js.Value
}

type Options struct {
	Texture   *texture.Texture
	VerticesX int
	VerticesY int
}

func (o Options) marshal() js.Value {
	m := map[string]interface{}{}

	if o.Texture != nil {
		m["texture"] = o.Texture.JSV()
	}

	if o.VerticesX != 0 {
		m["verticesX"] = o.VerticesX
	}

	if o.VerticesY != 0 {
		m["verticesY"] = o.VerticesY
	}

	return js.ValueOf(m)
}

func New(options Options) *MeshPlane {
	v := &MeshPlane{
		jsv: js.Global().Get("PIXI").Get("MeshPlane").New(options.marshal()),
	}

	return v
}

func (v *MeshPlane) SetJSV(value js.Value) { v.jsv = value }
func (v *MeshPlane) JSV() js.Value         { return v.jsv }
