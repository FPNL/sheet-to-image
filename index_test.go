package tableimage

import (
	"image/color"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkTableImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		commonTest(b)
	}
}

func TestTableImage(t *testing.T) {
	commonTest(t)
}

func commonTest(t require.TestingT) {
	dimBlack, err := GetColorByHex("#171717")
	require.NoError(t, err)

	red, err := GetColorByHex("#d03136")
	require.NoError(t, err)

	blue, err := GetColorByHex("#0075e2")
	require.NoError(t, err)

	green, err := GetColorByHex("#00d841")
	require.NoError(t, err)

	ti, err := Init(dimBlack, "MicrosoftYahei.ttf")
	if err != nil {
		panic(err)
	}

	ti.AddTH(
		TR{
			Tds: []TD{
				{
					Text: Text{"Order\nM28422C060AU", color.White},
				},
				{
					Text: Text{"item", red},
				},
				{
					Text: Text{"price", blue},
				},
				{
					Text: Text{"tax", green},
				},
				{
					Text: Text{"quantity", red},
				},
				{
					Text: Text{"total", blue},
				},
			},
		},
	)

	var tableRows []TR

	for _ = range 1 {
		tableRows = append(tableRows,
			TR{
				BackgroundColor: color.White,
				Tds: []TD{
					{
						Text: Text{S: "1"},
					},
					{
						Text: Text{S: "2"},
					},
					{
						Text: Text{S: "3"},
					},
					{
						Text: Text{S: "4"},
					},
					{
						Text: Text{S: "5"},
					},
					{
						Text: Text{S: "6"},
					},
				},
			},
			TR{
				Tds: []TD{
					{
						Text: Text{"/", color.White},
					},
					{
						Text: Text{"8", color.White},
					},
					{
						Text: Text{"9", color.White},
					},
					{
						Text: Text{"10", color.White},
					},
					{
						Text: Text{"11", color.White},
					},
					{
						Text: Text{"12", color.White},
					},
				},
			})
	}

	ti.AddTRs(tableRows)

	f, err := os.Create("./test.png")
	require.NoError(t, err)
	defer f.Close()

	err = ti.Write(PNG, f)
	require.NoError(t, err)
}
