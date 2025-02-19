package tableimage

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// addString 文字預設給黑色
func addString(img draw.Image, fontFace font.Face, point image.Point, text []string, clr color.Color) {
	var x, y = point.X, point.Y

	if clr == nil {
		clr = color.Black
	}

	for i, s := range text {
		d := &font.Drawer{
			Dst:  img,
			Src:  image.NewUniform(clr),
			Face: fontFace,
			Dot: fixed.Point26_6{
				X: fixed.Int26_6(x * 64),
				Y: fixed.Int26_6(((i + 1) * y) * 64),
			},
		}

		d.DrawString(s)
	}
}

// Thx to https://github.com/StephaneBunel/bresenham
func addLine(img draw.Image, p1 image.Point, p2 image.Point, clr color.Color) {
	var dx, dy, e, slope int
	var x1, y1, x2, y2 = p1.X, p1.Y, p2.X, p2.Y

	// Because drawing p1 -> p2 is equivalent to draw p2 -> p1,
	// I sort points in x-axis order to handle only half of possible cases.
	if x1 > x2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}

	dx, dy = x2-x1, y2-y1
	// Because point is x-axis ordered, dx cannot be negative
	if dy < 0 {
		dy = -dy
	}

	switch {

	// Is line a point ?
	case x1 == x2 && y1 == y2:
		img.Set(x1, y1, clr)

	// Is line a horizontal ?
	case y1 == y2:
		for ; dx != 0; dx-- {
			img.Set(x1, y1, clr)
			x1++
		}
		img.Set(x1, y1, clr)

	// Is line a vertical ?
	case x1 == x2:
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for ; dy != 0; dy-- {
			img.Set(x1, y1, clr)
			y1++
		}
		img.Set(x1, y1, clr)

	// Is line a diagonal ?
	case dx == dy:
		if y1 < y2 {
			for ; dx != 0; dx-- {
				img.Set(x1, y1, clr)
				x1++
				y1++
			}
		} else {
			for ; dx != 0; dx-- {
				img.Set(x1, y1, clr)
				x1++
				y1--
			}
		}
		img.Set(x1, y1, clr)

	// wider than high ?
	case dx > dy:
		if y1 < y2 {
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				img.Set(x1, y1, clr)
				x1++
				e -= dy
				if e < 0 {
					y1++
					e += slope
				}
			}
		} else {
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				img.Set(x1, y1, clr)
				x1++
				e -= dy
				if e < 0 {
					y1--
					e += slope
				}
			}
		}
		img.Set(x2, y2, clr)

	// higher than wide.
	default:
		if y1 < y2 {
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				img.Set(x1, y1, clr)
				y1++
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		} else {
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				img.Set(x1, y1, clr)
				y1--
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		}
		img.Set(x2, y2, clr)
	}
}

func addSurface(img draw.Image, rectangle image.Rectangle, clr color.Color) {
	var x1, y1, x2, y2 = rectangle.Min.X, rectangle.Min.Y, rectangle.Max.X, rectangle.Max.Y

	for x1 <= x2 {
		addLine(
			img,
			image.Point{X: x1, Y: y1},
			image.Point{X: x1, Y: y2},
			clr,
		)
		x1++
	}
}

func addBorder(img draw.Image, rectangle image.Rectangle, clr color.Color) {
	/*
		// 左上點
		leftHigh := image.Point{X: rectangle.Min.X, Y: rectangle.Min.Y}
		// 左下點
		leftLow := image.Point{X: rectangle.Min.X, Y: rectangle.Max.Y - 1}
		// 右上點
		rightHigh := image.Point{X: rectangle.Max.X - 1, Y: rectangle.Min.Y}
		// 右下點
		rightLow := image.Point{X: rectangle.Max.X - 1, Y: rectangle.Max.Y - 1}
	*/

	// 左上點
	leftHigh := image.Point{X: rectangle.Min.X, Y: rectangle.Min.Y}
	// 左下點
	leftLow := image.Point{X: rectangle.Min.X, Y: rectangle.Max.Y}
	// 右上點
	rightHigh := image.Point{X: rectangle.Max.X, Y: rectangle.Min.Y}
	// 右下點
	rightLow := image.Point{X: rectangle.Max.X, Y: rectangle.Max.Y}

	// 上面的線
	addLine(img, leftHigh, rightHigh, clr)
	// 下面的線
	addLine(img, leftLow, rightLow, clr)
	// 左邊的線
	addLine(img, leftHigh, leftLow, clr)
	// 右邊的線
	addLine(img, rightHigh, rightLow, clr)
}

func addColumn(img draw.Image, rectangle image.Rectangle, text []string, fontFace font.Face, txtColor, bgColor, borderColor color.Color) {
	if bgColor == nil {
		bgColor = color.Black
	}

	if borderColor == nil {
		borderColor = color.Black
	}

	var maxString string

	for _, s := range text {
		if len(s) > len(maxString) {
			maxString = s
		}
	}

	// 字串要從哪個點開始繪製，會往右下繪製
	stringDot := rectangle.Max.
		Add(rectangle.Min).
		Div(2). // 代表兩個點的中心點
		Sub(image.Point{
			X: font.MeasureString(fontFace, maxString).Ceil() / 2,
			Y: -fontSize / 3,
		}) // 計算字的長度，把它置中

	addSurface(img, rectangle, bgColor)
	addString(img, fontFace, stringDot, text, txtColor)
	addBorder(img, rectangle, borderColor)
}

func (ti *tableImage) addColumn(rectangle image.Rectangle, text Text, bgColor, borderColor color.Color) {
	wrappedTexts := wrapText(text.S)

	if bgColor == nil {
		bgColor = ti.backgroundColor
	}

	addColumn(
		ti.img,
		rectangle,
		wrappedTexts,
		ti.fontFace,
		text.C,
		bgColor,
		borderColor,
	)
}

func (ti *tableImage) drawTH() (err error) {
	for colNo, td := range ti.th.Tds {
		bgColor := ti.th.BackgroundColor
		if td.BackgroundColor != nil {
			bgColor = td.BackgroundColor
		}

		ti.addColumn(
			image.Rect(colNo*width, 0, (colNo+1)*width, height),
			td.Text,
			bgColor,
			ti.th.BorderColor,
		)
	}

	return nil
}

func (ti *tableImage) drawTR() (err error) {
	totalRowHeight := height

	for _, tds := range ti.trs {
		// start with a row space
		maxRowHeight := height

	DrawRow:
		for i, td := range tds.Tds {
			currRowHeight := len(wrapText(td.Text.S)) * height

			if maxRowHeight < currRowHeight {
				maxRowHeight = currRowHeight
				goto DrawRow
			}

			x0 := i * width
			y0 := totalRowHeight
			x1 := (i + 1) * width
			y1 := maxRowHeight + totalRowHeight

			var bgColor = td.BackgroundColor
			if td.BackgroundColor == nil {
				bgColor = tds.BackgroundColor
			}

			ti.addColumn(
				image.Rect(x0, y0, x1, y1),
				td.Text,
				bgColor,
				tds.BorderColor,
			)
		}

		// if the row has less than columns than the header
		if len(ti.th.Tds) > len(tds.Tds) {
			for a, b := len(tds.Tds), len(ti.th.Tds); a <= b; a++ {
				x0 := a * width
				y0 := totalRowHeight
				x1 := (a + 1) * width
				y1 := maxRowHeight + totalRowHeight

				ti.addColumn(
					image.Rect(x0, y0, x1, y1),
					Text{},
					tds.BackgroundColor,
					tds.BorderColor,
				)
			}
		}

		totalRowHeight += maxRowHeight

	}

	return nil
}
