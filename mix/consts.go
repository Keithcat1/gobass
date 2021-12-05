package mix
import (

)
/*
#include "bassmix.h"
*/
import "C"
const (
	MixerEnd = C.BASS_MIXER_END
	MixerNonstop = C.BASS_MIXER_NONSTOP
	MixerPosex = C.BASS_MIXER_POSEX
	MixerResume = C.BASS_MIXER_RESUME
	MixerChanBuffer = C.BASS_MIXER_CHAN_BUFFER
	MixerChanDownmix = C.BASS_MIXER_CHAN_DOWNMIX
	MixerChanLimit = C.BASS_MIXER_CHAN_LIMIT
	MixerChanMatrix = C.BASS_MIXER_CHAN_MATRIX
	MixerChanNorampin = C.BASS_MIXER_CHAN_NORAMPIN
	MixerChanPause = C.BASS_MIXER_CHAN_PAUSE
	AttribMixerThreads = C.BASS_ATTRIB_MIXER_THREADS
)