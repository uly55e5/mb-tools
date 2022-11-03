package mzmlReader

import "C"

//#cgo CFLAGS: -I../pwiz-wrapper
//#cgo LDFLAGS: -L../pwiz-wrapper/lib -lpwiz_wrapper -Wl,-rpath=../pwiz-wrapper/lib -lstdc++ -lpwiz_all -lm -ldl
//#include "cpwiz.h"
//#define _GLIBCXX_USE_CXX11_ABI 0
import "C"
import (
	"errors"
	"github.com/uly55e5/mb-tools/common"
	"unsafe"
)

type InstrumentInfo struct {
	Manufacturer string
	Model        string
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
		return nil, errors.New("could not open file")
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
		info.Model = C.GoString(cinfo.model)
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
	return data.instrumentInfo.Model
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
		scans = data.GetAllScans()
	}
	cScans, length := gSlice2CArrayInt(scans)
	cheaderPtr := C.getScanHeaderInfo(data.msData, cScans, C.int(length))
	cheader := *cheaderPtr
	var cseqNumPtr *C.int = cheader.values.seqNum
	var cacquisitionNumPtr *C.int = cheader.values.acquisitionNum
	var cmsLevelPtr *C.int = cheader.values.msLevel
	var cpolarityPtr *C.int = cheader.values.polarity
	var cpeaksCountPtr *C.int = cheader.values.peaksCount
	var ctotIonCurrentPtr *C.double = cheader.values.totIonCurrent
	var cretentionTimePtr *C.double = cheader.values.retentionTime
	var cbasePeakMZPtr *C.double = cheader.values.basePeakMZ
	var cbasePeakIntensityPtr *C.double = cheader.values.basePeakIntensity
	var ccollisionEnergyPtr *C.double = cheader.values.collisionEnergy
	var cionisationEnergyPtr *C.double = cheader.values.ionisationEnergy
	var clowMZPtr *C.double = cheader.values.lowMZ
	var chighMZPtr *C.double = cheader.values.highMZ
	var cprecursorScanNumPtr *C.int = cheader.values.precursorScanNum
	var cprecursorMZPtr *C.double = cheader.values.precursorMZ
	var cprecursorChargePtr *C.int = cheader.values.precursorCharge
	var cprecursorIntensityPtr *C.double = cheader.values.precursorIntensity
	var cmergedScanPtr *C.int = cheader.values.mergedScan
	var cmergedResultScanNumPtr *C.int = cheader.values.mergedResultScanNum
	var cmergedResultStartScanNumPtr *C.int = cheader.values.mergedResultStartScanNum
	var cmergedResultEndScanNumPtr *C.int = cheader.values.mergedResultEndScanNum
	var cionInjectionTimePtr *C.double = cheader.values.ionInjectionTime
	var cfilterStringPtr **C.char = cheader.values.filterString
	var cspectrumIdPtr **C.char = cheader.values.spectrumId
	var ccentroidedPtr *C.char = cheader.values.centroided
	var cionMobilityDriftTimePtr *C.double = cheader.values.ionMobilityDriftTime
	var cisolationWindowTargetMZPtr *C.double = cheader.values.isolationWindowTargetMZ
	var cisolationWindowLowerOffsetPtr *C.double = cheader.values.isolationWindowLowerOffset
	var cisolationWindowUpperOffsetPtr *C.double = cheader.values.isolationWindowUpperOffset
	var cscanWindowLowerLimitPtr *C.double = cheader.values.scanWindowLowerLimit
	var cscanWindowUpperLimitPtr *C.double = cheader.values.scanWindowUpperLimit
	size := int(cheader.size)
	header := HeaderInfo{}
	header.SeqNum = cArray2GoSliceInt(cseqNumPtr, size)
	header.AcquisitionNum = cArray2GoSliceInt(cacquisitionNumPtr, size)
	header.MsLevel = cArray2GoSliceInt(cmsLevelPtr, size)
	header.Polarity = cArray2GoSliceInt(cpolarityPtr, size)
	header.PeaksCount = cArray2GoSliceInt(cpeaksCountPtr, size)
	header.TotIonCurrent = cArray2GoSliceDouble(ctotIonCurrentPtr, size)
	header.RetentionTime = cArray2GoSliceDouble(cretentionTimePtr, size)
	header.BasePeakMZ = cArray2GoSliceDouble(cbasePeakMZPtr, size)
	header.BasePeakIntensity = cArray2GoSliceDouble(cbasePeakIntensityPtr, size)
	header.CollisionEnergy = cArray2GoSliceDouble(ccollisionEnergyPtr, size)
	header.IonisationEnergy = cArray2GoSliceDouble(cionisationEnergyPtr, size)
	header.LowMZ = cArray2GoSliceDouble(clowMZPtr, size)
	header.HighMZ = cArray2GoSliceDouble(chighMZPtr, size)
	header.PrecursorScanNum = cArray2GoSliceInt(cprecursorScanNumPtr, size)
	header.PrecursorMZ = cArray2GoSliceDouble(cprecursorMZPtr, size)
	header.PrecursorCharge = cArray2GoSliceInt(cprecursorChargePtr, size)
	header.PrecursorIntensity = cArray2GoSliceDouble(cprecursorIntensityPtr, size)
	header.MergedScan = cArray2GoSliceInt(cmergedScanPtr, size)
	header.MergedResultScanNum = cArray2GoSliceInt(cmergedResultScanNumPtr, size)
	header.MergedResultStartScanNum = cArray2GoSliceInt(cmergedResultStartScanNumPtr, size)
	header.MergedResultEndScanNum = cArray2GoSliceInt(cmergedResultEndScanNumPtr, size)
	header.IonInjectionTime = cArray2GoSliceDouble(cionInjectionTimePtr, size)
	header.FilterString = cArray2GoSliceStr(cfilterStringPtr, size)
	header.SpectrumId = cArray2GoSliceStr(cspectrumIdPtr, size)
	header.Centroided = cArray2GoSliceBool(ccentroidedPtr, size)
	header.IonMobilityDriftTime = cArray2GoSliceDouble(cionMobilityDriftTimePtr, size)
	header.IsolationWindowTargetMZ = cArray2GoSliceDouble(cisolationWindowTargetMZPtr, size)
	header.IsolationWindowLowerOffset = cArray2GoSliceDouble(cisolationWindowLowerOffsetPtr, size)
	header.IsolationWindowUpperOffset = cArray2GoSliceDouble(cisolationWindowUpperOffsetPtr, size)
	header.ScanWindowLowerLimit = cArray2GoSliceDouble(cscanWindowLowerLimitPtr, size)
	header.ScanWindowUpperLimit = cArray2GoSliceDouble(cscanWindowUpperLimitPtr, size)
	/*errorM := C.GoString(cheader.error)
	if errorM != "" {
		println(errorM)
		return HeaderInfo{}
	}*/
	C.deleteScanHeader(cheaderPtr)
	return header
}

