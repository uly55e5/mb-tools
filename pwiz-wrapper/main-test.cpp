//
// Created by david on 04.08.22.
//
#include "cpwiz.h"
#include <iostream>

int main() {
    const char* error= "";
    auto file = MSDataOpenFile("../data/examples/small.pwiz.1.1.mzML", &error);
    if(error != "") {
        std::cout << error << std::endl;
    }
    int scans[] = {0, 1, 2, 3, 4, 5};
    int chroms[] = {0};
    auto scanCount = getLastScan(file);
    auto chromCount = getLastChromatogram(file);
    auto instInfo = getInstrumentInfo(file);
    auto headerInfo = getScanHeaderInfo(file, scans, 6);
    deleteScanHeader(headerInfo);
    auto chromHeaderInfo = getChromatogramHeaderInfo(file, chroms, 1);
    deleteChromatogramHeader(chromHeaderInfo);
    auto chromInfo = getChromatogramInfo(file, 0);
    deleteChromatogramInfo(chromInfo);
    auto isolationWindow = getIsolationWindow(file);
    deleteIsolationWindow(isolationWindow);
    auto startTime = getRunStartTimeStamp(file);
    auto peaks = getPeakList(file, scans, 6);
    deletePeakList(peaks);
    auto map3D = get3DMap(file, scans, 6, 0, 2000, 1);
    delete3DMap(map3D);
    MSDataClose(file);
}