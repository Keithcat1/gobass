package bass
import (
	"runtime/cgo"
	"unsafe"
)

/*
#include "bass.h"
#cgo CFLAGS: -Wno-pointer-to-int-cast
extern RECORDPROC* _get_GoBassRecordCallbackStreamPutData();
extern SYNCPROC* _get_GoSyncprocCallbackWrapper();
extern SYNCPROC* _get_SyncprocFree();
*/
import "C"
type GoSyncproc = func(sync Sync, channel Channel, data int)


//export _GoSyncprocCallback
func _GoSyncprocCallback(sync C.HSYNC, channel C.HCHANNEL, data C.DWORD, userdata unsafe.Pointer) {
fn := cgo.Handle(uintptr(userdata)).Value().(GoSyncproc)
	fn(Sync(sync), Channel{handle: channel}, int(data))
}
var (
	STREAMPROC_DEVICE *C.STREAMPROC = C.STREAMPROC_DEVICE
	STREAMPROC_PUSH *C.STREAMPROC = C.STREAMPROC_PUSH
	STREAMPROC_DUMMY *C.STREAMPROC = C.STREAMPROC_DUMMY
	RecordCallbackStreamPutData *C.RECORDPROC = C._get_GoBassRecordCallbackStreamPutData()
	goSyncprocCallback *C.SYNCPROC = C._get_GoSyncprocCallbackWrapper()
SyncprocFree *C.SYNCPROC = C._get_SyncprocFree()
)