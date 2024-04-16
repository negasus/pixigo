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
	"github.com/negasus/pixigo/container"
	"github.com/negasus/pixigo/rectangle"
	"github.com/negasus/pixigo/sprite"
	"github.com/negasus/pixigo/ticker"
)

type maggot struct {
	s            *sprite.Sprite
	direction    float64
	turningSpeed float64
	speed        float64
	offset       float64
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

	txt, errLoad := assets.LoadTextures("maggot_tiny.png")
	if errLoad != nil {
		panic(errLoad)
	}

	sprites := container.New()

	app.Stage().AddChild(sprites)

	var maggots []*maggot

	totalSprites := 10000

	for i := 0; i < totalSprites; i++ {
		d := &maggot{}

		d.s = sprite.FromTexture(txt[0])

		d.s.Anchor().Set(0.5, 0.5)

		d.s.Scale().Set(0.8+rand.Float64()*0.3, 0.8+rand.Float64()*0.3)

		d.s.SetX(rand.Float64() * app.Screen().Width())
		d.s.SetY(rand.Float64() * app.Screen().Height())

		d.s.SetTint(rand.Float64() * float64(0x808080))

		d.direction = rand.Float64() * math.Pi * 2
		d.turningSpeed = rand.Float64() - 0.8

		d.speed = (2 + rand.Float64()*2) * 0.2
		d.offset = rand.Float64() * 100

		maggots = append(maggots, d)

		app.Stage().AddChild(d.s)
	}

	dudeBoundsPadding := 100.0

	dudeBounds := rectangle.New(
		-dudeBoundsPadding,
		-dudeBoundsPadding,
		app.Screen().Width()+dudeBoundsPadding*2,
		app.Screen().Height()+dudeBoundsPadding*2,
	)

	var tick float64

	app.Ticker().Add(func(t *ticker.Ticker) {
		for _, d := range maggots {
			//dude.scale.y = 0.95 + Math.sin(tick + dude.offset) * 0.05;
			d.s.Scale().SetY(0.95 + math.Sin(tick+d.offset)*0.05)

			//dude.direction += dude.turningSpeed * 0.01;
			d.direction += d.turningSpeed * 0.01
			//dude.x += Math.sin(dude.direction) * (dude.speed * dude.scale.x);
			d.s.SetX(d.s.X() + math.Sin(d.direction)*(d.speed*d.s.Scale().X()))
			//dude.y += Math.cos(dude.direction) * (dude.speed * dude.scale.y);
			d.s.SetY(d.s.Y() + math.Cos(d.direction)*(d.speed*d.s.Scale().Y()))
			//dude.rotation = -dude.direction + Math.PI;
			d.s.SetRotation(-d.direction + math.Pi)

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

		tick += 0.1
	})

	<-ctx.Done()
}
