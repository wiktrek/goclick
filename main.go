package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/go-vgo/robotgo"
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(app.Title("Auto clicker"))
		window.Option(app.Size(unit.Dp(400), unit.Dp(400)))
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	
	app.Main()
}

func check_mouse_location() {
  fmt.Println(robotgo.Location())
}
func click_location(button string, x int, y int) {
	robotgo.Move(x, y)
}
func click(button string) {
	robotgo.Click(button, true)
}
func hold() {
  robotgo.Toggle("left") // toggle
  robotgo.Toggle("left", "up") // free
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	var startButton widget.Clickable
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			title := material.H1(theme, "Auto clicker")

			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon

			title.Alignment = text.Middle
			
			title.Layout(gtx)
			
			button := material.Button(theme, &startButton, "Start")
			button.Layout(gtx)
			
			e.Frame(gtx.Ops)
		}
	}
}