package main

import (
	"syscall/js"

	"github.com/negasus/pixigo/application"
	"github.com/negasus/pixigo/graphics"
)

func main() {
	app := application.New()

	errInit := app.Init(&application.InitOptions{
		Background: 0x1099BB,
		ResizeTo:   js.Global().Get("window"),
	})
	if errInit != nil {
		panic(errInit)
	}

	js.Global().Get("document").Get("body").Call("appendChild", app.Canvas())

	g := graphics.New()

	g.Rect(50, 50, 100, 100)
	g.Fill(0xDE3249)

	g.Rect(200, 50, 100, 100)
	g.Fill(0x650A5A)
	g.Stroke(&graphics.FillStyleInputs{Width: 2, Color: 0xFEEB77})

	g.Rect(350, 50, 100, 100)
	g.Fill(0xC34288)
	g.Stroke(&graphics.FillStyleInputs{Width: 10, Color: 0xFFBD01})

	g.Rect(530, 50, 140, 100)
	g.Fill(0xAA4F08)
	g.Stroke(&graphics.FillStyleInputs{Width: 2, Color: 0xFFFFFF})

	g.Circle(100, 250, 50)
	g.Fill(0xDE3249) // todo ,1

	g.Circle(250, 250, 50)
	g.Fill(0x640A5A) // todo ,1
	g.Stroke(&graphics.FillStyleInputs{Width: 2, Color: 0xFEEB77})

	app.Stage().AddChild(g)
}
