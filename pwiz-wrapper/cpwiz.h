#ifndef __cpwiz_H__
#define __cpwiz_H__
#define _GLIBCXX_USE_CXX11_ABI 0
#ifdef __cplusplus
extern "C" {
#endif
typedef void *MSDataFile;
typedef struct  {
    const char * manufacturer;
    const char * model;
	const char * ionisation;
	const char * analyzer;
	const char * detector;
	const char * software;
	const char * sample;
	const char * source;
} InstrumentInfo;


typedef struct {
    const char** names;
    const void * * values;
    long unsigned int numRows;
    long unsigned int numCols;
    const char * error;
} Header;

typedef struct {
    double* time;
    double* intensity;
    const char * id;
    const char * error;
    long unsigned int size;
} ChromatogramInfo;


MSDataFile MSDataOpenFile(const char *fileName);
//MSDataFile[] MSDataOpenFiles(char** fileNames);
void MSDataClose(MSDataFile msdata);

//void writeMSfile(const string& filenames, const string& format);
/*void writeSpectrumList(const string& file, const string& format,
           Rcpp::DataFrame spctr_header, Rcpp::List spctr_data,
           bool rtime_seconds,
           Rcpp::List software_processing);
*/
/*void copyWriteMSfile(const string& file, const string& format,
         const string& originalFile,
         Rcpp::DataFrame spctr_header,
         Rcpp::List spctr_data,
         bool rtime_seconds,
         Rcpp::List software_processing);
 */
//string getFilename();

unsigned long getLastScan(MSDataFile);

int getLastChrom(MSDataFile);

InstrumentInfo getInstrumentInfo(MSDataFile file);

//Rcpp::List getRunInfo();

/**
 * Reads the scan header for the provided scan(s). Note that this function
 * no longer returns a List, but a DataFrame, even if length whichScan is 1.
 * @return The scan header info is returned as a Rcpp::DataFrame
 **/
    Header getScanHeaderInfo(MSDataFile file, const int* scans, int size);

Header *getChromatogramHeaderInfo(MSDataFile file, const int *scans, int scansSize);

    ChromatogramInfo * getChromatogramInfo(MSDataFile file, int chromIdx);

/*    Rcpp::DataFrame getAllScanHeaderInfo();

    Rcpp::DataFrame getAllChromatogramHeaderInfo();

    Rcpp::List getPeakList(Rcpp::IntegerVector whichScan);

    Rcpp::NumericMatrix get3DMap(std::vector<int> scanNumbers, double whichMzLow, double whichMzHigh, double resMz);

    string getRunStartTimeStamp();
*/


#ifdef __cplusplus
}


#endif

#endif