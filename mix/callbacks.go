package mix
import (
	"runtime/cgo"
	"unsafe"
	"github.com/keithcat1/gobass"
)

/*
#include "bass.h"
extern SYNCPROC* _get_MixSyncprocCallbackWrapper();
*/
import "C"



//export _MixSyncprocCallback
func _MixSyncprocCallback(sync C.HSYNC, channel C.HCHANNEL, data C.DWORD, userdata unsafe.Pointer) {
fn := cgo.Handle(uintptr(userdata)).Value().(bass.GoSyncproc)
	fn(bass.Sync(sync), bass.ChannelFromHandle(uint32(channel)), int(data))
}
var (
	mixSyncprocCallback *C.SYNCPROC = C._get_MixSyncprocCallbackWrapper()
)