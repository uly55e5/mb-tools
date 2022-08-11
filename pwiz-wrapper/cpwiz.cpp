# include "cpwiz.h"
# include "pwiz/data/msdata/MSDataFile.hpp"
# include "pwiz/data/msdata/LegacyAdapter.hpp"
# include "pwiz/data/common/CVTranslator.hpp"
# include <variant>
# include <iostream>

#define _GLIBCXX_USE_CXX11_ABI 0

typedef std::vector<int> IntegerVector;
typedef std::vector<double> NumericVector;
typedef std::vector<std::string> StringVector;
typedef std::vector<char> LogicalVector;
typedef std::map<std::string, std::variant<IntegerVector, NumericVector, StringVector, LogicalVector>> HeaderMap;

void deletePeakList(PeakList *list) {
    for(int i=0; i<list->scanNum;i++){
        delete[] list->values[i][0];
        delete[] list->values[i][1];
        delete[] list->values[i];
    }
    delete[] list->valSizes;
    delete[] list->values;
    delete list;
}

void delete3DMap(Map3d * map) {
    for(int i=0; i<map->scanSize;i++){
        delete[] map->values[i];
    }
    delete[] map->values;
}

void deleteHeader(Header*header) {
    for(int i=0; i<header->numCols;i++) {
        delete header->values[i];
    }
    delete[] header->values;
    delete[] header->names;
    delete header;
}

void deleteChromatogramInfo(ChromatogramInfo * info) {
    delete[] info->intensity;
    delete[] info->time;
    delete info;
}

void deleteIsolationWindow(IsolationWindows* windows) {
    delete[] windows->high;
    delete[] windows->low;
    delete windows;
}

int getAcquisitionNumber(MSDataFile file, std::string id, size_t index);


Header *convertHeader(HeaderMap &headerM);

MSDataFile MSDataOpenFile(const char *fileName, const char **errorMessage) {
    try {
        auto ms = new pwiz::msdata::MSDataFile(std::string(fileName));
        return ms;
    } catch (std::runtime_error e) {
        *errorMessage = e.what();
    }
    return nullptr;
}

void MSDataClose(MSDataFile file) {
    if (file != nullptr) {
        delete (pwiz::msdata::MSDataFile *) file;
    }
}

int getLastChromatogram(MSDataFile file) {
    auto chromP = ((pwiz::msdata::MSDataFile *) file)->run.chromatogramListPtr;
    return (int) chromP->size();
}

InstrumentInfo getInstrumentInfo(MSDataFile file) {
    auto info = InstrumentInfo();
    auto ffile = (pwiz::msdata::MSDataFile *) file;
    auto iConfigP = ffile->instrumentConfigurationPtrs;
    auto softwareP = ffile->softwarePtrs;
    auto sampleP = ffile->samplePtrs;
    auto scansettingP = ffile->scanSettingsPtrs;
    pwiz::data::CVTranslator cvTranslator;
    if (!iConfigP.empty()) {
        pwiz::msdata::LegacyAdapter_Instrument adapter(*iConfigP[0], cvTranslator);
        info.ionisation = adapter.ionisation().c_str();
        info.analyzer = adapter.analyzer().c_str();
        info.detector = adapter.detector().c_str();
        info.manufacturer = adapter.manufacturer().c_str();
        info.model = adapter.model().c_str();
    }
    info.sample = (!sampleP.empty() ? sampleP[0]->name + sampleP[0]->id : "").c_str();
    info.software = (!softwareP.empty() ? softwareP[0]->id + " " + softwareP[0]->version : "").c_str();
    info.source = (!scansettingP.empty() ? scansettingP[0]->sourceFilePtrs[0]->location : "").c_str();
    return info;

}

unsigned long getLastScan(MSDataFile file) {
    auto ffile = (pwiz::msdata::MSDataFile *) file;
    auto listP = ffile->run.spectrumListPtr;
    return listP->size();
}

typedef struct {
    std::vector<double> high;
    std::vector<double> low;
} IWindows;

