package application

import (
	"fmt"
	"syscall/js"

	"github.com/negasus/pixigo"
)

type Application struct {
	v js.Value
}

func New() *Application {
	app := &Application{}

	app.v = js.Global().Get("PIXI").Get("Application").New()

	return app
}

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

func (app *Application) Init(options *InitOptions) error {
	_, catch := pixigo.Await(app.v, "init", options.marshal())
	if catch != nil {
		return fmt.Errorf("failed to init: %v", catch)
	}

	js.Global().Get("document").Get("body").Call("appendChild", app.v.Get("canvas"))

	return nil

}
