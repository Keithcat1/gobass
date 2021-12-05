package bass

/*
#include "bass.h"
#include "stdlib.h"
*/
import "C"

import (
	"unsafe"
)
type Sample struct {
	handle C.DWORD
}
func (self Sample) cint() C.DWORD {
	return self.handle
}
func SampleLoad(data interface{}, offset, max, flags Flags) (Sample, error) {
	var ch C.DWORD
	switch data := data.(type) {
	case CBytes:
		ch = C.BASS_SampleLoad(1, data.Data, culong(offset), cuint(data.Length), cuint(max), cuint(flags))
	case string:
		cstring := unsafe.Pointer(ToUTF16(data))
		ch = C.BASS_SampleLoad(0, cstring, culong(offset), 0, cuint(max), cuint(flags))
	case []byte:
		// According to BASS documentation, BASS_SampleLoad makes an internal copy of the passed-in memory, so we don't need to worry about CGO restrictions and can just pass a pointer to Go memory
		ch = C.BASS_SampleLoad(1, unsafe.Pointer(&data[0]), culong(offset), cuint(len(data)), cuint(max), cuint(flags))
	}
	return sampleToError(ch)
}
func (self Sample) Free() error {
	return boolToError(C.BASS_SampleFree(self.cint()))
}
func (self Sample) GetChannel(flags Flags) (Channel, error) {
	return channelToError(C.BASS_SampleGetChannel(self.cint(), C.DWORD(flags)))
}
func (self Sample) Stop() error {
	return boolToError(C.BASS_SampleStop(self.cint()))
}
