package bass

/*
#include "bass.h"
*/
import "C"

import (
)
type FX struct {
	handle C.DWORD
}
func (self FX) ToError() (FX, error) {
	if self.handle == 0 {
		return FX{}, GetLastError()
	} else {
		return self, nil
	}
}
func (self FX) GetHandle() uint32 {
	return uint32(self.handle)
}
func (self FX) cint() C.DWORD {
	return C.DWORD(self.GetHandle())
}
func FXFromHandle(handle uint32) FX {
	return FX{handle: C.DWORD(handle)}
}
func (self Channel) SetFX(effectType, priority int) (FX, error) {
	return FXFromHandle(uint32(C.BASS_ChannelSetFX(self.cint(), C.DWORD(effectType), C.int(priority)))).ToError()
}
func (self Channel) RemoveFX(fx FX) error {
	return boolToError(C.BASS_ChannelRemoveFX(self.cint(), fx.cint()))
}