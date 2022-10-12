package main

import (
	_ "embed"
	"fmt"
	"image"
	"image/color"
	"os"
	"strings"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"golang.org/x/image/math/fixed"

	"github.com/loov/noteviz/internal/font/bravura"
	"github.com/loov/noteviz/internal/smufl"
)

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

/*
var input = `
# b3bb bb_b_ b_ b_ b_
2.34432.2.2.2.

# x x e4eggb dc+baf dg
`
*/

func (ui *UI) Layout(gtx layout.Context) layout.Dimensions {
	staves := []Staff{
		{
			One:          fixed.I(16),
			SpaceAboveSP: fixed.I(4),
			SpaceBelowSP: fixed.I(4),

			Notes: []int{-3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
		},
		{
			One:          fixed.I(48),
			SpaceAboveSP: fixed.I(4),
			SpaceBelowSP: fixed.I(4),

			Notes: []int{-3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
		},
	}
	return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		list := layout.List{
			Axis: layout.Vertical,
		}
		return list.Layout(gtx, len(staves), func(gtx layout.Context, i int) layout.Dimensions {
			return staves[i].Layout(gtx, ui.Font)
		})
	})
}

type Staff struct {
	One fixed.Int26_6 // staff space

	SpaceAboveSP fixed.Int26_6
	SpaceBelowSP fixed.Int26_6

	Notes []int
}

func (staff *Staff) Layout(gtx layout.Context, font *smufl.Font) layout.Dimensions {
	width := float32(gtx.Constraints.Max.X)

	lineHeight := staff.One * 4
	firstStaffLine := staff.SpaceAboveSP.Mul(staff.One) + 4*staff.One
	defer op.Offset(image.Pt(0, firstStaffLine.Round())).Push(gtx.Ops).Pop()

	{ // draw staff lines
		staffLineThickness := font.EngravingDefaults.StaffLineThickness.Px(lineHeight)
		halfStaffLineThicknessPx := float32((staffLineThickness / 2).Round())

		var builder clip.Path
		builder.Begin(gtx.Ops)
		for i := 0; i < 5; i++ {
			y := float32(staff.One.Mul(fixed.I(i)).Round())
			builder.MoveTo(f32.Pt(0, -y-halfStaffLineThicknessPx-0.5))
			builder.LineTo(f32.Pt(width, -y-halfStaffLineThicknessPx-0.5))
			builder.LineTo(f32.Pt(width, -y+halfStaffLineThicknessPx+0.5))
			builder.LineTo(f32.Pt(0, -y+halfStaffLineThicknessPx+0.5))
			builder.Close()
		}
		paint.FillShape(gtx.Ops, color.NRGBA{A: 0xFF}, clip.Outline{Path: builder.End()}.Op())
	}

	{
		advance := font.GlyphAdvanceWidths[smufl.NoteheadBlack].Px(lineHeight) + staff.One*12/10

		noteheadText, err := font.Face.Layout(lineHeight, 100000, gtx.Locale, strings.NewReader(string(smufl.NoteheadBlack)))
		if err != nil {
			panic(err)
		}
		noteheadClip := font.Face.Shape(lineHeight, noteheadText[0].Layout)

		legerLineText, err := font.Face.Layout(lineHeight, 100000, gtx.Locale, strings.NewReader(string(smufl.LegerLine)))
		if err != nil {
			panic(err)
		}
		legerLineClip := font.Face.Shape(lineHeight, legerLineText[0].Layout)

		x := staff.One
		for _, note := range staff.Notes {
			y := staff.One.Mul(-fixed.I(note) / 2)
			stack := op.Offset(image.Pt(x.Round(), y.Round())).Push(gtx.Ops)
			paint.FillShape(gtx.Ops, color.NRGBA{A: 0xFF}, clip.Outline{Path: noteheadClip}.Op())
			paint.FillShape(gtx.Ops, color.NRGBA{A: 0xFF}, clip.Outline{Path: legerLineClip}.Op())
			x += advance
			stack.Pop()
		}
	}

	return layout.Dimensions{
		Size: image.Point{
			X: gtx.Constraints.Max.X,
			Y: (firstStaffLine + staff.SpaceBelowSP.Mul(staff.One)).Round(),
		},
	}
}
