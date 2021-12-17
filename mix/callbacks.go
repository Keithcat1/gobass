package mix
import (
	"runtime/cgo"
	"unsafe"
	"github.com/keithcat1/gobass"
)

/*
#include "bass.h"
#include "_gobass_callbacks.h"

*/
import "C"
type GoSyncproc = func(bass.Sync, bass.Channel, int)


//export _GoSyncprocCallback
func _GoSyncprocCallback(sync C.HSYNC, channel C.HCHANNEL, data C.DWORD, userdata unsafe.Pointer) {
fn := cgo.Handle(uintptr(userdata)).Value().(GoSyncproc)
	fn(bass.Sync(sync), bass.ChannelFromHandle(uint32(channel)), int(data))
}
var (
	goSyncprocCallback *C.SYNCPROC = C._get_GoSyncprocCallbackWrapper()
)