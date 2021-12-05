package mp3
import (
	"github.com/keithcat1/gobass"
"github.com/keithcat1/gobass/enc"
)

/*
#include "bassenc_MP3.h"
#include "bass.h"
*/
import "C"
func NewMP3EncoderFile(channel bass.Channel, options string, flags bass.Flags, file string) (enc.Encoder, error) {
	coptions := (*C.char)(bass.ToUTF16(options))
	cfile := (*C.char)(bass.ToUTF16(file))

	return enc.EncoderFromHandle(uint32(C.BASS_Encode_MP3_StartFile(C.DWORD(channel.GetHandle()), coptions, C.DWORD(flags)|C.BASS_UNICODE, cfile))).ToError()
}