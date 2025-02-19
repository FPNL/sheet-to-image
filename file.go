package tableimage

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
)

func (ti *tableImage) saveFile(fileType FileType, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer f.Close()

	switch fileType {
	case JPEG, "jpeg":
		if err = jpeg.Encode(f, ti.img, nil); err != nil {
			return fmt.Errorf("jpeg.Encode: %w", err)
		}
	case PNG:
		if err = png.Encode(f, ti.img); err != nil {
			return fmt.Errorf("png.Encode: %w", err)
		}
	default:
		return fmt.Errorf("unsupported file type: %s", fileType)
	}

	return nil
}
