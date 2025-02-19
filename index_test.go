package tableimage

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkTableImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dimBlack, err := GetColorByHex("#171717")
		require.NoError(b, err)

		red, err := GetColorByHex("#d03136")
		require.NoError(b, err)

		blue, err := GetColorByHex("#0075e2")
		require.NoError(b, err)

		green, err := GetColorByHex("#00d841")
		require.NoError(b, err)

		ti, err := Init(dimBlack, "MicrosoftYahei.ttf")
		if err != nil {
			panic(err)
		}

		ti.AddTH(
			TR{
				Tds: []TD{
					{
						Text: Text{"XXXXXXX", color.White},
					},
					{
						Text: Text{"BBBBB", red},
					},
					{
						Text: Text{"AAAAA", blue},
					},
					{
						Text: Text{"NNNNNN", green},
					},
					{
						Text: Text{"BBBBB对", red},
					},
					{
						Text: Text{"SDFASDFASDF", blue},
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

		err = ti.Save(PNG, "./test.png")
		require.NoError(b, err)
	}
}

func TestTableImage(t *testing.T) {
	dimBlack, err := GetColorByHex("#171717")
	require.NoError(t, err)

	red, err := GetColorByHex("#d03136")
	require.NoError(t, err)

	blue, err := GetColorByHex("#0075e2")
	require.NoError(t, err)

	green, err := GetColorByHex("#00d841")
	require.NoError(t, err)

	// anotherGreen, err := GetColorByHex("#32a852")
	// require.NoError(t, err)

	// purple, err := GetColorByHex("#7732a8")
	// require.NoError(t, err)

	// pink, err := GetColorByHex("#f0f")
	// require.NoError(t, err)

	ti, err := Init(dimBlack, "MicrosoftYahei.ttf")
	if err != nil {
		panic(err)
	}

	ti.AddTH(
		TR{
			// BackgroundColor: color.White,
			// BorderColor:     anotherGreen,
			Tds: []TD{
				{
					Text: Text{"XXXXXXX", color.White},
				},
				{
					Text: Text{"BBBBB", red},
					// BackgroundColor: purple,
				},
				{
					Text: Text{"AAAAA", blue},
				},
				{
					Text: Text{"NNNNNN", green},
				},
				{
					Text: Text{"庄对", red},
				},
				{
					Text: Text{"SDFASDFASDF", blue},
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

	err = ti.Save(PNG, "./test.png")
	require.NoError(t, err)
}
