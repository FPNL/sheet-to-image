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

	// Split string into array of words
	words := strings.Fields(input)
	wordsLength := len(words)

	if wordsLength == 0 {
		return wrapped
	}

	var lineText string

	for i, word := range words {
		if len(lineText)+len(word)+1 >= wrapWordsLen {
			wrapped = append(wrapped, lineText)
			lineText = word
		} else {
			if lineText == "" {
				lineText += word
			} else {
				lineText += " " + word
			}
			// if it is the last word
			if i == wordsLength-1 {
				wrapped = append(wrapped, lineText)
			}
		}
	}

	return wrapped
}
