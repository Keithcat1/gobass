package mix
import (
	"github.com/keithcat1/gobass"
	"unsafe"
	"runtime/cgo"
)

/*
#include "bassmix.h"
#include "_gobass_callbacks.h"
*/
import "C"
type Mixer struct {
	bass.Channel
}
func (self Mixer) ToError() (Mixer, error) {
	_, err := self.Channel.ToError()
	return self, err
}
func StreamCreate(freq, chans int, flags bass.Flags) (Mixer, error) {
	return Mixer{Channel: bass.ChannelFromHandle(uint32(C.BASS_Mixer_StreamCreate(C.DWORD(freq), C.DWORD(chans), C.DWORD(flags))))}.ToError()
}
func (self Mixer) AddChannel(channel bass.Channel, flags bass.Flags) error {
	return boolToError(C.BASS_Mixer_StreamAddChannel(C.DWORD(self.GetHandle()), C.DWORD(channel.GetHandle()), C.DWORD(flags)))
}
func (self Mixer) AddChannelEx(channel bass.Channel, flags bass.Flags, start, length int64) error {
	return boolToError(C.BASS_Mixer_StreamAddChannelEx(C.DWORD(self.GetHandle()), C.DWORD(channel.GetHandle()), C.DWORD(flags), C.QWORD(start), C.QWORD(length)))
}
func (self Mixer) GetChannels(channels []bass.Channel) (int, error) {
	var ptr *C.DWORD
		memory := make([]C.DWORD, len(channels))
	if len(memory)>0 {
		ptr = &memory[0]
	}
	count := int(C.BASS_Mixer_StreamGetChannels(C.DWORD(self.GetHandle()), ptr, C.DWORD(len(memory))))
	if count +1 == 0 {
		return 0, bass.GetLastError()
	} else {
		for i, v := range memory {
			channels[i] = bass.ChannelFromHandle(uint32(v))
		}
		return count, nil
	}
}
func ChannelGetData(channel bass.Channel, buf []byte) (int, error) {
	var ptr unsafe.Pointer
	if len(buf)>0 {
		ptr = unsafe.Pointer(&buf[0])
	}
	count := int(C.BASS_Mixer_ChannelGetData(C.DWORD(channel.GetHandle()), ptr, C.DWORD(len(buf))))
	if count +1 == 0 {
		return 0, bass.GetLastError()
	} else {
		return count, nil
	}
}
func ChannelRemove(channel bass.Channel) error {
	return boolToError(C.BASS_Mixer_ChannelRemove(C.DWORD(channel.GetHandle())))
}
func ChannelFlags(channel bass.Channel, flags, mask bass.Flags) (bass.Flags, error) {
	flags = bass.Flags(C.BASS_Mixer_ChannelFlags(C.DWORD(channel.GetHandle()), C.DWORD(flags), C.DWORD(mask)))
	if flags +1 == 0 {
		return 0, bass.GetLastError()
	} else {
		return flags, nil
	}
}
func ChannelSetSync(channel bass.Channel, syncType int, flags bass.Flags, param int, callback cgo.Handle) (bass.Sync, error) {
	result := C.BASS_Mixer_ChannelSetSync(C.DWORD(channel.GetHandle()), C.DWORD(syncType)|C.DWORD(flags), C.QWORD(param), C._get_GoSyncprocCallbackWrapper(), unsafe.Pointer(callback))
	if result == 0 {
		return 0, bass.GetLastError()
	} else {
		return bass.Sync(result), nil
	}
}
func ChannelSetSyncCCallback(channel bass.Channel, syncType int, flags bass.Flags, param int, callback unsafe.Pointer, userdata unsafe.Pointer) (bass.Sync, error) {
	result := C.BASS_Mixer_ChannelSetSync(C.DWORD(channel.GetHandle()), C.DWORD(syncType)|C.DWORD(flags), C.QWORD(param), (*C.SYNCPROC)(callback), userdata)
	if result == 0 {
		return 0, bass.GetLastError()
	} else {
		return bass.Sync(result), nil
	}
}
func ChannelSetPosition(channel bass.Channel, pos int, posType int) error {
	return boolToError(C.BASS_Mixer_ChannelSetPosition(C.DWORD(channel.GetHandle()), C.QWORD(pos), C.DWORD(posType)))
}