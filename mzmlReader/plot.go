package mzmlReader

import (
	"bytes"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
)

func PlotChromatogram(chromatogram Chromatogram) (string, error) {
	time := chromatogram.Time
	intens := chromatogram.Intensity
	values := [2][]float64{time, intens}
	pl := PlotMSSpectrum(values, "Peaks")
	renderer := render.NewChartRender(pl)
	buf := new(bytes.Buffer)
	err := renderer.Render(buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func PlotScan(peaks PeakList, ScanId int64) (string, error) {
	pl := PlotMSSpectrum(peaks.Values[ScanId], "Peaks")
	renderer := render.NewChartRender(pl)
	buf := new(bytes.Buffer)
	err := renderer.Render(buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil

}

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
