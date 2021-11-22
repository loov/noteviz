package smufl

import (
	"encoding/json"

	"gioui.org/font/opentype"
)

type Font struct {
	Face *opentype.Font

	Name    string
	Version string

	EngravingDefaults  EngravingDefaults
	GlyphAdvanceWidths map[rune]Decimal
	GlyphBBoxes        map[rune]GlyphBBox
	GlyphAnchors       map[rune]GlyphAnchors
}

func ParseFont(fontData, fontMetadata []byte) (*Font, error) {
	font := &Font{}

	face, err := opentype.Parse(fontData)
	if err != nil {
		return nil, err
	}
	font.Face = face

	var metadata metadata
	if err = json.Unmarshal(fontMetadata, &metadata); err != nil {
		return nil, err
	}

	font.Name = metadata.FontName
	font.Version = metadata.FontVersion.String()
	font.EngravingDefaults = metadata.EngravingDefaults

	font.GlyphAdvanceWidths = map[rune]Decimal{}
	for name, advance := range metadata.GlyphAdvanceWidths {
		r, ok := NameToCodepoint(name)
		if !ok {
			// TODO: add fallbacks for non-standard names
			// return nil, fmt.Errorf("unknown name %q", name)
			continue
		}
		font.GlyphAdvanceWidths[r] = advance
	}

	font.GlyphBBoxes = map[rune]GlyphBBox{}
	for name, bbox := range metadata.GlyphBBoxes {
		r, ok := NameToCodepoint(name)
		if !ok {
			continue
		}
		font.GlyphBBoxes[r] = bbox
	}

	font.GlyphAnchors = map[rune]GlyphAnchors{}
	for name, anchors := range metadata.GlyphsWithAnchors {
		r, ok := NameToCodepoint(name)
		if !ok {
			continue
		}
		font.GlyphAnchors[r] = anchors
	}

	return font, err
}

func MustParseFont(fontData, fontMetadata []byte) *Font {
	font, err := ParseFont(fontData, fontMetadata)
	if err != nil {
		panic(err)
	}
	return font
}

type GlyphBBox struct {
	NE DecimalPoint `json:"bBoxNE"`
	SW DecimalPoint `json:"bBoxSW"`
}

