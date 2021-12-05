// bass project bass.go
// http://www.un4seen.com/doc/#bass/

package bass

/*
#include "bass.h"
#include "stdlib.h"
*/
import "C"

import (
	"unsafe"
	"runtime/cgo"
	"unicode/utf16"
)

//Stores a C byte array, with a pointer to the data and its length.
type CBytes struct {
	Data   unsafe.Pointer
	Length int
}
type cuint = C.DWORD
type culong = C.QWORD

type SyncProc C.SYNCPROC



// An error code returned by any function.
type Error int
func (self Error) Error() string {
	return "BASS error: "+codes[self]
}


type Handle interface {
	GetHandle() uint32
}

var (
	codes        = map[Error]string{OK: "all is OK", ERROR_MEM: "memory error", ERROR_FILEOPEN: "can't open the file", ERROR_DRIVER: "can't find a free/valid driver", ERROR_BUFLOST: "the sample buffer was lost", ERROR_HANDLE: "invalid handle", ERROR_FORMAT: "unsupported sample format", ERROR_POSITION: "invalid position", ERROR_INIT: "BASS_Init has not been successfully called", ERROR_START: "BASS_Start has not been successfully called", ERROR_SSL: "SSL/HTTPS support isn't available", ERROR_ALREADY: "already initialized/paused/whatever", ERROR_NOTAUDIO: "NOTAUDIO", ERROR_NOCHAN: "can't get a free channel", ERROR_ILLTYPE: "an illegal type was specified", ERROR_ILLPARAM: "an illegal parameter was specified", ERROR_NO3D: "no 3D support", ERROR_NOEAX: "no EAX support", ERROR_DEVICE: "illegal device number", ERROR_NOPLAY: "not playing", ERROR_FREQ: "illegal sample rate", ERROR_NOTFILE: "the stream is not a file stream", ERROR_NOHW: "no hardware voices available", ERROR_EMPTY: "the MOD music has no sequence data", ERROR_NONET: "no internet connection could be opened", ERROR_CREATE: "couldn't create the file", ERROR_NOFX: "effects are not available", ERROR_NOTAVAIL: "requested data is not available", ERROR_DECODE: "the channel is/isn't a 'decoding channel", ERROR_DX: "a sufficient DirectX version is not installed", ERROR_TIMEOUT: "connection timedout", ERROR_FILEFORM: "unsupported file format", ERROR_SPEAKER: "unavailable speaker", ERROR_VERSION: "invalid BASS version (used by add-ons)", ERROR_CODEC: "codec is not available/supported", ERROR_ENDED: "the channel/file has ended", ERROR_BUSY: "the device is busy", ERROR_UNSTREAMABLE: "Error_UNSTREAMABLE", ERROR_UNKNOWN: "some other mystery problem"}
)
func Init(device int, freq int, flags Flags, hwnd, clsid uintptr) error {
	window := (bassInitPointer)(unsafe.Pointer(hwnd))
	gid:=unsafe.Pointer(clsid)
	return boolToError(C.BASS_Init(C.int(device), C.DWORD(freq), C.DWORD(flags), window, (gid)))
}

/*
Free
BOOL BASS_Free();
*/
func Free() error {
	return boolToError(C.BASS_Free())
}

/*
GetConfig
DWORD BASSDEF(BASS_GetConfig)(DWORD option);
*/
func GetConfig(option int) (int64, error) {
	return longPairToError(C.BASS_GetConfig(C.DWORD(option)))
}

/*
SetConfig
BOOL BASSDEF(BASS_SetConfig)(DWORD option, DWORD value);
*/
func SetConfig(option, value int) error {
	return boolToError(C.BASS_SetConfig(C.DWORD(option), C.DWORD(value)))
}

// GetVol
// float BASSDEF(BASS_GetVolume)();
func GetVolume() (float64, error) {
	return floatPairToError(C.BASS_GetVolume())
}

// SetVol
func SetVolume(v float64) error {
	return boolToError(C.BASS_SetVolume(C.float(v)))
}



// PluginLoad
func PluginLoad(file string, flags Flags) (handle uint32, err error) {
	cfile := ToUTF16(file)
	plugin := C.BASS_PluginLoad(cfile, cuint(flags)|C.BASS_UNICODE)
	return uint32(plugin), errMsg()
}

// PluginFree
func PluginFree(handle uint32) error {
	return boolToError(C.BASS_PluginFree(cuint(handle)))
}
func errMsg() error {
	c := Error(C.BASS_ErrorGetCode())
	if c == 0 {
		return nil
	} else {
		return c
	}
}
// Returns the last error, if any, that was caused by a call to a BASS function. You should not normally call this function, all BASS functions in this package handle and return errors.
func GetLastError() error {
	return errMsg()
}

