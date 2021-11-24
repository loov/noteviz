package main

import (
	_ "embed"
	"fmt"
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
	"image/color"
	"os"

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
			fmt.Fprintln(os.Stderr, err)
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

	repeated := func(count int, r rune, advance fixed.Int26_6) text.Layout {
		var lay text.Layout
		for i := 0; i < count; i++ {
			lay.Text += string(r)
			lay.Advances = append(lay.Advances, advance)
		}
		return lay
	}

	size := fixed.I(16)
	for k := 0; k < 3; k++ {
		op.Offset(f32.Pt(0, 256)).Add(gtx.Ops)

		size *= 2
		face := ui.Font.Face
		{
			advance := ui.Font.GlyphAdvanceWidths[smufl.Staff5Lines].Px(size)
			clip := face.Shape(size, repeated(5, smufl.Staff5Lines, advance))
			paint.FillShape(gtx.Ops, color.NRGBA{A: 0xFF}, clip)
		}
		{
			advance := ui.Font.GlyphAdvanceWidths[smufl.NoteWhole].Px(size)
			clip := face.Shape(size, repeated(5, smufl.NoteWhole, advance))
			paint.FillShape(gtx.Ops, color.NRGBA{A: 0xFF}, clip)
		}
		{
			op.Offset(f32.Pt(0, float32((-size/4).Round()))).Add(gtx.Ops)

			advance := ui.Font.GlyphAdvanceWidths[smufl.GClef].Px(size)
			clip := face.Shape(size, repeated(5, smufl.GClef, advance))
			paint.FillShape(gtx.Ops, color.NRGBA{A: 0xFF}, clip)
		}
	}

	return layout.Dimensions{
		Size: gtx.Constraints.Max,
	}
}
