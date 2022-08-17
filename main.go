package main

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/uly55e5/mb-tools/msVisualizer"
	"github.com/uly55e5/mb-tools/mzmlReader"
	"net/http"
)

type pageHandler struct{}

func (handler pageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := mzmlReader.OpenMSData("data/examples/small.pwiz.1.1.mzML")
	if err != nil {
		println(err.Error())
		return
	}
	page := components.NewPage()
	peaks := data.Peaks()
	for i, v := range peaks.Values {
		title := data.Header(i).SpectrumId[0]
		rt := data.Header(i).RetentionTime[0]

		pl := msVisualizer.PlotMSSpectrum(v, title+" RT: "+fmt.Sprintf("%.3f", rt))
		page.AddCharts(pl)
	}

	page.Render(w)
}

func main() {
	err := http.ListenAndServe("localhost:8090", pageHandler{})
	if err != nil {
		println(err.Error())
	}

}
