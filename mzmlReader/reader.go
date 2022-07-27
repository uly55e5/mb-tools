package mzmlReader

import "C"
import (
	"encoding/xml"
	"github.com/uly55e5/readMZmlGo/schema"
	"os"
)

// # include "cpwiz.h"

func Read(filename string) (schema.MzML, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var mzML = schema.MzMLType{}
	err = xml.Unmarshal(file, &mzML)
	if err != nil {
		return nil, err
	}
	return &mzML, nil
}
