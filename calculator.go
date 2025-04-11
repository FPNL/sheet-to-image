package tableimage

func maxRows(tr TR) int {
	maxRowHeight := 1
	for _, td := range tr.Tds {
		wrappedText := wrapText(td.Text.S)
		// in case we have a multi line text
		if len(wrappedText) > maxRowHeight {
			maxRowHeight = len(wrappedText)
		}
	}

	return maxRowHeight
}

func (ti *tableImage) calculateHeight() {
	var totalRowNo int

	totalRowNo += maxRows(ti.th)

	for _, tr := range ti.trs {
		totalRowNo += maxRows(tr)
	}

	ti.height = totalRowNo * height
}

func (ti *tableImage) calculateWidth() {
	totalColumnNo := len(ti.th.Tds)

	ti.width = totalColumnNo * width
}
