//
// Created by david on 04.08.22.
//
#include "cpwiz.h"

int main() {
    auto file = MSDataOpenFile("../data/examples/small.pwiz.1.1.mzML");
    int scans[] = {0};
    getScanHeaderInfo(file, scans,1);
}