func (data *MSData) GetAllScans() []int {
	size := data.Length()
	scans := make([]int, size)
	for i := range scans {
		scans[i] = i
	}
	return scans
}

func (data *MSData) Peaks(scans ...int) PeakList {
	cScans, scanLen := gSlice2CArrayInt(scans)
	cPeakListPtr := C.getPeakList(data.msData, cScans, C.int(scanLen))
	cPeakList := *cPeakListPtr
	var peakList PeakList
	peakList.Scans = scans
	names := cArray2GoSliceStr(cPeakList.colNames, int(cPeakList.colNum))
	peakList.ColNames = names
	valSizes := cArray2GoSliceULongInt(cPeakList.valSizes, int(cPeakList.scanNum))
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
	C.deletePeakList(cPeakListPtr)
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
	var iWin []IsolationWindow
	high := cArray2GoSliceDouble(cWin.high, int(cWin.size))
	low := cArray2GoSliceDouble(cWin.low, int(cWin.size))
	for i := range high {
		iWin = append(iWin, IsolationWindow{high[i], low[i]})
	}

	if uniqueVals {
		return common.Unique(iWin)
	}
	C.deleteIsolationWindow(cWin)
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
	var chroms []Chromatogram
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
	C.deleteChromatogramInfo(cInfo)
	return chromatogram
}

