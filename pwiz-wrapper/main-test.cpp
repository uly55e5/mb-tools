//
// Created by david on 04.08.22.
//
#include "cpwiz.h"

int main() {
    auto file = MSDataOpenFile("../data/examples/small.pwiz.1.1.mzML");
    int scans[] = {0,1,2,3,4,5};
    getPeakList(file,scans,6);
    get3DMap(file, scans,6,0,2000,1);
}