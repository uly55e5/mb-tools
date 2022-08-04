package mzmlReader

import "C"

//#cgo CFLAGS: -I../pwiz-wrapper
//#cgo LDFLAGS: -L../pwiz-wrapper/lib -lpwiz_wrapper -Wl,-rpath=../pwiz-wrapper/lib -lstdc++ -lpwiz_all -lm -ldl
//#include "cpwiz.h"
//#define _GLIBCXX_USE_CXX11_ABI 0
import "C"
import (
	"reflect"
	"strings"
	"unsafe"
)

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

type HeaderInfo struct {
	SeqNum                     *[]int
	AcquisitionNum             *[]int
	MsLevel                    *[]int
	Polarity                   *[]int
	PeaksCount                 *[]int
	TotIonCurrent              *[]float32
	RetentionTime              *[]float32
	BasePeakMZ                 *[]float32
	BasePeakIntensity          *[]float32
	CollisionEnergy            *[]float32
	IonisationEnergy           *[]float32
	LowMZ                      *[]float32
	HighMZ                     *[]float32
	PrecursorScanNum           *[]int
	PrecursorMZ                *[]float32
	PrecursorCharge            *[]int
	PrecursorIntensity         *[]float32
	MergedScan                 *[]int
	MergedResultScanNum        *[]int
	MergedResultStartScanNum   *[]int
	MergedResultEndScanNum     *[]int
	IonInjectionTime           *[]float32
	FilterString               *[]string
	SpectrumId                 *[]string
	Centroided                 *[]bool
	IonMobilityDriftTime       *[]float32
	IsolationWindowTargetMZ    *[]float32
	IsolationWindowLowerOffset *[]float32
	IsolationWindowUpperOffset *[]float32
	ScanWindowLowerLimit       *[]float32
	ScanWindowUpperLimit       *[]float32
}

type MSData struct {
	msData         C.MSDataFile
	fileName       string
	instrumentInfo *InstrumentInfo
}

func (data *MSData) Get3DMap(scans int, lowMz float32, highMz float32, resMZ float32) {

}

func (data *MSData) WriteMSFile(fileName string, format string) {

}

func (data *MSData) GetLength() int {
	return int(C.getLastScan(data.msData))

}

func (data *MSData) GetManufacturer() string {
	if data.instrumentInfo == nil {
		data.GetInstrumentInfo()
	}
	return data.instrumentInfo.manufacturer
}

func (data *MSData) GetModel() string {
	if data.instrumentInfo == nil {
		data.GetInstrumentInfo()
	}
	return data.instrumentInfo.model
}

func (data *MSData) GetIonisation() string {
	if data.instrumentInfo == nil {
		data.GetInstrumentInfo()
	}
	return data.instrumentInfo.ionisation
}

func (data *MSData) GetAnalyzer() string {
	if data.instrumentInfo == nil {
		data.GetInstrumentInfo()
	}
	return data.instrumentInfo.analyzer
}

func (data *MSData) GetDetector() string {
	data.GetInstrumentInfo()
	return data.instrumentInfo.detector
}

func (data *MSData) GetHeader(scans []int) {
	cScans, length := gSlice2CArrayInt(scans)
	cheader := C.getScanHeaderInfo(data.msData, cScans, C.int(length))
	println(cheader.numCols)
	errorM := C.GoString(cheader.error)
	println(errorM)
	names := cArray2GoSliceStr(cheader.names, int(cheader.numCols))
	println(names)
	header := HeaderInfo{}
	for i, n := range names {
		nameB := []byte(n)
		nameB[0] = strings.ToUpper(string(nameB[0]))[0]
		n = string(nameB)
		val := reflect.ValueOf(&header).Elem().FieldByName(string(n))
		cVals := unsafe.Slice((*unsafe.Pointer)(cheader.values), int(cheader.numCols))
		if val.Type() == reflect.TypeOf(&[]int{}) {
			v := cArray2GoSliceInt((*C.int)(cVals[i]), int(cheader.numRows))
			val.Set(reflect.ValueOf(&v))
		} else if val.Type() == reflect.TypeOf(&[]float32{}) {
			v := cArray2GoSliceDouble((*C.double)(cVals[i]), int(cheader.numRows))
			val.Set(reflect.ValueOf(&v))
		} else if val.Type() == reflect.TypeOf(&[]string{}) {
			v := cArray2GoSliceStr((**C.char)(cVals[i]), int(cheader.numRows))
			val.Set(reflect.ValueOf(&v))
		} else if val.Type() == reflect.TypeOf(&[]bool{}) {
			v := cArray2GoSliceBool((*C.char)(cVals[i]), int(cheader.numRows))
			val.Set(reflect.ValueOf(&v))
		}
	}
	println(&header)
}

func cArray2GoSliceInt(array *C.int, length int) []int {
	cSlice := []C.int{}
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&cSlice)))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(array))
	var gSlice = []int{}
	for _, ci := range cSlice {
		gSlice = append(gSlice, int(ci))
	}
	return gSlice
}

func cArray2GoSliceBool(array *C.char, length int) []bool {
	cSlice := []C.char{}
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&cSlice)))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(array))
	var gSlice = []bool{}
	for _, ci := range cSlice {
		gSlice = append(gSlice, uint8(ci) > 0)
	}
	return gSlice
}

func cArray2GoSliceDouble(array *C.double, length int) []float32 {
	cSlice := []C.double{}
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&cSlice)))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(array))
	var gSlice = []float32{}
	for _, ci := range cSlice {
		gSlice = append(gSlice, float32(ci))
	}
	return gSlice
}

func cArray2GoSliceStr(array **C.char, length int) []string {
	cSlice := unsafe.Slice(array, length)
	var gSlice = []string{}
	for _, ci := range cSlice {
		gSlice = append(gSlice, C.GoString(ci))
	}
	return gSlice
}

func gSlice2CArrayInt(gSlice []int) (*C.int, int) {
	var cSlice = []C.int{}
	for _, gi := range gSlice {
		cSlice = append(cSlice, C.int(gi))
	}
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&cSlice))
	return (*C.int)(unsafe.Pointer(sliceHeader.Data)), len(gSlice)
}

func (data *MSData) GetPeaks(scans int) {

}

func (data *MSData) GetSpectra(scans int) {

}

func (data *MSData) GetPeaksCount(scans int) {

}

func (data *MSData) GetRunInfo() {

}

func (data *MSData) GetSoftwareInfo() {

}

func (data *MSData) GetSampleInfo() {

}

func (data *MSData) GetSourceInfo() {

}

func (data *MSData) String() string {
	return ""
}

func (data *MSData) isolationWindow() {

}

func (data *MSData) tic() {
}

func (data *MSData) chromatograms() {}

func (data *MSData) chromatogramHeader() {}

func OpenMSData(fileName string) MSData {
	var file MSData
	file.fileName = fileName
	file.msData = C.MSDataOpenFile(C.CString(fileName))
	return file
}

func (data *MSData) CloseMSData() {
	C.MSDataClose(data.msData)
}

func (data *MSData) GetInstrumentInfo() *InstrumentInfo {
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
	return data.instrumentInfo
}

func (data *MSData) getFileName() string {
	return data.fileName
}
