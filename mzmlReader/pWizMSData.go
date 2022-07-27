package mzmlReader

//#include "cpwiz.h"
import "C"

type MSData struct {
	msData C.MSData
}

func OpenMSData(fileName string) {
	var file MSData
	file.msData = C.ReadMSDataFile(fileName)
}