IsolationWindows *getIsolationWindow(MSDataFile file) {
    auto ffile = (pwiz::msdata::MSDataFile *) file;
    auto SpectrumListP = ffile->run.spectrumListPtr;
    auto cIWindows = new IsolationWindows{};
    auto scanSize = SpectrumListP->size();
    cIWindows->high = new double[scanSize];
    cIWindows->low = new double[scanSize];
    int ms2Count=0;
    for (int i = 0; i < scanSize; i++) {
        auto SpectrumP = SpectrumListP->spectrum(i, pwiz::msdata::DetailLevel_FullMetadata);
        if (!SpectrumP->precursors.empty()) {
            auto iwin = SpectrumP->precursors[0].isolationWindow;
            cIWindows->high[ms2Count]=iwin.cvParam(pwiz::cv::MS_isolation_window_upper_offset).value.empty() ? NAN :
                                     iwin.cvParam(pwiz::cv::MS_isolation_window_upper_offset).valueAs<double>();
            cIWindows->low[ms2Count] =iwin.cvParam(pwiz::cv::MS_isolation_window_lower_offset).value.empty() ? NAN :
                                    iwin.cvParam(pwiz::cv::MS_isolation_window_lower_offset).valueAs<double>();
            ms2Count++;

        }
    }
    cIWindows->size=ms2Count;
    return cIWindows;
}

