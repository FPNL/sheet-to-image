package tableimage

import (
	"errors"
	"fmt"
	"image/color"
	"strings"
)

func GetColorByHex(hexColor string) (color.Color, error) {
	var format string

	hexColor = strings.TrimPrefix(hexColor, "#")
	rgba := color.RGBA{}
	rgba.A = 255

	switch len(hexColor) {
	case 3:
		format = "%1x%1x%1x"
	case 6:
		format = "%02x%02x%02x"
	case 8:
		format = "%02x%02x%02x%02x"
	default:
		return nil, errors.New("invalid length")
	}

	_, err := fmt.Sscanf(hexColor, format, &rgba.R, &rgba.G, &rgba.B)
	if err != nil {
		return nil, fmt.Errorf("fmt.Sscanf: %w", err)
	}

	if len(hexColor) == 3 {
		rgba.R |= rgba.R << 4
		rgba.G |= rgba.G << 4
		rgba.B |= rgba.B << 4
	}

	return rgba, nil
}
