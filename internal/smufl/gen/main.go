package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

//go:embed glyphnames.json
var glyphnamesData []byte

func main() {
	flag.Parse()

	glyphs := map[string]Glyph{}
	if err := json.Unmarshal(glyphnamesData, &glyphs); err != nil {
		log.Fatal(err)
	}

	names := []string{}
	for name := range glyphs {
		names = append(names, name)
	}
	sort.Strings(names)

	var b bytes.Buffer
	fmt.Fprintf(&b, "package smufl\n\n")
	fmt.Fprintf(&b, "const (\n")
	for _, name := range names {
		glyph := glyphs[name]

		fmt.Fprintf(&b, "\t// %s\n", glyph.Description)
		fmt.Fprintf(&b, "\t%s = '%s'\n",
			CanonicalGlyphName(name),
			strings.ReplaceAll(glyph.Codepoint, "U+", "\\u"),
		)
	}
	fmt.Fprintf(&b, ")\n\n")

	fmt.Fprintf(&b, "var nameToCodepoint = map[string]rune{\n")
	for _, name := range names {
		fmt.Fprintf(&b, "\t%q: %s,\n", name, CanonicalGlyphName(name))
	}
	fmt.Fprintf(&b, "}\n")

	fmt.Fprintf(&b, "func NameToCodepoint(name string) (rune, bool) {\n")
	fmt.Fprintf(&b, "\tr, ok := nameToCodepoint[name]\n")
	fmt.Fprintf(&b, "\treturn r, ok\n")
	fmt.Fprintf(&b, "}\n")

	os.WriteFile("classes.go", b.Bytes(), 0644)
}

type Glyph struct {
	Codepoint   string // U+E3AA
	Description string
}

func CanonicalGlyphName(name string) string {
	if unicode.IsDigit(rune(name[0])) {
		switch name[0] {
		case '0':
			name = "Zero" + strings.Title(name[1:])
		case '1':
			name = "One" + strings.Title(name[1:])
		case '2':
			name = "Two" + strings.Title(name[1:])
		case '3':
			name = "Three" + strings.Title(name[1:])
		case '4':
			name = "Four" + strings.Title(name[1:])
		case '5':
			name = "Five" + strings.Title(name[1:])
		case '6':
			name = "Six" + strings.Title(name[1:])
		case '7':
			name = "Seven" + strings.Title(name[1:])
		case '8':
			name = "Eight" + strings.Title(name[1:])
		case '9':
			name = "Nine" + strings.Title(name[1:])
		}
	}
	return strings.Title(name)
}
