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

func TestGetInstrumentInfo(t *testing.T) {
	type args struct {
		data MSData
	}
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name string
		args args
		want *InstrumentInfo
	}{
		{"small.pwiz", args{file}, &InstrumentInfo{"Xcalibur ", "LTQ FT", "electrospray ionization", "fourier transform ion cyclotron resonance mass spectrometer", "inductive detector", "Xcalibur 1.1 Beta 7", "", ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInstrumentInfo(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstrumentInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloseMSData(t *testing.T) {
	type args struct {
		data MSData
	}
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name string
		args args
	}{
		{"small.pwiz", args{file}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CloseMSData(tt.args.data)
		})
	}
}

func Test_getLastChrom(t *testing.T) {
	type args struct {
		data MSData
	}
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name string
		args args
		want int
	}{
		{"small.pwiz", args{file}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastChrom(tt.args.data); got != tt.want {
				t.Errorf("getLastChrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFileName(t *testing.T) {
	type args struct {
		data MSData
	}
	file := OpenMSData("../data/examples/small.pwiz.1.1.mzML")
	tests := []struct {
		name string
		args args
		want string
	}{
		{"small.pwiz", args{file}, "../data/examples/small.pwiz.1.1.mzML"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFileName(tt.args.data); got != tt.want {
				t.Errorf("getFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
