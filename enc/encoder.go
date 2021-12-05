package enc
import (
	"github.com/keithcat1/gobass"
)
/*
#include "bassenc.h"
*/
import "C"
type Encoder struct {
	handle uint32
}
func (self Encoder) cint() C.DWORD {
	return C.DWORD(self.handle)
}
func (self Encoder) GetHandle() uint32 {
	return self.handle
}
func (self Encoder) SetChannel(channel bass.Channel) error {
	res := C.BASS_Encode_SetChannel(self.cint(), C.DWORD(channel.GetHandle()))
	if res==0 {
		return bass.GetLastError()
	} else {
		return nil
	}
}
func EncoderFromHandle(handle uint32) Encoder {
	return Encoder{handle: handle}
}
func (self Encoder) Free() error {
	res := C.BASS_Encode_Stop(self.cint())
	if res==0 {
		return bass.GetLastError()
	} else {
		return nil
	}
}
func (self Encoder) ToError() (Encoder, error) {
	if self.cint() == 0 {
		return Encoder{}, bass.GetLastError()
	} else {
		return self, nil
	}
}