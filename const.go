// const.go
package bass

/*
#include "bass.h"
*/
import "C"

const (
	// BASSVERSION as defined in bass.h:41
	BASSVERSION = C.BASSVERSION
	// BASSVERSIONTEXT as defined in bass.h:42
	BASSVERSIONTEXT = "2.4"
	// BASS_OK as defined in bass.h:61
	OK = Error(C.BASS_OK)
	// BASS_ERROR_MEM as defined in bass.h:62
	ERROR_MEM = Error(C.BASS_ERROR_MEM)
	// BASS_ERROR_FILEOPEN as defined in bass.h:63
	ERROR_FILEOPEN = Error(C.BASS_ERROR_FILEOPEN)
	// BASS_ERROR_DRIVER as defined in bass.h:64
	ERROR_DRIVER = Error(C.BASS_ERROR_DRIVER)
	// BASS_ERROR_BUFLOST as defined in bass.h:65
	ERROR_BUFLOST = Error(C.BASS_ERROR_BUFLOST)
	// BASS_ERROR_HANDLE as defined in bass.h:66
	ERROR_HANDLE = Error(C.BASS_ERROR_HANDLE)
	// BASS_ERROR_FORMAT as defined in bass.h:67
	ERROR_FORMAT = Error(C.BASS_ERROR_FORMAT)
	// BASS_ERROR_POSITION as defined in bass.h:68
	ERROR_POSITION = Error(C.BASS_ERROR_POSITION)
	// BASS_ERROR_INIT as defined in bass.h:69
	ERROR_INIT = Error(C.BASS_ERROR_INIT)
	// BASS_ERROR_START as defined in bass.h:70
	ERROR_START = Error(C.BASS_ERROR_START)
	// BASS_ERROR_SSL as defined in bass.h:71
	ERROR_SSL = Error(C.BASS_ERROR_SSL)
	// BASS_ERROR_ALREADY as defined in bass.h:72
	ERROR_ALREADY = Error(C.BASS_ERROR_ALREADY)
	// BASS_ERROR_NOCHAN as defined in bass.h:73
	ERROR_NOTAUDIO=Error(C.BASS_ERROR_NOTAUDIO)
	ERROR_NOCHAN = Error(C.BASS_ERROR_NOCHAN)
	// BASS_ERROR_ILLTYPE as defined in bass.h:74)
	ERROR_ILLTYPE = Error(C.BASS_ERROR_ILLTYPE)
	// BASS_ERROR_ILLPARAM as defined in bass.h:75
	ERROR_ILLPARAM = Error(C.BASS_ERROR_ILLPARAM)
	// BASS_ERROR_NO3D as defined in bass.h:76
	ERROR_NO3D = Error(C.BASS_ERROR_NO3D)
	// BASS_ERROR_NOEAX as defined in bass.h:77
	ERROR_NOEAX = Error(C.BASS_ERROR_NOEAX)
	// BASS_ERROR_DEVICE as defined in bass.h:78
	ERROR_DEVICE = Error(C.BASS_ERROR_DEVICE)
	// ERROR_NOPLAY as defined in bass.h:79
	ERROR_NOPLAY = Error(C.BASS_ERROR_NOPLAY)
	// BASS_ERROR_FREQ as defined in bass.h:80
	ERROR_FREQ = Error(C.BASS_ERROR_FREQ)
	// BASS_ERROR_NOTFILE as defined in bass.h:81
	ERROR_NOTFILE = Error(C.BASS_ERROR_NOTFILE)
	// BASS_ERROR_NOHW as defined in bass.h:82
	ERROR_NOHW = Error(C.BASS_ERROR_NOHW)
	// BASS_ERROR_EMPTY as defined in bass.h:83
	ERROR_EMPTY = Error(C.BASS_ERROR_EMPTY)
	// BASS_ERROR_NONET as defined in bass.h:84
	ERROR_NONET = Error(C.BASS_ERROR_NONET)
	// BASS_ERROR_CREATE as defined in bass.h:85
	ERROR_CREATE = Error(C.BASS_ERROR_CREATE)
	// BASS_ERROR_NOFX as defined in bass.h:86
	ERROR_NOFX = Error(C.BASS_ERROR_NOFX)
	// BASS_ERROR_NOTAVAIL as defined in bass.h:87
	ERROR_NOTAVAIL = Error(C.BASS_ERROR_NOTAVAIL)
	// BASS_ERROR_DECODE as defined in bass.h:88
	ERROR_DECODE = Error(C.BASS_ERROR_DECODE)
	// BASS_ERROR_DX as defined in bass.h:89
	ERROR_DX = Error(C.BASS_ERROR_DX)
	// BASS_ERROR_TIMEOUT as defined in bass.h:90
	ERROR_TIMEOUT = Error(C.BASS_ERROR_TIMEOUT)
	// BASS_ERROR_FILEFORM as defined in bass.h:91
	ERROR_FILEFORM = Error(C.BASS_ERROR_FILEFORM)
	// BASS_ERROR_SPEAKER as defined in bass.h:92
	ERROR_SPEAKER = Error(C.BASS_ERROR_SPEAKER)
	// BASS_ERROR_VERSION as defined in bass.h:93
	ERROR_VERSION = Error(C.BASS_ERROR_VERSION)
	// BASS_ERROR_CODEC as defined in bass.h:94
	ERROR_CODEC = Error(C.BASS_ERROR_CODEC)
	// BASS_ERROR_ENDED as defined in bass.h:95
	ERROR_ENDED = Error(C.BASS_ERROR_ENDED)
	// BASS_ERROR_BUSY as defined in bass.h:96
	ERROR_BUSY = Error(C.BASS_ERROR_BUSY)
	ERROR_UNSTREAMABLE = Error(C.BASS_ERROR_UNSTREAMABLE)
	// BASS_ERROR_UNKNOWN as defined in bass.h:97
	ERROR_UNKNOWN = Error(C.BASS_ERROR_UNKNOWN)
	// BASS_CONFIG_BUFFER as defined in bass.h:100
	CONFIG_BUFFER = C.BASS_CONFIG_BUFFER
	// BASS_CONFIG_UPDATEPERIOD as defined in bass.h:101
	CONFIG_UPDATEPERIOD = C.BASS_CONFIG_UPDATEPERIOD
	// BASS_CONFIG_GVOL_SAMPLE as defined in bass.h:102
	CONFIG_GVOL_SAMPLE = C.BASS_CONFIG_GVOL_SAMPLE
	// BASS_CONFIG_GVOL_STREAM as defined in bass.h:103
	CONFIG_GVOL_STREAM = C.BASS_CONFIG_GVOL_STREAM
	// BASS_CONFIG_GVOL_MUSIC as defined in bass.h:104
	CONFIG_GVOL_MUSIC = C.BASS_CONFIG_GVOL_MUSIC
	// BASS_CONFIG_CURVE_VOL as defined in bass.h:105
	CONFIG_CURVE_VOL = C.BASS_CONFIG_CURVE_VOL
	// BASS_CONFIG_CURVE_PAN as defined in bass.h:106
	CONFIG_CURVE_PAN = C.BASS_CONFIG_CURVE_PAN
	// BASS_CONFIG_FLOATDSP as defined in bass.h:107
	CONFIG_FLOATDSP = C.BASS_CONFIG_FLOATDSP
	// BASS_CONFIG_3DALGORITHM as defined in bass.h:108
	CONFIG_3DALGORITHM = C.BASS_CONFIG_3DALGORITHM
	// BASS_CONFIG_NET_TIMEOUT as defined in bass.h:109
	CONFIG_NET_TIMEOUT = C.BASS_CONFIG_NET_TIMEOUT
	// BASS_CONFIG_NET_BUFFER as defined in bass.h:110
	CONFIG_NET_BUFFER = C.BASS_CONFIG_NET_BUFFER
	// BASS_CONFIG_PAUSE_NOPLAY as defined in bass.h:111
	CONFIG_PAUSE_NOPLAY = C.BASS_CONFIG_PAUSE_NOPLAY
	// CONFIG_NET_PREBUF as defined in bass.h:112
	CONFIG_NET_PREBUF = C.BASS_CONFIG_NET_PREBUF
	// BASS_CONFIG_NET_PASSIVE as defined in bass.h:113
	CONFIG_NET_PASSIVE = C.BASS_CONFIG_NET_PASSIVE
	// BASS_CONFIG_REBASS_BUFFER as defined in bass.h:114
	CONFIG_REBASS_BUFFER = C.BASS_CONFIG_REC_BUFFER
	// BASS_CONFIG_NET_PLAYLIST as defined in bass.h:115
	CONFIG_NET_PLAYLIST = C.BASS_CONFIG_NET_PLAYLIST
	// BASS_CONFIG_MUSIBASS_VIRTUAL as defined in bass.h:116
	CONFIG_MUSIBASS_VIRTUAL = C.BASS_CONFIG_MUSIC_VIRTUAL
	// BASS_CONFIG_VERIFY as defined in bass.h:117
	CONFIG_VERIFY = C.BASS_CONFIG_VERIFY
	// BASS_CONFIG_UPDATETHREADS as defined in bass.h:118
	CONFIG_UPDATETHREADS = C.BASS_CONFIG_UPDATETHREADS
	// BASS_CONFIG_DEV_BUFFER as defined in bass.h:119
	CONFIG_DEV_BUFFER = C.BASS_CONFIG_DEV_BUFFER
	// BASS_CONFIG_VISTA_TRUEPOS as defined in bass.h:120
	CONFIG_VISTA_TRUEPOS = C.BASS_CONFIG_VISTA_TRUEPOS
	// BASS_CONFIG_IOS_MIXAUDIO as defined in bass.h:121
	CONFIG_IOS_MIXAUDIO = C.BASS_CONFIG_IOS_MIXAUDIO
	// BASS_CONFIG_DEV_DEFAULT as defined in bass.h:122
	CONFIG_DEV_DEFAULT = C.BASS_CONFIG_DEV_DEFAULT
	// BASS_CONFIG_NET_READTIMEOUT as defined in bass.h:123
	CONFIG_NET_READTIMEOUT = C.BASS_CONFIG_NET_READTIMEOUT
	// BASS_CONFIG_VISTA_SPEAKERS as defined in bass.h:124
	CONFIG_VISTA_SPEAKERS = C.BASS_CONFIG_VISTA_SPEAKERS
	// BASS_CONFIG_IOS_SPEAKER as defined in bass.h:125
	CONFIG_IOS_SPEAKER = C.BASS_CONFIG_IOS_SPEAKER
	// BASS_CONFIG_MF_DISABLE as defined in bass.h:126
	CONFIG_MF_DISABLE = C.BASS_CONFIG_MF_DISABLE
	// BASS_CONFIG_HANDLES as defined in bass.h:127
	CONFIG_HANDLES = C.BASS_CONFIG_HANDLES
	// BASS_CONFIG_UNICODE as defined in bass.h:128
	CONFIG_UNICODE = C.BASS_CONFIG_UNICODE
	// BASS_CONFIG_SRC as defined in bass.h:129
	CONFIG_SRC = C.BASS_CONFIG_SRC
	// BASS_CONFIG_SRBASS_SAMPLE as defined in bass.h:130
	CONFIG_SRBASS_SAMPLE = C.BASS_CONFIG_SRC_SAMPLE
	// BASS_CONFIG_ASYNCFILE_BUFFER as defined in bass.h:131
	CONFIG_ASYNCFILE_BUFFER = C.BASS_CONFIG_ASYNCFILE_BUFFER
	// BASS_CONFIG_OGG_PRESCAN as defined in bass.h:132
	CONFIG_OGG_PRESCAN = C.BASS_CONFIG_OGG_PRESCAN
	// BASS_CONFIG_MF_VIDEO as defined in bass.h:133
	CONFIG_MF_VIDEO = C.BASS_CONFIG_MF_VIDEO
	// BASS_CONFIG_AIRPLAY as defined in bass.h:134
	CONFIG_AIRPLAY = C.BASS_CONFIG_AIRPLAY
	// BASS_CONFIG_DEV_NONSTOP as defined in bass.h:135
	CONFIG_DEV_NONSTOP = C.BASS_CONFIG_DEV_NONSTOP
	// BASS_CONFIG_IOS_NOCATEGORY as defined in bass.h:136
	CONFIG_IOS_NOCATEGORY = C.BASS_CONFIG_IOS_NOCATEGORY
	// BASS_CONFIG_VERIFY_NET as defined in bass.h:137
	CONFIG_VERIFY_NET = C.BASS_CONFIG_VERIFY_NET
	// BASS_CONFIG_DEV_PERIOD as defined in bass.h:138
	CONFIG_DEV_PERIOD = C.BASS_CONFIG_DEV_PERIOD
	// BASS_CONFIG_FLOAT as defined in bass.h:139
	CONFIG_FLOAT = C.BASS_CONFIG_FLOAT
	// BASS_CONFIG_NET_SEEK as defined in bass.h:140
	CONFIG_NET_SEEK = C.BASS_CONFIG_NET_SEEK
	// BASS_CONFIG_NET_AGENT as defined in bass.h:143
	CONFIG_NET_AGENT = C.BASS_CONFIG_NET_AGENT
	// BASS_CONFIG_NET_PROXY as defined in bass.h:144
	CONFIG_NET_PROXY = C.BASS_CONFIG_NET_PROXY
	// BASS_CONFIG_IOS_NOTIFY as defined in bass.h:145
	CONFIG_IOS_NOTIFY = C.BASS_CONFIG_IOS_NOTIFY
	// DEVICE_8BITS as defined in bass.h:148
	DEVICE_8BITS = 1
	// DEVICE_MONO as defined in bass.h:149
	DEVICE_MONO = C.BASS_DEVICE_MONO
	// DEVICE_3D as defined in bass.h:150
	DEVICE_3D = 4
	// DEVICE_16BITS as defined in bass.h:151
	DEVICE_16BITS = 8
	// DEVICE_LATENCY as defined in bass.h:152
	DEVICE_LATENCY = C.BASS_DEVICE_LATENCY
	// DEVICE_CPSPEAKERS as defined in bass.h:153
	DEVICE_CPSPEAKERS = C.BASS_DEVICE_CPSPEAKERS
	// DEVICE_SPEAKERS as defined in bass.h:154
	DEVICE_SPEAKERS = C.BASS_DEVICE_SPEAKERS
	// DEVICE_NOSPEAKER as defined in bass.h:155
	DEVICE_NOSPEAKER = C.BASS_DEVICE_NOSPEAKER
	// DEVICE_DMIX as defined in bass.h:156
	DEVICE_DMIX = C.BASS_DEVICE_DMIX
	// DEVICE_FREQ as defined in bass.h:157
	DEVICE_FREQ = C.BASS_DEVICE_FREQ
	// DEVICE_STEREO as defined in bass.h:158
	DEVICE_STEREO = C.BASS_DEVICE_STEREO
	// BASS_OBJECT_DS as defined in bass.h:161
	BASS_OBJECT_DS = C.BASS_OBJECT_DS
	// BASS_OBJECT_DS3DL as defined in bass.h:162
	BASS_OBJECT_DS3DL = 2
	// DEVICE_ENABLED as defined in bass.h:177
	DEVICE_ENABLED = C.BASS_DEVICE_ENABLED
	// DEVICE_DEFAULT as defined in bass.h:178
	DEVICE_DEFAULT = C.BASS_DEVICE_DEFAULT
	// DEVICE_INIT as defined in bass.h:179
	DEVICE_INIT = C.BASS_DEVICE_INIT
	// DEVICE_TYPE_MASK as defined in bass.h:181
	DEVICE_TYPE_MASK = C.BASS_DEVICE_TYPE_MASK
	// DEVICE_TYPE_NETWORK as defined in bass.h:182
	DEVICE_TYPE_NETWORK = C.BASS_DEVICE_TYPE_NETWORK
	// DEVICE_TYPE_SPEAKERS as defined in bass.h:183
	DEVICE_TYPE_SPEAKERS = C.BASS_DEVICE_TYPE_SPEAKERS
	// DEVICE_TYPE_LINE as defined in bass.h:184
	DEVICE_TYPE_LINE = C.BASS_DEVICE_TYPE_LINE
	// DEVICE_TYPE_HEADPHONES as defined in bass.h:185
	DEVICE_TYPE_HEADPHONES = C.BASS_DEVICE_TYPE_HEADPHONES
	// DEVICE_TYPE_MICROPHONE as defined in bass.h:186
	DEVICE_TYPE_MICROPHONE = C.BASS_DEVICE_TYPE_MICROPHONE
	// DEVICE_TYPE_HEADSET as defined in bass.h:187
	DEVICE_TYPE_HEADSET = C.BASS_DEVICE_TYPE_HEADSET
	// DEVICE_TYPE_HANDSET as defined in bass.h:188
	DEVICE_TYPE_HANDSET = C.BASS_DEVICE_TYPE_HANDSET
	// DEVICE_TYPE_DIGITAL as defined in bass.h:189
	DEVICE_TYPE_DIGITAL = C.BASS_DEVICE_TYPE_DIGITAL
	// DEVICE_TYPE_SPDIF as defined in bass.h:190
	DEVICE_TYPE_SPDIF = C.BASS_DEVICE_TYPE_SPDIF
	// DEVICE_TYPE_HDMI as defined in bass.h:191
	DEVICE_TYPE_HDMI = C.BASS_DEVICE_TYPE_HDMI
	// DEVICE_TYPE_DISPLAYPORT as defined in bass.h:192
	DEVICE_TYPE_DISPLAYPORT = C.BASS_DEVICE_TYPE_DISPLAYPORT
	// DEVICES_AIRPLAY as defined in bass.h:195
	DEVICES_AIRPLAY = C.BASS_DEVICES_AIRPLAY
	// DSCAPS_CONTINUOUSRATE as defined in bass.h:215
	// DSCAPS_EMULDRIVER as defined in bass.h:216
	DSCAPS_EMULDRIVER = C.DSCAPS_EMULDRIVER
	// DSCAPS_CERTIFIED as defined in bass.h:217
	DSCAPS_CERTIFIED = C.DSCAPS_CERTIFIED
	// DSCAPS_SECONDARYMONO as defined in bass.h:218
	// DSCAPS_SECONDARYSTEREO as defined in bass.h:219
	// DSCAPS_SECONDARY8BIT as defined in bass.h:220
	DSCAPS_SECONDARY8BIT = 0x00000400
	// DSCAPS_SECONDARY16BIT as defined in bass.h:221
	DSCAPS_SECONDARY16BIT = 0x00000800
	// WAVE_FORMAT_1M08 as defined in bass.h:238
	WAVE_FORMAT_1M08 = C.WAVE_FORMAT_1M08
	// WAVE_FORMAT_1S08 as defined in bass.h:239
	WAVE_FORMAT_1S08 = C.WAVE_FORMAT_1S08
	// WAVE_FORMAT_1M16 as defined in bass.h:240
	WAVE_FORMAT_1M16 = C.WAVE_FORMAT_1M16
	// WAVE_FORMAT_1S16 as defined in bass.h:241
	WAVE_FORMAT_1S16 = C.WAVE_FORMAT_1S16
	// WAVE_FORMAT_2M08 as defined in bass.h:242
	WAVE_FORMAT_2M08 = C.WAVE_FORMAT_2M08
	// WAVE_FORMAT_2S08 as defined in bass.h:243
	WAVE_FORMAT_2S08 = C.WAVE_FORMAT_2S08
	// WAVE_FORMAT_2M16 as defined in bass.h:244
	WAVE_FORMAT_2M16 = C.WAVE_FORMAT_2M16
	// WAVE_FORMAT_2S16 as defined in bass.h:245
	WAVE_FORMAT_2S16 = C.WAVE_FORMAT_2S16
	// WAVE_FORMAT_4M08 as defined in bass.h:246
	WAVE_FORMAT_4M08 = C.WAVE_FORMAT_4M08
	// WAVE_FORMAT_4S08 as defined in bass.h:247
	WAVE_FORMAT_4S08 = C.WAVE_FORMAT_4S08
	// WAVE_FORMAT_4M16 as defined in bass.h:248
	WAVE_FORMAT_4M16 = C.WAVE_FORMAT_4M16
	// WAVE_FORMAT_4S16 as defined in bass.h:249
	WAVE_FORMAT_4S16 = C.WAVE_FORMAT_4M16
	// BASS_SAMPLE_8BITS as defined in bass.h:273
	SAMPLE_8BITS = C.BASS_SAMPLE_8BITS
	// BASS_SAMPLE_FLOAT as defined in bass.h:274
	SAMPLE_FLOAT = C.BASS_SAMPLE_FLOAT
	// BASS_SAMPLE_MONO as defined in bass.h:275
	SAMPLE_MONO = C.BASS_SAMPLE_MONO


	// BASS_SAMPLE_3D as defined in bass.h:277
	SAMPLE_3D = 8
	// BASS_SAMPLE_SOFTWARE as defined in bass.h:278
	SAMPLE_SOFTWARE = C.BASS_SAMPLE_SOFTWARE
	// BASS_SAMPLE_MUTEMAX as defined in bass.h:279
	SAMPLE_MUTEMAX = C.BASS_SAMPLE_MUTEMAX
	// BASS_SAMPLE_VAM as defined in bass.h:280
	SAMPLE_VAM = C.BASS_SAMPLE_VAM
	// BASS_SAMPLE_FX as defined in bass.h:281
	SAMPLE_FX = C.BASS_SAMPLE_FX
	// BASS_SAMPLE_OVER_POS as defined in bass.h:283
	// BASS_STREAM_PRESCAN as defined in bass.h:286
	STREAM_PRESCAN = C.BASS_STREAM_PRESCAN
	// BASS_MP3_SETPOS as defined in bass.h:287
	MP3_SETPOS = C.BASS_MP3_SETPOS
	// BASS_STREAM_AUTOFREE as defined in bass.h:288
	STREAM_AUTOFREE = C.BASS_STREAM_AUTOFREE
	// BASS_STREAM_RESTRATE as defined in bass.h:289
	STREAM_RESTRATE = C.BASS_STREAM_RESTRATE
	// BASS_STREAM_BLOCK as defined in bass.h:290
	STREAM_BLOCK = C.BASS_STREAM_BLOCK
	// BASS_STREAM_DECODE as defined in bass.h:291
	STREAM_DECODE = C.BASS_STREAM_DECODE
	// BASS_STREAM_STATUS as defined in bass.h:292
	STREAM_STATUS = C.BASS_STREAM_STATUS
	// BASS_SPEAKER_FRONT as defined in bass.h:318
	BASS_SPEAKER_FRONT = C.BASS_SPEAKER_FRONT
	// BASS_SPEAKER_REAR as defined in bass.h:319
	BASS_SPEAKER_REAR = C.BASS_SPEAKER_REAR
	// BASS_SPEAKER_CENLFE as defined in bass.h:320
	BASS_SPEAKER_CENLFE = C.BASS_SPEAKER_CENLFE
	// BASS_SPEAKER_REAR2 as defined in bass.h:321
	BASS_SPEAKER_REAR2 = 0x4000000
	// BASS_SPEAKER_LEFT as defined in bass.h:323
	BASS_SPEAKER_LEFT = C.BASS_SPEAKER_LEFT
	// BASS_SPEAKER_RIGHT as defined in bass.h:324
	BASS_SPEAKER_RIGHT = C.BASS_SPEAKER_RIGHT
	// BASS_SPEAKER_FRONTLEFT as defined in bass.h:325
	BASS_SPEAKER_FRONTLEFT = BASS_SPEAKER_FRONT | BASS_SPEAKER_LEFT
	// BASS_SPEAKER_FRONTRIGHT as defined in bass.h:326
	BASS_SPEAKER_FRONTRIGHT = BASS_SPEAKER_FRONT | BASS_SPEAKER_RIGHT
	// BASS_SPEAKER_REARLEFT as defined in bass.h:327
	BASS_SPEAKER_REARLEFT = BASS_SPEAKER_REAR | BASS_SPEAKER_LEFT
	// BASS_SPEAKER_REARRIGHT as defined in bass.h:328
	BASS_SPEAKER_REARRIGHT = BASS_SPEAKER_REAR | BASS_SPEAKER_RIGHT
	// BASS_SPEAKER_CENTER as defined in bass.h:329
	BASS_SPEAKER_CENTER = BASS_SPEAKER_CENLFE | BASS_SPEAKER_LEFT
	// BASS_SPEAKER_LFE as defined in bass.h:330
	BASS_SPEAKER_LFE = BASS_SPEAKER_CENLFE | BASS_SPEAKER_RIGHT
	// BASS_SPEAKER_REAR2LEFT as defined in bass.h:331
	BASS_SPEAKER_REAR2LEFT = BASS_SPEAKER_REAR2 | BASS_SPEAKER_LEFT
	// BASS_SPEAKER_REAR2RIGHT as defined in bass.h:332
	BASS_SPEAKER_REAR2RIGHT = BASS_SPEAKER_REAR2 | BASS_SPEAKER_RIGHT
	// BASS_ASYNCFILE as defined in bass.h:334
	BASS_ASYNCFILE = C.BASS_ASYNCFILE

	// BASS_RECORD_PAUSE as defined in bass.h:337
	BASS_RECORD_PAUSE = C.BASS_RECORD_PAUSE
	// BASS_RECORD_ECHOCANCEL as defined in bass.h:338
	BASS_RECORD_ECHOCANCEL = C.BASS_RECORD_ECHOCANCEL
	// BASS_RECORD_AGC as defined in bass.h:339
	BASS_RECORD_AGC = C.BASS_RECORD_AGC
	// BASS_VAM_HARDWARE as defined in bass.h:342
	BASS_VAM_HARDWARE = C.BASS_VAM_HARDWARE
	// BASS_VAM_SOFTWARE as defined in bass.h:343
	BASS_VAM_SOFTWARE = C.BASS_VAM_SOFTWARE
	// BASS_VAM_TERM_TIME as defined in bass.h:344
	BASS_VAM_TERM_TIME = C.BASS_VAM_TERM_TIME
	// BASS_VAM_TERM_DIST as defined in bass.h:345
	BASS_VAM_TERM_DIST = C.BASS_VAM_TERM_DIST
	// BASS_VAM_TERM_PRIO as defined in bass.h:346
	BASS_VAM_TERM_PRIO = C.BASS_VAM_TERM_PRIO
	// BASS_CTYPE_SAMPLE as defined in bass.h:361
	BASS_CTYPE_SAMPLE = C.BASS_CTYPE_SAMPLE
	// BASS_CTYPE_RECORD as defined in bass.h:362
	BASS_CTYPE_RECORD = C.BASS_CTYPE_RECORD
	// BASS_CTYPE_STREAM as defined in bass.h:363
	BASS_CTYPE_STREAM = C.BASS_CTYPE_STREAM
	// BASS_CTYPE_STREAM_OGG as defined in bass.h:364
	BASS_CTYPE_STREAM_OGG = C.BASS_CTYPE_STREAM_OGG
	// BASS_CTYPE_STREAM_MP1 as defined in bass.h:365
	BASS_CTYPE_STREAM_MP1 = 0x10003
	// BASS_CTYPE_STREAM_MP2 as defined in bass.h:366
	BASS_CTYPE_STREAM_MP2 = 0x10004
	// BASS_CTYPE_STREAM_MP3 as defined in bass.h:367
	BASS_CTYPE_STREAM_MP3 = 0x10005
	// BASS_CTYPE_STREAM_AIFF as defined in bass.h:368
	BASS_CTYPE_STREAM_AIFF = C.BASS_CTYPE_STREAM_AIFF
	// BASS_CTYPE_STREAM_CA as defined in bass.h:369
	BASS_CTYPE_STREAM_CA = C.BASS_CTYPE_STREAM_CA
	// BASS_CTYPE_STREAM_MF as defined in bass.h:370
	BASS_CTYPE_STREAM_MF = C.BASS_CTYPE_STREAM_MF
	// BASS_CTYPE_STREAM_WAV as defined in bass.h:371
	BASS_CTYPE_STREAM_WAV = C.BASS_CTYPE_STREAM_WAV
	// BASS_CTYPE_STREAM_WAV_PCM as defined in bass.h:372
	BASS_CTYPE_STREAM_WAV_PCM = C.BASS_CTYPE_STREAM_WAV_PCM
	// BASS_CTYPE_STREAM_WAV_FLOAT as defined in bass.h:373
	BASS_CTYPE_STREAM_WAV_FLOAT = C.BASS_CTYPE_STREAM_WAV_FLOAT
	// BASS_CTYPE_MUSIBASS_MOD as defined in bass.h:374
	BASS_CTYPE_MUSIBASS_MOD = C.BASS_CTYPE_MUSIC_MOD
	// BASS_CTYPE_MUSIBASS_MTM as defined in bass.h:375
	BASS_CTYPE_MUSIBASS_MTM = C.BASS_CTYPE_MUSIC_MTM
	// BASS_CTYPE_MUSIBASS_S3M as defined in bass.h:376
	BASS_CTYPE_MUSIBASS_S3M = 0x20002
	// BASS_CTYPE_MUSIBASS_XM as defined in bass.h:377
	BASS_CTYPE_MUSIBASS_XM = C.BASS_CTYPE_MUSIC_XM
	// BASS_CTYPE_MUSIBASS_IT as defined in bass.h:378
	BASS_CTYPE_MUSIBASS_IT = C.BASS_CTYPE_MUSIC_IT
	// BASS_CTYPE_MUSIBASS_MO3 as defined in bass.h:379
	BASS_CTYPE_MUSIBASS_MO3 = 0x00100
	// BASS_3DMODE_NORMAL as defined in bass.h:410
	BASS_3DMODE_NORMAL = 0
	// BASS_3DMODE_RELATIVE as defined in bass.h:411
	BASS_3DMODE_RELATIVE = 1
	// BASS_3DMODE_OFF as defined in bass.h:412
	BASS_3DMODE_OFF = 2
	// BASS_3DALG_DEFAULT as defined in bass.h:415
	BASS_3DALG_DEFAULT = 0
	// BASS_3DALG_OFF as defined in bass.h:416
	BASS_3DALG_OFF = 1
	// BASS_3DALG_FULL as defined in bass.h:417
	BASS_3DALG_FULL = 2
	// BASS_3DALG_LIGHT as defined in bass.h:418
	BASS_3DALG_LIGHT = 3
	// BASS_STREAMPROBASS_END as defined in bass.h:491
	BASS_STREAMPROBASS_END = C.BASS_STREAMPROC_END
	// STREAMFILE_NOBUFFER as defined in bass.h:498
	STREAMFILE_NOBUFFER = C.STREAMFILE_NOBUFFER
	// STREAMFILE_BUFFER as defined in bass.h:499
	STREAMFILE_BUFFER = C.STREAMFILE_BUFFER
	// STREAMFILE_BUFFERPUSH as defined in bass.h:500
	STREAMFILE_BUFFERPUSH = C.STREAMFILE_BUFFERPUSH
	// BASS_FILEDATA_END as defined in bass.h:516
	BASS_FILEDATA_END = C.BASS_FILEDATA_END
	// BASS_FILEPOS_CURRENT as defined in bass.h:519
	BASS_FILEPOS_CURRENT = C.BASS_FILEPOS_CURRENT
	// BASS_FILEPOS_DECODE as defined in bass.h:520
	BASS_FILEPOS_DECODE = BASS_FILEPOS_CURRENT
	// BASS_FILEPOS_DOWNLOAD as defined in bass.h:521
	BASS_FILEPOS_DOWNLOAD = C.BASS_FILEPOS_DOWNLOAD
	// BASS_FILEPOS_END as defined in bass.h:522
	BASS_FILEPOS_END = C.BASS_FILEPOS_END
	// BASS_FILEPOS_START as defined in bass.h:523
	BASS_FILEPOS_START = C.BASS_FILEPOS_START
	// BASS_FILEPOS_CONNECTED as defined in bass.h:524
	BASS_FILEPOS_CONNECTED = C.BASS_FILEPOS_CONNECTED
	// BASS_FILEPOS_BUFFER as defined in bass.h:525
	BASS_FILEPOS_BUFFER = C.BASS_FILEPOS_BUFFER
	// BASS_FILEPOS_SOCKET as defined in bass.h:526
	BASS_FILEPOS_SOCKET = C.BASS_FILEPOS_SOCKET
	// BASS_FILEPOS_ASYNCBUF as defined in bass.h:527
	BASS_FILEPOS_ASYNCBUF = C.BASS_FILEPOS_ASYNCBUF
	// BASS_FILEPOS_SIZE as defined in bass.h:528
	BASS_FILEPOS_SIZE = C.BASS_FILEPOS_SIZE
	// BASS_SYNBASS_POS as defined in bass.h:537
	BASS_SYNBASS_POS = C.BASS_SYNC_POS
	// BASS_SYNBASS_END as defined in bass.h:538
	BASS_SYNBASS_END = C.BASS_SYNC_END
	// BASS_SYNBASS_META as defined in bass.h:539
	BASS_SYNBASS_META = C.BASS_SYNC_META
	// BASS_SYNBASS_SLIDE as defined in bass.h:540
	BASS_SYNBASS_SLIDE = C.BASS_SYNC_SLIDE
	// BASS_SYNBASS_STALL as defined in bass.h:541
	BASS_SYNBASS_STALL = C.BASS_SYNC_STALL
	// BASS_SYNBASS_DOWNLOAD as defined in bass.h:542
	BASS_SYNBASS_DOWNLOAD = C.BASS_SYNC_DOWNLOAD
	// BASS_SYNBASS_FREE as defined in bass.h:543
	BASS_SYNBASS_FREE = C.BASS_SYNC_FREE
	// BASS_SYNBASS_SETPOS as defined in bass.h:544
	BASS_SYNBASS_SETPOS = C.BASS_SYNC_SETPOS
	// BASS_SYNBASS_MUSICPOS as defined in bass.h:545
	BASS_SYNBASS_MUSICPOS = C.BASS_SYNC_MUSICPOS
	// BASS_SYNBASS_MUSICINST as defined in bass.h:546
	BASS_SYNBASS_MUSICINST = C.BASS_SYNC_MUSICINST
	// BASS_SYNBASS_MUSICFX as defined in bass.h:547
	BASS_SYNBASS_MUSICFX = C.BASS_SYNC_MUSICFX
	// BASS_SYNBASS_OGG_CHANGE as defined in bass.h:548
	BASS_SYNBASS_OGG_CHANGE = C.BASS_SYNC_OGG_CHANGE
	// BASS_SYNBASS_MIXTIME as defined in bass.h:549
	BASS_SYNBASS_MIXTIME = C.BASS_SYNC_MIXTIME
	// BASS_SYNBASS_ONETIME as defined in bass.h:550
	BASS_SYNBASS_ONETIME = C.BASS_SYNC_ONETIME
	// BASS_ACTIVE_STOPPED as defined in bass.h:581
	ACTIVE_STOPPED = C.BASS_ACTIVE_STOPPED
	// BASS_ACTIVE_PLAYING as defined in bass.h:582
	ACTIVE_PLAYING = C.BASS_ACTIVE_PLAYING
	// BASS_ACTIVE_STALLED as defined in bass.h:583
	ACTIVE_STALLED = C.BASS_ACTIVE_STALLED
	// BASS_ACTIVE_PAUSED as defined in bass.h:584
	ACTIVE_PAUSED = C.BASS_ACTIVE_PAUSED
	// BASS_ATTRIB_FREQ as defined in bass.h:587
	ATTRIB_FREQ = C.BASS_ATTRIB_FREQ
	// BASS_ATTRIB_VOL as defined in bass.h:588
	ATTRIB_VOL = C.BASS_ATTRIB_VOL
	// BASS_ATTRIB_PAN as defined in bass.h:589
	ATTRIB_PAN = C.BASS_ATTRIB_PAN
	// BASS_ATTRIB_EAXMIX as defined in bass.h:590
	ATTRIB_EAXMIX = C.BASS_ATTRIB_EAXMIX
	// BASS_ATTRIB_NOBUFFER as defined in bass.h:591
	ATTRIB_NOBUFFER = C.BASS_ATTRIB_NOBUFFER
	// BASS_ATTRIB_VBR as defined in bass.h:592
	ATTRIB_VBR = C.BASS_ATTRIB_VBR
	// BASS_ATTRIB_CPU as defined in bass.h:593
	ATTRIB_CPU = C.BASS_ATTRIB_CPU
	// BASS_ATTRIB_SRC as defined in bass.h:594
	ATTRIB_SRC = C.BASS_ATTRIB_SRC
	// BASS_ATTRIB_NET_RESUME as defined in bass.h:595
	ATTRIB_NET_RESUME = C.BASS_ATTRIB_NET_RESUME
	// BASS_ATTRIB_SCANINFO as defined in bass.h:596
	ATTRIB_SCANINFO = C.BASS_ATTRIB_SCANINFO
	// BASS_ATTRIB_NORAMP as defined in bass.h:597
	ATTRIB_NORAMP = C.BASS_ATTRIB_NORAMP
	// BASS_ATTRIB_BITRATE as defined in bass.h:598
	ATTRIB_BITRATE = C.BASS_ATTRIB_BITRATE
	// BASS_ATTRIB_MUSIBASS_AMPLIFY as defined in bass.h:599
	ATTRIB_MUSIBASS_AMPLIFY = C.BASS_ATTRIB_MUSIC_AMPLIFY
	// BASS_ATTRIB_MUSIBASS_PANSEP as defined in bass.h:600
	ATTRIB_MUSIBASS_PANSEP = C.BASS_ATTRIB_MUSIC_PANSEP
	// BASS_ATTRIB_MUSIBASS_PSCALER as defined in bass.h:601
	ATTRIB_MUSIBASS_PSCALER = C.BASS_ATTRIB_MUSIC_PSCALER
	// BASS_ATTRIB_MUSIBASS_BPM as defined in bass.h:602
	ATTRIB_MUSIBASS_BPM = C.BASS_ATTRIB_MUSIC_BPM
	// BASS_ATTRIB_MUSIBASS_SPEED as defined in bass.h:603
	ATTRIB_MUSIBASS_SPEED = C.BASS_ATTRIB_MUSIC_SPEED
	// BASS_ATTRIB_MUSIBASS_VOL_GLOBAL as defined in bass.h:604
	ATTRIB_MUSIBASS_VOL_GLOBAL = C.BASS_ATTRIB_MUSIC_VOL_GLOBAL
	// BASS_ATTRIB_MUSIBASS_ACTIVE as defined in bass.h:605
	ATTRIB_MUSIBASS_ACTIVE = C.BASS_ATTRIB_MUSIC_ACTIVE
	// BASS_ATTRIB_MUSIBASS_VOL_CHAN as defined in bass.h:606
	ATTRIB_MUSIBASS_VOL_CHAN = C.BASS_ATTRIB_MUSIC_VOL_CHAN
	// BASS_ATTRIB_MUSIBASS_VOL_INST as defined in bass.h:607
	ATTRIB_MUSIBASS_VOL_INST = C.BASS_ATTRIB_MUSIC_VOL_INST
	// BASS_DATA_AVAILABLE as defined in bass.h:610
	BASS_DATA_AVAILABLE = C.BASS_DATA_AVAILABLE
	// BASS_DATA_FIXED as defined in bass.h:611
	BASS_DATA_FIXED = C.BASS_DATA_FIXED
	// BASS_DATA_FLOAT as defined in bass.h:612
	BASS_DATA_FLOAT = C.BASS_DATA_FLOAT
	// BASS_DATA_FFT256 as defined in bass.h:613
	BASS_DATA_FFT256 = 0x80000000
	// BASS_DATA_FFT512 as defined in bass.h:614
	BASS_DATA_FFT512 = 0x80000001
	// BASS_DATA_FFT1024 as defined in bass.h:615
	BASS_DATA_FFT1024 = 0x80000002
	// BASS_DATA_FFT2048 as defined in bass.h:616
	BASS_DATA_FFT2048 = 0x80000003
	// BASS_DATA_FFT4096 as defined in bass.h:617
	BASS_DATA_FFT4096 = 0x80000004
	// BASS_DATA_FFT8192 as defined in bass.h:618
	BASS_DATA_FFT8192 = 0x80000005
	// BASS_DATA_FFT16384 as defined in bass.h:619
	BASS_DATA_FFT16384 = 0x80000006
	// BASS_DATA_FFT32768 as defined in bass.h:620
	BASS_DATA_FFT32768 = 0x80000007
	// BASS_DATA_FFT_INDIVIDUAL as defined in bass.h:621
	BASS_DATA_FFT_INDIVIDUAL = C.BASS_DATA_FFT_INDIVIDUAL
	// BASS_DATA_FFT_NOWINDOW as defined in bass.h:622
	BASS_DATA_FFT_NOWINDOW = C.BASS_DATA_FFT_NOWINDOW
	// BASS_DATA_FFT_REMOVEDC as defined in bass.h:623
	BASS_DATA_FFT_REMOVEDC = C.BASS_DATA_FFT_REMOVEDC
	// BASS_DATA_FFT_COMPLEX as defined in bass.h:624
	BASS_DATA_FFT_COMPLEX = C.BASS_DATA_FFT_COMPLEX
	// BASS_LEVEL_MONO as defined in bass.h:627
	BASS_LEVEL_MONO = C.BASS_LEVEL_MONO
	// BASS_LEVEL_STEREO as defined in bass.h:628
	BASS_LEVEL_STEREO = C.BASS_LEVEL_STEREO
	// BASS_LEVEL_RMS as defined in bass.h:629
	BASS_LEVEL_RMS = C.BASS_LEVEL_RMS
	// BASS_TAG_ID3 as defined in bass.h:632
	BASS_TAG_ID3 = 0
	// BASS_TAG_ID3V2 as defined in bass.h:633
	BASS_TAG_ID3V2 = 1
	// BASS_TAG_OGG as defined in bass.h:634
	BASS_TAG_OGG = C.BASS_TAG_OGG
	// BASS_TAG_HTTP as defined in bass.h:635
	BASS_TAG_HTTP = C.BASS_TAG_HTTP
	// BASS_TAG_ICY as defined in bass.h:636
	BASS_TAG_ICY = C.BASS_TAG_ICY
	// BASS_TAG_META as defined in bass.h:637
	BASS_TAG_META = C.BASS_TAG_META
	// BASS_TAG_APE as defined in bass.h:638
	BASS_TAG_APE = C.BASS_TAG_APE
	// BASS_TAG_MP4 as defined in bass.h:639
	BASS_TAG_MP4 = 7
	// BASS_TAG_WMA as defined in bass.h:640
	BASS_TAG_WMA = C.BASS_TAG_WMA
	// BASS_TAG_VENDOR as defined in bass.h:641
	BASS_TAG_VENDOR = C.BASS_TAG_VENDOR
	// BASS_TAG_LYRICS3 as defined in bass.h:642
	BASS_TAG_LYRICS3 = 10
	// BASS_TAG_CA_CODEC as defined in bass.h:643
	BASS_TAG_CA_CODEC = C.BASS_TAG_CA_CODEC
	// BASS_TAG_MF as defined in bass.h:644
	BASS_TAG_MF = C.BASS_TAG_MF
	// BASS_TAG_WAVEFORMAT as defined in bass.h:645
	BASS_TAG_WAVEFORMAT = C.BASS_TAG_WAVEFORMAT
	// BASS_TAG_RIFF_INFO as defined in bass.h:646
	BASS_TAG_RIFF_INFO = C.BASS_TAG_RIFF_INFO
	// BASS_TAG_RIFF_BEXT as defined in bass.h:647
	BASS_TAG_RIFF_BEXT = C.BASS_TAG_RIFF_BEXT
	// BASS_TAG_RIFF_CART as defined in bass.h:648
	BASS_TAG_RIFF_CART = C.BASS_TAG_RIFF_CART
	// BASS_TAG_RIFF_DISP as defined in bass.h:649
	BASS_TAG_RIFF_DISP = C.BASS_TAG_RIFF_DISP
	// BASS_TAG_APE_BINARY as defined in bass.h:650
	BASS_TAG_APE_BINARY = C.BASS_TAG_APE_BINARY
	// BASS_TAG_MUSIBASS_NAME as defined in bass.h:651
	BASS_TAG_MUSIBASS_NAME = C.BASS_TAG_MUSIC_NAME
	// BASS_TAG_MUSIBASS_MESSAGE as defined in bass.h:652
	BASS_TAG_MUSIBASS_MESSAGE = C.BASS_TAG_MUSIC_MESSAGE
	// BASS_TAG_MUSIBASS_ORDERS as defined in bass.h:653
	BASS_TAG_MUSIBASS_ORDERS = C.BASS_TAG_MUSIC_ORDERS
	// BASS_TAG_MUSIBASS_AUTH as defined in bass.h:654
	BASS_TAG_MUSIBASS_AUTH = C.BASS_TAG_MUSIC_AUTH
	// BASS_TAG_MUSIBASS_INST as defined in bass.h:655
	BASS_TAG_MUSIBASS_INST = C.BASS_TAG_MUSIC_INST
	// BASS_TAG_MUSIBASS_SAMPLE as defined in bass.h:656
	BASS_TAG_MUSIBASS_SAMPLE = C.BASS_TAG_MUSIC_SAMPLE
	// BASS_POS_BYTE as defined in bass.h:767
	POS_BYTE = C.BASS_POS_BYTE
	// BASS_POS_MUSIBASS_ORDER as defined in bass.h:768
	POS_MUSIBASS_ORDER = C.BASS_POS_MUSIC_ORDER
	// BASS_POS_OGG as defined in bass.h:769
	POS_OGG = C.BASS_POS_OGG
	// BASS_POS_INEXACT as defined in bass.h:770
	POS_INEXACT = C.BASS_POS_INEXACT
	// BASS_POS_DECODE as defined in bass.h:771
	POS_DECODE = C.BASS_POS_DECODE
	// BASS_POS_DECODETO as defined in bass.h:772
	POS_DECODETO = C.BASS_POS_DECODETO
	// BASS_POS_SCAN as defined in bass.h:773
	POS_SCAN = C.BASS_POS_SCAN
	// BASS_INPUT_OFF as defined in bass.h:776
	BASS_INPUT_OFF = C.BASS_INPUT_OFF
	// BASS_INPUT_ON as defined in bass.h:777
	BASS_INPUT_ON = C.BASS_INPUT_ON
	// BASS_INPUT_TYPE_MASK as defined in bass.h:779
	BASS_INPUT_TYPE_MASK = C.BASS_INPUT_TYPE_MASK
	// BASS_INPUT_TYPE_UNDEF as defined in bass.h:780
	BASS_INPUT_TYPE_UNDEF = C.BASS_INPUT_TYPE_UNDEF
	// BASS_INPUT_TYPE_DIGITAL as defined in bass.h:781
	BASS_INPUT_TYPE_DIGITAL = C.BASS_INPUT_TYPE_DIGITAL
	// BASS_INPUT_TYPE_LINE as defined in bass.h:782
	BASS_INPUT_TYPE_LINE = C.BASS_INPUT_TYPE_LINE
	// BASS_INPUT_TYPE_MIC as defined in bass.h:783
	BASS_INPUT_TYPE_MIC = C.BASS_INPUT_TYPE_MIC
	// BASS_INPUT_TYPE_SYNTH as defined in bass.h:784
	BASS_INPUT_TYPE_SYNTH = C.BASS_INPUT_TYPE_SYNTH
	// BASS_INPUT_TYPE_CD as defined in bass.h:785
	BASS_INPUT_TYPE_CD = C.BASS_INPUT_TYPE_CD
	// BASS_INPUT_TYPE_PHONE as defined in bass.h:786
	BASS_INPUT_TYPE_PHONE = C.BASS_INPUT_TYPE_PHONE
	// BASS_INPUT_TYPE_SPEAKER as defined in bass.h:787
	BASS_INPUT_TYPE_SPEAKER = C.BASS_INPUT_TYPE_SPEAKER
	// BASS_INPUT_TYPE_WAVE as defined in bass.h:788
	BASS_INPUT_TYPE_WAVE = C.BASS_INPUT_TYPE_WAVE
	// BASS_INPUT_TYPE_AUX as defined in bass.h:789
	BASS_INPUT_TYPE_AUX = C.BASS_INPUT_TYPE_AUX
	// BASS_INPUT_TYPE_ANALOG as defined in bass.h:790
	BASS_INPUT_TYPE_ANALOG = C.BASS_INPUT_TYPE_ANALOG
	// BASS_DX8_PHASE_NEG_180 as defined in bass.h:884
	BASS_DX8_PHASE_NEG_180 = C.BASS_DX8_PHASE_NEG_180
	// BASS_DX8_PHASE_NEG_90 as defined in bass.h:885
	BASS_DX8_PHASE_NEG_90 = C.BASS_DX8_PHASE_NEG_90
	// BASS_DX8_PHASE_ZERO as defined in bass.h:886
	BASS_DX8_PHASE_ZERO = C.BASS_DX8_PHASE_ZERO
	// BASS_DX8_PHASE_90 as defined in bass.h:887
	BASS_DX8_PHASE_90 = C.BASS_DX8_PHASE_90
	// BASS_DX8_PHASE_180 as defined in bass.h:888
	BASS_DX8_PHASE_180 = C.BASS_DX8_PHASE_180
	// BASS_IOSNOTIFY_INTERRUPT as defined in bass.h:894
	BASS_IOSNOTIFY_INTERRUPT = C.BASS_IOSNOTIFY_INTERRUPT
	// BASS_IOSNOTIFY_INTERRUPT_END as defined in bass.h:895
	BASS_IOSNOTIFY_INTERRUPT_END = C.BASS_IOSNOTIFY_INTERRUPT_END

	// BASS_FX_DX8_CHORUS as declared in bass.h:795
	BASS_FX_DX8_CHORUS = C.BASS_FX_DX8_CHORUS
	// BASS_FX_DX8_COMPRESSOR as declared in bass.h:796
	BASS_FX_DX8_COMPRESSOR = C.BASS_FX_DX8_COMPRESSOR
	// BASS_FX_DX8_DISTORTION as declared in bass.h:797
	BASS_FX_DX8_DISTORTION = C.BASS_FX_DX8_DISTORTION
	// BASS_FX_DX8_ECHO as declared in bass.h:798
	BASS_FX_DX8_ECHO = C.BASS_FX_DX8_ECHO
	// BASS_FX_DX8_FLANGER as declared in bass.h:799
	BASS_FX_DX8_FLANGER = C.BASS_FX_DX8_FLANGER
	// BASS_FX_DX8_GARGLE as declared in bass.h:800
	BASS_FX_DX8_GARGLE = C.BASS_FX_DX8_GARGLE
	// BASS_FX_DX8_I3DL2REVERB as declared in bass.h:801
	BASS_FX_DX8_I3DL2REVERB = C.BASS_FX_DX8_I3DL2REVERB
	// BASS_FX_DX8_PARAMEQ as declared in bass.h:802
	BASS_FX_DX8_PARAMEQ = C.BASS_FX_DX8_PARAMEQ
	// BASS_FX_DX8_REVERB as declared in bass.h:803
	BASS_FX_DX8_REVERB = C.BASS_FX_DX8_REVERB
	SYNC_DEV_FAIL=C.BASS_SYNC_DEV_FAIL
	SYNC_DEV_FORMAT=C.BASS_SYNC_DEV_FORMAT
	SYNC_DOWNLOAD=C.BASS_SYNC_DOWNLOAD
	SYNC_END=C.BASS_SYNC_END
	SYNC_FREE=C.BASS_SYNC_FREE
	SYNC_META=C.BASS_SYNC_META

	SYNC_MUSICFX=C.BASS_SYNC_MUSICFX
	SYNC_MUSICINST=C.BASS_SYNC_MUSICINST
	SYNC_MUSICPOS=C.BASS_SYNC_MUSICPOS
	SYNC_OGG_CHANGE=C.BASS_SYNC_OGG_CHANGE
	SYNC_POS=C.BASS_SYNC_POS
	SYNC_SETPOS=C.BASS_SYNC_SETPOS
	SYNC_SLIDE=C.BASS_SYNC_SLIDE
	SYNC_STALL=C.BASS_SYNC_STALL
	// BASS_SampleGetChannel flags
	SamchanNew = C.BASS_SAMCHAN_NEW
	SAMCHAN_STREAM = C.BASS_SAMCHAN_STREAM
	SAMPLE_LOOP = C.BASS_SAMPLE_LOOP
	SAMPLE_OVER_VOL = C.BASS_SAMPLE_OVER_VOL
	SAMPLE_OVER_POS = C.BASS_SAMPLE_OVER_POS
	SAMPLE_OVER_DIST = C.BASS_SAMPLE_OVER_DIST
	CONFIG_REC_WASAPI = C.BASS_CONFIG_REC_WASAPI
	CONFIG_REC_BUFFER = C.BASS_CONFIG_REC_BUFFER
	DATA_AVAILABLE = C.BASS_DATA_AVAILABLE
	ATTRIB_BUFFER = C.BASS_ATTRIB_BUFFER
	SYNC_MIXTIME = C.BASS_SYNC_MIXTIME
	SYNC_ONETIME = C.BASS_SYNC_ONETIME
	SYNC_THREAD = C.BASS_SYNC_THREAD
)