package main

import (
	_ "embed"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"golang.org/x/image/math/fixed"

	"github.com/loov/noteviz/internal/font/bravura"
	"github.com/loov/noteviz/internal/smufl"
)

//go:embed internal/font/bravura/otf/Bravura.otf
var bravuraFontData []byte

func main() {
	ui := NewUI()

	go func() {
		w := app.NewWindow(
			app.Title("Music Font"),
			app.Size(unit.Dp(720), unit.Dp(720)),
		)
		if err := ui.Run(w); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	app.Main()
}

var defaultMargin = unit.Dp(10)

type UI struct {
	Theme *material.Theme

	Font *smufl.Font
}

func NewUI() *UI {
	ui := &UI{}
	ui.Theme = material.NewTheme(gofont.Collection())
	ui.Font = bravura.Font()

	return ui
}

func (ui *UI) Run(w *app.Window) error {
	var ops op.Ops
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				ui.Layout(gtx)
				e.Frame(gtx.Ops)

			case key.Event:
				switch e.Name {
				case key.NameEscape:
					return nil
				}

			case system.DestroyEvent:
				return e.Err
			}
		}
	}
}

func (ui *UI) Layout(gtx layout.Context) layout.Dimensions {
	op.Offset(f32.Pt(256, 256)).Add(gtx.Ops)

	face := ui.Font.Face
	size := fixed.I(64)
	{
		clip := face.Shape(size, text.Layout{
			Text:     string(smufl.Staff5Lines),
			Advances: []fixed.Int26_6{fixed.I(0)},
		})
		paint.FillShape(gtx.Ops, color.NRGBA{A: 0xFF}, clip)
	}
	{
		clip := face.Shape(size, text.Layout{
			Text:     string(smufl.NoteWhole),
			Advances: []fixed.Int26_6{fixed.I(0)},
		})
		paint.FillShape(gtx.Ops, color.NRGBA{A: 0xFF}, clip)
	}
	{
		op.Offset(f32.Pt(0, -64*1/4)).Add(gtx.Ops)
		clip := face.Shape(size, text.Layout{
			Text:     string(smufl.GClef),
			Advances: []fixed.Int26_6{fixed.I(0)},
		})
		paint.FillShape(gtx.Ops, color.NRGBA{A: 0xFF}, clip)
	}
	return layout.Dimensions{
		Size: gtx.Constraints.Max,
	}
}
