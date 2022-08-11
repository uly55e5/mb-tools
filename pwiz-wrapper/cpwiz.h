#ifndef __cpwiz_H__
#define __cpwiz_H__
#define _GLIBCXX_USE_CXX11_ABI 0
#ifdef __cplusplus
extern "C" {
#endif
typedef void *MSDataFile;
typedef struct {
    const char *manufacturer;
    const char *model;
    const char *ionisation;
    const char *analyzer;
    const char *detector;
    const char *software;
    const char *sample;
    const char *source;
} InstrumentInfo;


typedef struct {
    const char **names;
    void **values;
    long unsigned int numRows;
    long unsigned int numCols;
    const char *error;
} Header;

typedef struct {
    double *time;
    double *intensity;
    const char *id;
    const char *error;
    long unsigned int size;
} ChromatogramInfo;

typedef struct {
    double *high;
    double *low;
    long unsigned int size;
} IsolationWindows;

typedef struct {
    const char *error;
    const char **colNames;
    int colNum;
    int scanNum;
    double ***values;
    int *valSizes;
    int *scans;

} PeakList;

typedef struct {
    int *scans;
    int scanSize;
    double **values;
    int valueSize;

} Map3d;

MSDataFile MSDataOpenFile(const char *fileName, const char **errorMessage);
void MSDataClose(MSDataFile msdata);
void deletePeakList(PeakList* list);
void delete3DMap(Map3d * map);
void deleteChromatogramInfo(ChromatogramInfo * info);
void deleteIsolationWindow(IsolationWindows* windows);
void deleteHeader(Header*header);



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

unsigned long getLastScan(MSDataFile);

int getLastChromatogram(MSDataFile);

InstrumentInfo getInstrumentInfo(MSDataFile file);

Header *getScanHeaderInfo(MSDataFile file, const int *scans, int size);

Header *getChromatogramHeaderInfo(MSDataFile file, const int *scans, int scansSize);

ChromatogramInfo *getChromatogramInfo(MSDataFile file, int chromIdx);

IsolationWindows *getIsolationWindow(MSDataFile file);

const char *getRunStartTimeStamp(MSDataFile file);

PeakList *getPeakList(MSDataFile file, int *scans, int size);

Map3d get3DMap(MSDataFile file, int *scans, int scanSize, double whichMzLow, double whichMzHigh, double resMz);

#ifdef __cplusplus
}


#endif

#endif