Header * getScanHeaderInfo(MSDataFile file, const int *scans, int scansSize) {
    auto ffile = (pwiz::msdata::MSDataFile *) file;
    auto headerM = HeaderMap{
            {"seqNum",                     IntegerVector{}},
            {"acquisitionNum",             IntegerVector{}},
            {"msLevel",                    IntegerVector{}},
            {"polarity",                   IntegerVector{}},
            {"peaksCount",                 IntegerVector{}},
            {"totIonCurrent",              NumericVector{}},
            {"retentionTime",              NumericVector{}},
            {"basePeakMZ",                 NumericVector{}},
            {"basePeakIntensity",          NumericVector{}},
            {"collisionEnergy",            NumericVector{}},
            {"ionisationEnergy",           NumericVector{}},
            {"lowMZ",                      NumericVector{}},
            {"highMZ",                     NumericVector{}},
            {"precursorScanNum",           IntegerVector{}},
            {"precursorMZ",                NumericVector{}},
            {"precursorCharge",            IntegerVector{}},
            {"precursorIntensity",         NumericVector{}},
            {"mergedScan",                 IntegerVector{}},
            {"mergedResultScanNum",        IntegerVector{}},
            {"mergedResultStartScanNum",   IntegerVector{}},
            {"mergedResultEndScanNum",     IntegerVector{}},
            {"ionInjectionTime",           NumericVector{}},
            {"filterString",               StringVector{}},
            {"spectrumId",                 StringVector{}},
            {"centroided",                 LogicalVector{}},
            {"ionMobilityDriftTime",       NumericVector{}},
            {"isolationWindowTargetMZ",    NumericVector{}},
            {"isolationWindowLowerOffset", NumericVector{}},
            {"isolationWindowUpperOffset", NumericVector{}},
            {"scanWindowLowerLimit",       NumericVector{}},
            {"scanWindowUpperLimit",       NumericVector{}},
    };
    auto SpectrumListP = ffile->run.spectrumListPtr;
    for (size_t i = 0; i < scansSize; i++) {
        int current_scan = scans[i];
        auto current_index = static_cast<size_t>(current_scan);
        auto SpectrumP = SpectrumListP->spectrum(current_index, pwiz::msdata::DetailLevel_FullMetadata);
        auto &scan = SpectrumP->scanList.scans[0];
        std::get<IntegerVector>((headerM)["seqNum"]).push_back(current_scan);
        std::get<IntegerVector>((headerM)["acquisitionNum"]).push_back(
                getAcquisitionNumber(file, SpectrumP->id, current_index));
        std::get<StringVector>((headerM)["spectrumId"]).push_back(SpectrumP->id);
        std::get<IntegerVector>((headerM)["msLevel"]).push_back(
                SpectrumP->cvParam(pwiz::cv::MS_ms_level).valueAs<int>());
        std::get<IntegerVector>((headerM)["peaksCount"]).push_back(static_cast<int>(SpectrumP->defaultArrayLength));
        std::get<NumericVector>((headerM)["totIonCurrent"]).push_back(
                SpectrumP->cvParam(pwiz::cv::MS_total_ion_current).valueAs<double>());
        std::get<NumericVector>((headerM)["basePeakMZ"]).push_back(
                SpectrumP->cvParam(pwiz::cv::MS_base_peak_m_z).valueAs<double>());
        std::get<NumericVector>((headerM)["basePeakIntensity"]).push_back(
                SpectrumP->cvParam(pwiz::cv::MS_base_peak_intensity).valueAs<double>());
        std::get<NumericVector>((headerM)["ionisationEnergy"]).push_back(
                SpectrumP->cvParam(pwiz::cv::MS_ionization_energy_OBSOLETE).valueAs<double>());
        std::get<NumericVector>((headerM)["lowMZ"]).push_back(
                SpectrumP->cvParam(pwiz::cv::MS_lowest_observed_m_z).valueAs<double>());
        std::get<NumericVector>((headerM)["highMZ"]).push_back(
                SpectrumP->cvParam(pwiz::cv::MS_highest_observed_m_z).valueAs<double>());
        auto param = SpectrumP->cvParamChild(pwiz::cv::MS_scan_polarity);
        std::get<IntegerVector>((headerM)["polarity"]).push_back(
                param.cvid == pwiz::cv::MS_negative_scan ? 0 : (param.cvid == pwiz::cv::MS_positive_scan ? +1 : -1));
        param = SpectrumP->cvParamChild(pwiz::cv::MS_spectrum_representation);
        std::get<LogicalVector>((headerM)["centroided"]).push_back(
                param.cvid == pwiz::cv::MS_centroid_spectrum);
        std::get<NumericVector>((headerM)["retentionTime"]).push_back(
                scan.cvParam(pwiz::cv::MS_scan_start_time).timeInSeconds());
        std::get<NumericVector>((headerM)["ionInjectionTime"]).push_back(
                (scan.cvParam(pwiz::cv::MS_ion_injection_time).timeInSeconds() * 1000));
        std::get<StringVector>((headerM)["filterString"]).push_back(
                scan.cvParam(pwiz::cv::MS_filter_string).value.empty() ? "" :
                scan.cvParam(pwiz::cv::MS_filter_string).value);
        std::get<NumericVector>((headerM)["ionMobilityDriftTime"]).push_back(
                scan.cvParam(pwiz::cv::MS_ion_mobility_drift_time).value.empty() ? NAN : (
                        scan.cvParam(pwiz::cv::MS_ion_mobility_drift_time).timeInSeconds() * 1000));
        if (!scan.scanWindows.empty()) {
            std::get<NumericVector>((headerM)["scanWindowLowerLimit"]).push_back(
                    scan.scanWindows[0].cvParam(pwiz::cv::MS_scan_window_lower_limit).valueAs<double>());
            std::get<NumericVector>((headerM)["scanWindowUpperLimit"]).push_back(
                    scan.scanWindows[0].cvParam(pwiz::cv::MS_scan_window_upper_limit).valueAs<double>());
        } else {
            std::get<NumericVector>((headerM)["scanWindowLowerLimit"]).push_back(NAN);
            std::get<NumericVector>((headerM)["scanWindowUpperLimit"]).push_back(NAN);
        }
        std::get<IntegerVector>((headerM)["mergedScan"]).push_back(-1);
        std::get<IntegerVector>((headerM)["mergedResultScanNum"]).push_back(-1);
        std::get<IntegerVector>((headerM)["mergedResultStartScanNum"]).push_back(-1);
        std::get<IntegerVector>((headerM)["mergedResultEndScanNum"]).push_back(-1);

        const auto &precursor = !SpectrumP->precursors.empty() ? SpectrumP->precursors[0] : pwiz::msdata::Precursor{};
        std::get<NumericVector>((headerM)["collisionEnergy"]).push_back(
                precursor.activation.cvParam(pwiz::cv::MS_collision_energy).valueAs<double>());
        size_t precursorIndex = SpectrumListP->find(precursor.spectrumID);
        if (precursorIndex < SpectrumListP->size()) {
            std::get<IntegerVector>((headerM)["precursorScanNum"]).push_back(
                    getAcquisitionNumber(file, precursor.spectrumID, precursorIndex));
        } else {
            std::get<IntegerVector>((headerM)["precursorScanNum"]).push_back(-1);
        }
        const auto &selectedIon = !precursor.selectedIons.empty() ? precursor.selectedIons[0]
                                                                  : pwiz::msdata::SelectedIon{};
        std::get<NumericVector>((headerM)["precursorMZ"]).push_back(
                selectedIon.cvParam(pwiz::cv::MS_selected_ion_m_z).value.empty()
                ? selectedIon.cvParam(pwiz::cv::MS_m_z).valueAs<double>()
                : selectedIon.cvParam(pwiz::cv::MS_selected_ion_m_z).valueAs<double>());
        std::get<IntegerVector>((headerM)["precursorCharge"]).push_back(
                selectedIon.cvParam(pwiz::cv::MS_charge_state).valueAs<int>());
        std::get<NumericVector>((headerM)["precursorIntensity"]).push_back(
                selectedIon.cvParam(pwiz::cv::MS_peak_intensity).valueAs<double>());

        auto iwin = !SpectrumP->precursors.empty() ? SpectrumP->precursors[0].isolationWindow
                                                   : pwiz::msdata::IsolationWindow{};
        std::get<NumericVector>((headerM)["isolationWindowTargetMZ"]).push_back(
                iwin.cvParam(pwiz::cv::MS_isolation_window_target_m_z).value.empty() ? NAN
                                                                                     : iwin.cvParam(
                        pwiz::cv::MS_isolation_window_target_m_z).valueAs<double>());
        std::get<NumericVector>((headerM)["isolationWindowLowerOffset"]).push_back(
                iwin.cvParam(pwiz::cv::MS_isolation_window_lower_offset).value.empty() ? NAN
                                                                                       : iwin.cvParam(
                        pwiz::cv::MS_isolation_window_lower_offset).valueAs<double>());
        std::get<NumericVector>((headerM)["isolationWindowUpperOffset"]).push_back(
                iwin.cvParam(pwiz::cv::MS_isolation_window_upper_offset).value.empty() ? NAN
                                                                                       : iwin.cvParam(
                        pwiz::cv::MS_isolation_window_upper_offset).valueAs<double>());
    }
    Header *cMap = convertHeader(headerM);
    cMap->numRows = std::get<IntegerVector>((headerM)["seqNum"]).size();
    return cMap;
}