func (data *MSData) ChromatogramHeader(scans ...int) ChromatogramHeaderInfo {
	cScans, length := gSlice2CArrayInt(scans)
	cheader := C.getChromatogramHeaderInfo(data.msData, cScans, C.int(length))
	chromInfo := ChromatogramHeaderInfo{}
	size := int(cheader.size)
	var cChromIdPtr **C.char = cheader.values.chromatogramId
	var cChromIdxPtr *C.int = cheader.values.chromatogramIndex
	var cPolarityPtr *C.int = cheader.values.polarity
	var cPrecIWMZPtr *C.double = cheader.values.precursorIsolationWindowTargetMZ
	var cPrecIWLowPtr *C.double = cheader.values.precursorIsolationWindowLowerOffset
	var cPrecIWHighPtr *C.double = cheader.values.precursorIsolationWindowUpperOffset
	var cPrecCollEPtr *C.double = cheader.values.precursorCollisionEnergy
	var cProdIWMZ *C.double = cheader.values.productIsolationWindowTargetMZ
	var cProdIWLowPtr *C.double = cheader.values.productIsolationWindowLowerOffset
	var cProdIWHighPtr *C.double = cheader.values.productIsolationWindowUpperOffset
	chromInfo.ChromatogramId = cArray2GoSliceStr(cChromIdPtr, size)
	chromInfo.ChromatogramIndex = cArray2GoSliceInt(cChromIdxPtr, size)
	chromInfo.Polarity = cArray2GoSliceInt(cPolarityPtr, size)
	chromInfo.PrecursorIsolationWindowTargetMZ = cArray2GoSliceDouble(cPrecIWMZPtr, size)
	chromInfo.PrecursorIsolationWindowLowerOffset = cArray2GoSliceDouble(cPrecIWLowPtr, size)
	chromInfo.PrecursorIsolationWindowUpperOffset = cArray2GoSliceDouble(cPrecIWHighPtr, size)
	chromInfo.PrecursorCollisionEnergy = cArray2GoSliceDouble(cPrecCollEPtr, size)
	chromInfo.ProductIsolationWindowTargetMZ = cArray2GoSliceDouble(cProdIWMZ, size)
	chromInfo.ProductIsolationWindowLowerOffset = cArray2GoSliceDouble(cProdIWLowPtr, size)
	chromInfo.ProductIsolationWindowUpperOffset = cArray2GoSliceDouble(cProdIWHighPtr, size)
	C.deleteChromatogramHeader(cheader)
	return chromInfo
}

func (data *MSData) Get3DMap(lowMz float64, highMz float64, resMZ float64, scans ...int) Map3D {
	if len(scans) == 0 {
		scans = data.GetAllScans()
	}
	cScans, length := gSlice2CArrayInt(scans)
	cMap3d := C.get3DMap(data.msData, cScans, C.int(length), C.double(lowMz), C.double(highMz), C.double(resMZ))
	var cValuesPtr **C.double = cMap3d.values
	c3dSlice := unsafe.Slice(cValuesPtr, int(cMap3d.scanSize))
	var map3D = make(Map3D, int(cMap3d.scanSize))
	for i := range c3dSlice {
		map3D[i] = cArray2GoSliceDouble(c3dSlice[i], int(cMap3d.valueSize))
	}
	C.delete3DMap(cMap3d)
	return map3D
}