type EngravingDefaults struct {
	// This list may also use the generic font family values defined in CSS,
	// i.e. serif, sans-serif, cursive, fantasy, and monospace. Generic font
	// family names should be listed after specific font families. An array
	// containing the text font family (or families, in descending order of
	// preference) that are ideally paired with this music
	TextFontFamily []string `json:"textFontFamily"`
	// The thickness of each staff line
	StaffLineThickness Decimal `json:"staffLineThickness"`
	// The thickness of a stem
	StemThickness Decimal `json:"stemThickness"`
	// The thickness of a beam
	BeamThickness Decimal `json:"beamThickness"`
	// The distance between the inner edge of the primary and outer edge of
	// subsequent secondary beams
	BeamSpacing Decimal `json:"beamSpacing"`
	// The thickness of a leger line (normally somewhat thicker than a staff
	// line)
	LegerLineThickness Decimal `json:"legerLineThickness"`
	// The amount by which a leger line should extend either side of a
	// notehead, scaled proportionally with the notehead's size, e.g.  when
	// scaled down as a grace note
	LegerLineExtension Decimal `json:"legerLineExtension"`
	// The thickness of the end of a slur
	SlurEndpointThickness Decimal `json:"slurEndpointThickness"`
	// The thickness of the mid-point of a slur (i.e. its thickest point)
	SlurMidpointThickness Decimal `json:"slurMidpointThickness"`
	// The thickness of the end of a tie
	TieEndpointThickness Decimal `json:"tieEndpointThickness"`
	// The thickness of the mid-point of a tie
	TieMidpointThickness Decimal `json:"tieMidpointThickness"`
	// The thickness of a thin barline, e.g. a normal barline, or each of the
	// lines of a double barline
	ThinBarlineThickness Decimal `json:"thinBarlineThickness"`
	// The thickness of a thick barline, e.g. in a final barline or a repeat
	// barline
	ThickBarlineThickness Decimal `json:"thickBarlineThickness"`
	// The thickness of a dashed barline
	DashedBarlineThickness Decimal `json:"dashedBarlineThickness"`
	// The length of the dashes to be used in a dashed barline
	DashedBarlineDashLength Decimal `json:"dashedBarlineDashLength"`
	// The length of the gap between dashes in a dashed barline
	DashedBarlineGapLength Decimal `json:"dashedBarlineGapLength"`
	// The default distance between multiple thin barlines when locked
	// together, e.g. between two thin barlines making a double barline,
	// measured from the right-hand edge of the left barline to the left-hand
	// edge of the right barline.
	BarlineSeparation Decimal `json:"barlineSeparation"`
	// The default distance between a pair of thin and thick barlines when
	// locked together, e.g. between the thin and thick barlines making a
	// final barline, or between the thick and thin barlines making a start
	// repeat barline.
	ThinThickBarlineSeparation Decimal `json:"thinThickBarlineSeparation"`
	// The default horizontal distance between the dots and the inner barline
	// of a repeat barline, measured from the edge of  the dots to the edge
	// of the barline.
	RepeatBarlineDotSeparation Decimal `json:"repeatBarlineDotSeparation"`
	// The thickness of the vertical line of a bracket grouping staves together
	BracketThickness Decimal `json:"bracketThickness"`
	// The thickness of the vertical line of a sub-bracket grouping staves
	// belonging to the same instrument together
	SubBracketThickness Decimal `json:"subBracketThickness"`
	// The thickness of a crescendo/diminuendo hairpin
	HairpinThickness Decimal `json:"hairpinThickness"`
	// The thickness of the dashed line used for an octave line
	OctaveLineThickness Decimal `json:"octaveLineThickness"`
	// The thickness of the line used for piano pedaling
	PedalLineThickness Decimal `json:"pedalLineThickness"`
	// The thickness of the brackets drawn to indicate repeat endings
	RepeatEndingLineThickness Decimal `json:"repeatEndingLineThickness"`
	// The thickness of the line used for the shaft of an arrow
	ArrowShaftThickness Decimal `json:"arrowShaftThickness"`
	// The thickness of the lyric extension line to indicate a melisma in
	// vocal music
	LyricLineThickness Decimal `json:"lyricLineThickness"`
	// The thickness of a box drawn around text instructions (e.g. rehearsal
	// marks)
	TextEnclosureThickness Decimal `json:"textEnclosureThickness"`
	// The thickness of the brackets drawn either side of tuplet numbers
	TupletBracketThickness Decimal `json:"tupletBracketThickness"`
	// The thickness of the horizontal line drawn between two vertical lines,
	// known as the H-bar, in a multi-bar rest
	HBarThickness Decimal `json:"hBarThickness"`
}

