package smufl

import (
	"encoding/json"

	"golang.org/x/image/math/fixed"
)

// Metadata defines json struct for metadata definition.
type Metadata struct {
	FontName    string
	FontVersion string

	EngravingDefaults  EngravingDefaults
	GlyphAdvanceWidths map[rune]fixed.Int26_6
	GlyphBBoxes        map[rune]GlyphBBox
	GlyphAlternates    map[rune][]rune
}

type GlyphBBox struct {
	NE fixed.Point26_6
	SW fixed.Point26_6
}

type GlyphAnchors struct {
	// The exact position at which the bottom right-hand (south-east) corner of an angled upward-pointing stem connecting the right-hand side of a notehead to a vertical stem to its left should start, relative to the glyph origin, expressed as Cartesian coordinates in staff spaces.
	SplitStemUpSE fixed.Point26_6
	// The exact position at which the bottom left-hand (south-west) corner of an angled upward-pointing stem connecting the left-hand side of a notehead to a vertical stem to its right should start, relative to the glyph origin, expressed as Cartesian coordinates in staff spaces.
	SplitStemUpSW fixed.Point26_6
	// The exact position at which the top right-hand (north-east) corner of an angled downward-pointing stem connecting the right-hand side of a notehead to a vertical stem to its left should start, relative to the glyph origin, expressed as Cartesian coordinates in staff spaces.
	SplitStemDownNE fixed.Point26_6
	// The exact position at which the top left-hand (north-west) corner of an angled downward-pointing stem connecting the left-hand side of a notehead to a vertical stem to its right should start, relative to the glyph origin, expressed as Cartesian coordinates in staff spaces.
	SplitStemDownNW fixed.Point26_6
	// The exact position at which the bottom right-hand (south-east) corner of an upward-pointing stem rectangle should start, relative to the glyph origin, expressed as Cartesian coordinates in staff spaces.
	StemUpSE fixed.Point26_6
	// The exact position at which the top left-hand (north-west) corner of a downward-pointing stem rectangle should start, relative to the glyph origin, expressed as Cartesian coordinates in staff spaces.
	StemDownNW fixed.Point26_6
	// The amount by which an up-stem should be lengthened from its nominal unmodified length in order to ensure a good connection with a flag, in spaces.11
	StemUpNW fixed.Point26_6
	// The amount by which a down-stem should be lengthened from its nominal unmodified length in order to ensure a good connection with a flag, in spaces.
	StemDownSW fixed.Point26_6
	// The width in staff spaces of a given glyph that should be used for e.g. positioning leger lines correctly.12
	NominalWidth fixed.Point26_6
	// The position in staff spaces that should be used to position numerals relative to clefs with ligated numbers where those numbers hang from the bottom of the clef, corresponding horizontally to the center of the numeral’s bounding box.
	NumeralTop fixed.Point26_6
	// The position in staff spaces that should be used to position numerals relative to clefs with ligatured numbers where those numbers sit on the baseline or at the north-east corner of the G clef, corresponding horizontally to the center of the numeral’s bounding box.
	NumeralBottom fixed.Point26_6
	// The Cartesian coordinates in staff spaces of the bottom left corner of a nominal rectangle that intersects the top right corner of the glyph’s bounding box. This rectangle, together with those in the other four corners of the glyph’s bounding box, can be cut out to produce a more detailed bounding box (of abutting rectangles), useful for kerning or interlocking symbols such as accidentals.
	CutOutNE fixed.Point26_6
	// The Cartesian coordinates in staff spaces of the top left corner of a nominal rectangle that intersects the bottom right corner of the glyph’s bounding box.
	CutOutSE fixed.Point26_6
	// The Cartesian coordinates in staff spaces of the top right corner of a nominal rectangle that intersects the bottom left corner of the glyph’s bounding box.
	CutOutSW fixed.Point26_6
	// The Cartesian coordinates in staff spaces of the bottom right corner of a nominal rectangle that intersects the top left corner of the glyph’s bounding box.
	CutOutNW fixed.Point26_6
	// The Cartesian coordinates in staff spaces of the position at which the glyph graceNoteSlashStemUp should be positioned relative to the stem-up flag of an unbeamed grace note; alternatively, the bottom left corner of a diagonal line drawn instead of using the above glyph.
	GraceNoteSlashSW fixed.Point26_6
	// The Cartesian coordinates in staff spaces of the top right corner of a diagonal line drawn instead of using the glyph graceNoteSlashStemUp for a stem-up flag of an unbeamed grace note.
	GraceNoteSlashNE fixed.Point26_6
	// The Cartesian coordinates in staff spaces of the position at which the glyph graceNoteSlashStemDown should be positioned relative to the stem-down flag of an unbeamed grace note; alternatively, the top left corner of a diagonal line drawn instead of using the above glyph.
	GraceNoteSlashNW fixed.Point26_6
	// The Cartesian coordinates in staff spaces of the bottom right corner of a diagonal line drawn instead of using the glyph graceNoteSlashStemDown for a stem-down flag of an unbeamed grace note.
	GraceNoteSlashSE fixed.Point26_6
	// The Cartesian coordinates in staff spaces of the horizontal position at which a glyph repeats, i.e. the position at which the same glyph or another of the same group should be positioned to ensure correct tessellation. This is used for e.g. multi-segment lines and the component glyphs that make up trills and mordents.
	RepeatOffset fixed.Point26_6
	// The Cartesian coordinates in staff spaces of the left-hand edge of a notehead with a non-zero left-hand side bearing (e.g. a double whole, or breve, notehead with two vertical lines at each side), to assist in the correct horizontal alignment of these noteheads with other noteheads with zero-width left-side bearings.
	NoteheadOrigin fixed.Point26_6
	// The Cartesian coordinates in staff spaces of the optical center of the glyph, to assist in the correct horizontal alignment of the glyph relative to a notehead or stem. Currently recommended for use with glyphs in the Dynamics range.
	OpticalCenter fixed.Point26_6
}

