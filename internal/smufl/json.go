package smufl

import (
	"encoding/json"
)

type metadata struct {
	FontName             string                     `json:"fontName"`
	FontVersion          json.Number                `json:"fontVersion"`
	EngravingDefaults    EngravingDefaults          `json:"engravingDefaults"`
	GlyphAdvanceWidths   map[string]Decimal         `json:"glyphAdvanceWidths"`
	GlyphBBoxes          map[string]GlyphBBox       `json:"glyphBBoxes"`
	GlyphsWithAlternates map[string]glyphAlternates `json:"glyphsWithAlternates"`
	GlyphsWithAnchors    map[string]GlyphAnchors    `json:"glyphsWithAnchors"`
	Ligatures            map[string]glyphLigatures  `json:"ligatures"`
	Sets                 map[string]glyphSet        `json:"sets"`
	OptionalGlyphs       map[string]optionalGlyph   `json:"optionalGlyphs"`
}

type glyphAlternates struct {
	Alternates []glyph `json:"alternates"`
}

type glyphLigatures struct {
	Codepoint       string   `json:"codepoint"`
	ComponentGlyphs []string `json:"componentGlyphs"`
}

type glyph struct {
	Codepoint   string `json:"codepoint"`
	Description string `json:"description"`
}

type glyphSet struct {
	// "opticalVariantsSmall" - Glyphs designed for use on smaller staff sizes.
	// "flagsShort" - Alternate shorter flags for notes with augmentation dots.
	// "flagsStraight" - Alternate flags that are straight rather than curved.
	// "timeSigsLarge" - Alternate time signature digits for use outside the staff.
	// "noteheadsLarge" - Alternate oversized noteheads.
	Type        string          `json:"type"`
	Description string          `json:"description"`
	Glyphs      []glyphSetEntry `json:"glyphs"`
}

type glyphSetEntry struct {
	Codepoint    string `json:"codepoint"`
	Description  string `json:"description"`
	AlternateFor string `json:"alternateFor"`
}

type optionalGlyph struct {
	Classes   []string `json:"classes"`
	Codepoint string   `json:"codepoint"`
}
