# include "cpwiz.h"
# include "pwiz/data/msdata/MSDataFile.hpp"
# include "pwiz/data/msdata/LegacyAdapter.hpp"
# include "pwiz/data/common/CVTranslator.hpp"
# include <variant>
# include <iostream>

#define _GLIBCXX_USE_CXX11_ABI 0

void deletePeakList(PeakList *list) {
    for (int i = 0; i < list->scanNum; i++) {
        delete[] list->values[i][0];
        delete[] list->values[i][1];
        delete[] list->values[i];
    }
    delete[] list->valSizes;
    delete[] list->values;
    delete[] list->colNames;
    delete list;
}

void delete3DMap(Map3d *map) {
    for (int i = 0; i < map->scanSize; i++) {
        delete[] map->values[i];
    }
    delete[] map->values;
    delete map;
}

void deleteChromatogramInfo(ChromatogramInfo *info) {
    delete[] info->intensity;
    delete[] info->time;
    delete info;
}

void deleteIsolationWindow(IsolationWindows *windows) {
    delete[] windows->high;
    delete[] windows->low;
    delete windows;
}

void deleteChromatogramHeader(ChromatogramHeader *header) {
    delete[] header->values.chromatogramId;
    delete[] header->values.chromatogramIndex;
    delete[] header->values.polarity;
    delete[] header->values.precursorIsolationWindowTargetMZ;
    delete[] header->values.precursorIsolationWindowLowerOffset;
    delete[] header->values.precursorIsolationWindowUpperOffset;
    delete[] header->values.precursorCollisionEnergy;
    delete[] header->values.productIsolationWindowTargetMZ;
    delete[] header->values.productIsolationWindowLowerOffset;
    delete[] header->values.productIsolationWindowUpperOffset;
    delete header;
}

void deleteScanHeader(ScanHeader *header) {
    delete[] header->values.seqNum;
    delete[] header->values.acquisitionNum;
    delete[] header->values.msLevel;
    delete[] header->values.polarity;
    delete[] header->values.peaksCount;
    delete[] header->values.totIonCurrent;
    delete[] header->values.retentionTime;
    delete[] header->values.basePeakMZ;
    delete[] header->values.basePeakIntensity;
    delete[] header->values.collisionEnergy;
    delete[] header->values.ionisationEnergy;
    delete[] header->values.lowMZ;
    delete[] header->values.highMZ;
    delete[] header->values.precursorScanNum;
    delete[] header->values.precursorMZ;
    delete[] header->values.precursorCharge;
    delete[] header->values.precursorIntensity;
    delete[] header->values.mergedScan;
    delete[] header->values.mergedResultScanNum;
    delete[] header->values.mergedResultStartScanNum;
    delete[] header->values.mergedResultEndScanNum;
    delete[] header->values.ionInjectionTime;
    delete[] header->values.filterString;
    delete[] header->values.spectrumId;
    delete[] header->values.centroided;
    delete[] header->values.ionMobilityDriftTime;
    delete[] header->values.isolationWindowTargetMZ;
    delete[] header->values.isolationWindowLowerOffset;
    delete[] header->values.isolationWindowUpperOffset;
    delete[] header->values.scanWindowLowerLimit;
    delete[] header->values.scanWindowUpperLimit;
    delete header;
}

int getAcquisitionNumber(MSDataFile file, const std::string &id, size_t index);

