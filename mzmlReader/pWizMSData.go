package mzmlReader

import "C"

//#cgo CFLAGS: -I../pwiz-wrapper
//#cgo LDFLAGS: -L../pwiz-wrapper/lib -lpwiz_wrapper -Wl,-rpath=../pwiz-wrapper/lib -lstdc++ -lpwiz_all -lm -ldl
//#include "cpwiz.h"
//#define _GLIBCXX_USE_CXX11_ABI 0
import "C"
import (
	"golang.org/x/exp/constraints"
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

type IsolationWindow struct {
	high float64
	low  float64
}

type PeakCount struct {
	count []int
	scans []int
}

type Map3D [][]float64

type RunInfo struct {
	scanCount      int
	lowMz          float64
	highMz         float64
	dStartTime     float64
	dEndTime       float64
	msLevels       []int
	startTimeStamp string
}

type PeakList struct {
	Values   [][][]float64
	ColNames []string
	Scans    []int
}

func (data *MSData) Get3DMap(scans []int, lowMz float64, highMz float64, resMZ float64) Map3D {
	cScans, length := gSlice2CArrayInt(scans)
	cMap3d := C.get3DMap(data.msData, cScans, C.int(length), C.double(lowMz), C.double(highMz), C.double(resMZ))
	c3dSlice := []*C.double{}
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&c3dSlice))
	sliceHeader.Cap = int(cMap3d.scanSize)
	sliceHeader.Len = int(cMap3d.scanSize)
	sliceHeader.Data = uintptr(unsafe.Pointer(cMap3d.values))
	var map3D = make(Map3D, int(cMap3d.scanSize))
	for i := range c3dSlice {
		map3D[i] = cArray2GoSliceDouble(c3dSlice[i], int(cMap3d.valueSize))
	}
	return map3D
}

func (data *MSData) WriteMSFile(fileName string, format string) {

}

func (data *MSData) Length() int {
	return int(C.getLastScan(data.msData))

}

func (data *MSData) Manufacturer() string {
	if data.instrumentInfo == nil {
		data.InstrumentInfo()
	}
	return data.instrumentInfo.manufacturer
}

func (data *MSData) Model() string {
	if data.instrumentInfo == nil {
		data.InstrumentInfo()
	}
	return data.instrumentInfo.model
}

func (data *MSData) Ionisation() string {
	if data.instrumentInfo == nil {
		data.InstrumentInfo()
	}
	return data.instrumentInfo.ionisation
}

func (data *MSData) Analyzer() string {
	if data.instrumentInfo == nil {
		data.InstrumentInfo()
	}
	return data.instrumentInfo.analyzer
}

func (data *MSData) Detector() string {
	data.InstrumentInfo()
	return data.instrumentInfo.detector
}

func (data *MSData) Header() HeaderInfo {
	size := data.Length()
	allScans := make([]int, size)
	for i := range allScans {
		allScans[i] = i
	}
	return data.HeaderForScans(allScans)
}

func (data *MSData) HeaderForScans(scans []int) HeaderInfo {
	cScans, length := gSlice2CArrayInt(scans)
	cheader := C.getScanHeaderInfo(data.msData, cScans, C.int(length))
	header := HeaderInfo{}

	errorM := C.GoString(cheader.error)
	if errorM != "" {
		println(errorM)
		return HeaderInfo{}
	}
	convertHeaderData(&header, cheader.names, cheader.values, cheader.numCols, cheader.numRows)
	return header
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
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&cSlice))
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
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&cSlice))
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
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&cSlice))
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

func (data *MSData) Peaks(scans []int) PeakList {
	cScans, scanLen := gSlice2CArrayInt(scans)
	cPeakList := C.getPeakList(data.msData, cScans, C.int(scanLen))
	var peakList PeakList
	peakList.Scans = scans
	names := cArray2GoSliceStr(cPeakList.colnames, int(cPeakList.colNum))
	peakList.ColNames = names
	valSizes := cArray2GoSliceInt(cPeakList.valSizes, int(cPeakList.scanNum))
	cVals := []**C.double{}
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&cVals))
	sliceHeader.Cap = int(cPeakList.scanNum)
	sliceHeader.Len = int(cPeakList.scanNum)
	sliceHeader.Data = uintptr(unsafe.Pointer(cPeakList.values))
	peakList.Values = make([][][]float64, int(cPeakList.scanNum))
	for i := 0; i < int(cPeakList.scanNum); i++ {
		peakList.Values[i] = make([][]float64, 2)
		cScanVals := []*C.double{}
		cScanValsH := (*reflect.SliceHeader)(unsafe.Pointer(&cScanVals))
		cScanValsH.Cap = int(cPeakList.colNum)
		cScanValsH.Len = int(cPeakList.colNum)
		cScanValsH.Data = uintptr(unsafe.Pointer(cVals[i]))
		mzs := cArray2GoSliceDouble(cScanVals[0], valSizes[i])
		ints := cArray2GoSliceDouble(cScanVals[1], valSizes[i])
		peakList.Values[i][0] = mzs
		peakList.Values[i][1] = ints
	}
	return peakList
}

