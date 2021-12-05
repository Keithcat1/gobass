package bass

/*
#include "bass.h"
#include "stdlib.h"
*/
import "C"

import (
	"unsafe"
)
type Record struct {
	Channel
}
func (self Record) ToError() (Record, error) {
	_, err := self.Channel.ToError()
	return self, err
}

/* RECORD */
/*
RecordInit
BOOL BASSDEF(BASS_RecordInit)(int device);
*/
func RecordInit(device int) error {
	return boolToError(C.BASS_RecordInit(C.int(device)))
}

/*
RecordFree
BOOL BASSDEF(BASS_RecordFree)();
*/
func RecordFree() error {
	return boolToError(C.BASS_RecordFree())
}

func RecordStart(freq, chans int, flags Flags, streamproc *C.RECORDPROC, userdata unsafe.Pointer) (Record, error) {
	channel := C.BASS_RecordStart(cuint(freq), cuint(chans), cuint(flags), streamproc, userdata)
	return Record{Channel: ChannelFromHandle(uint32(channel))}.ToError()
}

func RecordGetDeviceInfo(device int) (DeviceInfo, error) {
	var info C.BASS_DEVICEINFO
	err := boolToError(C.BASS_RecordGetDeviceInfo(C.DWORD(device), &info))
	if err!=nil {
		return DeviceInfo{}, err
	} else {
		return DeviceInfo{Name: C.GoString(info.name), Driver: C.GoString(info.driver), Flags: Flags(info.flags), Kind: getHighWord(info.flags)}, nil
	}
}
func RecordGetDeviceInfoFlags(device int) (Flags, error) {
	var info C.BASS_DEVICEINFO
	err := boolToError(C.BASS_RecordGetDeviceInfo(C.DWORD(device), &info))
	if err!=nil {
		return 0, err
	} else {
		return Flags(info.flags), nil
	}
}
type RecordInfo struct {
	Flags, Formats, Inputs, Freq, Channels int
	SingleIn bool
}
func RecordGetInfo() (RecordInfo, error) {
	var info C.BASS_RECORDINFO
	err := boolToError(C.BASS_RecordGetInfo(&info))
	if err!=nil {
		return RecordInfo{}, err
	} else {
		 // For some reason the channels are stored in the last byte of formats
		formats := info.formats>>8
		ptr := (*uint8)(unsafe.Pointer(&formats))
		return RecordInfo{Flags: int(info.flags), Formats: int(info.formats), Inputs: int(info.inputs), SingleIn: info.singlein!=0, Freq: int(info.freq), Channels: int(*ptr)}, nil
	}
}

func RecordSetDevice(device int) error {
	return boolToError(C.BASS_RecordSetDevice(C.DWORD(device)))
}
func RecordGetDevice() (int, error) {
	return intPairToError(C.BASS_RecordGetDevice())
}