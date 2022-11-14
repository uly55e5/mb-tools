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
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service      DefaultApiServicer
	errorHandler ErrorHandler
}

// DefaultApiOption for how the controller is set up.
type DefaultApiOption func(*DefaultApiController)

// WithDefaultApiErrorHandler inject ErrorHandler into controller
func WithDefaultApiErrorHandler(h ErrorHandler) DefaultApiOption {
	return func(c *DefaultApiController) {
		c.errorHandler = h
	}
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer, opts ...DefaultApiOption) Router {
	controller := &DefaultApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{
		{
			"Get3dMap",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/3dmap/{scanId}",
			c.Get3dMap,
		},
		{
			"GetChromatogramData",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/chromatograms/{chromatogramId}",
			c.GetChromatogramData,
		},
		{
			"GetChromatogramHeader",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/chromatograms/{chromatogramId}/header",
			c.GetChromatogramHeader,
		},
		{
			"GetChromatogramImage",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/chromatograms/{chromatogramId}/image",
			c.GetChromatogramImage,
		},
		{
			"GetChromatograms",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/chromatograms",
			c.GetChromatograms,
		},
		{
			"GetFileName",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/filename",
			c.GetFileName,
		},
		{
			"GetInstrumentAnalyzer",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/instrument/analyzer",
			c.GetInstrumentAnalyzer,
		},
		{
			"GetInstrumentData",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/instrument",
			c.GetInstrumentData,
		},
		{
			"GetInstrumentDetector",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/instrument/detector",
			c.GetInstrumentDetector,
		},
		{
			"GetInstrumentManufacturer",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/instrument/manufacturer",
			c.GetInstrumentManufacturer,
		},
		{
			"GetInstrumentModel",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/instrument/model",
			c.GetInstrumentModel,
		},
		{
			"GetIonisationMethod",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/instrument/ionisation",
			c.GetIonisationMethod,
		},
		{
			"GetIsolationWindows",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/IsolationWindows",
			c.GetIsolationWindows,
		},
		{
			"GetRunData",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/run",
			c.GetRunData,
		},
		{
			"GetSampleData",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/instrument/sample",
			c.GetSampleData,
		},
		{
			"GetScanHeader",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/scan/{scanId}/header",
			c.GetScanHeader,
		},
		{
			"GetScanImage",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/scan/{scanId}/image",
			c.GetScanImage,
		},
		{
			"GetScanPeaks",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/scan/{scanId}/peaks",
			c.GetScanPeaks,
		},
		{
			"GetScansData",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/scans",
			c.GetScansData,
		},
		{
			"GetSoftware",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/instrument/software",
			c.GetSoftware,
		},
		{
			"GetSource",
			strings.ToUpper("Get"),
			"/msdata/{msDataId}/instrument/source",
			c.GetSource,
		},
		{
			"PostFile",
			strings.ToUpper("Post"),
			"/file",
			c.PostFile,
		},
	}
}

