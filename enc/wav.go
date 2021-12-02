package enc
import (
	"github.com/keithcat1/gobass"
)

/*
#include "bassenc.h"
#include "bass.h"
*/
import "C"
func NewWAVEncoderFile(channel bass.Channel, flags bass.Flags, file string) (Encoder, error) {
	cfile := (*C.char)(bass.ToUTF16(file))
	return Encoder(C.BASS_Encode_Start(C.DWORD(channel), cfile, EncodePCM|C.BASS_UNICODE, nil, nil)).ToError()
}