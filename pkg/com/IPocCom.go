package com

import (
	"unsafe"

	"github.com/moutend/PocCom/pkg/types"

	"github.com/go-ole/go-ole"
)

type MSAAEventHandler func(eventId types.MSAAEvent, childId int64, pInterface uintptr) int64

type UIAEventHandler func(eventId types.UIAEvent, pInterface uintptr) int64

type IPocCom struct {
	ole.IDispatch
}

type IPocComVtbl struct {
	ole.IDispatchVtbl
	Start                   uintptr
	Stop                    uintptr
	SetMSAAEventHandler     uintptr
	SetUIAEventHandler      uintptr
	GetIUIAutomationElement uintptr
	UpdateTreeWalker        uintptr
}

func (v *IPocCom) VTable() *IPocComVtbl {
	return (*IPocComVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IPocCom) Start(logServerAddr string, logLevel int) error {
	return pcStart(v, logServerAddr, logLevel)
}

func (v *IPocCom) Stop() error {
	return pcStop(v)
}

func (v *IPocCom) SetMSAAEventHandler(handleFunc MSAAEventHandler) error {
	return pcSetMSAAEventHandler(v, handleFunc)
}

func (v *IPocCom) SetUIAEventHandler(handleFunc UIAEventHandler) error {
	return pcSetUIAEventHandler(v, handleFunc)
}

func (v *IPocCom) GetIUIAutomationElement(direction types.TreeWalkerDirection, pRootElement uintptr) (pElement *IUIAutomationElement, err error) {
	return pcGetIUIAutomationElement(v, direction, pRootElement)
}

func (v *IPocCom) UpdateTreeWalker() error {
	return pcUpdateTreeWalker(v)
}
