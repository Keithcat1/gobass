package bass

/*
#include "bass.h"
#include "stdlib.h"
*/
import "C"

import (
	"unsafe"
)
type Stream struct {
	Channel
}
func StreamCreateCCallback(freq, chans int, flags Flags, streamproc *C.STREAMPROC, userdata unsafe.Pointer) (Stream, error) {
	channel := C.BASS_StreamCreate(cuint(freq), cuint(chans), cuint(flags), streamproc, userdata)
	return streamToError(channel)
}

func (self Stream) ToError() (Stream, error) {
	_, err := self.Channel.ToError()
	return self, err
}
// StreamCreateURL
// HSTREAM BASSDEF(BASS_StreamCreateURL)(const char *url, DWORD offset, DWORD flags, DOWNLOADPROC *proc, void *user);
func StreamCreateURL(url string, flags Flags) (Stream, error) {
	curl := ToUTF16(url)
	ch := C.BASS_StreamCreateURL(curl, 0, cuint(flags)|C.BASS_UNICODE, nil, nil)
	return streamToError(ch)
}
// BASS_StreamCreateFile
// HSTREAM BASSDEF(BASS_StreamCreateFile)(BOOL mem, const void *file, QWORD offset, QWORD length, DWORD flags);
func StreamCreateFile(data interface{}, offset int, flags Flags) (Stream, error) {
	var ch C.DWORD
	switch data := data.(type) {
	case CBytes:
		ch = C.BASS_StreamCreateFile(1, data.Data, culong(offset), culong(data.Length), cuint(flags))
	case string:
		cstring := unsafe.Pointer(ToUTF16(data))
		ch = C.BASS_StreamCreateFile(0, cstring, culong(offset), 0, cuint(flags)|C.BASS_UNICODE)
	case []byte:
		cbytes := C.CBytes(data)
		ch = C.BASS_StreamCreateFile(1, cbytes, culong(offset), culong(len(data)), cuint(flags))
		// unlike BASS_SampleLoad, BASS won't make a copy of the sample data internally, which means we can't just pass a pointer to the Go bytes. Instead we need to set a sync to free the bytes when the stream it's associated with is freed
		if ch != 0 {
			channel := Channel{handle: ch}.toStream()
			_, err := channel.SetSyncCCallback(SYNC_FREE, SYNC_ONETIME, 0, SyncprocFree, cbytes)
			if err != nil {
				return Stream{}, err
			}
		}
	}
	return streamToError(ch)
}
func (self Stream) StreamFree() error {
	return boolToError(C.BASS_StreamFree(self.cint()))
}
func (self Stream) PutData(data []byte, flags Flags) (int, error) {
	var ptr unsafe.Pointer
	if len(data)>0 {
		ptr = unsafe.Pointer(&data[0])
	}
	val := C.BASS_StreamPutData(self.cint(), ptr, C.DWORD(len(data)|int(flags)))
	if val +1 == 0 { // -1 indicates an error, but this is an unsigned integer
		return 0, errMsg()
	} else {
		return int(val), nil
	}
}