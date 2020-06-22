// +build windows

package types

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func bstrString(p BSTR) string {
	if p == 0 {
		return ""
	}

	length := ole.SysStringLen((*int16)(unsafe.Pointer(p)))

	if length <= 0 {
		return ""
	}

	u16s := make([]uint16, length)

	for i := 0; i < int(length); i++ {
		u16s[i] = *(*uint16)(unsafe.Pointer(uintptr(p) + uintptr(i*2)))
	}

	return syscall.UTF16ToString(u16s)
}
