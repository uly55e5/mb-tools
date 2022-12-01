package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	mzserver2 "github.com/uly55e5/mb-tools/cmd/openapi-server/src"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

const SERVER = "http://localhost:49043/"

func TestMain(m *testing.M) {
	DefaultApiService := mzserver2.NewDefaultApiService()
	DefaultApiController := mzserver2.NewDefaultApiController(DefaultApiService)
	router := mzserver2.NewRouter(DefaultApiController)
	go func() {
		err := http.ListenAndServe("localhost:49043", router)
		if err != nil {
			panic("Could not start server")
		}
	}()
	m.Run()
}

func TestPostFile(t *testing.T) {
	res, err := getPostFileResponse("../../data/examples/small.pwiz.1.1.mzML", "Test1.mzML")
	assert.NoError(t, err, "Error while getting response")
	assert.Equal(t, 200, res.StatusCode, "Status not OK")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Error while reading body")
	assert.NotEmpty(t, b, "Body is empty")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not Unmarshal json")
	assert.Contains(t, js, "id", "Response does not inlcude id")
	id := js["id"].(string)
	assert.NotEmpty(t, id, "id is empty")
}

func TestGetFileName(t *testing.T) {
	res, err := getMsDataRequest("filename")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "filename", "No filename in response")
	assert.Equal(t, "Test1.mzML", js["filename"])

}
func TestGetInstrument(t *testing.T) {
	res, err := getMsDataRequest("instrument")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "Analyzer", "No analyzer in response")
	assert.Equal(t, "fourier transform ion cyclotron resonance mass spectrometer", js["Analyzer"])
	assert.Contains(t, js, "Detector", "No detector in response")
	assert.Equal(t, "inductive detector", js["Detector"])

	assert.Contains(t, js, "Ionisation", "No ionisation in response")
	assert.Equal(t, "electrospray ionization", js["Ionisation"])

	assert.Contains(t, js, "Manufacturer", "No manufacturer in response")
	assert.Equal(t, "Xcalibur ", js["Manufacturer"])
	assert.Contains(t, js, "Model", "No model in response")
	assert.Equal(t, "LTQ FT", js["Model"])
	assert.Contains(t, js, "Sample", "No sample in response")
	assert.Equal(t, "", js["Sample"])
	assert.Contains(t, js, "Software", "No software in response")
	assert.Equal(t, "Xcalibur 1.1 Beta 7", js["Software"])
	assert.Contains(t, js, "Source", "No source in response")
	assert.Equal(t, "", js["Source"])
}