MSDataFile MSDataOpenFile(const char *fileName, const char **errorMessage) {
    try {
        auto ms = new pwiz::msdata::MSDataFile(std::string(fileName));
        return ms;
    } catch (std::runtime_error &e) {
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

IsolationWindows *getIsolationWindow(MSDataFile file) {
    auto ffile = (pwiz::msdata::MSDataFile *) file;
    auto SpectrumListP = ffile->run.spectrumListPtr;
    auto cIWindows = new IsolationWindows{};
    auto scanSize = SpectrumListP->size();
    cIWindows->high = new double[scanSize];
    cIWindows->low = new double[scanSize];
    int ms2Count = 0;
    for (int i = 0; i < scanSize; i++) {
        auto SpectrumP = SpectrumListP->spectrum(i, pwiz::msdata::DetailLevel_FullMetadata);
        if (!SpectrumP->precursors.empty()) {
            auto iwin = SpectrumP->precursors[0].isolationWindow;
            cIWindows->high[ms2Count] = iwin.cvParam(pwiz::cv::MS_isolation_window_upper_offset).value.empty() ? NAN :
                                        iwin.cvParam(pwiz::cv::MS_isolation_window_upper_offset).valueAs<double>();
            cIWindows->low[ms2Count] = iwin.cvParam(pwiz::cv::MS_isolation_window_lower_offset).value.empty() ? NAN :
                                       iwin.cvParam(pwiz::cv::MS_isolation_window_lower_offset).valueAs<double>();
            ms2Count++;

        }
    }
    cIWindows->size = ms2Count;
    return cIWindows;
}

ScanHeader *getScanHeaderInfo(MSDataFile file, const int *scans, int scansSize) {
    auto ffile = (pwiz::msdata::MSDataFile *) file;
    auto header = new ScanHeader{
            {
                    new int[scansSize],
                    new int[scansSize],
                    new int[scansSize],
                    new int[scansSize],
                    new int[scansSize],
                    new double[scansSize],
                    new double[scansSize],
                    new double[scansSize],
                    new double[scansSize],
                    new double[scansSize],
                    new double[scansSize],
                    new double[scansSize],
                    new double[scansSize],
                    new int[scansSize],
                    new double[scansSize],
                    new int[scansSize],
                    new double[scansSize],
                    new int[scansSize],
                    new int[scansSize],
                    new int[scansSize],
                    new int[scansSize],
                    new double[scansSize],
                    new const char *[scansSize],
                    new const char *[scansSize],
                    new char[scansSize],
                    new double[scansSize],
                    new double[scansSize],
                    new double[scansSize],
                    new double[scansSize],
                    new double[scansSize],
                    new double[scansSize],
            },
            scansSize
    };

    auto SpectrumListP = ffile->run.spectrumListPtr;
    for (size_t i = 0; i < scansSize; i++) {
        int current_scan = scans[i];
        auto current_index = static_cast<size_t>(current_scan);
        auto SpectrumP = SpectrumListP->spectrum(current_index, pwiz::msdata::DetailLevel_FullMetadata);
        auto &scan = SpectrumP->scanList.scans[0];
        header->values.seqNum[i] = current_scan;
        header->values.acquisitionNum[i] =
                getAcquisitionNumber(file, SpectrumP->id, current_index);
        header->values.spectrumId[i] = SpectrumP->id.c_str();
        header->values.msLevel[i] =
                SpectrumP->cvParam(pwiz::cv::MS_ms_level).valueAs<int>();
        header->values.peaksCount[i] = static_cast<int>(SpectrumP->defaultArrayLength);
        header->values.totIonCurrent[i] =
                SpectrumP->cvParam(pwiz::cv::MS_total_ion_current).valueAs<double>();
        header->values.basePeakMZ[i] =
                SpectrumP->cvParam(pwiz::cv::MS_base_peak_m_z).valueAs<double>();
        header->values.basePeakIntensity[i] =
                SpectrumP->cvParam(pwiz::cv::MS_base_peak_intensity).valueAs<double>();
        header->values.ionisationEnergy[i] =
                SpectrumP->cvParam(pwiz::cv::MS_ionization_energy_OBSOLETE).valueAs<double>();
        header->values.lowMZ[i] =
                SpectrumP->cvParam(pwiz::cv::MS_lowest_observed_m_z).valueAs<double>();
        header->values.highMZ[i] =
                SpectrumP->cvParam(pwiz::cv::MS_highest_observed_m_z).valueAs<double>();
        auto param = SpectrumP->cvParamChild(pwiz::cv::MS_scan_polarity);
        header->values.polarity[i] =
                param.cvid == pwiz::cv::MS_negative_scan ? 0 : (param.cvid == pwiz::cv::MS_positive_scan ? +1 : -1);
        param = SpectrumP->cvParamChild(pwiz::cv::MS_spectrum_representation);
        header->values.centroided[i] =
                param.cvid == pwiz::cv::MS_centroid_spectrum ? 1 : 0;
        header->values.retentionTime[i] =
                scan.cvParam(pwiz::cv::MS_scan_start_time).timeInSeconds();
        header->values.ionInjectionTime[i] =
                (scan.cvParam(pwiz::cv::MS_ion_injection_time).timeInSeconds() * 1000);
        header->values.filterString[i] =
                scan.cvParam(pwiz::cv::MS_filter_string).value.empty() ? "" :
                scan.cvParam(pwiz::cv::MS_filter_string).value.c_str();
        header->values.ionMobilityDriftTime[i] =
                scan.cvParam(pwiz::cv::MS_ion_mobility_drift_time).value.empty() ? NAN : (
                        scan.cvParam(pwiz::cv::MS_ion_mobility_drift_time).timeInSeconds() * 1000);
        if (!scan.scanWindows.empty()) {
            header->values.scanWindowLowerLimit[i] =
                    scan.scanWindows[0].cvParam(pwiz::cv::MS_scan_window_lower_limit).valueAs<double>();
            header->values.scanWindowUpperLimit[i] =
                    scan.scanWindows[0].cvParam(pwiz::cv::MS_scan_window_upper_limit).valueAs<double>();
        } else {
            header->values.scanWindowLowerLimit[i] = NAN;
            header->values.scanWindowUpperLimit[i] = NAN;
        }
        header->values.mergedScan[i] = -1;
        header->values.mergedResultScanNum[i] = -1;
        header->values.mergedResultStartScanNum[i] = -1;
        header->values.mergedResultEndScanNum[i] = -1;

        const auto &precursor = !SpectrumP->precursors.empty() ? SpectrumP->precursors[0] : pwiz::msdata::Precursor{};
        header->values.collisionEnergy[i] =
                precursor.activation.cvParam(pwiz::cv::MS_collision_energy).valueAs<double>();
        size_t precursorIndex = SpectrumListP->find(precursor.spectrumID);
        if (precursorIndex < SpectrumListP->size()) {
            header->values.precursorScanNum[i] =
                    getAcquisitionNumber(file, precursor.spectrumID, precursorIndex);
        } else {
            header->values.precursorScanNum[i] = -1;
        }
        const auto &selectedIon = !precursor.selectedIons.empty() ? precursor.selectedIons[0]
                                                                  : pwiz::msdata::SelectedIon{};
        header->values.precursorMZ[i] =
                selectedIon.cvParam(pwiz::cv::MS_selected_ion_m_z).value.empty()
                ? selectedIon.cvParam(pwiz::cv::MS_m_z).valueAs<double>()
                : selectedIon.cvParam(pwiz::cv::MS_selected_ion_m_z).valueAs<double>();
        header->values.precursorCharge[i] =
                selectedIon.cvParam(pwiz::cv::MS_charge_state).valueAs<int>();
        header->values.precursorIntensity[i] =
                selectedIon.cvParam(pwiz::cv::MS_peak_intensity).valueAs<double>();

        auto iwin = !SpectrumP->precursors.empty() ? SpectrumP->precursors[0].isolationWindow
                                                   : pwiz::msdata::IsolationWindow{};
        header->values.isolationWindowTargetMZ[i] =
                iwin.cvParam(pwiz::cv::MS_isolation_window_target_m_z).value.empty() ? NAN
                                                                                     : iwin.cvParam(
                        pwiz::cv::MS_isolation_window_target_m_z).valueAs<double>();
        header->values.isolationWindowLowerOffset[i] =
                iwin.cvParam(pwiz::cv::MS_isolation_window_lower_offset).value.empty() ? NAN
                                                                                       : iwin.cvParam(
                        pwiz::cv::MS_isolation_window_lower_offset).valueAs<double>();
        header->values.isolationWindowUpperOffset[i] =
                iwin.cvParam(pwiz::cv::MS_isolation_window_upper_offset).value.empty() ? NAN
                                                                                       : iwin.cvParam(
                        pwiz::cv::MS_isolation_window_upper_offset).valueAs<double>();
    }

    return header;
}

ChromatogramInfo *getChromatogramInfo(MSDataFile file, int chromIdx) {
    auto ffile = (pwiz::msdata::MSDataFile *) file;
    auto info = new ChromatogramInfo{};
    auto chromListPtr = ffile->run.chromatogramListPtr;
    if (chromListPtr.get() == nullptr) {
        info->error = "The direct support for chromatogram info is only available in mzML format.";
        return info;
    } else if (chromListPtr->empty()) {
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


ChromatogramHeader *getChromatogramHeaderInfo(MSDataFile file, const int *scans, int scansSize) {
    auto ffile = (pwiz::msdata::MSDataFile *) file;

    // CVID nativeIdFormat_ = id::getDefaultNativeIDFormat(*msd);
    auto clp = ffile->run.chromatogramListPtr;
    if (clp.get() == nullptr) {
        std::cout << "The direct support for chromatogram cInfo is only available in mzML format.";
        return new ChromatogramHeader;
    } else if (clp->empty()) {
        std::cout << "No available chromatogram cInfo.";
        return new ChromatogramHeader;
    }
    auto info = new ChromatogramHeader{
            {
                    new const char *[scansSize],
                    new int[scansSize],
                    new int[scansSize],
                    new double[scansSize],
                    new double[scansSize],
                    new double[scansSize],
                    new double[scansSize],
                    new double[scansSize],
                    new double[scansSize],
                    new double[scansSize],
            },
            scansSize,
    };
    for (int i = 0; i < scansSize; i++) {
        int current_chrom = scans[i];
        if (current_chrom < 0 || current_chrom > clp->size()) {
            std::cout << "Provided index out of bounds.";
            return new ChromatogramHeader;
        }
        auto ch = clp->chromatogram(current_chrom, false);
        info->values.chromatogramId[i] = ch->id.c_str();
        info->values.chromatogramIndex[i] = current_chrom;
        auto param = ch->cvParamChild(pwiz::cv::MS_scan_polarity);
        info->values.polarity[i] = (param.cvid == pwiz::cv::MS_negative_scan ? 0 : (
                param.cvid == pwiz::cv::MS_positive_scan ? +1 : -1));
        if (!ch->precursor.empty()) {

            info->values.precursorIsolationWindowTargetMZ[i] = ch->precursor.isolationWindow.cvParam(
                    pwiz::cv::MS_isolation_window_target_m_z).value.empty() ? NAN
                                                                            : ch->precursor.isolationWindow.cvParam(
                            pwiz::cv::MS_isolation_window_target_m_z).valueAs<double>();

            info->values.precursorIsolationWindowLowerOffset[i] = ch->precursor.isolationWindow.cvParam(
                    pwiz::cv::MS_isolation_window_lower_offset).value.empty() ? NAN
                                                                              : ch->precursor.isolationWindow.cvParam(
                            pwiz::cv::MS_isolation_window_lower_offset).valueAs<double>();

            info->values.precursorIsolationWindowUpperOffset[i] = ch->precursor.isolationWindow.cvParam(
                    pwiz::cv::MS_isolation_window_upper_offset).value.empty() ? NAN
                                                                              : ch->precursor.isolationWindow.cvParam(
                            pwiz::cv::MS_isolation_window_upper_offset).valueAs<double>();
            info->values.precursorCollisionEnergy[i] = ch->precursor.activation.cvParam(
                    pwiz::cv::MS_collision_energy).value.empty() ? NAN : ch->precursor.activation.cvParam(
                    pwiz::cv::MS_collision_energy).valueAs<double>();
        } else {
            info->values.precursorIsolationWindowTargetMZ[i] = NAN;
            info->values.precursorIsolationWindowLowerOffset[i] = NAN;
            info->values.precursorIsolationWindowUpperOffset[i] = NAN;
            info->values.precursorCollisionEnergy[i] = NAN;
        }
        if (!ch->product.empty()) {
            info->values.productIsolationWindowTargetMZ[i] = ch->product.isolationWindow.cvParam(
                    pwiz::cv::MS_isolation_window_target_m_z).value.empty() ? NAN : ch->product.isolationWindow.cvParam(
                    pwiz::cv::MS_isolation_window_target_m_z).valueAs<double>();
            info->values.productIsolationWindowLowerOffset[i] = ch->product.isolationWindow.cvParam(
                    pwiz::cv::MS_isolation_window_lower_offset).value.empty() ? NAN
                                                                              : ch->product.isolationWindow.cvParam(
                            pwiz::cv::MS_isolation_window_lower_offset).valueAs<double>();
            info->values.productIsolationWindowUpperOffset[i] = ch->product.isolationWindow.cvParam(
                    pwiz::cv::MS_isolation_window_upper_offset).value.empty() ? NAN
                                                                              : ch->product.isolationWindow.cvParam(
                            pwiz::cv::MS_isolation_window_upper_offset).valueAs<double>();
        } else {
            info->values.productIsolationWindowTargetMZ[i] = NAN;
            info->values.productIsolationWindowLowerOffset[i] = NAN;
            info->values.productIsolationWindowUpperOffset[i] = NAN;
        }
    }
    return info;
}

int getAcquisitionNumber(MSDataFile file, const std::string &id, size_t index) {
    // const SpectrumIdentity& si = msd->run.spectrumListPtr->spectrumIdentity(index);
    auto scanNumber = pwiz::msdata::id::translateNativeIDToScanNumber(
            pwiz::msdata::id::getDefaultNativeIDFormat(*(pwiz::msdata::MSDataFile *) file), id);
    if (scanNumber.empty())
        return static_cast<int>(index) + 1;
    else
        return boost::lexical_cast<int>(scanNumber);
}

const char *getRunStartTimeStamp(MSDataFile file) {
    auto ffile = (pwiz::msdata::MSDataFile *) file;
    return ffile->run.startTimeStamp.c_str();

}


typedef std::vector<std::vector<double>> Matrix;

PeakList *getPeakList(MSDataFile file, int *scans, int size) {
    auto ffile = (pwiz::msdata::MSDataFile *) file;
    auto result = new PeakList{};
    result->colNames = new const char*[2]{"mz", "intensity"};
    result->colNum = 2;

    auto slp = ffile->run.spectrumListPtr;

    int current_scan;
    auto res = std::vector<Matrix>(size);
    for (size_t i = 0; i < size; i++) {
        current_scan = scans[i];
        if (current_scan < 0 || current_scan >= slp->size()) {
            result->error = ("Index whichScan out of bounds [1 ..." + std::to_string(size) + "%d].\n").c_str();
            return result;
        }
        auto current_index = static_cast<size_t>(current_scan);
        auto sp = slp->spectrum(current_index, pwiz::msdata::DetailLevel_FullData);
        auto mzs = sp->getMZArray();
        auto ints = sp->getIntensityArray();
        if (!mzs.get() || !ints.get()) {
            res.at(i) = Matrix(2);
            continue;
        }
        if (mzs->data.size() != ints->data.size())
            result->error = "Sizes of mz and intensity arrays don't match.";
        auto data_matrix = Matrix{mzs->data, ints->data};
        res.at(i) = data_matrix;
    }


    result->scanNum = size;
    result->values = new double **[size];
    result->valSizes = new unsigned long int[size];
    result->scans = scans;
    for (int i = 0; i < size; i++) {
        auto valSize = res[i][0].size();
        result->valSizes[i] = valSize;
        result->values[i] = new double *[2];
        result->values[i][0] = new double[valSize];
        result->values[i][1] = new double[valSize];
        for (int j = 0; j < valSize; j++) {
            result->values[i][0][j] = res.at(i).at(0).at(j);
            result->values[i][1][j] = res.at(i).at(1).at(j);
        }
    }

    return result;
}


Map3d *get3DMap(MSDataFile file, int *scans, int scanSize, double whichMzLow, double whichMzHigh, double resMz) {
    auto ffile = (pwiz::msdata::MSDataFile *) file;


    auto slp = ffile->run.spectrumListPtr;
    double f = 1 / resMz;
    int low = round(whichMzLow * f);
    int high = round(whichMzHigh * f);
    int dmz = high - low + 1;

    auto *map3d = new Map3d;
    map3d->values = new double *[scanSize];
    map3d->scans = scans;
    map3d->scanSize = scanSize;
    map3d->valueSize = dmz;
    for (int i = 0; i < scanSize; i++) {
        map3d->values[i] = new double[dmz];
        for (int j = 0; j < dmz; j++) {
            map3d->values[i][j] = 0.0;
        }
    }

    int j;
    for (int i = 0; i < scanSize; i++) {
        auto s = slp->spectrum(scans[i], pwiz::msdata::DetailLevel_FullData);
        std::vector<pwiz::msdata::MZIntensityPair> pairs;
        s->getMZIntensityPairs(pairs);

        for (auto p: pairs) {
            j = round(p.mz * f) - low;
            if ((j >= 0) & (j < dmz)) {
                if (p.intensity > map3d->values[i][j]) {
                    map3d->values[i][j] = p.intensity;
                }
            }
        }
    }
    return map3d;
}