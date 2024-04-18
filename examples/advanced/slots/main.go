package main

import (
	"context"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"syscall/js"
	"time"

	"github.com/negasus/pixigo/application"
	"github.com/negasus/pixigo/assets"
	"github.com/negasus/pixigo/container"
	"github.com/negasus/pixigo/graphics"
	"github.com/negasus/pixigo/sprite"
	"github.com/negasus/pixigo/text"
	"github.com/negasus/pixigo/ticker"
)

type reel struct {
	c                *container.Container
	symbols          []*sprite.Sprite
	position         int
	previousPosition int
	//blur
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

	txt, _ := assets.LoadTextures("eggHead.png", "flowerTop.png", "helmlok.png", "skully.png")

	reelWidth := 160.0
	symbolSize := 150.0

	var reels []*reel

	reelContainer := container.New()

	for i := 0; i < 5; i++ {
		rc := container.New()

		rc.SetX(float64(i) * reelWidth)
		reelContainer.AddChild(rc)

		r := &reel{
			c:                rc,
			symbols:          nil,
			position:         0,
			previousPosition: 0,
		}

		for j := 0; j < 4; j++ {
			symbol := sprite.FromTexture(txt[j])
			symbol.SetY(float64(j) * symbolSize)

			ss := min(symbolSize/symbol.Width(), symbolSize/symbol.Height())
			symbol.Scale().SetX(ss)
			symbol.Scale().SetY(ss)

			symbol.SetX(math.Round((symbolSize - symbol.Width()) / 2))

			r.symbols = append(r.symbols, symbol)
			rc.AddChild(symbol)
		}

		reels = append(reels, r)
	}

	app.Stage().AddChild(reelContainer)

	margin := (app.Screen().Height() - symbolSize*3) / 2

	reelContainer.SetY(margin)
	reelContainer.SetX(math.Round(app.Screen().Width() - reelWidth*5))

	top := graphics.New()
	top.Rect(0, 0, app.Screen().Width(), margin)
	top.Fill(0x000000)

	bottom := graphics.New()
	bottom.Rect(0, symbolSize*3+margin, app.Screen().Width(), margin)
	bottom.Fill(0x000000)

	// fill gradient

	style := &text.Style{
		FontFamily:    "Arial",
		FontSize:      36,
		FontStyle:     "italic",
		FontWeight:    "bold",
		WordWrap:      true,
		WordWrapWidth: 440,
	}

	playText := text.New("Spin the wheels!", style)

	playText.SetX(math.Round((bottom.Width() - playText.Width()) / 2))
	playText.SetY(app.Screen().Height() - margin + math.Round((margin-playText.Height())/2))
	bottom.AddChild(playText)

	headerText := text.New("PIXI MONSTER SLOTS!", style)
	headerText.SetX(math.Round((top.Width() - headerText.Width()) / 2))
	headerText.SetY(math.Round((margin - headerText.Height()) / 2))
	top.AddChild(headerText)

	app.Stage().AddChild(top, bottom)

	running := false

	startPlay := func() {
		if running {
			return
		}

		running = true

		for i := 0; i < len(reels); i++ {
			r := reels[i]
			extra := int(math.Floor(rand.Float64() * 3))
			target := r.position + 10 + i*5 + extra
			tim := 2500 + i*600 + extra*600

			tweenTo(r, "position", target, tim, backout(0.5), nil, nil)
		}
	}

	reelsComplete := func() {
		running = false
	}

	type tween struct {
		start time.Time
		time  time.Time
	}

	tweenTo := func() {

	}

	bottom.SetEventMode("static")
	bottom.SetCursor("pointer")
	bottom.On("pointerdown", func() {
		startPlay()
	})

	app.Ticker().Add(func(t *ticker.Ticker) {
		for i := 0; i < len(reels); i++ {
			r := reels[i]

			// r.blur.blurY = (r.position - r.previousPosition) * 8;
			// todo

			r.previousPosition = r.position

			for j := 0; j < len(r.symbols); j++ {
				s := r.symbols[j]

				var prevy = s.Y()

				s.SetY(math.Round(float64((r.position+j)%len(r.symbols))*symbolSize - symbolSize))

				if s.Y() < 0 && prevy > symbolSize {
					// s.texture = slotTextures[Math.floor(Math.random() * slotTextures.length)];
					// todo

					sc := min(symbolSize/s.Width(), symbolSize/s.Height())
					s.Scale().Set(sc, sc)
					s.SetX(math.Round((symbolSize - s.Width()) / 2))
				}
			}
		}
	})

	var tweening []*tween

	app.Ticker().Add(func(t *ticker.Ticker) {
		now := time.Now()
		remove := []any{}

		for i := 0; i < len(tweening); i++ {
			tw := tweening[i]

			phase := min(1, int(now.Sub(tw.start).Nanoseconds())/tw.time.Nanosecond())
		}
	})

	<-ctx.Done()
}
