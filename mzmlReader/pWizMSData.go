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
	SeqNum                     []int
	AcquisitionNum             []int
	MsLevel                    []int
	Polarity                   []int
	PeaksCount                 []int
	TotIonCurrent              []float64
	RetentionTime              []float64
	BasePeakMZ                 []float64
	BasePeakIntensity          []float64
	CollisionEnergy            []float64
	IonisationEnergy           []float64
	LowMZ                      []float64
	HighMZ                     []float64
	PrecursorScanNum           []int
	PrecursorMZ                []float64
	PrecursorCharge            []int
	PrecursorIntensity         []float64
	MergedScan                 []int
	MergedResultScanNum        []int
	MergedResultStartScanNum   []int
	MergedResultEndScanNum     []int
	IonInjectionTime           []float64
	FilterString               []string
	SpectrumId                 []string
	Centroided                 []bool
	IonMobilityDriftTime       []float64
	IsolationWindowTargetMZ    []float64
	IsolationWindowLowerOffset []float64
	IsolationWindowUpperOffset []float64
	ScanWindowLowerLimit       []float64
	ScanWindowUpperLimit       []float64
}

type ChromatogramHeaderInfo struct {
	ChromatogramId                      []string
	ChromatogramIndex                   []int
	Polarity                            []int
	PrecursorIsolationWindowTargetMZ    []float64
	PrecursorIsolationWindowLowerOffset []float64
	PrecursorIsolationWindowUpperOffset []float64
	PrecursorCollisionEnergy            []float64
	ProductIsolationWindowTargetMZ      []float64
	ProductIsolationWindowLowerOffset   []float64
	ProductIsolationWindowUpperOffset   []float64
}

type Chromatogram struct {
	Id        string
	time      []float64
	intensity []float64
}

type MSData struct {
	msData         C.MSDataFile
	fileName       string
	instrumentInfo *InstrumentInfo
}

func (data *MSData) Get3DMap(scans int, lowMz float64, highMz float64, resMZ float64) {

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

func (data *MSData) GetHeader(scans []int) *HeaderInfo {
	cScans, length := gSlice2CArrayInt(scans)
	cheader := C.getScanHeaderInfo(data.msData, cScans, C.int(length))
	header := HeaderInfo{}

	errorM := C.GoString(cheader.error)
	if errorM != "" {
		println(errorM)
		return nil
	}
	convertHeaderData(&header, cheader.names, cheader.values, cheader.numCols, cheader.numRows)
	return &header
}

func convertHeaderData(header interface{}, cNames **C.char, cVals *unsafe.Pointer, numCols C.ulong, numRows C.ulong) {
	names := cArray2GoSliceStr(cNames, int(numCols))
	for i, n := range names {
		nameB := []byte(n)
		nameB[0] = strings.ToUpper(string(nameB[0]))[0]
		n = string(nameB)
		val := reflect.ValueOf(header).Elem().FieldByName(string(n))
		cVals := unsafe.Slice((*unsafe.Pointer)(cVals), int(numCols))
		if val.Type() == reflect.TypeOf([]int{}) {
			v := cArray2GoSliceInt((*C.int)(cVals[i]), int(numRows))
			for _, vc := range v {
				reflect.Append(val, reflect.ValueOf(vc))
			}
			val.Set(reflect.ValueOf(v))
		} else if val.Type() == reflect.TypeOf([]float64{}) {
			v := cArray2GoSliceDouble((*C.double)(cVals[i]), int(numRows))
			for _, vc := range v {
				reflect.Append(val, reflect.ValueOf(vc))
			}
			val.Set(reflect.ValueOf(v))
		} else if val.Type() == reflect.TypeOf([]string{}) {
			v := cArray2GoSliceStr((**C.char)(cVals[i]), int(numRows))
			for _, vc := range v {
				reflect.Append(val, reflect.ValueOf(vc))
			}
			val.Set(reflect.ValueOf(v))
		} else if val.Type() == reflect.TypeOf([]bool{}) {
			v := cArray2GoSliceBool((*C.char)(cVals[i]), int(numRows))
			for _, vc := range v {
				reflect.Append(val, reflect.ValueOf(vc))
			}
			val.Set(reflect.ValueOf(v))
		}
	}
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

func cArray2GoSliceDouble(array *C.double, length int) []float64 {
	cSlice := []C.double{}
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&cSlice)))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(array))
	var gSlice = []float64{}
	for _, ci := range cSlice {
		gSlice = append(gSlice, float64(ci))
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

func (data *MSData) tic() Chromatogram {
	return data.chromatogram(0)
}

func (data *MSData) chromatogram(chromIdx int) Chromatogram {
	cInfo := C.getChromatogramInfo(data.msData, C.int(chromIdx))
	var chromatogram = Chromatogram{}
	chromatogram.intensity = cArray2GoSliceDouble(cInfo.intensity, int(cInfo.size))
	chromatogram.time = cArray2GoSliceDouble(cInfo.time, int(cInfo.size))
	chromatogram.Id = C.GoString(cInfo.id)
	var errorM string = C.GoString(cInfo.error)
	if errorM != "" {
		println(errorM)
	}
	return chromatogram
}

func (data *MSData) chromatograms(chromIdxs []int) []Chromatogram {
	var chroms = []Chromatogram{}
	for _, idx := range chromIdxs {
		chroms = append(chroms, data.chromatogram(idx))
	}
	return chroms
}

func (data *MSData) chromatogramHeader(scans []int) ChromatogramHeaderInfo {
	cScans, length := gSlice2CArrayInt(scans)
	cheader := C.getChromatogramHeaderInfo(data.msData, cScans, C.int(length))
	chromInfo := ChromatogramHeaderInfo{}
	convertHeaderData(&chromInfo, cheader.names, cheader.values, cheader.numCols, cheader.numRows)
	return chromInfo

}

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
