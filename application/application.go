package application

import (
	"syscall/js"

	"github.com/negasus/pixigo"
	"github.com/negasus/pixigo/container"
	"github.com/negasus/pixigo/rectangle"
	"github.com/negasus/pixigo/ticker"
)

type Application struct {
	jsv    js.Value
	stage  *container.Container
	screen *rectangle.Rectangle
	ticker *ticker.Ticker
}

func New() *Application {
	v := &Application{
		jsv:    pixigo.PIXI().Get("Application").New(),
		stage:  container.New(),
		screen: &rectangle.Rectangle{},
		ticker: &ticker.Ticker{},
	}

	return v
}

func (v *Application) Canvas() js.Value {
	return v.jsv.Get("canvas")
}

func (v *Application) Stage() *container.Container {
	return v.stage
}

func (v *Application) Screen() *rectangle.Rectangle {
	return v.screen
}

func (v *Application) Ticker() *ticker.Ticker {
	return v.ticker
}