type EngravingDefaults struct {
	// This list may also use the generic font family values defined in CSS, i.e. serif, sans-serif, cursive, fantasy, and monospace. Generic font family names should be listed after specific font families.
	// An array containing the text font family (or families, in descending order of preference) that are ideally paired with this music
	TextFontFamily []string `json:"textFontFamily"`
	// The thickness of each staff line
	StaffLineThickness fixed.Int26_6 `json:"staffLineThickness"`
	// The thickness of a stem
	StemThickness fixed.Int26_6 `json:"stemThickness"`
	// The thickness of a beam
	BeamThickness fixed.Int26_6 `json:"beamThickness"`
	// The distance between the inner edge of the primary and outer edge of subsequent secondary beams
	BeamSpacing fixed.Int26_6 `json:"beamSpacing"`
	// The thickness of a leger line (normally somewhat thicker than a staff line)
	LegerLineThickness fixed.Int26_6 `json:"legerLineThickness"`
	// The amount by which a leger line should extend either side of a notehead, scaled proportionally with the notehead's size, e.g.  when scaled down as a grace note
	LegerLineExtension fixed.Int26_6 `json:"legerLineExtension"`
	// The thickness of the end of a slur
	SlurEndpointThickness fixed.Int26_6 `json:"slurEndpointThickness"`
	// The thickness of the mid-point of a slur (i.e. its thickest point)
	SlurMidpointThickness fixed.Int26_6 `json:"slurMidpointThickness"`
	// The thickness of the end of a tie
	TieEndpointThickness fixed.Int26_6 `json:"tieEndpointThickness"`
	// The thickness of the mid-point of a tie
	TieMidpointThickness fixed.Int26_6 `json:"tieMidpointThickness"`
	// The thickness of a thin barline, e.g. a normal barline, or each of the lines of a double barline
	ThinBarlineThickness fixed.Int26_6 `json:"thinBarlineThickness"`
	// The thickness of a thick barline, e.g. in a final barline or a repeat barline
	ThickBarlineThickness fixed.Int26_6 `json:"thickBarlineThickness"`
	// The thickness of a dashed barline
	DashedBarlineThickness fixed.Int26_6 `json:"dashedBarlineThickness"`
	// The length of the dashes to be used in a dashed barline
	DashedBarlineDashLength fixed.Int26_6 `json:"dashedBarlineDashLength"`
	// The length of the gap between dashes in a dashed barline
	DashedBarlineGapLength fixed.Int26_6 `json:"dashedBarlineGapLength"`
	// The default distance between multiple thin barlines when locked together, e.g. between two thin barlines making a double barline, measured from the right-hand edge of the left barline to the left-hand edge of the right barline.
	BarlineSeparation fixed.Int26_6 `json:"barlineSeparation"`
	// The default distance between a pair of thin and thick barlines when locked together, e.g. between the thin and thick barlines making a final barline, or between the thick and thin barlines making a start repeat barline.
	ThinThickBarlineSeparation fixed.Int26_6 `json:"thinThickBarlineSeparation"`
	// The default horizontal distance between the dots and the inner barline of a repeat barline, measured from the edge of  the dots to the edge of the barline.
	RepeatBarlineDotSeparation fixed.Int26_6 `json:"repeatBarlineDotSeparation"`
	// The thickness of the vertical line of a bracket grouping staves together
	BracketThickness fixed.Int26_6 `json:"bracketThickness"`
	// The thickness of the vertical line of a sub-bracket grouping staves belonging to the same instrument together
	SubBracketThickness fixed.Int26_6 `json:"subBracketThickness"`
	// The thickness of a crescendo/diminuendo hairpin
	HairpinThickness fixed.Int26_6 `json:"hairpinThickness"`
	// The thickness of the dashed line used for an octave line
	OctaveLineThickness fixed.Int26_6 `json:"octaveLineThickness"`
	// The thickness of the line used for piano pedaling
	PedalLineThickness fixed.Int26_6 `json:"pedalLineThickness"`
	// The thickness of the brackets drawn to indicate repeat endings
	RepeatEndingLineThickness fixed.Int26_6 `json:"repeatEndingLineThickness"`
	// The thickness of the line used for the shaft of an arrow
	ArrowShaftThickness fixed.Int26_6 `json:"arrowShaftThickness"`
	// The thickness of the lyric extension line to indicate a melisma in vocal music
	LyricLineThickness fixed.Int26_6 `json:"lyricLineThickness"`
	// The thickness of a box drawn around text instructions (e.g. rehearsal marks)
	TextEnclosureThickness fixed.Int26_6 `json:"textEnclosureThickness"`
	// The thickness of the brackets drawn either side of tuplet numbers
	TupletBracketThickness fixed.Int26_6 `json:"tupletBracketThickness"`
	// The thickness of the horizontal line drawn between two vertical lines, known as the H-bar, in a multi-bar rest
	HBarThickness fixed.Int26_6 `json:"hBarThickness"`
}

