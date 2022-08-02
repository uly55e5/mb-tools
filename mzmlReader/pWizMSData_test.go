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
		want InstrumentInfo
	}{
		{"small.pwiz", args{file}, InstrumentInfo{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInstrumentInfo(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstrumentInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
