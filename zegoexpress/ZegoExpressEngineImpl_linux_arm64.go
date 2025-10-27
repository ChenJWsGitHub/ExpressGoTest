//go:build linux && arm64

package zegoexpress

/*
#cgo CFLAGS: -I${SRCDIR}/lib/include
#cgo LDFLAGS: -L${SRCDIR}/lib/linux-arm64 -lZegoExpressEngine -Wl,-rpath=${SRCDIR}/lib/linux-arm64
#include "zego-express-engine.h"
*/
import "C"

func getVersion() string {
	var cVersion *C.char
	C.zego_express_get_version(&cVersion)
	return C.GoString(cVersion)
}
