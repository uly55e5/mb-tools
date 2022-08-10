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
			_, err := OpenMSData(tt.args.fileName)
			if err != nil {
				t.Errorf("Could not open file: %s", tt.args.fileName)
			}
		})
	}
}

func TestMSData_InstrumentInfo(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     *InstrumentInfo
	}{
		{"small.pwiz", pwizSmall_FileName, &InstrumentInfo{"Xcalibur ", "LTQ FT", "electrospray ionization", "fourier transform ion cyclotron resonance mass spectrometer", "inductive detector", "Xcalibur 1.1 Beta 7", "", ""}},
	}
	for _, tt := range tests {
		file, _ := OpenMSData(tt.fileName)
		t.Run(tt.name, func(t *testing.T) {
			if got := file.InstrumentInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstrumentInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_CloseMSData(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
	}{
		{"small.pwiz", pwizSmall_FileName},
	}
	for _, tt := range tests {
		file, _ := OpenMSData(tt.fileName)
		t.Run(tt.name, func(t *testing.T) {
			file.CloseMSData()
		})
	}
}

func TestMSData_FileName(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{"small.pwiz", pwizSmall_FileName, "../data/examples/small.pwiz.1.1.mzML"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := OpenMSData(tt.fileName)
			if got := file.FileName(); got != tt.want {
				t.Errorf("FileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Length(t *testing.T) {

	tests := []struct {
		name     string
		fileName string
		want     int
	}{
		{"small.pwiz", pwizSmall_FileName, 48},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := OpenMSData(tt.fileName)
			if got := file.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Manufacturer(t *testing.T) {

	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{"small.pwiz", pwizSmall_FileName, "Xcalibur "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := OpenMSData(tt.fileName)
			if got := file.Manufacturer(); got != tt.want {
				t.Errorf("Manufacturer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Model(t *testing.T) {

	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{"small.pwiz", pwizSmall_FileName, "LTQ FT"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := OpenMSData(tt.fileName)
			if got := file.Model(); got != tt.want {
				t.Errorf("Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Ionisation(t *testing.T) {

	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{"small.pwiz", pwizSmall_FileName, "electrospray ionization"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := OpenMSData(tt.fileName)
			if got := file.Ionisation(); got != tt.want {
				t.Errorf("Ionisation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Analyzer(t *testing.T) {

	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{"small.pwiz", pwizSmall_FileName, "fourier transform ion cyclotron resonance mass spectrometer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := OpenMSData(tt.fileName)
			if got := file.Analyzer(); got != tt.want {
				t.Errorf("Analyzer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Detector(t *testing.T) {

	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{"small.pwiz", pwizSmall_FileName, "inductive detector"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := OpenMSData(tt.fileName)
			if got := file.Detector(); got != tt.want {
				t.Errorf("Detector() = %v, want %v", got, tt.want)
			}
		})
	}
}

var nanEqualOpt = cmp.Comparer(func(x, y float64) bool {
	return (math.IsNaN(x) && math.IsNaN(y)) || x == y
})

func TestMSData_Header(t *testing.T) {

	type args struct {
		scans []int
	}
	tests := []struct {
		name     string
		fileName string
		args     args
		want     HeaderInfo
	}{
		{"small.pwiz", pwizSmall_FileName, args{[]int{0}},
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
			file, _ := OpenMSData(tt.fileName)
			if got := file.Header(tt.args.scans...); !cmp.Equal(got, tt.want, nanEqualOpt) {
				t.Errorf("Header() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_ChromatogramHeader(t *testing.T) {
	type args struct {
		scans []int
	}

	tests := []struct {
		name     string
		fileName string
		args     args
		want     ChromatogramHeaderInfo
	}{
		{"small.pwiz", pwizSmall_FileName, args{[]int{0}}, ChromatogramHeaderInfo{
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
			file, _ := OpenMSData(tt.fileName)
			if got := file.ChromatogramHeader(tt.args.scans...); !cmp.Equal(got, tt.want, nanEqualOpt) {
				t.Errorf("ChromatogramHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Chromatogram(t *testing.T) {

	type args struct {
		chromIdx int
	}
	tests := []struct {
		name     string
		fileName string
		args     args
		want     Chromatogram
	}{
		{"small.pwiz", pwizSmall_FileName, args{0}, pwizSmall_Chromatogram_0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := OpenMSData(tt.fileName)
			if got := file.Chromatogram(tt.args.chromIdx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chromatogram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_IsolationWindow(t *testing.T) {

	result := make([]IsolationWindow, 34)
	for i := range result {
		result[i] = IsolationWindow{0.5, 0.5}
	}
	type args struct {
		uniqe bool
	}
	tests := []struct {
		name     string
		fileName string
		args     args
		want     []IsolationWindow
	}{
		{"small.pwiz", pwizSmall_FileName, args{false}, result},
		{"small.pwiz.Unique", pwizSmall_FileName, args{true}, []IsolationWindow{{0.5, 0.5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := OpenMSData(tt.fileName)
			if got := file.IsolationWindow(tt.args.uniqe); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsolationWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Peaks(t *testing.T) {

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
			file, _ := OpenMSData(tt.fileName)
			if got := file.Peaks(tt.args.scans...); !cmp.Equal(got, tt.want) {
				t.Errorf("Peaks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_PeaksCount(t *testing.T) {

	type args struct {
		scans []int
	}
	tests := []struct {
		name     string
		fileName string
		args     args
		want     PeakCount
	}{
		{"small.pwiz", pwizSmall_FileName, args{[]int{0}}, PeakCount{[]int{19914}, []int{0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := OpenMSData(tt.fileName)
			if got := file.PeaksCount(tt.args.scans...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PeakCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Get3DMap(t *testing.T) {

	type args struct {
		scans  []int
		lowMz  float64
		highMz float64
		resMZ  float64
	}
	tests := []struct {
		name     string
		fileName string
		args     args
		want     Map3D
	}{
		{"small.pwiz", pwizSmall_FileName, args{[]int{0, 1, 2, 3}, 0, 2000, 0.5}, pwizSmall_3DMap_0_3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := OpenMSData(tt.fileName)
			if got := file.Get3DMap(tt.args.lowMz, tt.args.highMz, tt.args.resMZ, tt.args.scans...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get3DMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_GetRunInfo(t *testing.T) {

	tests := []struct {
		name     string
		fileName string
		want     RunInfo
	}{
		{"small.pwiz", pwizSmall_FileName, RunInfo{
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
			file, _ := OpenMSData(tt.fileName)
			if got := file.GetRunInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRunInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_Chromatograms(t *testing.T) {

	type args struct {
		chromIdxs []int
	}
	tests := []struct {
		name     string
		fileName string
		args     args
		want     []Chromatogram
	}{
		{"small.pwiz", pwizSmall_FileName, args{[]int{0}}, []Chromatogram{pwizSmall_Chromatogram_0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := OpenMSData(tt.fileName)
			if got := file.Chromatograms(tt.args.chromIdxs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chromatograms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_SourceInfo(t *testing.T) {

	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{"small.pwiz", pwizSmall_FileName, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := OpenMSData(tt.fileName)
			if got := file.SourceInfo(); got != tt.want {
				t.Errorf("SourceInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_SampleInfo(t *testing.T) {

	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{"small.pwiz", pwizSmall_FileName, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := OpenMSData(tt.fileName)
			if got := file.SampleInfo(); got != tt.want {
				t.Errorf("SampleInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMSData_SoftwareInfo(t *testing.T) {

	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{"small.pwiz", pwizSmall_FileName, "Xcalibur 1.1 Beta 7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := OpenMSData(tt.fileName)
			if got := file.SoftwareInfo(); got != tt.want {
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
			file, _ := OpenMSData(tt.fileName)
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
			file, _ := OpenMSData(tt.fileName)
			if got := file.TIC(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TIC() = %v, want %v", got, tt.want)
			}
		})
	}
}
