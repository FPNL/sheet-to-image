Fork from https://github.com/Techbinator/go-table-image

---

### Usage

```go
package main

import (
	"image/color"
	"os"

	"tableimage"
)

func main() {
	dimBlack, err := tableimage.GetColorByHex("#171717")
	red, err := tableimage.GetColorByHex("#d03136")
	blue, err := tableimage.GetColorByHex("#0075e2")
	green, err := tableimage.GetColorByHex("#00d841")
	if err != nil {
		panic(err)
	}

	ti, err := tableimage.Init(dimBlack, "MicrosoftYahei.ttf", &tableimage.Config{
		fontSize:     14,
		hPadding:     0,
		vPadding:     2,
		wrapWordsLen: 12,
	})
	if err != nil {
		panic(err)
	}

	ti.AddTH(
		tableimage.TR{
			Tds: []tableimage.TD{
				{
					Text: tableimage.Text{"That", color.White},
				},
				{
					Text: tableimage.Text{"Hello", red},
					// BackgroundColor: purple,
				},
				{
					Text: tableimage.Text{"Beach", blue},
				},
				{
					Text: tableimage.Text{"Peach", green},
				},
				{
					Text: tableimage.Text{"hello", red},
				},
				{
					Text: tableimage.Text{"leisure", blue},
				},
			},
		},
	)

	ti.AddTRs(
		[]tableimage.TR{
			{
				BackgroundColor: color.White,
				Tds: []tableimage.TD{
					{
						Text: tableimage.Text{S: "2223"},
					},
					{
						Text: tableimage.Text{S: "Really cool product on two lines"},
					},
					{},
					{
						Text: tableimage.Text{S: "2000$"},
					},
				},
			},
			{
				// BackgroundColor: purple,
				Tds: []tableimage.TD{
					{},
					{
						Text: tableimage.Text{S: "11"},
					},
					{
						Text: tableimage.Text{S: "A more cooler product this time on 3 lines"},
					},
					{
						Text: tableimage.Text{S: "200$"},
					},
				},
			},
			{
				BackgroundColor: color.White,
				Tds: []tableimage.TD{
					{
						Text: tableimage.Text{S: "2231"},
					},
					{
						Text: tableimage.Text{S: "Lenovo"},
					},
					{
						Text: tableimage.Text{S: "20400$"},
					},
				},
			},
		},
	)

	f, err := os.Create("./test.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = ti.Write(tableimage.PNG, f)
	if err != nil {
		panic(err)
	}
}
```

Outputs:

![Example](test.png "Example")
