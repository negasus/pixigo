package text

import (
	"syscall/js"

	"github.com/negasus/pixigo"
	"github.com/negasus/pixigo/observable"
)

//go:generate go run ../cmd/gen --src_dir ../container --src_type Container --dest_type Text --dest_package text

type Style struct {
	FontFamily string
	FontSize   int
	FontStyle  string
	FontWeight string
	//Fill
	//Stroke
	//DropShadow
	WordWrap      bool
	WordWrapWidth int
}

func (s *Style) marshal() map[string]any {
	result := make(map[string]any)

	if s.FontFamily != "" {
		result["fontFamily"] = s.FontFamily
	}
	if s.FontSize != 0 {
		result["fontSize"] = s.FontSize
	}
	if s.FontStyle != "" {
		result["fontStyle"] = s.FontStyle
	}
	if s.FontWeight != "" {
		result["fontWeight"] = s.FontWeight
	}
	if s.WordWrap {
		result["wordWrap"] = s.WordWrap
	}
	if s.WordWrapWidth != 0 {
		result["wordWrapWidth"] = s.WordWrapWidth
	}

	return result
}

type Text struct {
	jsv js.Value

	pivot    *observable.Point // Container
	position *observable.Point // Container
}

func New(s string, style *Style) *Text {
	v := &Text{}

	v.jsv = pixigo.PIXI().Get("Text").New(s, style.marshal())

	return v
}

func (v *Text) JSV() js.Value         { return v.jsv }
func (v *Text) SetJSV(value js.Value) { v.jsv = value }
