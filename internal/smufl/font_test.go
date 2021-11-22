package smufl_test

import (
	"testing"

	"github.com/loov/noteviz/internal/font/bravura"
)

func TestBravura(t *testing.T) {
	font := bravura.Font()
	t.Logf("%#v\n", font)
}
