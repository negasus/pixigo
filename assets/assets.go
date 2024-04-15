package assets

import (
	"fmt"

	"github.com/negasus/pixigo"
	"github.com/negasus/pixigo/texture"
)

func LoadTextures(args ...any) ([]*texture.Texture, error) {
	var result []*texture.Texture

	a := pixigo.PIXI().Get("Assets")

	res, err := pixigo.Await(a, "load", args...)
	if err != nil {
		return nil, fmt.Errorf("failed to load assets: %v", err)
	}

	for _, v := range res {
		t := &texture.Texture{}
		t.SetJSV(v)
		result = append(result, t)
	}

	return result, nil
}
