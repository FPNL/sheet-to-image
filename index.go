package tableimage

import (
	"fmt"
	"image"
	"image/color"
	"io"

	"golang.org/x/image/font"
)

// FileType the image format png or jpg
type FileType string

type Text struct {
	S string
	C color.Color
}

// TD a table data container
type TD struct {
	Text            Text
	BackgroundColor color.Color
}

// TR the table row
type TR struct {
	BorderColor     color.Color
	BackgroundColor color.Color
	Tds             []TD
}

type tableImage struct {
	width           int
	height          int
	th              TR
	trs             []TR
	backgroundColor color.Color
	img             *image.RGBA
	fontFace        font.Face
}

const (
	PNG  FileType = "png"
	JPEG FileType = "jpg"
)

// Init initialise the table image receiver
func Init(backgroundColor color.Color, fontPath string, config *Config) (tableImage, error) {
	if config != nil {
		DefaultConfig = *config
	}

	face, err := initFontFace(fontPath, float64(DefaultConfig.fontSize))
	if err != nil {
		return tableImage{}, err
	}

	ti := tableImage{
		backgroundColor: backgroundColor,
		fontFace:        face,
	}

	return ti, nil
}

// AddTH adds the table header
func (ti *tableImage) AddTH(th TR) {
	ti.th = th
}

// AddTRs add the table rows
func (ti *tableImage) AddTRs(trs []TR) {
	ti.trs = trs
}

// Write saves the table
func (ti *tableImage) Write(fileType FileType, w io.Writer) error {
	ti.calculateHeight()
	ti.calculateWidth()
	ti.setRgba()

	ti.drawTH()

	ti.drawTR()

	err := ti.write(fileType, w)
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}

	return nil
}
