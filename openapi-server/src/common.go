package mzserver

import (
	"encoding/json"
	"github.com/uly55e5/mb-tools/mzmlReader"
	"os"
)

type fileData struct {
	fileName  string
	file      *os.File
	msdata    *mzmlReader.MSData
	timestamp int64
	size      uintptr
}

var files = map[string]fileData{}

func getMapFromStruct(structData interface{}) map[string]interface{} {
	var dataMap map[string]interface{}
	jsonData, _ := json.Marshal(structData)
	json.Unmarshal(jsonData, &dataMap)
	return dataMap
}

func getFileData(msDataId string) (*fileData, ImplResponse) {

	if msDataId == "" {
		return nil, Response(400, ErrorMsg{"MS Data ID is not set"})
	}
	var fileInfo fileData
	var ok bool
	if fileInfo, ok = files[msDataId]; ok {
		var err error
		if fileInfo.msdata == nil && fileInfo.file != nil {
			fileInfo.msdata, err = mzmlReader.OpenMSData(fileInfo.file.Name())
			if err != nil {
				return nil, Response(400, ErrorMsg{"Could not open file."})
			}
		}

	} else {
		return nil, Response(404, ErrorMsg{"Dataset not found"})
	}
	return &fileInfo, Response(500, "This should never be returned")
}
