package tableimage

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
)

func (ti *tableImage) write(fileType FileType, w io.Writer) error {
	switch fileType {
	case JPEG, "jpeg":
		if err := jpeg.Encode(w, ti.img, nil); err != nil {
			return fmt.Errorf("jpeg.Encode: %w", err)
		}
	case PNG:
		if err := png.Encode(w, ti.img); err != nil {
			return fmt.Errorf("png.Encode: %w", err)
		}
	default:
		return fmt.Errorf("unsupported file type: %s", fileType)
	}

	return nil
}