func boolToInt(val bool) C.int {
	if val {
		return 1
	} else {
		return 0
	}
}
func IsStarted() bool {
	return C.BASS_IsStarted()!=0
}
func intToBool(val C.int) bool {
	return val != 0
}
//Allocates C memory and coppies data to that C memory.
func CopyBytes(data []byte) CBytes {
	return CBytes{Data: C.CBytes(data), Length: len(data)}
}
func GetDevice() (int64, error) {
	return longPairToError(C.BASS_GetDevice())
}
func SetDevice(device int) error {
	return boolToError(C.BASS_SetDevice(C.DWORD(device)))
}
func boolToError(value C.BOOL) error {
	if value == 0 {
		return errMsg()
	} else {
		return nil
	}
}
func pairToError(value C.int) (int, error) {
	return int(value), boolToError(value)
}
func floatPairToError(value C.float) (float64, error) {
	return float64(value), boolToError(C.int(value))
}
func longPairToError(value C.DWORD) (int64, error) {
	return int64(value), boolToError(C.int(value))
}
func channelToError(ch C.DWORD) (Channel, error) {
	return Channel{handle: ch}, boolToError(C.int(ch))
}
func sampleToError(ch C.DWORD) (Sample, error) {
	return Sample{handle: ch}, boolToError(C.int(ch))
}
func longlongPairToError(value C.QWORD) (int64, error) {
	return int64(value), boolToError(C.int(value))
}
func intToError(value cuint) error {
	if value==0 {
		return errMsg()
	} else {
		return nil
	}
}
func intPairToError(value C.DWORD) (int, error) {
	return int(value), intToError(value)
}
func longToError(value culong) error {
	if value!=0 {
		return nil
	} else {
		return errMsg()
	}
}


type DeviceInfo struct {
	Name, Driver string
	Flags Flags
	Kind uint8
}
func GetDeviceInfo(device int) (DeviceInfo, error) {
	var info C.BASS_DEVICEINFO
	err := boolToError(C.BASS_GetDeviceInfo(C.DWORD(device), &info))
	if err!=nil {
		return DeviceInfo{}, err
	} else {
		return DeviceInfo{Name: C.GoString(info.name), Driver: C.GoString(info.driver), Flags: Flags(info.flags), Kind: getHighWord(info.flags)}, nil
	}
}

func GetDeviceInfoFlags(device int) (Flags, error) {
	var info C.BASS_DEVICEINFO
	err := boolToError(C.BASS_GetDeviceInfo(C.DWORD(device), &info))
	if err!=nil {
		return 0, err
	} else {
		return Flags(info.flags), nil
	}
}
func streamToError(v C.DWORD) (Stream, error) {
	channel, err := channelToError(v)
	return channel.toStream(), err
}

// a helper type to allow setting / clearing BASS flags easily
type Flags int64
func (self Flags) Add(flag int) Flags {
	return Flags(int(self) | flag)
}
func (self Flags) Has(flag int) bool {
	return int(self) & flag == flag
}
// Some of BASS functions like to put multiple values into a single integer
func getHighWord(v C.DWORD) uint8 {
	v = v >> 8
	ptr := (*uint8)(unsafe.Pointer(&v))
	return *ptr
}
type Sync uint32
func (self Channel) RemoveSync(sync Sync) error {
	return boolToError(C.BASS_ChannelRemoveSync(self.cint(), C.DWORD(sync)))
}

// creates a new runtime/cgo.Handle and converts it to unsafe.Pointer, ready to be pased into C land
func NewCGOHandle(value interface{}) unsafe.Pointer {
	return unsafe.Pointer(cgo.NewHandle(value))
}
// frees an unneeded handle created by NewCGOHandle
func DestroyCGOHandle(handle unsafe.Pointer) {
	cgo.Handle(handle).Delete()
}
// if condition is true, adds the specified flags and returns the result, else just returns self without modifying it
func (self Flags) AddIf(flag Flags, condition bool) Flags {
	if condition {
		return self | flag
	} else {
		return self
	}
}
func GetCPU() float64 {
	return float64(C.BASS_GetCPU())
}
// converts a UTF8 string to a UTF16 string, then returns a pointer to that UTF16 data suitable to pass to BASS along with the C.BASS_UNICODE flag
// only exposed so that BASS plugin bindings can use it!
func ToUTF16(s string) *C.char {
	if len(s) == 0 {
		return nil
	}
	data := utf16.Encode([]rune(s))
	ptr := unsafe.Pointer(&data[0])
	return (*C.char)(ptr)
}