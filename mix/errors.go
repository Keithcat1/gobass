package mix
import (
	"github.com/keithcat1/gobass"
)
/*
#include "bassmix.h"
*/
import "C"

func (self Mixer) toError() (Mixer, error) {
	if self.GetHandle() == 0 {
		return Mixer{}, bass.GetLastError()
	} else {
		return self, nil
	}
}
func boolToError(value C.BOOL) error {
	if value == 0 {
		return bass.GetLastError()
	} else {
		return nil
	}
}