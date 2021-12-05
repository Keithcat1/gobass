package ogg
import (
	"github.com/keithcat1/gobass"
	"github.com/keithcat1/gobass/enc"
)

/*
#include "bassenc_OGG.h"
#include "bass.h"
*/
import "C"
func NewOGGEncoderFile(channel bass.Channel, options string, flags bass.Flags, file string) (enc.Encoder, error) {
	coptions := (*C.char)(bass.ToUTF16(options))
	cfile := (*C.char)(bass.ToUTF16(file))
	return enc.EncoderFromHandle(uint32(C.BASS_Encode_OGG_StartFile(C.DWORD(channel.GetHandle()), coptions, C.DWORD(flags)|C.BASS_UNICODE, cfile))).ToError()
}