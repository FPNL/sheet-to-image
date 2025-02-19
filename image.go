package tableimage

import (
	"image"
	"image/draw"
	"os"
)

func (ti *tableImage) setRgba() {
	img := image.NewRGBA(image.Rect(0, 0, ti.width, ti.height))

	// set image background
	draw.Draw(
		img,
		img.Bounds(),
		image.NewUniform(ti.backgroundColor),
		image.Point{},
		draw.Src,
	)
	ti.img = img
}

func addImage(baseImage *image.RGBA, path string, point image.Point) error {
	templateFile, err := os.Open(path)
	if err != nil {
		return err
	}

	template, _, err := image.Decode(templateFile)

	if err != nil {
		return err
	}

	draw.Draw(baseImage, baseImage.Bounds(), template, point, draw.Over)

	return nil
}
