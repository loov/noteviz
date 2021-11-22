package smufl

import "gioui.org/font/opentype"

type Font struct {
	Face *opentype.Font
}

func ParseFont(fontData, fontMetadata []byte) (*Font, error) {
	font := &Font{}

	face, err := opentype.Parse(fontData)
	if err != nil {
		return nil, err
	}
	font.Face = face

	return font, err
}

func MustParseFont(fontData, fontMetadata []byte) *Font {
	font, err := ParseFont(fontData, fontMetadata)
	if err != nil {
		panic(err)
	}
	return font
}
