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
	// start from 1 since we have th
	totalRowNo := 1
	for _, tr := range ti.trs {
		totalRowNo += maxRows(tr)
	}

	ti.height = totalRowNo * height
}

func (ti *tableImage) calculateWidth() {
	totalColumnNo := len(ti.th.Tds)

	ti.width = totalColumnNo * width
}
