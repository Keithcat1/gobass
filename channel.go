package bass

/*
#include "bass.h"
#include "stdlib.h"
*/
import "C"

import (
	"unsafe"
)
type Channel struct {
	handle C.DWORD
}
func (self Channel) cint() C.DWORD {
	return self.handle
}
func (self Channel) GetHandle() uint32 {
	return uint32(self.handle)
}
func (self Channel) toStream() Stream {
	return Stream{Channel: self}
}
// internal function, do not use
func (self Channel) ToError() (Channel, error) {
	if self.handle == 0 {
		return Channel{}, GetLastError()
	} else {
		return self, nil
	}
}
// somewhat unsafe function that creates a channel from a raw BASS handle.
func ChannelFromHandle(handle uint32) Channel {
	return Channel{handle: C.DWORD(handle)}
}
func (self Channel) SetSync(syncType int, flags Flags, param int, callback *C.SYNCPROC, userdata unsafe.Pointer) (Sync, error) {
	sync := Sync(C.BASS_ChannelSetSync(self.cint(), cuint(syncType), culong(param | int(flags)), callback, userdata))
	if sync == 0 {
return 0, errMsg()
	} else {
		return sync, nil
	}
}

// ChannelPlay
// BOOL BASSDEF(BASS_ChannelPlay)(DWORD handle, BOOL restart);
func (self Channel) Play(restart bool) error {
	return boolToError(C.BASS_ChannelPlay(self.cint(), boolToInt(restart)))
}

// ChannelPause
// BOOL BASSDEF(BASS_ChannelPause)(DWORD handle);
func (self Channel) Pause() error {
	return boolToError(C.BASS_ChannelPause(self.cint()))
}

// ChannelStop
// BOOL BASSDEF(BASS_ChannelStop)(DWORD handle);
func (self Channel) Stop() error {
	return boolToError(C.BASS_ChannelStop(self.cint()))
}

func (self Channel) IsActive() (int, error) {
	active := int(C.BASS_ChannelIsActive(self.cint()))
	if active==ACTIVE_STOPPED {
		return active, errMsg()
	} else {
		return active, nil
	}
}

// ChannelGetAttribute
// BOOL BASSDEF(BASS_ChannelGetAttribute)(DWORD handle, DWORD attrib, float *value);
func (self Channel) GetAttribute(attrib int) (float64, error) {
	var cvalue C.float
	result:=C.BASS_ChannelGetAttribute(self.cint(), C.DWORD(attrib), &cvalue)
	return float64(cvalue), boolToError(result)
}
// ChannelSetAttribute
// BOOL BASSDEF(BASS_ChannelSetAttribute)(DWORD handle, DWORD attrib, float value);
func (self Channel) SetAttribute(attrib int, value float64) error {
	return boolToError(C.BASS_ChannelSetAttribute(self.cint(), C.DWORD(attrib), C.float(value)))
}

//ChannelGetLevel
//DWORD BASSDEF(BASS_ChannelGetLevel)(DWORD handle);
func (self Channel) GetLevel() (c int, e error) {
	c = int(C.BASS_ChannelGetLevel(self.cint()))
	if c == -1 {
		return 0, errMsg()
	}
	return c, nil
}

// ChannelGetTags
// const char *BASSDEF(BASS_ChannelGetTags)(DWORD handle, DWORD tags);
func (self Channel) GetTags(tag int) string {
	return C.GoString(C.BASS_ChannelGetTags(self.cint(), C.DWORD(tag)))
}

func (self Channel) SlideAttribute(attrib uint64, value float64, time uint64) error {
	return boolToError(C.BASS_ChannelSlideAttribute(self.cint(), cuint(attrib), C.float(value), cuint(time)))
}

func (self Channel) SetPosition(pos, mode int) error {
	return boolToError(C.BASS_ChannelSetPosition(self.cint(), culong(pos), cuint(mode)))
}

func (self Channel) GetPosition(mode uint64) (int, error) {
	value := C.BASS_ChannelGetPosition(self.cint(), C.DWORD(mode))
	if value+1 == 0 {
		return 0, errMsg()
	} else {
		return int(value), nil
	}
}
func (self Channel) GetLength(mode uint64) (int, error) {
	result := C.BASS_ChannelGetLength(self.cint(), C.DWORD(mode))
	if result +1 == 0 {
		return 0, errMsg()
	} else {
		return int(result), nil
	}
}


func (self Channel) Bytes2Seconds(bytes int) (float64, error) {
	value:=float64(C.BASS_ChannelBytes2Seconds(self.cint(), C.QWORD(bytes)))
	if value<0 {
		return value, errMsg()
	} else {
		return value, nil
	}
}
func (self Channel) Seconds2Bytes(seconds float64) (int, error) {
	res := int(C.BASS_ChannelSeconds2Bytes(self.cint(), C.double(seconds)))
	if res == -1 {
		return 0, GetLastError()
	}
	return res, nil
}
func (self Channel) Flags(flags, mask Flags) (Flags, error) {
	return Flags(C.BASS_ChannelFlags(self.cint(), cuint(flags), cuint(mask))), errMsg()
}

func (self Channel) GetData(data []byte, flags Flags) (int, error) {
	var ptr unsafe.Pointer
	if len(data)>0 {
		ptr = unsafe.Pointer(&data[0])
	}
	val := C.BASS_ChannelGetData(self.cint(), ptr, C.DWORD(len(data)|int(flags)))
	if val +1 == 0 { // -1 indicates an error, but this is an unsigned integer
		return 0, errMsg()
	} else {
		return int(val), nil
	}
}
func (self Channel) Free() error {
	return boolToError(C.BASS_ChannelFree(self.cint()))
}
