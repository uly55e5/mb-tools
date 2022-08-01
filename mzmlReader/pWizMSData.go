package mzmlReader

//#cgo CFLAGS: -I../pwiz-wrapper
//#cgo LDFLAGS: -L../pwiz-wrapper/lib -lpwiz_wrapper -Wl,-rpath=../pwiz-wrapper/lib -lstdc++ -lpwiz_all -lm -ldl
//#include "cpwiz.h"
//#define _GLIBCXX_USE_CXX11_ABI 0
import "C"

type MSData struct {
	msData C.MSDataFile
}

func OpenMSData(fileName string) {
	var file MSData
	file.msData = C.MSDataOpenFile(C.CString(fileName))
}
