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

func TestMSData_GetInstrumentInfo(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
		want   *InstrumentInfo
	}{
		{"small.pwiz", &file, &InstrumentInfo{"Xcalibur ", "LTQ FT", "electrospray ionization", "fourier transform ion cyclotron resonance mass spectrometer", "inductive detector", "Xcalibur 1.1 Beta 7", "", ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.GetInstrumentInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstrumentInfo() = %v, want %v", got, tt.want)
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
		{"small.pwiz", &file},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.msdata.CloseMSData()
		})
	}
}

func TestMSData_getFileName(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
		want   string
	}{
		{"small.pwiz", &file, "../data/examples/small.pwiz.1.1.mzML"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.getFileName(); got != tt.want {
				t.Errorf("getFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_GetLength(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
		want   int
	}{
		{"small.pwiz", &file, 48},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.msdata.GetLength(); got != tt.want {
				t.Errorf("GetLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_GetManufacturer(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
		want   string
	}{
		{"small.pwiz", &file, "Xcalibur "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.GetManufacturer(); got != tt.want {
				t.Errorf("GetManufacturer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_GetModel(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
		want   string
	}{
		{"small.pwiz", &file, "LTQ FT"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.GetModel(); got != tt.want {
				t.Errorf("GetModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_GetIonisation(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
		want   string
	}{
		{"small.pwiz", &file, "electrospray ionization"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.GetIonisation(); got != tt.want {
				t.Errorf("GetIonisation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_GetAnalyzer(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")

	tests := []struct {
		name   string
		msdata *MSData
		want   string
	}{
		{"small.pwiz", &file, "fourier transform ion cyclotron resonance mass spectrometer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.GetAnalyzer(); got != tt.want {
				t.Errorf("GetAnalyzer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_GetDetector(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name   string
		msdata *MSData
		want   string
	}{
		{"small.pwiz", &file, "inductive detector"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.GetDetector(); got != tt.want {
				t.Errorf("GetDetector() = %v, want %v", got, tt.want)
			}
		})
	}
}

var nanEqualOpt = cmp.Comparer(func(x, y float64) bool {
	return (math.IsNaN(x) && math.IsNaN(y)) || x == y
})

func TestMSData_GetHeader(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	type args struct {
		scans []int
	}
	tests := []struct {
		name   string
		msdata *MSData
		args   args
		want   *HeaderInfo
	}{
		{"small.pwiz", &file, args{[]int{0}},
			&HeaderInfo{
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

			if got := tt.msdata.GetHeader(tt.args.scans); !cmp.Equal(got, tt.want, nanEqualOpt) {
				t.Errorf("GetHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_chromatogramHeader(t *testing.T) {
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
		{"small.pwiz", &file, args{[]int{0}}, ChromatogramHeaderInfo{
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
			if got := tt.msdata.chromatogramHeader(tt.args.scans); !cmp.Equal(got, tt.want, nanEqualOpt) {
				t.Errorf("chromatogramHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_chromatogram(t *testing.T) {
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
		{"small.pwiz", &file, args{0}, Chromatogram{
			"TIC", []float64{0.004935, 0.007896666666666666, 0.011218333333333334, 0.022838333333333332, 0.034925, 0.04862, 0.06192333333333334, 0.075015, 0.07778833333333333, 0.08120333333333334, 0.09290333333333332, 0.10480333333333333, 0.11721500000000001, 0.13002166666666667, 0.14345166666666667, 0.14640833333333333, 0.149755, 0.16144166666666668, 0.17337, 0.18665833333333332, 0.200695, 0.2136733333333333, 0.21674666666666667, 0.22007333333333332, 0.23292333333333332, 0.244745, 0.2591716666666667, 0.2726633333333333, 0.28548333333333337, 0.2888983333333333, 0.3037033333333333, 0.31565, 0.32852666666666663, 0.342915, 0.35855833333333337, 0.36142833333333335, 0.364755, 0.37657833333333335, 0.3886733333333333, 0.40196166666666666, 0.4151316666666667, 0.4284833333333333, 0.4332216666666666, 0.4365666666666667, 0.44832, 0.46056499999999995, 0.47310333333333326, 0.48723666666666665},
			[]float64{1.5245068e+07, 1.2901166e+07, 586279, 441570.15625, 114331.703125, 130427.3046875, 580561.0625, 1.5148302e+07, 1.0349958e+07, 848427.3125, 456143.4375, 124170.3828125, 104264.796875, 147409.234375, 1.8257344e+07, 1.1037852e+07, 1.102582125e+06, 360250.96875, 125874.828125, 142243.390625, 147414.578125, 1.7613074e+07, 1.5974105e+06, 990298.5, 447647.96875, 71677.03125, 119999.7421875, 152281.25, 2.2136832e+07, 1.243453e+07, 379009.78125, 120473.4296875, 113763.3515625, 73607.4921875, 1.6495375e+07, 6.5487065e+06, 1.04157375e+06, 626711.3125, 109042.7265625, 156294.984375, 79339.078125, 1.2015003e+07, 1.3332331e+07, 925073.25, 419351.46875, 88901.921875, 100616.1953125, 77939.0078125},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msdata.chromatogram(tt.args.chromIdx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chromatogram() = %v, want %v", got, tt.want)
			}
		})
	}
}
