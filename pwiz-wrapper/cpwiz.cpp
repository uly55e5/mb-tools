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


int getAcquisitionNumber(MSDataFile file,std::string id, size_t index);



MSDataFile MSDataOpenFile(const char *fileName) {
    auto ms = new pwiz::msdata::MSDataFile(std::string(fileName));
    return ms;
}

void MSDataClose(MSDataFile file) {
    delete (pwiz::msdata::MSDataFile *) file;
}

int getLastChrom(MSDataFile file) {
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

Header getScanHeaderInfo(MSDataFile file, const int *scans, int scansSize) {
    auto ffile = (pwiz::msdata::MSDataFile *) file;
    auto (*headerM) = new HeaderMap {
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
        std::cout << "Current index: " << current_index;
        auto SpectrumP = SpectrumListP->spectrum(current_index, pwiz::msdata::DetailLevel_FullMetadata);
        auto &scan = SpectrumP->scanList.scans[0];
        std::get<IntegerVector>((*headerM)["seqNum"]).push_back(current_scan);
        std::get<IntegerVector>((*headerM)["acquisitionNum"]).push_back(getAcquisitionNumber(file,SpectrumP->id, current_index));
        std::get<StringVector>((*headerM)["spectrumId"]).push_back(SpectrumP->id);
        std::get<IntegerVector>((*headerM)["msLevel"]).push_back(SpectrumP->cvParam(pwiz::cv::MS_ms_level).valueAs<int>());
        std::get<IntegerVector>((*headerM)["peaksCount"]).push_back(static_cast<int>(SpectrumP->defaultArrayLength));
        std::get<NumericVector>((*headerM)["totIonCurrent"]).push_back(
                SpectrumP->cvParam(pwiz::cv::MS_total_ion_current).valueAs<double>());
        std::get<NumericVector>((*headerM)["basePeakMZ"]).push_back(
                SpectrumP->cvParam(pwiz::cv::MS_base_peak_m_z).valueAs<double>());
        std::get<NumericVector>((*headerM)["basePeakIntensity"]).push_back(
                SpectrumP->cvParam(pwiz::cv::MS_base_peak_intensity).valueAs<double>());
        std::get<NumericVector>((*headerM)["ionisationEnergy"]).push_back(
                SpectrumP->cvParam(pwiz::cv::MS_ionization_energy_OBSOLETE).valueAs<double>());
        std::get<NumericVector>((*headerM)["lowMZ"]).push_back(
                SpectrumP->cvParam(pwiz::cv::MS_lowest_observed_m_z).valueAs<double>());
        std::get<NumericVector>((*headerM)["highMZ"]).push_back(
                SpectrumP->cvParam(pwiz::cv::MS_highest_observed_m_z).valueAs<double>());
        auto param = SpectrumP->cvParamChild(pwiz::cv::MS_scan_polarity);
        std::get<IntegerVector>((*headerM)["polarity"]).push_back(
                param.cvid == pwiz::cv::MS_negative_scan ? 0 : (param.cvid == pwiz::cv::MS_positive_scan ? +1 : -1));
        param = SpectrumP->cvParamChild(pwiz::cv::MS_spectrum_representation);
        std::get<LogicalVector>((*headerM)["centroided"]).push_back(
                param.cvid == pwiz::cv::MS_centroid_spectrum);
        std::get<NumericVector>((*headerM)["retentionTime"]).push_back(
                scan.cvParam(pwiz::cv::MS_scan_start_time).timeInSeconds());
        std::get<NumericVector>((*headerM)["ionInjectionTime"]).push_back(
                (scan.cvParam(pwiz::cv::MS_ion_injection_time).timeInSeconds() * 1000));
        std::get<StringVector>((*headerM)["filterString"]).push_back(
                scan.cvParam(pwiz::cv::MS_filter_string).value.empty() ? "" :
                scan.cvParam(pwiz::cv::MS_filter_string).value);
        std::get<NumericVector>((*headerM)["ionMobilityDriftTime"]).push_back(
                scan.cvParam(pwiz::cv::MS_ion_mobility_drift_time).value.empty() ? NAN : (
                        scan.cvParam(pwiz::cv::MS_ion_mobility_drift_time).timeInSeconds() * 1000));
        if (!scan.scanWindows.empty()) {
            std::get<NumericVector>((*headerM)["scanWindowLowerLimit"]).push_back(
                    scan.scanWindows[0].cvParam(pwiz::cv::MS_scan_window_lower_limit).valueAs<double>());
            std::get<NumericVector>((*headerM)["scanWindowUpperLimit"]).push_back(
                    scan.scanWindows[0].cvParam(pwiz::cv::MS_scan_window_upper_limit).valueAs<double>());
        } else {
            std::get<NumericVector>((*headerM)["scanWindowLowerLimit"]).push_back(NAN);
            std::get<NumericVector>((*headerM)["scanWindowUpperLimit"]).push_back(NAN);
        }
        std::get<IntegerVector>((*headerM)["mergedScan"]).push_back(-1);
        std::get<IntegerVector>((*headerM)["mergedResultScanNum"]).push_back(-1);
        std::get<IntegerVector>((*headerM)["mergedResultStartScanNum"]).push_back(-1);
        std::get<IntegerVector>((*headerM)["mergedResultEndScanNum"]).push_back(-1);

            const auto &precursor = !SpectrumP->precursors.empty() ? SpectrumP->precursors[0] : pwiz::msdata::Precursor{};
            std::get<NumericVector>((*headerM)["collisionEnergy"]).push_back(
                    precursor.activation.cvParam(pwiz::cv::MS_collision_energy).valueAs<double>());
            size_t precursorIndex = SpectrumListP->find(precursor.spectrumID);
            if (precursorIndex < SpectrumListP->size()) {
                std::get<IntegerVector>((*headerM)["precursorScanNum"]).push_back(getAcquisitionNumber(file,precursor.spectrumID, precursorIndex));
            } else {
                std::get<IntegerVector>((*headerM)["precursorScanNum"]).push_back(-1);
            }
            const auto &selectedIon = !precursor.selectedIons.empty() ?  precursor.selectedIons[0] : pwiz::msdata::SelectedIon{};
                std::get<NumericVector>((*headerM)["precursorMZ"]).push_back(
                        selectedIon.cvParam(pwiz::cv::MS_selected_ion_m_z).value.empty()
                        ? selectedIon.cvParam(pwiz::cv::MS_m_z).valueAs<double>()
                        : selectedIon.cvParam(pwiz::cv::MS_selected_ion_m_z).valueAs<double>());
                std::get<IntegerVector>((*headerM)["precursorCharge"]).push_back(
                        selectedIon.cvParam(pwiz::cv::MS_charge_state).valueAs<int>());
                std::get<NumericVector>((*headerM)["precursorIntensity"]).push_back(
                        selectedIon.cvParam(pwiz::cv::MS_peak_intensity).valueAs<double>());

            auto iwin = !SpectrumP->precursors.empty() ? SpectrumP->precursors[0].isolationWindow : pwiz::msdata::IsolationWindow{};
                std::get<NumericVector>((*headerM)["isolationWindowTargetMZ"]).push_back(
                        iwin.cvParam(pwiz::cv::MS_isolation_window_target_m_z).value.empty() ? NAN
                                                                                             : iwin.cvParam(
                                pwiz::cv::MS_isolation_window_target_m_z).valueAs<double>());
                std::get<NumericVector>((*headerM)["isolationWindowLowerOffset"]).push_back(
                        iwin.cvParam(pwiz::cv::MS_isolation_window_lower_offset).value.empty() ? NAN
                                                                                               : iwin.cvParam(
                                pwiz::cv::MS_isolation_window_lower_offset).valueAs<double>());
                std::get<NumericVector>((*headerM)["isolationWindowUpperOffset"]).push_back(
                        iwin.cvParam(pwiz::cv::MS_isolation_window_upper_offset).value.empty() ? NAN
                                                                                               : iwin.cvParam(
                                pwiz::cv::MS_isolation_window_upper_offset).valueAs<double>());
    }
    auto cMap = new Header;
    cMap->names = new const char*[(*headerM).size()];
    cMap->values = new const void *[(*headerM).size()];
    cMap->numCols =0;
    cMap->numRows = std::get<IntegerVector>((*headerM)["seqNum"]).size();
    cMap->error = "";
    for(const auto& [key,value] : (*headerM)){
        cMap->names[cMap->numCols] = key.c_str();
        std::visit([key,&cMap](auto&& arg){
            if(arg.size() != cMap->numRows) {
                cMap->error = (std::string("ColSize does not match: ")+key+" : " + std::to_string(arg.size())+"/"+std::to_string(cMap->numRows)).c_str();
            }
            cMap->values[cMap->numCols++] = arg.data();},value);
        ;
    };
    if((*headerM).size() != cMap->numCols) {
        cMap->error = (std::string("ColSize does not match header: ") + std::to_string(cMap->numCols)+"/"+std::to_string((*headerM).size())).c_str();
    }
    return *cMap;
}

int getAcquisitionNumber(MSDataFile file,std::string id, size_t index)
{
  // const SpectrumIdentity& si = msd->run.spectrumListPtr->spectrumIdentity(index);
  auto scanNumber = pwiz::msdata::id::translateNativeIDToScanNumber(pwiz::msdata::id::getDefaultNativeIDFormat(*(pwiz::msdata::MSDataFile *)file), id);
  if (scanNumber.empty())
    return static_cast<int>(index) + 1;
  else
    return boost::lexical_cast<int>(scanNumber);
}

