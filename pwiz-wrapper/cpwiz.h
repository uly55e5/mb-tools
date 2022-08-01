#ifndef __cpwiz_H__
#define __cpwiz_H__
#define _GLIBCXX_USE_CXX11_ABI 0
#ifdef __cplusplus
extern "C" {
#endif
typedef void *MSDataFile;
MSDataFile MSDataOpenFile(char *fileName);
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

//int getLastScan() const;

int getLastChrom(MSDataFile);

//Rcpp::List getInstrumentInfo();

//Rcpp::List getRunInfo();

/**
 * Reads the scan header for the provided scan(s). Note that this function
 * no longer returns a List, but a DataFrame, even if length whichScan is 1.
 * @return The scan header info is returned as a Rcpp::DataFrame
 **/
/*  Rcpp::DataFrame getScanHeaderInfo(Rcpp::IntegerVector whichScan);

    Rcpp::DataFrame getChromatogramHeaderInfo(Rcpp::IntegerVector whichScan);

    Rcpp::DataFrame getChromatogramsInfo(int whichChrom);

    Rcpp::DataFrame getAllScanHeaderInfo();

    Rcpp::DataFrame getAllChromatogramHeaderInfo();

    Rcpp::List getPeakList(Rcpp::IntegerVector whichScan);

    Rcpp::NumericMatrix get3DMap(std::vector<int> scanNumbers, double whichMzLow, double whichMzHigh, double resMz);

    string getRunStartTimeStamp();
*/
#ifdef __cplusplus
}
#endif

#endif