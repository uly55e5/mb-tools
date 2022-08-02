# include "cpwiz.h"
# include "pwiz/data/msdata/MSDataFile.hpp"
# include "pwiz/data/msdata/LegacyAdapter.hpp"
# include "pwiz/data/common/CVTranslator.hpp"


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

InstrumentInfo getInstrumentInfo(MSDataFile file) {
    auto info = InstrumentInfo();
    auto ffile = (pwiz::msdata::MSDataFile *)file;
    auto iConfigP = ffile->instrumentConfigurationPtrs;
    auto softwareP = ffile->softwarePtrs;
	auto sampleP = ffile->samplePtrs;
	auto scansettingP = ffile->scanSettingsPtrs;
    pwiz::data::CVTranslator cvTranslator;
    if(iConfigP.size()>0) {
        pwiz::msdata::LegacyAdapter_Instrument adapter(*iConfigP[0], cvTranslator);
        info.analyzer = adapter.analyzer().data();
        info.detector = adapter.detector().data();
        info.ionisation = adapter.ionisation().data();
        info.manufacturer = adapter.manufacturer().data();
        info.model = adapter.model().data();
    }
    info.sample=(sampleP.size()>0?sampleP[0]->name + sampleP[0]->id:"").data();
    info.software=(softwareP.size()>0?softwareP[0]->id + " " + softwareP[0]->version:"").data();
    info.source=(scansettingP.size()>0?scansettingP[0]->sourceFilePtrs[0]->location:"").data();
    return info;

}

