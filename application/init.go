package application

import (
	"fmt"

	"github.com/negasus/pixigo"
)

type InitOptions struct {
	Background int
	ResizeTo   any
}

func (opt *InitOptions) marshal() map[string]any {
	args := map[string]any{}

	if opt.Background != 0 {
		args["background"] = opt.Background
	}

	if opt.ResizeTo != nil {
		args["resizeTo"] = opt.ResizeTo
	}

	return args
}

func (v *Application) Init(options *InitOptions) error {
	_, catch := pixigo.Await(v.jsv, "init", options.marshal())
	if catch != nil {
		return fmt.Errorf("failed to init: %v", catch)
	}

	v.stage.SetJSV(v.jsv.Get("stage"))
	v.screen.SetJSV(v.jsv.Get("screen"))
	v.ticker.SetJSV(v.jsv.Get("ticker"))

	return nil
}