Header *convertHeader(HeaderMap &headerM) {
    auto cMap = new Header{};
    cMap->names = new const char *[(headerM).size()];
    cMap->values = new void *[(headerM).size()];
    cMap->numCols = 0;
    cMap->error = "";
    for (const auto &[key, value]: (headerM)) {
        cMap->names[cMap->numCols] = key.c_str();
        std::visit([key, &cMap](auto &&arg) {
            if (arg.size() != cMap->numRows) {
                cMap->error = (std::string("ColSize does not match: ") + key + " : " + std::to_string(arg.size()) +
                               "/" + std::to_string(cMap->numRows)).c_str();
            }
            using T = std::decay_t<decltype(arg)>;
            if constexpr (std::is_same_v<T, NumericVector>) {
                auto cArray = new double[arg.size()];
                std::copy(arg.begin(),arg.end(),cArray);
                cMap->values[cMap->numCols++] = cArray;
            } else if constexpr (std::is_same_v<T, StringVector>) {
                auto cArray = new char*[arg.size()];
                for(int i=0; i<arg.size();i++) {
                    cArray[i] = new char[arg[i].size()+1];
                    strcpy(cArray[i],arg[i].c_str());
                }
                cMap->values[cMap->numCols++] = cArray;
            } else if constexpr (std::is_same_v<T, IntegerVector>) {
            auto cArray = new int[arg.size()];
            std::copy(arg.begin(),arg.end(),cArray);
            cMap->values[cMap->numCols++] = cArray;
        } else if constexpr (std::is_same_v<T, LogicalVector>) {
            auto cArray = new char[arg.size()];
            std::copy(arg.begin(),arg.end(),cArray);
            cMap->values[cMap->numCols++] = cArray;
        } else {
                std::cout << "Error" << std::endl;
            }



        }, value);;
    };
    if ((headerM).size() != cMap->numCols) {
        cMap->error = (std::string("ColSize does not match header: ") + std::to_string(cMap->numCols) + "/" +
                       std::to_string((headerM).size())).c_str();
    }
    return cMap;
}

