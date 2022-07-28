package mzmlReader

//#cgo CFLAGS: -I../pwiz-wrapper
//#cgo LDFLAGS: -L../pwiz-wrapper/lib -lpwiz_wrapper -L/home/david/Projekte/Reference/pwiz/build-linux-x86_64/pwiz/data/msdata/gcc-9/release/link-static/runtime-link-static/threading-multi -Wl,-rpath=../pwiz-wrapper/lib -lstdc++ -lpwiz_data_msdata
//#include "cpwiz.h"
import "C"

type MSData struct {
	msData C.MSDataFile
}

func OpenMSData(fileName string) {
	var file MSData
	file.msData = C.MSDataOpenFile(fileName)
}