type GlyphAnchors struct {
	// The exact position at which the bottom right-hand (south-east) corner
	// of an angled upward-pointing stem connecting the right-hand side of a
	// notehead to a vertical stem to its left should start, relative to the
	// glyph origin, expressed as Cartesian coordinates in staff spaces.
	SplitStemUpSE DecimalPoint `json:"splitStemUpSE"`
	// The exact position at which the bottom left-hand (south-west) corner of
	// an angled upward-pointing stem connecting the left-hand side of a
	// notehead to a vertical stem to its right should start, relative to the
	// glyph origin, expressed as Cartesian coordinates in staff spaces.
	SplitStemUpSW DecimalPoint `json:"splitStemUpSW"`
	// The exact position at which the top right-hand (north-east) corner of
	// an angled downward-pointing stem connecting the right-hand side of a
	// notehead to a vertical stem to its left should start, relative to the
	// glyph origin, expressed as Cartesian coordinates in staff spaces.
	SplitStemDownNE DecimalPoint `json:"splitStemDownNE"`
	// The exact position at which the top left-hand (north-west) corner of an
	// angled downward-pointing stem connecting the left-hand side of a
	// notehead to a vertical stem to its right should start, relative to the
	// glyph origin, expressed as Cartesian coordinates in staff spaces.
	SplitStemDownNW DecimalPoint `json:"splitStemDownNW"`
	// The exact position at which the bottom right-hand (south-east) corner
	// of an upward-pointing stem rectangle should start, relative to the
	// glyph origin, expressed as Cartesian coordinates in staff spaces.
	StemUpSE DecimalPoint `json:"stemUpSE"`
	// The exact position at which the top left-hand (north-west) corner of a
	// downward-pointing stem rectangle should start, relative to the glyph
	// origin, expressed as Cartesian coordinates in staff spaces.
	StemDownNW DecimalPoint `json:"stemDownNW"`
	// The amount by which an up-stem should be lengthened from its nominal
	// unmodified length in order to ensure a good connection with a flag, in
	// spaces.11
	StemUpNW DecimalPoint `json:"stemUpNW"`
	// The amount by which a down-stem should be lengthened from its nominal
	// unmodified length in order to ensure a good connection with a flag, in
	// spaces.
	StemDownSW DecimalPoint `json:"stemDownSW"`
	// The width in staff spaces of a given glyph that should be used for e.g.
	// positioning leger lines correctly.12
	NominalWidth DecimalPoint `json:"nominalWidth"`
	// The position in staff spaces that should be used to position numerals
	// relative to clefs with ligated numbers where those numbers hang from
	// the bottom of the clef, corresponding horizontally to the center of
	// the numeral’s bounding box.
	NumeralTop DecimalPoint `json:"numeralTop"`
	// The position in staff spaces that should be used to position numerals
	// relative to clefs with ligatured numbers where those numbers sit on
	// the baseline or at the north-east corner of the G clef, corresponding
	// horizontally to the center of the numeral’s bounding box.
	NumeralBottom DecimalPoint `json:"numeralBottom"`
	// The Cartesian coordinates in staff spaces of the bottom left corner of
	// a nominal rectangle that intersects the top right corner of the
	// glyph’s bounding box. This rectangle, together with those in the other
	// four corners of the glyph’s bounding box, can be cut out to produce a
	// more detailed bounding box (of abutting rectangles), useful for
	// kerning or interlocking symbols such as accidentals.
	CutOutNE DecimalPoint `json:"cutOutNE"`
	// The Cartesian coordinates in staff spaces of the top left corner of a
	// nominal rectangle that intersects the bottom right corner of the
	// glyph’s bounding box.
	CutOutSE DecimalPoint `json:"cutOutSE"`
	// The Cartesian coordinates in staff spaces of the top right corner of a
	// nominal rectangle that intersects the bottom left corner of the
	// glyph’s bounding box.
	CutOutSW DecimalPoint `json:"cutOutSW"`
	// The Cartesian coordinates in staff spaces of the bottom right corner of
	// a nominal rectangle that intersects the top left corner of the glyph’s
	// bounding box.
	CutOutNW DecimalPoint `json:"cutOutNW"`
	// The Cartesian coordinates in staff spaces of the position at which the
	// glyph graceNoteSlashStemUp should be positioned relative to the
	// stem-up flag of an unbeamed grace note; alternatively, the bottom left
	// corner of a diagonal line drawn instead of using the above glyph.
	GraceNoteSlashSW DecimalPoint `json:"graceNoteSlashSW"`
	// The Cartesian coordinates in staff spaces of the top right corner of a
	// diagonal line drawn instead of using the glyph graceNoteSlashStemUp
	// for a stem-up flag of an unbeamed grace note.
	GraceNoteSlashNE DecimalPoint `json:"graceNoteSlashNE"`
	// The Cartesian coordinates in staff spaces of the position at which the
	// glyph graceNoteSlashStemDown should be positioned relative to the
	// stem-down flag of an unbeamed grace note; alternatively, the top left
	// corner of a diagonal line drawn instead of using the above glyph.
	GraceNoteSlashNW DecimalPoint `json:"graceNoteSlashNW"`
	// The Cartesian coordinates in staff spaces of the bottom right corner of
	// a diagonal line drawn instead of using the glyph
	// graceNoteSlashStemDown for a stem-down flag of an unbeamed grace note.
	GraceNoteSlashSE DecimalPoint `json:"graceNoteSlashSE"`
	// The Cartesian coordinates in staff spaces of the horizontal position at
	// which a glyph repeats, i.e. the position at which the same glyph or
	// another of the same group should be positioned to ensure correct
	// tessellation. This is used for e.g. multi-segment lines and the
	// component glyphs that make up trills and mordents.
	RepeatOffset DecimalPoint `json:"repeatOffset"`
	// The Cartesian coordinates in staff spaces of the left-hand edge of a
	// notehead with a non-zero left-hand side bearing (e.g. a double whole,
	// or breve, notehead with two vertical lines at each side), to assist in
	// the correct horizontal alignment of these noteheads with other
	// noteheads with zero-width left-side bearings.
	NoteheadOrigin DecimalPoint `json:"noteheadOrigin"`
	// The Cartesian coordinates in staff spaces of the optical center of the
	// glyph, to assist in the correct horizontal alignment of the glyph
	// relative to a notehead or stem. Currently recommended for use with
	// glyphs in the Dynamics range.
	OpticalCenter DecimalPoint `json:"opticalCenter"`
}