ChromatogramInfo *getChromatogramInfo(MSDataFile file, int chromIdx) {
    auto ffile = (pwiz::msdata::MSDataFile *) file;
    auto info = new ChromatogramInfo{};
    auto chromListPtr = ffile->run.chromatogramListPtr;
    if (chromListPtr.get() == 0) {
        info->error = "The direct support for chromatogram info is only available in mzML format.";
        return info;
    } else if (chromListPtr->size() == 0) {
        info->error = "No available chromatogram info.";
        return info;
    } else if ((chromIdx < 0) || (chromIdx > chromListPtr->size())) {
        info->error = "Index whichChrom out of bounds [0 ... %d].\n", (chromListPtr->size()) - 1;
        return info;
    } else {

        auto chrom = chromListPtr->chromatogram(chromIdx, true);
        std::vector<pwiz::msdata::TimeIntensityPair> pairs;
        chrom->getTimeIntensityPairs(pairs);
        info->time = new double[pairs.size()];
        info->intensity = new double[pairs.size()];
        for (int i = 0; i < pairs.size(); i++) {
            auto p = pairs.at(i);
            info->time[i] = p.time;
            info->intensity[i] = p.intensity;
        }
        info->id = chrom->id.c_str();
        info->size = pairs.size();
    }
    return info;
}

