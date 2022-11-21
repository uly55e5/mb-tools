package mzmlReader

import (
	"bytes"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	tpls "github.com/go-echarts/go-echarts/v2/templates"
	"io"
	"regexp"
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
	pl := PlotMSSpectrum(peaks.Values[0], "Peaks")
	pl.AssetsHost = "https://go-echarts.github.io/go-echarts-assets/assets/"
	renderer := NewPartialRender(pl)
	buf := new(bytes.Buffer)
	err := renderer.Render(buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil

}

type partialRender struct {
	c      interface{}
	before []func()
}

const ModPartial = "partial"

var pat = regexp.MustCompile(`(__f__")|("__f__)|(__f__)`)

var partialTpl = `
{{- define "partial" }}
    {{- template "base" . }}
<style>
    .container {margin-top:30px; display: flex;justify-content: center;align-items: center;}
    .item {margin: auto;}
</style>
{{ end }}
`

// NewPartialRender returns a render implementation for Chart.
func NewPartialRender(c interface{}, before ...func()) render.Renderer {
	return &partialRender{c: c, before: before}
}

// Render renders the chart into the given io.Writer.
func (r *partialRender) Render(w io.Writer) error {
	for _, fn := range r.before {
		fn()
	}

	contents := []string{tpls.HeaderTpl, tpls.BaseTpl, partialTpl}
	tpl := render.MustTemplate(ModPartial, contents)

	var buf bytes.Buffer
	r.c.(components.Charter).Validate()
	if err := tpl.ExecuteTemplate(&buf, ModPartial, r.c); err != nil {
		return err
	}

	content := pat.ReplaceAll(buf.Bytes(), []byte(""))

	_, err := w.Write(content)
	return err
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