func TestGetRun(t *testing.T) {
	res, err := getMsDataRequest("run")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "DEndTime", "No end time in response")
	assert.InDelta(t, 29.234199999999998, js["DEndTime"], 0.0000000001, "End time not equal")
	assert.Contains(t, js, "DStartTime", "No start time in response")
	assert.InDelta(t, 0.29610000000000003, js["DStartTime"], 0.0000000001, "Start time not equal")
	assert.Contains(t, js, "HighMz", "No high Mz in response")
	assert.InDelta(t, 2000.0099466203771, js["HighMz"], 0.0000000001, "High Mz not equal")
	assert.Contains(t, js, "LowMz", "No lowMz in response")
	assert.InDelta(t, 162.24594116210938, js["LowMz"], 0.0000000001, "Low Mz not equal")
	assert.Contains(t, js, "MsLevels", "No msLevels in response")
	var msLevels []int
	for _, d := range js["MsLevels"].([]interface{}) {
		msLevels = append(msLevels, int(d.(float64)))
	}
	assert.Equal(t, []int{1, 2}, msLevels)
	assert.Contains(t, js, "ScanCount", "No scan count in response")
	assert.Equal(t, 48.0, js["ScanCount"])
	assert.Contains(t, js, "StartTimeStamp", "No start time stamp in response")
	assert.Equal(t, "2005-07-20T14:44:22", js["StartTimeStamp"])
}
func TestGetDetector(t *testing.T) {
	res, err := getMsDataRequest("instrument/detector")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "detector", "No detector in response")
	assert.Equal(t, "inductive detector", js["detector"])
}
func TestGetAnalyzer(t *testing.T) {
	res, err := getMsDataRequest("instrument/analyzer")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "analyzer", "No analyzer in response")
	assert.Equal(t, "fourier transform ion cyclotron resonance mass spectrometer", js["analyzer"])
}
func TestGetIonisation(t *testing.T) {
	res, err := getMsDataRequest("instrument/ionisation")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "ionisation", "No ionisation in response")
	assert.Equal(t, "electrospray ionization", js["ionisation"])
}
func TestGetManufacturer(t *testing.T) {
	res, err := getMsDataRequest("instrument/manufacturer")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "manufacturer", "No manufacturer in response")
	assert.Equal(t, "Xcalibur ", js["manufacturer"])
}
func TestGetModel(t *testing.T) {
	res, err := getMsDataRequest("instrument/model")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "model", "No model in response")
	assert.Equal(t, "LTQ FT", js["model"])
}
func TestGetSample(t *testing.T) {
	res, err := getMsDataRequest("instrument/sample")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.NotContains(t, js, "sample", "Found sample in response")
}
func TestGetSoftware(t *testing.T) {
	res, err := getMsDataRequest("instrument/software")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "software", "No software in response")
	assert.Equal(t, "Xcalibur 1.1 Beta 7", js["software"])
}
func TestGetSource(t *testing.T) {
	res, err := getMsDataRequest("instrument/source")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.NotContains(t, js, "source", "Found Source in response")
}
func TestGetScans(t *testing.T) {
	res, err := getMsDataRequest("scans")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "scans", "No scans in response")
	var scans []int
	for _, d := range js["scans"].([]interface{}) {
		scans = append(scans, int(d.(float64)))
	}
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47}, scans)
}
func TestGetScanHeader(t *testing.T) {
	res, err := getMsDataRequest("scan/0/header")
	b, err := io.ReadAll(res.Body)
	println(b)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "AcquisitionNum", "No filename in response")
	assert.Equal(t, 1.0, js["AcquisitionNum"])
	assert.Contains(t, js, "BasePeakIntensity", "No filename in response")
	assert.InDelta(t, 1.471973875e+06, js["BasePeakIntensity"], 0.000000000001, "Base peak intensity not in delta")
	assert.Contains(t, js, "BasePeakMZ", "No filename in response")
	assert.InDelta(t, 810.415283203125, js["BasePeakMZ"], 0.000000000001, "Base peak Mz not in delta")
	assert.Contains(t, js, "Centroided", "No filename in response")
	assert.Equal(t, false, js["Centroided"])
	assert.Contains(t, js, "CollisionEnergy", "No filename in response")
	assert.Equal(t, 0.0, js["CollisionEnergy"])
	assert.Contains(t, js, "FilterString", "No filename in response")
	assert.Equal(t, "FTMS + p ESI Full ms [200.00-2000.00]", js["FilterString"])
	assert.Contains(t, js, "HighMZ", "No filename in response")
	assert.InDelta(t, 2000.0099466203771, js["HighMZ"], 0.000000000001, "Base peak Mz not in delta")
	assert.Contains(t, js, "IonInjectionTime", "No filename in response")
	assert.Equal(t, 0.0, js["IonInjectionTime"])
	assert.Contains(t, js, "IonMobilityDriftTime", "No filename in response")
	assert.Equal(t, "NaN", js["IonMobilityDriftTime"])
	assert.Contains(t, js, "IonisationEnergy", "No filename in response")
	assert.Equal(t, 0.0, js["IonisationEnergy"])
	assert.Contains(t, js, "IsolationWindowLowerOffset", "No filename in response")
	assert.Equal(t, "NaN", js["IsolationWindowLowerOffset"])
	assert.Contains(t, js, "IsolationWindowTargetMZ", "No filename in response")
	assert.Equal(t, "NaN", js["IsolationWindowTargetMZ"])
	assert.Contains(t, js, "IsolationWindowUpperOffset", "No filename in response")
	assert.Equal(t, "NaN", js["IsolationWindowUpperOffset"])
	assert.Contains(t, js, "LowMZ", "No filename in response")
	assert.InDelta(t, 200.00018816645022, js["LowMZ"], 0.000000000001, "Base peak Mz not in delta")
	assert.Contains(t, js, "MergedResultEndScanNum", "No filename in response")
	assert.Equal(t, -1.0, js["MergedResultEndScanNum"])
	assert.Contains(t, js, "MergedResultScanNum", "No filename in response")
	assert.Equal(t, -1.0, js["MergedResultScanNum"])
	assert.Contains(t, js, "MergedResultStartScanNum", "No filename in response")
	assert.Equal(t, -1.0, js["MergedResultStartScanNum"])
	assert.Contains(t, js, "MergedScan", "No filename in response")
	assert.Equal(t, -1.0, js["MergedScan"])
	assert.Contains(t, js, "MsLevel", "No filename in response")
	assert.Equal(t, 1.0, js["MsLevel"])
	assert.Contains(t, js, "PeaksCount", "No filename in response")
	assert.Equal(t, 19914.0, js["PeaksCount"])
	assert.Contains(t, js, "Polarity", "No filename in response")
	assert.Equal(t, 1.0, js["Polarity"])
	assert.Contains(t, js, "PrecursorCharge", "No filename in response")
	assert.Equal(t, 0.0, js["PrecursorCharge"])
	assert.Contains(t, js, "PrecursorIntensity", "No filename in response")
	assert.Equal(t, 0.0, js["PrecursorIntensity"])
	assert.Contains(t, js, "PrecursorMZ", "No filename in response")
	assert.Equal(t, 0.0, js["PrecursorMZ"])
	assert.Contains(t, js, "PrecursorScanNum", "No filename in response")
	assert.Equal(t, -1.0, js["PrecursorScanNum"])
	assert.Contains(t, js, "RetentionTime", "No filename in response")
	assert.InDelta(t, 0.29610000000000003, js["RetentionTime"], 0.000000000001, "Base peak intensity not in delta")
	assert.Contains(t, js, "ScanWindowLowerLimit", "No filename in response")
	assert.Equal(t, 200.0, js["ScanWindowLowerLimit"])
	assert.Contains(t, js, "ScanWindowUpperLimit", "No filename in response")
	assert.Equal(t, 2000.0, js["ScanWindowUpperLimit"])
	assert.Contains(t, js, "SeqNum", "No filename in response")
	assert.Equal(t, 0.0, js["SeqNum"])
	assert.Contains(t, js, "SpectrumId", "No filename in response")
	assert.Equal(t, "controllerType=0 controllerNumber=1 scan=1", js["SpectrumId"])
	assert.Contains(t, js, "TotIonCurrent", "No filename in response")
	assert.InDelta(t, 1.5245068e+07, js["TotIonCurrent"], 0.000000000001, "Base peak intensity not in delta")

}
func TestGetScanPeaks(t *testing.T) {
	res, err := getMsDataRequest("scan/0/peaks")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "source", "Found Source in response")
}
func TestGetIsolationWindows(t *testing.T) {
	res, err := getMsDataRequest("IsolationWindows")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "IsolationWindows", "Found Source in response")
	assert.Len(t, js["IsolationWindows"], 1)
	assert.Contains(t, js["IsolationWindows"].([]interface{})[0], "low")
	assert.Contains(t, js["IsolationWindows"].([]interface{})[0], "high")
}
func TestGetChromatograms(t *testing.T) {
	res, err := getMsDataRequest("chromatograms")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "source", "Found Source in response")
}
func TestGetChromatogram(t *testing.T) {
	res, err := getMsDataRequest("chromatograms/0")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "Id")
	assert.Equal(t, "TIC", js["Id"])
	assert.Contains(t, js, "Intensity")
	assert.Contains(t, js["Intensity"], 114331.703125)
	assert.Contains(t, js, "Time")
	assert.Contains(t, js["Time"], 0.06192333333333334)

}
func TestGetChromatogramHeader(t *testing.T) {
	res, err := getMsDataRequest("chromatograms/0/header")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "source", "Found Source in response")
}
func TestGet3dMap(t *testing.T) {
	res, err := getMsDataRequest("3dmap/0")
	b, err := io.ReadAll(res.Body)
	assert.NoError(t, err, "Could not read response body")
	js := map[string]interface{}{}
	err = json.Unmarshal(b, &js)
	assert.NoError(t, err, "Could not decode json")
	assert.Contains(t, js, "source", "Found Source in response")
}
func TestGetScanImage(t *testing.T) {
	res, err := getMsDataRequest("scan/0/image")
	b, err := io.ReadAll(res.Body)
	s := string(b)
	assert.NoError(t, err, "Could not read response body")
	assert.Contains(t, s, "<div class=\"container\">")
	assert.Contains(t, s, "<script type=\"text/javascript\">")
	assert.Contains(t, s, "{\"value\":3772.0205078125},{\"value\":2250.034912109375}")

}
func TestGetChromatogramImage(t *testing.T) {
	res, err := getMsDataRequest("chromatograms/0/image")
	b, err := io.ReadAll(res.Body)
	s := string(b)
	assert.NoError(t, err, "Could not read response body")
	assert.Contains(t, s, "<div class=\"container\">")
	assert.Contains(t, s, "<script type=\"text/javascript\">")
	assert.Contains(t, s, "{\"value\":11037852},{\"value\":1102582.125}")
}

func getMsDataRequest(apiPath string) (*http.Response, error) {
	id, err := getTestDataFileId()
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	return client.Get(SERVER + "msdata/" + id + "/" + apiPath)
}

func getTestDataFileId() (string, error) {
	res, err := getPostFileResponse("../../data/examples/small.pwiz.1.1.mzML", "Test1.mzML")
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", errors.New("response")
	}
	js := map[string]interface{}{}
	b, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(b, &js)
	if err != nil {
		return "", err
	}
	return js["id"].(string), nil

}

func getPostFileResponse(filepath string, filename string) (*http.Response, error) {
	file, err := os.Open(filepath)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	err = file.Close()
	if err != nil {
		println("Could not close file: ", err.Error())
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	_, err = part.Write(fileContents)
	if err != nil {
		return nil, err
	}

	err = writer.WriteField("filename", filename)

	err = writer.Close()
	if err != nil {
		println(err.Error())
		return nil, err
	}

	req, err := http.NewRequest("POST", SERVER+"file", body)
	req.Header.Set("Content-Type", "multipart/form-data; boundary="+writer.Boundary())
	if err != nil {
		println(err.Error())
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	return res, nil
}
