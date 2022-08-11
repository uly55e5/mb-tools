package mzmlReader

import "C"

//#cgo CFLAGS: -I../pwiz-wrapper
//#cgo LDFLAGS: -L../pwiz-wrapper/lib -lpwiz_wrapper -Wl,-rpath=../pwiz-wrapper/lib -lstdc++ -lpwiz_all -lm -ldl
//#include "cpwiz.h"
//#define _GLIBCXX_USE_CXX11_ABI 0
import "C"
import (
	"errors"
	"github.com/uly55e5/readMZmlGo/common"
	"reflect"
	"strings"
	"unsafe"
)

type InstrumentInfo struct {
	Manufacturer string
	Mmodel       string
	Ionisation   string
	Analyzer     string
	Detector     string
	Software     string
	Sample       string
	Source       string
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
	Time      []float64
	Intensity []float64
}

type MSData struct {
	msData         C.MSDataFile
	fileName       string
	instrumentInfo *InstrumentInfo
}

type IsolationWindow struct {
	High float64
	Low  float64
}

type PeakCount struct {
	Count []int
	Scans []int
}

type Map3D [][]float64

type RunInfo struct {
	ScanCount      int
	LowMz          float64
	HighMz         float64
	DStartTime     float64
	DEndTime       float64
	MsLevels       []int
	StartTimeStamp string
}

type PeakList struct {
	Values   [][][]float64
	ColNames []string
	Scans    []int
}

func OpenMSData(fileName string) (*MSData, error) {
	var cErrorMsg *C.char
	cMsData := C.MSDataOpenFile(C.CString(fileName), &cErrorMsg)
	if len(C.GoString(cErrorMsg)) > 0 {
		return nil, errors.New(C.GoString(cErrorMsg))
	} else if cMsData == nil {
		return nil, errors.New("Could not open file.")
	}
	var file MSData
	file.msData = cMsData
	file.fileName = fileName
	return &file, nil
}

func (data *MSData) CloseMSData() {
	C.MSDataClose(data.msData)
}

func (data *MSData) FileName() string {
	return data.fileName
}

func (data *MSData) InstrumentInfo() *InstrumentInfo {
	if data.instrumentInfo == nil {
		cinfo := C.getInstrumentInfo(data.msData)
		info := InstrumentInfo{}
		info.Manufacturer = C.GoString(cinfo.manufacturer)
		info.Mmodel = C.GoString(cinfo.model)
		info.Ionisation = C.GoString(cinfo.ionisation)
		info.Analyzer = C.GoString(cinfo.analyzer)
		info.Detector = C.GoString(cinfo.detector)
		info.Software = C.GoString(cinfo.software)
		info.Sample = C.GoString(cinfo.sample)
		info.Source = C.GoString(cinfo.source)
		data.instrumentInfo = &info
	}
	return data.instrumentInfo
}

func (data *MSData) GetRunInfo() RunInfo {
	header := data.Header()
	var runInfo = RunInfo{}
	runInfo.ScanCount = data.Length()
	runInfo.LowMz = common.Min(header.LowMZ)
	runInfo.HighMz = common.Max(header.HighMZ)
	runInfo.DStartTime = common.Min(header.RetentionTime)
	runInfo.DEndTime = common.Max(header.RetentionTime)
	runInfo.MsLevels = common.Unique(header.MsLevel)
	timeStamp := C.getRunStartTimeStamp(data.msData)
	runInfo.StartTimeStamp = C.GoString(timeStamp)
	return runInfo
}

func (data *MSData) Length() int {
	return int(C.getLastScan(data.msData))

}

func (data *MSData) Analyzer() string {
	if data.instrumentInfo == nil {
		data.InstrumentInfo()
	}
	return data.instrumentInfo.Analyzer
}

func (data *MSData) Detector() string {
	data.InstrumentInfo()
	return data.instrumentInfo.Detector
}

func (data *MSData) Ionisation() string {
	if data.instrumentInfo == nil {
		data.InstrumentInfo()
	}
	return data.instrumentInfo.Ionisation
}

func (data *MSData) Manufacturer() string {
	if data.instrumentInfo == nil {
		data.InstrumentInfo()
	}
	return data.instrumentInfo.Manufacturer
}

func (data *MSData) Model() string {
	if data.instrumentInfo == nil {
		data.InstrumentInfo()
	}
	return data.instrumentInfo.Mmodel
}

func (data *MSData) SampleInfo() string {
	if data.instrumentInfo == nil {
		data.InstrumentInfo()
	}
	return data.instrumentInfo.Sample
}

func (data *MSData) SoftwareInfo() string {
	if data.instrumentInfo == nil {
		data.InstrumentInfo()
	}
	return data.instrumentInfo.Software
}

