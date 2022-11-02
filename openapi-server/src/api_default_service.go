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
	"net/http"
	"errors"
	"os"
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
	// TODO - update Get3dMap with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(200, map[string]interface{}{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("Get3dMap method not implemented")
}

// GetChromatogramCount - 
func (s *DefaultApiService) GetChromatogramCount(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetChromatogramCount with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetChromatogramCount200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetChromatogramCount200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetChromatogramCount method not implemented")
}

// GetChromatogramData - 
func (s *DefaultApiService) GetChromatogramData(ctx context.Context, msDataId string, chromatogramId int32) (ImplResponse, error) {
	// TODO - update GetChromatogramData with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(200, map[string]interface{}{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetChromatogramData method not implemented")
}

// GetChromatogramHeader - 
func (s *DefaultApiService) GetChromatogramHeader(ctx context.Context, msDataId string, chromatogramId int32) (ImplResponse, error) {
	// TODO - update GetChromatogramHeader with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(200, map[string]interface{}{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetChromatogramHeader method not implemented")
}

// GetChromatograms - 
func (s *DefaultApiService) GetChromatograms(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetChromatograms with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(200, map[string]interface{}{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetChromatograms method not implemented")
}

// GetFileName - 
func (s *DefaultApiService) GetFileName(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetFileName with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetFileName200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetFileName200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetFileName method not implemented")
}

// GetInstrumentAnalyzer - 
func (s *DefaultApiService) GetInstrumentAnalyzer(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetInstrumentAnalyzer with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetInstrumentAnalyzer200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetInstrumentAnalyzer200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetInstrumentAnalyzer method not implemented")
}

// GetInstrumentData - 
func (s *DefaultApiService) GetInstrumentData(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetInstrumentData with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(200, map[string]interface{}{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetInstrumentData method not implemented")
}

// GetInstrumentDetector - 
func (s *DefaultApiService) GetInstrumentDetector(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetInstrumentDetector with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetInstrumentDetector200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetInstrumentDetector200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetInstrumentDetector method not implemented")
}

// GetInstrumentManufacturer - 
func (s *DefaultApiService) GetInstrumentManufacturer(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetInstrumentManufacturer with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetInstrumentManufacturer200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetInstrumentManufacturer200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetInstrumentManufacturer method not implemented")
}

// GetInstrumentModel - 
func (s *DefaultApiService) GetInstrumentModel(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetInstrumentModel with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetInstrumentModel200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetInstrumentModel200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetInstrumentModel method not implemented")
}

// GetIonisationMethod - 
func (s *DefaultApiService) GetIonisationMethod(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetIonisationMethod with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetIonisationMethod200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetIonisationMethod200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetIonisationMethod method not implemented")
}

// GetIsolationWindows - 
func (s *DefaultApiService) GetIsolationWindows(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetIsolationWindows with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetIsolationWindows200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetIsolationWindows200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetIsolationWindows method not implemented")
}

// GetSampleData - 
func (s *DefaultApiService) GetSampleData(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetSampleData with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetSampleData200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetSampleData200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetSampleData method not implemented")
}

// GetScanCount - 
func (s *DefaultApiService) GetScanCount(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetScanCount with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetScanCount200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetScanCount200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetScanCount method not implemented")
}

// GetScanHeader - 
func (s *DefaultApiService) GetScanHeader(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetScanHeader with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(200, map[string]interface{}{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetScanHeader method not implemented")
}

// GetScanPeakCount - 
func (s *DefaultApiService) GetScanPeakCount(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetScanPeakCount with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetScanPeakCount200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetScanPeakCount200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetScanPeakCount method not implemented")
}

// GetScanPeaks - 
func (s *DefaultApiService) GetScanPeaks(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetScanPeaks with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetScanPeaks200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetScanPeaks200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetScanPeaks method not implemented")
}

// GetScansData - 
func (s *DefaultApiService) GetScansData(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetScansData with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetScansData200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetScansData200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetScansData method not implemented")
}

// GetSoftware - 
func (s *DefaultApiService) GetSoftware(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetSoftware with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetSoftware200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetSoftware200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetSoftware method not implemented")
}

// GetSource - 
func (s *DefaultApiService) GetSource(ctx context.Context, msDataId string) (ImplResponse, error) {
	// TODO - update GetSource with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetSource200Response{}) or use other options such as http.Ok ...
	//return Response(200, GetSource200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetSource method not implemented")
}

// PostFile - 
func (s *DefaultApiService) PostFile(ctx context.Context, body *os.File) (ImplResponse, error) {
	// TODO - update PostFile with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PostFile200Response{}) or use other options such as http.Ok ...
	//return Response(200, PostFile200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PostFile method not implemented")
}
