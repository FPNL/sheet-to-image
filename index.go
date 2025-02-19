package tableimage

import (
	"fmt"
	"image"
	"image/color"

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
	letterPerPx           = 10
	wrapWordsLen          = 20
	height                = 26
	width                 = wrapWordsLen * letterPerPx
	fontSize              = 12
	PNG          FileType = "png"
	JPEG         FileType = "jpg"
)

// Init initialise the table image receiver
func Init(backgroundColor color.Color, fontPath string) (tableImage, error) {
	face, err := initFontFace(fontPath, fontSize)
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

// Save saves the table
func (ti *tableImage) Save(fileType FileType, filePath string) error {
	ti.calculateHeight()
	ti.calculateWidth()

	ti.setRgba()

	if err := ti.drawTH(); err != nil {
		return fmt.Errorf("drawTH: %w", err)
	}

	if err := ti.drawTR(); err != nil {
		return fmt.Errorf("drawTR: %w", err)
	}

	err := ti.saveFile(fileType, filePath)
	if err != nil {
		return fmt.Errorf("saveFile: %w", err)
	}

	return nil
}
