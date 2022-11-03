/*
 * Title
 *
 * Title
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mzserver

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/uly55e5/mb-tools/mzmlReader"
	"os"
	"strconv"
	"time"
	"unsafe"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// Get3dMap -
func (s *DefaultApiService) Get3dMap(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	// TODO implement parameters
	map3D := fileInfo.msdata.Get3DMap(0, 5000, 1)
	map3Dmap := getMapFromStruct(map3D)
	return Response(200, map3Dmap), nil

}

// GetChromatogramCount -
func (s *DefaultApiService) GetChromatogramCount(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	count := fileInfo.msdata.ChromatogramCount()
	return Response(200, GetChromatogramCount200Response{int64(count)}), nil
}

// GetChromatogramData -
func (s *DefaultApiService) GetChromatogramData(ctx context.Context, msDataId string, chromatogramId int64) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	chromatogramData := fileInfo.msdata.Chromatogram(int(chromatogramId))
	chromMap := getMapFromStruct(chromatogramData)
	return Response(200, chromMap), nil

}

// GetChromatogramHeader -
func (s *DefaultApiService) GetChromatogramHeader(ctx context.Context, msDataId string, chromatogramId int64) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	header := fileInfo.msdata.ChromatogramHeader(int(chromatogramId))
	headerMap := getMapFromStruct(header)
	return Response(200, headerMap), nil
}

// GetChromatograms -
func (s *DefaultApiService) GetChromatograms(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	chroms := fileInfo.msdata.Chromatograms()
	chromMap := getMapFromStruct(chroms)
	return Response(200, chromMap), nil

}

// GetFileName -
func (s *DefaultApiService) GetFileName(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	return Response(200, GetFileName200Response{fileInfo.fileName}), nil

}

// GetInstrumentAnalyzer -
func (s *DefaultApiService) GetInstrumentAnalyzer(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	analyzer := fileInfo.msdata.Analyzer()
	return Response(200, GetInstrumentAnalyzer200Response{analyzer}), nil
}

// GetInstrumentData -
func (s *DefaultApiService) GetInstrumentData(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	info := fileInfo.msdata.InstrumentInfo()
	infoMap := getMapFromStruct(info)
	return Response(200, infoMap), nil
}

// GetInstrumentDetector -
func (s *DefaultApiService) GetInstrumentDetector(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	detector := fileInfo.msdata.Detector()
	return Response(200, GetInstrumentDetector200Response{detector}), nil
}

// GetInstrumentManufacturer -
func (s *DefaultApiService) GetInstrumentManufacturer(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	manufacturer := fileInfo.msdata.Manufacturer()
	return Response(200, GetInstrumentManufacturer200Response{manufacturer}), nil

}

// GetInstrumentModel -
func (s *DefaultApiService) GetInstrumentModel(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	model := fileInfo.msdata.Model()
	return Response(200, GetInstrumentModel200Response{model}), nil
}

// GetIonisationMethod -
func (s *DefaultApiService) GetIonisationMethod(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	method := fileInfo.msdata.Ionisation()
	return Response(200, GetIonisationMethod200Response{method}), nil
}

// GetIsolationWindows -
func (s *DefaultApiService) GetIsolationWindows(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	windows := fileInfo.msdata.IsolationWindow(true)
	var result = []GetIsolationWindows200ResponseIsolationWindowsInner{}
	for _, w := range windows {
		result = append(result, GetIsolationWindows200ResponseIsolationWindowsInner{w.Low, w.High})
	}
	return Response(200, GetIsolationWindows200Response{result}), nil
}

// GetSampleData -
func (s *DefaultApiService) GetSampleData(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	data := fileInfo.msdata.SampleInfo()
	return Response(200, GetSampleData200Response{data}), nil

}

// GetScanCount -
func (s *DefaultApiService) GetScanCount(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	count := fileInfo.msdata.Length()
	return Response(200, GetScanCount200Response{int64(count)}), nil
}

// GetScanHeader -
func (s *DefaultApiService) GetScanHeader(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	header := fileInfo.msdata.Header()
	headerMap := getMapFromStruct(header)
	return Response(200, headerMap), nil
}

// GetScanPeakCount -
func (s *DefaultApiService) GetScanPeakCount(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	//count := fileInfo.msdata.PeaksCount()
	// TODO implement
	return Response(200, GetScanPeakCount200Response{}), nil

}

// GetScanPeaks -
func (s *DefaultApiService) GetScanPeaks(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	//peaks := fileInfo.msdata.Peaks()
	// TODO: implement
	return Response(200, GetScanPeaks200Response{}), nil

}

// GetScansData -
func (s *DefaultApiService) GetScansData(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	scans := fileInfo.msdata.GetAllScans()
	scans64 := []int64{}
	for _, e := range scans {
		scans64 = append(scans64, int64(e))
	}
	return Response(200, GetScansData200Response{scans64}), nil
}

// GetSoftware -
func (s *DefaultApiService) GetSoftware(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	software := fileInfo.msdata.SoftwareInfo()
	return Response(200, GetSoftware200Response{software}), nil

}

// GetSource -
func (s *DefaultApiService) GetSource(ctx context.Context, msDataId string) (ImplResponse, error) {
	fileInfo, errResponse := getFileData(msDataId)
	if fileInfo == nil {
		return errResponse, nil
	}
	source := fileInfo.msdata.SourceInfo()
	return Response(200, GetSource200Response{source}), nil
}

// PostFile -
func (s *DefaultApiService) PostFile(ctx context.Context, filename string, file *os.File) (ImplResponse, error) {
	if filename != "" && file != nil {
		timestamp := time.Now().Unix()
		id := fmt.Sprintf("%x", sha256.Sum256([]byte(strconv.FormatInt(timestamp, 10)+filename)))[:45]
		msData, err := mzmlReader.OpenMSData(file.Name())
		if err != nil {
			return Response(400, ErrorMsg{"Could not open file"}), nil
		}
		files[id] = fileData{filename, file, msData, timestamp, unsafe.Sizeof(*msData)}
		return Response(200, PostFile200Response{id}), nil
	}
	return Response(400, ErrorMsg{"File or filename not submitted"}), nil
}
