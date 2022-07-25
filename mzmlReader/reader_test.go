package mzmlReader

import (
	"github.com/uly55e5/readMZmlGo/schema"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    schema.MzML
		wantErr bool
	}{
		{"read a file", args{"data/examples/small.pwiz.1.1.mzML"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}