// Get3dMap -
func (c *DefaultApiController) Get3dMap(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	msDataIdParam := chi.URLParam(r, "msDataId")

	scanIdParam, err := parseInt64Parameter(chi.URLParam(r, "scanId"), true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	lowMzParam, err := parseFloatParameter(query.Get("lowMz"), false)
	highMzParam, err := parseFloatParameter(query.Get("highMz"), false)
	resMzParam, err := parseFloatParameter(query.Get("resMz"), false)
	result, err := c.service.Get3dMap(r.Context(), msDataIdParam, scanIdParam, lowMzParam, highMzParam, resMzParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetChromatogramData -
func (c *DefaultApiController) GetChromatogramData(w http.ResponseWriter, r *http.Request) {
	msDataIdParam := chi.URLParam(r, "msDataId")

	chromatogramIdParam, err := parseInt64Parameter(chi.URLParam(r, "chromatogramId"), true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetChromatogramData(r.Context(), msDataIdParam, chromatogramIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetChromatogramHeader -
func (c *DefaultApiController) GetChromatogramHeader(w http.ResponseWriter, r *http.Request) {
	msDataIdParam := chi.URLParam(r, "msDataId")

	chromatogramIdParam, err := parseInt64Parameter(chi.URLParam(r, "chromatogramId"), true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetChromatogramHeader(r.Context(), msDataIdParam, chromatogramIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetChromatogramImage -
func (c *DefaultApiController) GetChromatogramImage(w http.ResponseWriter, r *http.Request) {
	chromatogramIdParam, err := parseInt64Parameter(chi.URLParam(r, "chromatogramId"), true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	msDataIdParam := chi.URLParam(r, "msDataId")

	result, err := c.service.GetChromatogramImage(r.Context(), chromatogramIdParam, msDataIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	WriteResponse(result.Body.(string), &result.Code, "text.html", result.Headers, w)

}

// GetChromatograms -
func (c *DefaultApiController) GetChromatograms(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	msDataIdParam := chi.URLParam(r, "msDataId")

	limitParam, err := parseInt64Parameter(query.Get("limit"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	pageParam, err := parseInt64Parameter(query.Get("page"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	countOnlyParam, err := parseBoolParameter(query.Get("countOnly"), false)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	result, err := c.service.GetChromatograms(r.Context(), msDataIdParam, limitParam, pageParam, countOnlyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetFileName -
func (c *DefaultApiController) GetFileName(w http.ResponseWriter, r *http.Request) {
	msDataIdParam := chi.URLParam(r, "msDataId")

	result, err := c.service.GetFileName(r.Context(), msDataIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetInstrumentAnalyzer -
func (c *DefaultApiController) GetInstrumentAnalyzer(w http.ResponseWriter, r *http.Request) {
	msDataIdParam := chi.URLParam(r, "msDataId")

	result, err := c.service.GetInstrumentAnalyzer(r.Context(), msDataIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetInstrumentData -
func (c *DefaultApiController) GetInstrumentData(w http.ResponseWriter, r *http.Request) {
	msDataIdParam := chi.URLParam(r, "msDataId")

	result, err := c.service.GetInstrumentData(r.Context(), msDataIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetInstrumentDetector -
func (c *DefaultApiController) GetInstrumentDetector(w http.ResponseWriter, r *http.Request) {
	msDataIdParam := chi.URLParam(r, "msDataId")

	result, err := c.service.GetInstrumentDetector(r.Context(), msDataIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetInstrumentManufacturer -
func (c *DefaultApiController) GetInstrumentManufacturer(w http.ResponseWriter, r *http.Request) {
	msDataIdParam := chi.URLParam(r, "msDataId")

	result, err := c.service.GetInstrumentManufacturer(r.Context(), msDataIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetInstrumentModel -
func (c *DefaultApiController) GetInstrumentModel(w http.ResponseWriter, r *http.Request) {
	msDataIdParam := chi.URLParam(r, "msDataId")

	result, err := c.service.GetInstrumentModel(r.Context(), msDataIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetIonisationMethod -
func (c *DefaultApiController) GetIonisationMethod(w http.ResponseWriter, r *http.Request) {
	msDataIdParam := chi.URLParam(r, "msDataId")

	result, err := c.service.GetIonisationMethod(r.Context(), msDataIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetIsolationWindows -
func (c *DefaultApiController) GetIsolationWindows(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	msDataIdParam := chi.URLParam(r, "msDataId")

	uniqueParam, err := parseBoolParameter(query.Get("unique"), false)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	result, err := c.service.GetIsolationWindows(r.Context(), msDataIdParam, uniqueParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetRunData -
func (c *DefaultApiController) GetRunData(w http.ResponseWriter, r *http.Request) {
	msDataIdParam := chi.URLParam(r, "msDataId")

	result, err := c.service.GetRunData(r.Context(), msDataIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetSampleData -
func (c *DefaultApiController) GetSampleData(w http.ResponseWriter, r *http.Request) {
	msDataIdParam := chi.URLParam(r, "msDataId")

	result, err := c.service.GetSampleData(r.Context(), msDataIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetScanHeader -
func (c *DefaultApiController) GetScanHeader(w http.ResponseWriter, r *http.Request) {
	msDataIdParam := chi.URLParam(r, "msDataId")

	scanIdParam, err := parseInt64Parameter(chi.URLParam(r, "scanId"), true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetScanHeader(r.Context(), msDataIdParam, scanIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetScanImage -
func (c *DefaultApiController) GetScanImage(w http.ResponseWriter, r *http.Request) {
	scanIdParam, err := parseInt64Parameter(chi.URLParam(r, "scanId"), true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	msDataIdParam := chi.URLParam(r, "msDataId")

	result, err := c.service.GetScanImage(r.Context(), scanIdParam, msDataIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	WriteResponse(result.Body.(string), &result.Code, "text.html", result.Headers, w)

}

// GetScanPeaks -
func (c *DefaultApiController) GetScanPeaks(w http.ResponseWriter, r *http.Request) {
	msDataIdParam := chi.URLParam(r, "msDataId")

	scanIdParam, err := parseInt64Parameter(chi.URLParam(r, "scanId"), true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetScanPeaks(r.Context(), msDataIdParam, scanIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetScansData -
func (c *DefaultApiController) GetScansData(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	msDataIdParam := chi.URLParam(r, "msDataId")

	countOnlyParam, err := parseBoolParameter(query.Get("countOnly"), false)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	limitParam, err := parseInt64Parameter(query.Get("limit"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	pageParam, err := parseInt64Parameter(query.Get("page"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetScansData(r.Context(), msDataIdParam, countOnlyParam, limitParam, pageParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetSoftware -
func (c *DefaultApiController) GetSoftware(w http.ResponseWriter, r *http.Request) {
	msDataIdParam := chi.URLParam(r, "msDataId")

	result, err := c.service.GetSoftware(r.Context(), msDataIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetSource -
func (c *DefaultApiController) GetSource(w http.ResponseWriter, r *http.Request) {
	msDataIdParam := chi.URLParam(r, "msDataId")

	result, err := c.service.GetSource(r.Context(), msDataIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// PostFile -
func (c *DefaultApiController) PostFile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	filenameParam := r.FormValue("filename")

	fileParam, err := ReadFormFileToTempFile(r, "file")
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.PostFile(r.Context(), filenameParam, fileParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code

	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}
