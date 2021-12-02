package flac
import (
	"github.com/keithcat1/gobass"
"github.com/keithcat1/gobass/enc"
)

/*
#include "bassenc_FLAC.h"
#include "bass.h"
*/
import "C"

func NewFLACEncoderFile(channel bass.Channel, options string, flags bass.Flags, file string) (enc.Encoder, error) {
	coptions := (*C.char)(bass.ToUTF16(options))

	cfile := (*C.char)(bass.ToUTF16(file))

	return enc.Encoder(C.BASS_Encode_FLAC_StartFile(C.DWORD(channel), coptions, C.DWORD(flags)|C.BASS_UNICODE, cfile)).ToError()
}