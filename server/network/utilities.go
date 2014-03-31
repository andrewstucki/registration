package network

// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

func convertStringsToSlice(strings **C.char) []string {
	defer C.free(unsafe.Pointer(strings))	
	stringValues := (*[1 << 20]*C.char)(unsafe.Pointer(strings))
	count := 0
	var stringValue *C.char = stringValues[count]
	rStrings := make([]string, 0, 0)
	for stringValue != nil {
		rStrings = append(rStrings, C.GoString(stringValue))
		count++
		C.free(unsafe.Pointer(stringValue))
		stringValue = stringValues[count]
	}
	return rStrings
}
