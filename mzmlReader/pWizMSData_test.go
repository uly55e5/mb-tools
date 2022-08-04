package mzmlReader

import (
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

func TestMSData_GetHeader(t *testing.T) {
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	type args struct {
		scans []int
	}
	tests := []struct {
		name   string
		msdata *MSData
		args   args
	}{
		{"small.pwiz", &file, args{[]int{0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.msdata.GetHeader(tt.args.scans)
		})
	}
}
