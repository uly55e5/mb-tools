package msVisualizer

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func PlotMSSpectrum(data [2][]float64, title string) *charts.Bar {
	pl := charts.NewBar()
	pl.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title}),
		charts.WithXAxisOpts(opts.XAxis{Name: "m/z"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "intensity"}),
		charts.WithColorsOpts(opts.Colors{"#000055"}),
	)

	mzData := make([]opts.BarData, len(data[1]))
	for i, d := range data[1] {
		mzData[i].Value = d
	}

	pl.SetXAxis(data[0]).
		AddSeries("Category A", mzData)
	return pl
}