Header *getChromatogramHeaderInfo(MSDataFile file, const int *scans, int scansSize) {
    auto ffile = (pwiz::msdata::MSDataFile *) file;

    // CVID nativeIdFormat_ = id::getDefaultNativeIDFormat(*msd);
    auto clp = ffile->run.chromatogramListPtr;
    if (clp.get() == 0) {
        std::cout << "The direct support for chromatogram cInfo is only available in mzML format.";
        return new Header ;
    } else if (clp->size() == 0) {
        std::cout << "No available chromatogram cInfo.";
        return new Header;
    }
    auto infoM = HeaderMap{
            {"chromatogramId",                      StringVector(scansSize)},
            {"chromatogramIndex",                   IntegerVector(scansSize)},
            {"polarity",                            IntegerVector(scansSize)},
            {"precursorIsolationWindowTargetMZ",    NumericVector(scansSize)},
            {"precursorIsolationWindowLowerOffset", NumericVector(scansSize)},
            {"precursorIsolationWindowUpperOffset", NumericVector(scansSize)},
            {"precursorCollisionEnergy",            NumericVector(scansSize)},
            {"productIsolationWindowTargetMZ",      NumericVector(scansSize)},
            {"productIsolationWindowLowerOffset",   NumericVector(scansSize)},
            {"productIsolationWindowUpperOffset",   NumericVector(scansSize)},
    };
    for (int i = 0; i < scansSize; i++) {
        int current_chrom = scans[i];
        if (current_chrom < 0 || current_chrom > clp->size()) {
            std::cout << "Provided index out of bounds.";
            return new Header ;
        }
        auto ch = clp->chromatogram(current_chrom, false);
        std::get<StringVector>((infoM)["chromatogramId"])[i] = ch->id;
        std::get<IntegerVector>((infoM)["chromatogramIndex"])[i] = current_chrom;
        auto param = ch->cvParamChild(pwiz::cv::MS_scan_polarity);
        std::get<IntegerVector>((infoM)["polarity"])[i] = (param.cvid == pwiz::cv::MS_negative_scan ? 0 : (
                param.cvid == pwiz::cv::MS_positive_scan ? +1 : -1));
        if (!ch->precursor.empty()) {
            std::get<NumericVector>(
                    (infoM)["precursorIsolationWindowTargetMZ"])[i] = ch->precursor.isolationWindow.cvParam(
                    pwiz::cv::MS_isolation_window_target_m_z).value.empty() ? NAN
                                                                            : ch->precursor.isolationWindow.cvParam(
                            pwiz::cv::MS_isolation_window_target_m_z).valueAs<double>();
            std::get<NumericVector>(
                    (infoM)["precursorIsolationWindowLowerOffset"])[i] = ch->precursor.isolationWindow.cvParam(
                    pwiz::cv::MS_isolation_window_lower_offset).value.empty() ? NAN
                                                                              : ch->precursor.isolationWindow.cvParam(
                            pwiz::cv::MS_isolation_window_lower_offset).valueAs<double>();
            std::get<NumericVector>(
                    (infoM)["precursorIsolationWindowUpperOffset"])[i] = ch->precursor.isolationWindow.cvParam(
                    pwiz::cv::MS_isolation_window_upper_offset).value.empty() ? NAN
                                                                              : ch->precursor.isolationWindow.cvParam(
                            pwiz::cv::MS_isolation_window_upper_offset).valueAs<double>();
            std::get<NumericVector>((infoM)["precursorCollisionEnergy"])[i] = ch->precursor.activation.cvParam(
                    pwiz::cv::MS_collision_energy).value.empty() ? NAN : ch->precursor.activation.cvParam(
                    pwiz::cv::MS_collision_energy).valueAs<double>();
        } else {
            std::get<NumericVector>((infoM)["precursorIsolationWindowTargetMZ"])[i] = NAN;
            std::get<NumericVector>((infoM)["precursorIsolationWindowLowerOffset"])[i] = NAN;
            std::get<NumericVector>((infoM)["precursorIsolationWindowUpperOffset"])[i] = NAN;
            std::get<NumericVector>((infoM)["precursorCollisionEnergy"])[i] = NAN;
        }
        if (!ch->product.empty()) {
            std::get<NumericVector>(
                    (infoM)["productIsolationWindowTargetMZ"])[i] = ch->product.isolationWindow.cvParam(
                    pwiz::cv::MS_isolation_window_target_m_z).value.empty() ? NAN : ch->product.isolationWindow.cvParam(
                    pwiz::cv::MS_isolation_window_target_m_z).valueAs<double>();
            std::get<NumericVector>(
                    (infoM)["productIsolationWindowLowerOffset"])[i] = ch->product.isolationWindow.cvParam(
                    pwiz::cv::MS_isolation_window_lower_offset).value.empty() ? NAN
                                                                              : ch->product.isolationWindow.cvParam(
                            pwiz::cv::MS_isolation_window_lower_offset).valueAs<double>();
            std::get<NumericVector>(
                    (infoM)["productIsolationWindowUpperOffset"])[i] = ch->product.isolationWindow.cvParam(
                    pwiz::cv::MS_isolation_window_upper_offset).value.empty() ? NAN
                                                                              : ch->product.isolationWindow.cvParam(
                            pwiz::cv::MS_isolation_window_upper_offset).valueAs<double>();
        } else {
            std::get<NumericVector>((infoM)["productIsolationWindowTargetMZ"])[i] = NAN;
            std::get<NumericVector>((infoM)["productIsolationWindowLowerOffset"])[i] = NAN;
            std::get<NumericVector>((infoM)["productIsolationWindowUpperOffset"])[i] = NAN;
        }
    }
    auto cInfo = convertHeader(infoM);
    cInfo->numRows = std::get<StringVector>((infoM)["chromatogramId"]).size();
    return cInfo;
}

