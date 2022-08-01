# include "cpwiz.h"
# include "pwiz/data/msdata/MSDataFile.hpp"
#define _GLIBCXX_USE_CXX11_ABI 0

MSDataFile MSDataOpenFile(char *fileName) {
    auto ms = new pwiz::msdata::MSDataFile(std::string(fileName));
    return ms;
}

void MSDataClose(MSDataFile file) {
    delete (pwiz::msdata::MSDataFile *) file;
}

int getLastChrom(MSDataFile file) {
    auto clp = ((pwiz::msdata::MSDataFile *) file)->run.chromatogramListPtr;
    return (int) clp->size();
}