func (data *MSData) Spectra(scans []int) PeakList {
	return data.Peaks(scans)
}

func (data *MSData) PeaksCount(scans []int) PeakCount {
	peaks := data.Peaks(scans)
	peakCount := PeakCount{}
	peakCount.scans = scans
	for i := range peaks.Values {
		peakCount.count = append(peakCount.count, len(peaks.Values[i][0]))
	}
	return peakCount
}

func max[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m < v {
			m = v
		}
	}
	return m
}

func min[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m > v {
			m = v
		}
	}
	return m
}

func unique[T comparable](array []T) []T {
	var uniqeVals []T
	m := map[T]bool{}
	for _, v := range array {
		if !m[v] {
			m[v] = true
			uniqeVals = append(uniqeVals, v)
		}
	}
	return uniqeVals
}

func (data *MSData) GetRunInfo() RunInfo {
	header := data.Header()
	var runInfo = RunInfo{}
	runInfo.scanCount = data.Length()
	runInfo.lowMz = min(header.LowMZ)
	runInfo.highMz = max(header.HighMZ)
	runInfo.dStartTime = min(header.RetentionTime)
	runInfo.dEndTime = max(header.RetentionTime)
	runInfo.msLevels = unique(header.MsLevel)
	timeStamp := C.getRunStartTimeStamp(data.msData)
	runInfo.startTimeStamp = C.GoString(timeStamp)
	return runInfo
}

func (data *MSData) SoftwareInfo() string {
	if data.instrumentInfo == nil {
		data.InstrumentInfo()
	}
	return data.instrumentInfo.software
}

func (data *MSData) SampleInfo() string {
	if data.instrumentInfo == nil {
		data.InstrumentInfo()
	}
	return data.instrumentInfo.sample
}

func (data *MSData) SourceInfo() string {
	if data.instrumentInfo == nil {
		data.InstrumentInfo()
	}
	return data.instrumentInfo.source
}

func (data *MSData) IsolationWindow(uniqueVals bool) []IsolationWindow {
	cWin := C.getIsolationWindow(data.msData)
	iWin := []IsolationWindow{}
	high := cArray2GoSliceDouble(cWin.high, int(cWin.size))
	low := cArray2GoSliceDouble(cWin.low, int(cWin.size))
	for i := range high {
		iWin = append(iWin, IsolationWindow{high[i], low[i]})
	}

	if uniqueVals {
		return unique(iWin)
	}
	return iWin
}

func (data *MSData) TIC() Chromatogram {
	return data.Chromatogram(0)
}

func (data *MSData) Chromatogram(chromIdx int) Chromatogram {
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

func (data *MSData) Chromatograms(chromIdxs []int) []Chromatogram {
	var chroms = []Chromatogram{}
	for _, idx := range chromIdxs {
		chroms = append(chroms, data.Chromatogram(idx))
	}
	return chroms
}

func (data *MSData) ChromatogramHeader(scans []int) ChromatogramHeaderInfo {
	cScans, length := gSlice2CArrayInt(scans)
	cheader := C.getChromatogramHeaderInfo(data.msData, cScans, C.int(length))
	chromInfo := ChromatogramHeaderInfo{}
	convertHeaderData(&chromInfo, cheader.names, cheader.values, cheader.numCols, cheader.numRows)
	return chromInfo
}

func OpenMSData(fileName string) *MSData {
	var file MSData
	file.fileName = fileName
	file.msData = C.MSDataOpenFile(C.CString(fileName))
	return &file
}

func (data *MSData) CloseMSData() {
	C.MSDataClose(data.msData)
}

func (data *MSData) InstrumentInfo() *InstrumentInfo {
	if data.instrumentInfo == nil {
		cinfo := C.getInstrumentInfo(data.msData)
		info := InstrumentInfo{}
		info.manufacturer = C.GoString(cinfo.manufacturer)
		info.model = C.GoString(cinfo.model)
		info.ionisation = C.GoString(cinfo.ionisation)
		info.analyzer = C.GoString(cinfo.analyzer)
		info.detector = C.GoString(cinfo.detector)
		info.software = C.GoString(cinfo.software)
		info.sample = C.GoString(cinfo.sample)
		info.source = C.GoString(cinfo.source)
		data.instrumentInfo = &info
	}
	return data.instrumentInfo
}

func (data *MSData) FileName() string {
	return data.fileName
}
