package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/layout"
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
	var changeHotkeyButton widget.Clickable
	hotkey := "F6";
	// var spawned_thread bool

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
			
			hotkey_text := material.H1(theme, hotkey)



			hotkey_text.Color = maroon

			hotkey_text.Alignment = text.Middle
			

			title.Layout(gtx)
			start := material.Button(theme, &startButton, "Start")
			start.Layout(gtx)
			change_hotkey := material.Button(theme, &changeHotkeyButton, "Change hotkey")
			change_hotkey.Layout(gtx)

			e.Frame(gtx.Ops)
		}
	}
}
func open_key_press() {
	// open a small window for someone to bind the key
}

func get_key(gtx layout.Context) key.Name {
	for {
		ev, ok := gtx.Event(key.Filter{})
		if !ok {
				break
		}
		if x, ok := ev.(key.Event); ok {
			switch x.State {
				case key.Press:
					fmt.Printf("KEY   : %+v\n", x.Name)
					return x.Name
			}
		}
	}
	return key.NameF6
}