type metadata struct {
	FontName             string                     `json:"fontName"`
	FontVersion          json.RawMessage            `json:"fontVersion"`
	EngravingDefaults    engravingDefaults          `json:"engravingDefaults"`
	GlyphAdvanceWidths   map[string]json.Number     `json:"glyphAdvanceWidths"`
	GlyphBBoxes          map[string]glyphBBox       `json:"glyphBBoxes"`
	GlyphsWithAlternates map[string]glyphAlternates `json:"glyphsWithAlternates"`
	GlyphsWithAnchors    map[string]glyphAnchors    `json:"glyphsWithAnchors"`
	Ligatures            map[string]glyphLigatures  `json:"ligatures"`
	Sets                 map[string]glyphSet        `json:"sets"`
	OptionalGlyphs       map[string]optionalGlyph   `json:"optionalGlyphs"`
}

type glyphBBox struct {
	BBoxNE [2]json.Number `json:"bBoxNE"`
	BBoxSW [2]json.Number `json:"bBoxSW"`
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

type engravingDefaults struct {
	TextFontFamily             []string    `json:"textFontFamily"`
	StaffLineThickness         json.Number `json:"staffLineThickness"`
	StemThickness              json.Number `json:"stemThickness"`
	BeamThickness              json.Number `json:"beamThickness"`
	BeamSpacing                json.Number `json:"beamSpacing"`
	LegerLineThickness         json.Number `json:"legerLineThickness"`
	LegerLineExtension         json.Number `json:"legerLineExtension"`
	SlurEndpointThickness      json.Number `json:"slurEndpointThickness"`
	SlurMidpointThickness      json.Number `json:"slurMidpointThickness"`
	TieEndpointThickness       json.Number `json:"tieEndpointThickness"`
	TieMidpointThickness       json.Number `json:"tieMidpointThickness"`
	ThinBarlineThickness       json.Number `json:"thinBarlineThickness"`
	ThickBarlineThickness      json.Number `json:"thickBarlineThickness"`
	DashedBarlineThickness     json.Number `json:"dashedBarlineThickness"`
	DashedBarlineDashLength    json.Number `json:"dashedBarlineDashLength"`
	DashedBarlineGapLength     json.Number `json:"dashedBarlineGapLength"`
	BarlineSeparation          json.Number `json:"barlineSeparation"`
	ThinThickBarlineSeparation json.Number `json:"thinThickBarlineSeparation"`
	RepeatBarlineDotSeparation json.Number `json:"repeatBarlineDotSeparation"`
	BracketThickness           json.Number `json:"bracketThickness"`
	SubBracketThickness        json.Number `json:"subBracketThickness"`
	HairpinThickness           json.Number `json:"hairpinThickness"`
	OctaveLineThickness        json.Number `json:"octaveLineThickness"`
	PedalLineThickness         json.Number `json:"pedalLineThickness"`
	RepeatEndingLineThickness  json.Number `json:"repeatEndingLineThickness"`
	ArrowShaftThickness        json.Number `json:"arrowShaftThickness"`
	LyricLineThickness         json.Number `json:"lyricLineThickness"`
	TextEnclosureThickness     json.Number `json:"textEnclosureThickness"`
	TupletBracketThickness     json.Number `json:"tupletBracketThickness"`
	HBarThickness              json.Number `json:"hBarThickness"`
}

type glyphAnchors struct {
	SplitStemUpSE    [2]json.Number `json:"splitStemUpSE"`
	SplitStemUpSW    [2]json.Number `json:"splitStemUpSW"`
	SplitStemDownNE  [2]json.Number `json:"splitStemDownNE"`
	SplitStemDownNW  [2]json.Number `json:"splitStemDownNW"`
	StemUpSE         [2]json.Number `json:"stemUpSE"`
	StemDownNW       [2]json.Number `json:"stemDownNW"`
	StemUpNW         [2]json.Number `json:"stemUpNW"`
	StemDownSW       [2]json.Number `json:"stemDownSW"`
	NominalWidth     [2]json.Number `json:"nominalWidth"`
	NumeralTop       [2]json.Number `json:"numeralTop"`
	NumeralBottom    [2]json.Number `json:"numeralBottom"`
	CutOutNE         [2]json.Number `json:"cutOutNE"`
	CutOutSE         [2]json.Number `json:"cutOutSE"`
	CutOutSW         [2]json.Number `json:"cutOutSW"`
	CutOutNW         [2]json.Number `json:"cutOutNW"`
	GraceNoteSlashSW [2]json.Number `json:"graceNoteSlashSW"`
	GraceNoteSlashNE [2]json.Number `json:"graceNoteSlashNE"`
	GraceNoteSlashNW [2]json.Number `json:"graceNoteSlashNW"`
	GraceNoteSlashSE [2]json.Number `json:"graceNoteSlashSE"`
	RepeatOffset     [2]json.Number `json:"repeatOffset"`
	NoteheadOrigin   [2]json.Number `json:"noteheadOrigin"`
	OpticalCenter    [2]json.Number `json:"opticalCenter"`
}
