// +build windows

package com

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/moutend/PocCom/pkg/types"
)

func pcStart(v *IPocCom, logServerAddr string, logLevel int) error {
	u16LogServerAddr, err := syscall.UTF16FromString(logServerAddr)

	if err != nil {
		return err
	}

	hr, _, _ := syscall.Syscall6(
		v.VTable().Start,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&u16LogServerAddr[0])),
		uintptr(logLevel),
		0,
		0,
		0)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func pcStop(v *IPocCom) error {
	hr, _, _ := syscall.Syscall(
		v.VTable().Stop,
		0,
		uintptr(unsafe.Pointer(v)),
		0,
		0)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func pcSetMSAAEventHandler(v *IPocCom, handleFunc MSAAEventHandler) error {
	hr, _, _ := syscall.Syscall(
		v.VTable().SetMSAAEventHandler,
		2,
		uintptr(unsafe.Pointer(v)),
		syscall.NewCallback(handleFunc),
		0)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func pcSetUIAEventHandler(v *IPocCom, handleFunc UIAEventHandler) error {
	hr, _, _ := syscall.Syscall(
		v.VTable().SetUIAEventHandler,
		2,
		uintptr(unsafe.Pointer(v)),
		syscall.NewCallback(handleFunc),
		0)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func pcGetIUIAutomationElement(v *IPocCom, direction types.TreeWalkerDirection, pRootElement uintptr) (pElement *IUIAutomationElement, err error) {
	hr, _, _ := syscall.Syscall6(
		v.VTable().GetIUIAutomationElement,
		4,
		uintptr(unsafe.Pointer(v)),
		uintptr(direction),
		pRootElement,
		uintptr(unsafe.Pointer(&pElement)),
		0,
		0)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return pElement, nil
}

func pcUpdateTreeWalker(v *IPocCom) error {
	hr, _, _ := syscall.Syscall(
		v.VTable().UpdateTreeWalker,
		1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}
