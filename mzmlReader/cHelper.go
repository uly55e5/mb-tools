package mzmlReader

import "C"
import (
	"reflect"
	"unsafe"
)

func cArray2GoSliceInt(array *C.int, length int) []int {
	var cSlice []C.int
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&cSlice))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(array))
	var gSlice []int
	for _, ci := range cSlice {
		gSlice = append(gSlice, int(ci))
	}
	return gSlice
}

func cArray2GoSliceULongInt(array *C.ulong, length int) []int {
	var arrayPtr *C.ulong = array
	cSlice := unsafe.Slice(arrayPtr, length)
	var gSlice []int
	for _, ci := range cSlice {
		gSlice = append(gSlice, int(ci))
	}
	return gSlice
}

func cArray2GoSliceBool(array *C.char, length int) []bool {
	var cSlice []C.char
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&cSlice))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(array))
	var gSlice []bool
	for _, ci := range cSlice {
		gSlice = append(gSlice, uint8(ci) > 0)
	}
	return gSlice
}

func cArray2GoSliceDouble(array *C.double, length int) []float64 {
	var cSlice []C.double
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&cSlice))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(array))
	var gSlice []float64
	for _, ci := range cSlice {
		gSlice = append(gSlice, float64(ci))
	}
	return gSlice
}

func cArray2GoSliceStr(array **C.char, length int) []string {
	cSlice := unsafe.Slice(array, length)
	var gSlice []string
	for _, ci := range cSlice {
		gSlice = append(gSlice, C.GoString(ci))
	}
	return gSlice
}

func gSlice2CArrayInt(gSlice []int) (*C.int, int) {
	var cSlice []C.int
	for _, gi := range gSlice {
		cSlice = append(cSlice, C.int(gi))
	}
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&cSlice))
	return (*C.int)(unsafe.Pointer(sliceHeader.Data)), len(gSlice)
}