int getAcquisitionNumber(MSDataFile file, std::string id, size_t index) {
    // const SpectrumIdentity& si = msd->run.spectrumListPtr->spectrumIdentity(index);
    auto scanNumber = pwiz::msdata::id::translateNativeIDToScanNumber(
            pwiz::msdata::id::getDefaultNativeIDFormat(*(pwiz::msdata::MSDataFile *) file), id);
    if (scanNumber.empty())
        return static_cast<int>(index) + 1;
    else
        return boost::lexical_cast<int>(scanNumber);
}

const char* getRunStartTimeStamp(MSDataFile file) {
    auto ffile = (pwiz::msdata::MSDataFile *) file;
    return ffile->run.startTimeStamp.c_str();

}



typedef std::vector<std::vector<double>> Matrix;

PeakList *getPeakList(MSDataFile file, int * scans, int size) {
    auto ffile = (pwiz::msdata::MSDataFile *) file;
    auto result = new PeakList{};
    const char * names[2] = {"mz","intensity"};
    result->colNames = names;
    result->colNum =2;

    auto slp = ffile->run.spectrumListPtr;

        int current_scan;
        auto res = std::vector<Matrix>(size);
        for (size_t i = 0; i < size; i++) {
            current_scan = scans[i];
            if (current_scan < 0 || current_scan >= slp->size()) {
                result->error = ("Index whichScan out of bounds [1 ..."+ std::to_string(size)+"%d].\n").c_str();
                return result;
            }
            size_t current_index = static_cast<size_t>(current_scan);
            auto sp = slp->spectrum(current_index, pwiz::msdata::DetailLevel_FullData);
            auto mzs = sp->getMZArray();
            auto ints = sp->getIntensityArray();
            if (!mzs.get() || !ints.get()) {
                res.at(i) = Matrix(2);
                continue;
            }
            if (mzs->data.size() != ints->data.size())
                result->error = "Sizes of mz and intensity arrays don't match.";
            auto  data_matrix = Matrix {mzs->data,ints->data};
            res.at(i) = data_matrix;
        }


        result->scanNum = size;
        result->values = new double**[size];
        result->valSizes = new int[size];
        result->scans = scans;
        for (int i=0; i<size; i++) {
            auto valSize = res[i][0].size();
            result->valSizes[i] = valSize;
            result->values[i]=new double*[2];
            result->values[i][0] = new double[valSize];
            result->values[i][1] = new double[valSize];
            for (int j=0; j<valSize; j++) {
                result->values[i][0][j]=res.at(i).at(0).at(j);
                result->values[i][1][j]=res.at(i).at(1).at(j);
            }
        }

        return result;
}



Map3d get3DMap (MSDataFile file, int * scans, int scanSize, double whichMzLow, double whichMzHigh, double resMz )
{
        auto ffile = (pwiz::msdata::MSDataFile *) file;


      auto slp = ffile->run.spectrumListPtr;
      double f = 1 / resMz;
      int low = round(whichMzLow * f);
      int high = round(whichMzHigh * f);
      int dmz = high - low + 1;

      Map3d map3d;
      map3d.values = new double*[scanSize];
      map3d.scans = scans;
      map3d.scanSize = scanSize;
      map3d.valueSize = dmz;
      for (int i = 0; i < scanSize; i++)
        {
          map3d.values[i]=new double[dmz];
	  for (int j = 0; j < dmz; j++)
            {
	      map3d.values[i][j] = 0.0;
            }
        }

      int j=0;
      for (int i = 0; i < scanSize; i++)
        {
	  auto s = slp->spectrum(scans[i], pwiz::msdata::DetailLevel_FullData);
	  std::vector<pwiz::msdata::MZIntensityPair> pairs;
	  s->getMZIntensityPairs(pairs);

	  for (int k=0; k < pairs.size(); k++)
            {
	      auto p = pairs.at(k);
	      j = round(p.mz * f) - low;
	      if ((j >= 0) & (j < dmz))
                {
		  if (p.intensity > map3d.values[i][j])
                    {
		      map3d.values[i][j] = p.intensity;
                    }
                }
            }
        }
      return(map3d);
}