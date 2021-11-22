package bravura

import (
	_ "embed"

	"github.com/loov/noteviz/internal/smufl"
)

//go:embed otf/Bravura.otf
var otf []byte

//go:embed bravura_metadata.json
var metadata []byte

func Font() *smufl.Font {
	return smufl.MustParseFont(otf, metadata)
}
