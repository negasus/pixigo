package main

import (
	"context"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"syscall/js"

	"github.com/negasus/pixigo/application"
	"github.com/negasus/pixigo/assets"
	"github.com/negasus/pixigo/rectangle"
	"github.com/negasus/pixigo/sprite"
	"github.com/negasus/pixigo/ticker"
)

type dude struct {
	s            *sprite.Sprite
	direction    float64
	turningSpeed float64
	speed        float64
}

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

	txt, errLoad := assets.LoadTextures("eggHead.png")
	if errLoad != nil {
		panic(errLoad)
	}

	var aliens []*dude

	totalDudes := 20

	for i := 0; i < totalDudes; i++ {
		d := &dude{}

		d.s = sprite.FromTexture(txt[0])

		d.s.Anchor().Set(0.5, 0.5)

		d.s.Scale().Set(0.8+rand.Float64()*0.3, 0.8+rand.Float64()*0.3)

		d.s.SetX(rand.Float64() * app.Screen().Width())
		d.s.SetY(rand.Float64() * app.Screen().Height())

		d.s.SetTint(rand.Float64() * float64(0xFFFFFF))

		d.direction = rand.Float64() * math.Pi * 2
		d.turningSpeed = rand.Float64() - 0.8

		d.speed = 2 + rand.Float64()*2

		aliens = append(aliens, d)

		app.Stage().AddChild(d.s)
	}

	dudeBoundsPadding := 100.0

	dudeBounds := rectangle.New(
		-dudeBoundsPadding,
		-dudeBoundsPadding,
		app.Screen().Width()+dudeBoundsPadding*2,
		app.Screen().Height()+dudeBoundsPadding*2,
	)

	app.Ticker().Add(func(t *ticker.Ticker) {
		for _, d := range aliens {
			d.direction += d.turningSpeed * 0.01
			d.s.SetX(d.s.X() + math.Sin(d.direction)*d.speed)
			d.s.SetY(d.s.Y() + math.Cos(d.direction)*d.speed)
			d.s.SetRotation(-d.direction - math.Pi/2)

			if d.s.X() < dudeBounds.X() {
				d.s.SetX(d.s.X() + dudeBounds.Width())
			} else if d.s.X() > dudeBounds.X()+dudeBounds.Width() {
				d.s.SetX(d.s.X() - dudeBounds.Width())
			}

			if d.s.Y() < dudeBounds.Y() {
				d.s.SetY(d.s.Y() + dudeBounds.Height())
			} else if d.s.Y() > dudeBounds.Y()+dudeBounds.Height() {
				d.s.SetY(d.s.Y() - dudeBounds.Height())
			}
		}
	})

	<-ctx.Done()
}
