package mzmlReader

import "C"

//#cgo CFLAGS: -I../pwiz-wrapper
//#cgo LDFLAGS: -L../pwiz-wrapper/lib -lpwiz_wrapper -Wl,-rpath=../pwiz-wrapper/lib -lstdc++ -lpwiz_all -lm -ldl
//#include "cpwiz.h"
//#define _GLIBCXX_USE_CXX11_ABI 0
import "C"

type InstrumentInfo struct {
	manufacturer string
	model        string
	ionisation   string
	analyzer     string
	detector     string
	software     string
	sample       string
	source       string
}

type MSData struct {
	msData         C.MSDataFile
	fileName       string
	instrumentInfo *InstrumentInfo
}

func OpenMSData(fileName string) MSData {
	var file MSData
	file.fileName = fileName
	file.msData = C.MSDataOpenFile(C.CString(fileName))
	return file
}

func CloseMSData(data MSData) {
	C.MSDataClose(data.msData)
}

func GetInstrumentInfo(data MSData) *InstrumentInfo {
	if data.instrumentInfo == nil {
		cinfo := C.getInstrumentInfo(data.msData)
		info := &InstrumentInfo{}
		info.manufacturer = C.GoString(cinfo.manufacturer)
		info.model = C.GoString(cinfo.model)
		info.ionisation = C.GoString(cinfo.ionisation)
		info.analyzer = C.GoString(cinfo.analyzer)
		info.detector = C.GoString(cinfo.detector)
		info.software = C.GoString(cinfo.software)
		info.sample = C.GoString(cinfo.sample)
		info.source = C.GoString(cinfo.source)
		data.instrumentInfo = info
	}
	return data.instrumentInfo
}

func getLastChrom(data MSData) int {
	return int(C.getLastChrom(data.msData))
}

func getFileName(data MSData) string {
	return data.fileName
}
