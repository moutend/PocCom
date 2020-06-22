package com

import (
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/moutend/PocCom/pkg/types"
)

type IAccessible struct {
	ole.IDispatch
}

type IAccessibleVtbl struct {
	ole.IDispatchVtbl
	GetAccParent           uintptr
	GetAccChildCount       uintptr
	GetAccChild            uintptr
	GetAccName             uintptr
	GetAccValue            uintptr
	GetAccDescription      uintptr
	GetAccRole             uintptr
	GetAccState            uintptr
	GetAccHelp             uintptr
	GetAccHelpTopic        uintptr
	GetAccKeyboardShortcut uintptr
	GetAccFocus            uintptr
	GetAccSelection        uintptr
	GetAccDefaultAction    uintptr
	AccSelect              uintptr
	AccLocation            uintptr
	AccNavigate            uintptr
	AccHitTest             uintptr
	AccDoDefaultAction     uintptr
	PutAccName             uintptr
	PutAccValue            uintptr
}

func (v *IAccessible) VTable() *IAccessibleVtbl {
	return (*IAccessibleVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAccessible) GetAccParent() error {
	return accGetAccParent(v)
}

func (v *IAccessible) GetAccChildCount() error {
	return accGetAccChildCount(v)
}

func (v *IAccessible) GetAccChild() error {
	return accGetAccChild(v)
}

func (v *IAccessible) GetAccName(child ole.VARIANT) (types.BSTR, error) {
	return accGetAccName(v, child)
}

func (v *IAccessible) GetAccValue() error {
	return accGetAccValue(v)
}

func (v *IAccessible) GetAccDescription() error {
	return accGetAccDescription(v)
}

func (v *IAccessible) GetAccRole() error {
	return accGetAccRole(v)
}

func (v *IAccessible) GetAccState() error {
	return accGetAccState(v)
}

func (v *IAccessible) GetAccHelp() error {
	return accGetAccHelp(v)
}

func (v *IAccessible) GetAccHelpTopic() error {
	return accGetAccHelpTopic(v)
}

func (v *IAccessible) GetAccKeyboardShortcut() error {
	return accGetAccKeyboardShortcut(v)
}

func (v *IAccessible) GetAccFocus() error {
	return accGetAccFocus(v)
}

func (v *IAccessible) GetAccSelection() error {
	return accGetAccSelection(v)
}

func (v *IAccessible) GetAccDefaultAction() error {
	return accGetAccDefaultAction(v)
}

func (v *IAccessible) AccSelect() error {
	return accAccSelect(v)
}

func (v *IAccessible) AccLocation(child ole.VARIANT) (left, top, width, height int32, err error) {
	return accAccLocation(v, child)
}

func (v *IAccessible) AccNavigate() error {
	return accAccNavigate(v)
}

func (v *IAccessible) AccHitTest() error {
	return accAccHitTest(v)
}

func (v *IAccessible) AccDoDefaultAction() error {
	return accAccDoDefaultAction(v)
}

func (v *IAccessible) PutAccName() error {
	return accPutAccName(v)
}

func (v *IAccessible) PutAccValue() error {
	return accPutAccValue(v)
}
