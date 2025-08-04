package tableimage

import (
	"os"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

func initFontFace(path string, fontSize float64) (font.Face, error) {
	fontBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ttf, err := sfnt.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	face, err := opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}

	return face, nil
}

func wrapText(input string) []string {
	var wrapped []string

	maxLetterPerLine := DefaultConfig.LineLen()/DefaultConfig.fontSize + 1

	for _, word := range strings.Split(input, "\n") {
		runes := []rune(word)

		for i, j := 0, len(runes); i < j; i += maxLetterPerLine {
			end := i + maxLetterPerLine

			if end > j {
				end = j
			}

			wrapped = append(wrapped, string(runes[i:end]))
		}
	}

	return wrapped
}
