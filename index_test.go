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

	ti, err := Init(dimBlack, "MicrosoftYahei.ttf", &Config{
		fontSize:     14,
		hPadding:     0,
		vPadding:     4,
		wrapWordsLen: 12,
	})
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

	for range 1 {
		tableRows = append(tableRows,
			TR{
				BackgroundColor: color.White,
				Tds: []TD{
					{
						Text: Text{S: "1192837981723981123123123123123123723987"},
					},
					{
						Text: Text{S: "21192837981723981723987"},
					},
					{
						Text: Text{S: "."},
					},
					{
						Text: Text{S: "119283798172398112312312312312312"},
					},
					{
						Text: Text{S: "119283798172398112312312312312312372398712213123123"},
					},
					{
						Text: Text{S: "119283798172398112312312312312312372398712213123123119283798172398112312312312312312372398712213123123"},
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

func TestGetColorByHex(t *testing.T) {
	dimBlack, err := GetColorByHex("#171717")
	red, err := GetColorByHex("#d03136")
	blue, err := GetColorByHex("#0075e2")
	green, err := GetColorByHex("#00d841")
	if err != nil {
		panic(err)
	}

	ti, err := Init(dimBlack, "MicrosoftYahei.ttf", &Config{
		fontSize:     14,
		hPadding:     0,
		vPadding:     2,
		wrapWordsLen: 12,
	})
	if err != nil {
		panic(err)
	}

	ti.AddTH(
		TR{
			Tds: []TD{
				{
					Text: Text{"That", color.White},
				},
				{
					Text: Text{"Hello", red},
					// BackgroundColor: purple,
				},
				{
					Text: Text{"Beach", blue},
				},
				{
					Text: Text{"Peach", green},
				},
				{
					Text: Text{"hello", red},
				},
				{
					Text: Text{"leisure", blue},
				},
			},
		},
	)

	ti.AddTRs(
		[]TR{
			{
				BackgroundColor: color.White,
				Tds: []TD{
					{
						Text: Text{S: "2223"},
					},
					{
						Text: Text{S: "Really cool product on two lines"},
					},
					{},
					{
						Text: Text{S: "2000$"},
					},
				},
			},
			{
				// BackgroundColor: purple,
				Tds: []TD{
					{},
					{
						Text: Text{S: "11"},
					},
					{
						Text: Text{S: "A more cooler product this time on 3 lines"},
					},
					{
						Text: Text{S: "200$"},
					},
				},
			},
			{
				BackgroundColor: color.White,
				Tds: []TD{
					{
						Text: Text{S: "2231"},
					},
					{
						Text: Text{S: "Lenovo"},
					},
					{
						Text: Text{S: "20400$"},
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

	err = ti.Write(PNG, f)
	if err != nil {
		panic(err)
	}
}
