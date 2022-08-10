package mzmlReader

import (
	"github.com/google/go-cmp/cmp"
	"math"
	"reflect"
	"testing"
)

func TestOpenMSData(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Read Example", args{"../data/examples/small.pwiz.1.1.mzML"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OpenMSData(tt.args.fileName)
		})
	}
}

func TestMSData_InstrumentInfo(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
		want   *InstrumentInfo
	}{
		{"small.pwiz", file, &InstrumentInfo{"Xcalibur ", "LTQ FT", "electrospray ionization", "fourier transform ion cyclotron resonance mass spectrometer", "inductive detector", "Xcalibur 1.1 Beta 7", "", ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.InstrumentInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstrumentInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_CloseMSData(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
	}{
		{"small.pwiz", file},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.msdata.CloseMSData()
		})
	}
}

func TestMSData_FileName(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
		want   string
	}{
		{"small.pwiz", file, "../data/examples/small.pwiz.1.1.mzML"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.FileName(); got != tt.want {
				t.Errorf("FileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Length(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
		want   int
	}{
		{"small.pwiz", file, 48},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.msdata.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Manufacturer(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
		want   string
	}{
		{"small.pwiz", file, "Xcalibur "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.Manufacturer(); got != tt.want {
				t.Errorf("Manufacturer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Model(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
		want   string
	}{
		{"small.pwiz", file, "LTQ FT"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.Model(); got != tt.want {
				t.Errorf("Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Ionisation(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
		want   string
	}{
		{"small.pwiz", file, "electrospray ionization"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.Ionisation(); got != tt.want {
				t.Errorf("Ionisation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Analyzer(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")

	tests := []struct {
		name   string
		msdata *MSData
		want   string
	}{
		{"small.pwiz", file, "fourier transform ion cyclotron resonance mass spectrometer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.Analyzer(); got != tt.want {
				t.Errorf("Analyzer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Detector(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
		want   string
	}{
		{"small.pwiz", file, "inductive detector"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.Detector(); got != tt.want {
				t.Errorf("Detector() = %v, want %v", got, tt.want)
			}
		})
	}
}

var nanEqualOpt = cmp.Comparer(func(x, y float64) bool {
	return (math.IsNaN(x) && math.IsNaN(y)) || x == y
})

func TestMSData_Header(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	type args struct {
		scans []int
	}
	tests := []struct {
		name   string
		msdata *MSData
		args   args
		want   HeaderInfo
	}{
		{"small.pwiz", file, args{[]int{0}},
			HeaderInfo{
				[]int{0},
				[]int{1},
				[]int{1},
				[]int{1},
				[]int{19914},
				[]float64{1.5245068e+07},
				[]float64{0.29610000000000003},
				[]float64{810.415283203125},
				[]float64{1.471973875e+06},
				[]float64{0},
				[]float64{0},
				[]float64{200.00018816645022},
				[]float64{2000.0099466203771},
				[]int{-1},
				[]float64{0},
				[]int{0},
				[]float64{0},
				[]int{-1},
				[]int{-1},
				[]int{-1},
				[]int{-1},
				[]float64{0},
				[]string{"FTMS + p ESI Full ms [200.00-2000.00]"},
				[]string{"controllerType=0 controllerNumber=1 scan=1"},
				[]bool{false},
				[]float64{math.NaN()},
				[]float64{math.NaN()},
				[]float64{math.NaN()},
				[]float64{math.NaN()},
				[]float64{200},
				[]float64{2000},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.msdata.Header(tt.args.scans...); !cmp.Equal(got, tt.want, nanEqualOpt) {
				t.Errorf("Header() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_ChromatogramHeader(t *testing.T) {
	type args struct {
		scans []int
	}
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
		args   args
		want   ChromatogramHeaderInfo
	}{
		{"small.pwiz", file, args{[]int{0}}, ChromatogramHeaderInfo{
			[]string{"TIC"},
			[]int{0},
			[]int{-1},
			[]float64{math.NaN()},
			[]float64{math.NaN()},
			[]float64{math.NaN()},
			[]float64{math.NaN()},
			[]float64{math.NaN()},
			[]float64{math.NaN()},
			[]float64{math.NaN()},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.ChromatogramHeader(tt.args.scans...); !cmp.Equal(got, tt.want, nanEqualOpt) {
				t.Errorf("ChromatogramHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Chromatogram(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	type args struct {
		chromIdx int
	}
	tests := []struct {
		name   string
		msdata *MSData
		args   args
		want   Chromatogram
	}{
		{"small.pwiz", file, args{0}, pwizSmall_Chromatogram_0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.Chromatogram(tt.args.chromIdx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chromatogram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_IsolationWindow(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	result := make([]IsolationWindow, 34)
	for i := range result {
		result[i] = IsolationWindow{0.5, 0.5}
	}
	type args struct {
		uniqe bool
	}
	tests := []struct {
		name   string
		msdata *MSData
		args   args
		want   []IsolationWindow
	}{
		{"small.pwiz", file, args{false}, result},
		{"small.pwiz.Unique", file, args{true}, []IsolationWindow{{0.5, 0.5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.IsolationWindow(tt.args.uniqe); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsolationWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Peaks(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	type args struct {
		scans []int
	}
	tests := []struct {
		name   string
		msdata *MSData
		args   args
		want   PeakList
	}{
		{"small.pwiz", file, args{[]int{0}}, pwizSmall_PeakList_0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.Peaks(tt.args.scans...); !cmp.Equal(got, tt.want) {
				t.Errorf("Peaks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_PeaksCount(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	type args struct {
		scans []int
	}
	tests := []struct {
		name   string
		msdata *MSData
		args   args
		want   PeakCount
	}{
		{"small.pwiz", file, args{[]int{0}}, PeakCount{[]int{19914}, []int{0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.PeaksCount(tt.args.scans...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PeakCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Get3DMap(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	type args struct {
		scans  []int
		lowMz  float64
		highMz float64
		resMZ  float64
	}
	tests := []struct {
		name   string
		msdata *MSData
		args   args
		want   Map3D
	}{
		{"small.pwiz", file, args{[]int{0, 1, 2, 3}, 0, 2000, 0.5}, pwizSmall_3DMap_0_3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.Get3DMap(tt.args.lowMz, tt.args.highMz, tt.args.resMZ, tt.args.scans...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get3DMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_GetRunInfo(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
		want   RunInfo
	}{
		{"small.pwiz", file, RunInfo{
			48,
			162.24594116210938,
			2000.0099466203771,
			0.29610000000000003,
			29.234199999999998,
			[]int{1, 2},
			"2005-07-20T14:44:22",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.GetRunInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRunInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Chromatograms(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	type args struct {
		chromIdxs []int
	}
	tests := []struct {
		name   string
		msdata *MSData
		args   args
		want   []Chromatogram
	}{
		{"small.pwiz", file, args{[]int{0}}, []Chromatogram{pwizSmall_Chromatogram_0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.msdata.Chromatograms(tt.args.chromIdxs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chromatograms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_SourceInfo(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")

	tests := []struct {
		name   string
		msdata *MSData
		want   string
	}{
		{"small.pwiz", file, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.msdata.SourceInfo(); got != tt.want {
				t.Errorf("SourceInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_SampleInfo(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")

	tests := []struct {
		name   string
		msdata *MSData
		want   string
	}{
		{"small.pwiz", file, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.msdata.SampleInfo(); got != tt.want {
				t.Errorf("SampleInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_SoftwareInfo(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")

	tests := []struct {
		name   string
		msdata *MSData
		want   string
	}{
		{"small.pwiz", file, "Xcalibur 1.1 Beta 7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.msdata.SoftwareInfo(); got != tt.want {
				t.Errorf("SoftwareInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Spectra(t *testing.T) {
	type args struct {
		scans []int
	}
	tests := []struct {
		name     string
		fileName string
		args     args
		want     PeakList
	}{
		{"small.pwiz", pwizSmall_FileName, args{[]int{0}}, pwizSmall_PeakList_0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := OpenMSData(tt.fileName)
			if got := file.Spectra(tt.args.scans...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Spectra() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_TIC(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     Chromatogram
	}{
		{"small.pwiz", pwizSmall_FileName, pwizSmall_Chromatogram_0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := OpenMSData(tt.fileName)
			if got := file.TIC(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TIC() = %v, want %v", got, tt.want)
			}
		})
	}
}
