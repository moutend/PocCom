// +build windows

package com

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/moutend/PocCom/pkg/types"
)

func accGetAccParent(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accGetAccChildCount(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accGetAccChild(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accGetAccName(v *IAccessible, child ole.VARIANT) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().GetAccName,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&child)),
		uintptr(unsafe.Pointer(&bstr)))

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func accGetAccValue(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accGetAccDescription(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accGetAccRole(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accGetAccState(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accGetAccHelp(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accGetAccHelpTopic(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accGetAccKeyboardShortcut(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accGetAccFocus(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accGetAccSelection(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accGetAccDefaultAction(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accAccSelect(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accAccLocation(v *IAccessible, child ole.VARIANT) (left, top, width, height int32, err error) {
	hr, _, _ := syscall.Syscall6(
		v.VTable().AccLocation,
		6,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&left)),
		uintptr(unsafe.Pointer(&top)),
		uintptr(unsafe.Pointer(&width)),
		uintptr(unsafe.Pointer(&height)),
		uintptr(unsafe.Pointer(&child)))

	if hr != 0 {
		err = ole.NewError(hr)
	}

	return
}

func accAccNavigate(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accAccHitTest(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accAccDoDefaultAction(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accPutAccName(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accPutAccValue(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}