func (data *MSData) SourceInfo() string {
	if data.instrumentInfo == nil {
		data.InstrumentInfo()
	}
	return data.instrumentInfo.Source
}

func (data *MSData) Header(scans ...int) HeaderInfo {
	if len(scans) == 0 {
		scans = getAllScans(data)
	}
	cScans, length := gSlice2CArrayInt(scans)
	cheaderPtr := C.getScanHeaderInfo(data.msData, cScans, C.int(length))
	cheader := *cheaderPtr
	header := HeaderInfo{}

	errorM := C.GoString(cheader.error)
	if errorM != "" {
		println(errorM)
		return HeaderInfo{}
	}
	convertHeaderData(&header, cheader.names, cheader.values, cheader.numCols, cheader.numRows)
	return header
}

func getAllScans(data *MSData) []int {
	size := data.Length()
	scans := make([]int, size)
	for i := range scans {
		scans[i] = i
	}
	return scans
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

func (data *MSData) Peaks(scans ...int) PeakList {
	cScans, scanLen := gSlice2CArrayInt(scans)
	cPeakListPtr := C.getPeakList(data.msData, cScans, C.int(scanLen))
	cPeakList := *cPeakListPtr
	var peakList PeakList
	peakList.Scans = scans
	names := cArray2GoSliceStr(cPeakList.colNames, int(cPeakList.colNum))
	peakList.ColNames = names
	valSizes := cArray2GoSliceInt(cPeakList.valSizes, int(cPeakList.scanNum))
	var cValsPtr ***C.double = cPeakList.values
	cVals := unsafe.Slice(cValsPtr, int(cPeakList.scanNum))
	peakList.Values = make([][][]float64, int(cPeakList.scanNum))
	for i := 0; i < int(cPeakList.scanNum); i++ {
		peakList.Values[i] = make([][]float64, 2)
		var cScanValsPtr **C.double = cVals[i]
		cScanVals := unsafe.Slice(cScanValsPtr, int(cPeakList.colNum))
		mzs := cArray2GoSliceDouble(cScanVals[0], valSizes[i])
		ints := cArray2GoSliceDouble(cScanVals[1], valSizes[i])
		peakList.Values[i][0] = mzs
		peakList.Values[i][1] = ints
	}
	return peakList
}

func (data *MSData) Spectra(scans ...int) PeakList {
	return data.Peaks(scans...)
}

func (data *MSData) PeaksCount(scans ...int) PeakCount {
	peaks := data.Peaks(scans...)
	peakCount := PeakCount{}
	peakCount.Scans = scans
	for i := range peaks.Values {
		peakCount.Count = append(peakCount.Count, len(peaks.Values[i][0]))
	}
	return peakCount
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
		return common.Unique(iWin)
	}
	return iWin
}

func (data *MSData) ChromatogramCount() int {
	cCount := C.getLastChromatogram(data.msData)
	return int(cCount)
}

func (data *MSData) Chromatograms(chromIdxs ...int) []Chromatogram {
	if len(chromIdxs) == 0 {
		chromIdxs = make([]int, data.ChromatogramCount())
		for i := range chromIdxs {
			chromIdxs[i] = i
		}
	}
	var chroms = []Chromatogram{}
	for _, idx := range chromIdxs {
		chroms = append(chroms, data.Chromatogram(idx))
	}
	return chroms
}

func (data *MSData) TIC() Chromatogram {
	return data.Chromatogram(0)
}

func (data *MSData) Chromatogram(chromIdx int) Chromatogram {
	cInfo := C.getChromatogramInfo(data.msData, C.int(chromIdx))
	var chromatogram = Chromatogram{}
	chromatogram.Intensity = cArray2GoSliceDouble(cInfo.intensity, int(cInfo.size))
	chromatogram.Time = cArray2GoSliceDouble(cInfo.time, int(cInfo.size))
	chromatogram.Id = C.GoString(cInfo.id)
	var errorM string = C.GoString(cInfo.error)
	if errorM != "" {
		println(errorM)
	}
	return chromatogram
}

func (data *MSData) ChromatogramHeader(scans ...int) ChromatogramHeaderInfo {
	cScans, length := gSlice2CArrayInt(scans)
	cheader := C.getChromatogramHeaderInfo(data.msData, cScans, C.int(length))
	chromInfo := ChromatogramHeaderInfo{}
	convertHeaderData(&chromInfo, cheader.names, cheader.values, cheader.numCols, cheader.numRows)
	return chromInfo
}

func (data *MSData) Get3DMap(lowMz float64, highMz float64, resMZ float64, scans ...int) Map3D {
	if len(scans) == 0 {
		scans = getAllScans(data)
	}
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
