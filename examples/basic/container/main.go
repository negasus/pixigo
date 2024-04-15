package main

import (
	"context"
	"math"
	"os"
	"os/signal"
	"syscall/js"

	"github.com/negasus/pixigo/application"
	"github.com/negasus/pixigo/assets"
	"github.com/negasus/pixigo/container"
	"github.com/negasus/pixigo/sprite"
	"github.com/negasus/pixigo/ticker"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := application.New()

	errInit := app.Init(&application.InitOptions{
		Background: 0x1099BB,
		ResizeTo:   js.Global().Get("window"),
	})
	if errInit != nil {
		panic(errInit)
	}

	js.Global().Get("document").Get("body").Call("appendChild", app.Canvas())

	c := container.New()
	app.Stage().AddChild(c)

	txt, errLoad := assets.LoadTextures("bunny.png")
	if errLoad != nil {
		panic(errLoad)
	}

	for i := 0; i < 25; i++ {
		bunny := sprite.FromTexture(txt[0])

		bunny.SetX(float64(i%5) * 40)
		bunny.SetY(math.Floor(float64(i)/5) * 40)
		c.AddChild(bunny)
	}

	c.SetX(app.Screen().Width() / 2)
	c.SetY(app.Screen().Height() / 2)

	c.Pivot().SetX(c.Width() / 2)
	c.Pivot().SetY(c.Height() / 2)

	app.Ticker().Add(func(t *ticker.Ticker) {
		c.SetRotation(c.Rotation() - 0.01*t.DeltaTime())
	})

	<-ctx.Done()
}
