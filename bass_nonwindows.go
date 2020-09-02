//  +build !windows

package bass

/*

#include "bass.h"
#include "stdlib.h"
*/
import "C"


import (
	"unsafe"
)
func Init(device int, freq int, flags int, hwnd, clsid uintptr) (bool, error) {
	window:=unsafe.Pointer(hwnd)
	gid:=unsafe.Pointer(clsid)
	if C.BASS_Init(C.int(device), C.DWORD(freq), C.DWORD(flags), window, gid) != 0 {
		return true, nil
	} else {
		return false, errMsg()
	}
}



