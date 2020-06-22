// +build windows

package com

import (
	"syscall"
	"unsafe"

	"github.com/moutend/PocCom/pkg/types"

	"github.com/go-ole/go-ole"
)

func uiaeSetFocus(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetRuntimeId(v *IUIAutomationElement) error {
	var safearray *ole.SAFEARRAY

	hr, _, _ := syscall.Syscall(
		v.VTable().GetRuntimeId,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&safearray)),
		0)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func uiaeFindFirst(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeFindAll(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeFindFirstBuildCache(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeFindAllBuildCache(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeBuildUpdatedCache(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCurrentPropertyValue(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCurrentPropertyValueEx(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCachedPropertyValue(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCachedPropertyValueEx(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCurrentPatternAs(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCachedPatternAs(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCurrentPattern(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCachedPattern(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCachedParent(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCachedChildren(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentProcessId(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentControlType(v *IUIAutomationElement) (types.UIAControlType, error) {
	var controlType types.UIAControlType

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentControlType,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&controlType)),
		0)

	if hr != 0 {
		return controlType, ole.NewError(hr)
	}

	return controlType, nil
}

func uiaeCurrentLocalizedControlType(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentLocalizedControlType,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCurrentName(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentName,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCurrentAcceleratorKey(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentAcceleratorKey,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCurrentAccessKey(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentAccessKey,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCurrentHasKeyboardFocus(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsKeyboardFocusable(v *IUIAutomationElement) (bool, error) {
	var result int64

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentIsKeyboardFocusable,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&result)),
		0)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return result == 1, nil
}

func uiaeCurrentIsEnabled(v *IUIAutomationElement) (bool, error) {
	var result int64

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentIsEnabled,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&result)),
		0)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return result == 1, nil
}

func uiaeCurrentAutomationId(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentAutomationId,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCurrentClassName(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentClassName,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCurrentHelpText(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentHelpText,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCurrentCulture(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsControlElement(v *IUIAutomationElement) (bool, error) {
	var result int64

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentIsControlElement,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&result)),
		0)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return result == 1, nil
}

func uiaeCurrentIsContentElement(v *IUIAutomationElement) (bool, error) {
	var result int64

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentIsContentElement,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&result)),
		0)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return result == 1, nil
}

func uiaeCurrentIsPassword(v *IUIAutomationElement) (bool, error) {
	var result int64

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentIsPassword,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&result)),
		0)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return result == 1, nil
}

func uiaeCurrentNativeWindowHandle(v *IUIAutomationElement) (uintptr, error) {
	var hwnd uintptr

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentNativeWindowHandle,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&hwnd)),
		0)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return hwnd, nil
}

func uiaeCurrentItemType(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentItemType,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCurrentIsOffscreen(v *IUIAutomationElement) (bool, error) {
	var result int64

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentIsOffscreen,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&result)),
		0)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return result == 1, nil
}

func uiaeCurrentOrientation(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentFrameworkId(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentFrameworkId,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCurrentIsRequiredForForm(v *IUIAutomationElement) (bool, error) {
	var result int64

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentIsRequiredForForm,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&result)),
		0)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return result == 1, nil
}

func uiaeCurrentItemStatus(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentItemStatus,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCurrentBoundingRectangle(v *IUIAutomationElement) (*types.RECT, error) {
	var rect types.RECT

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentBoundingRectangle,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&rect)),
		0)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return &rect, nil
}

func uiaeCurrentLabeledBy(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentAriaRole(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentAriaRole,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCurrentAriaProperties(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentAriaProperties,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCurrentIsDataValidForForm(v *IUIAutomationElement) (bool, error) {
	var result int64

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentIsDataValidForForm,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&result)),
		0)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return result == 1, nil
}

func uiaeCurrentControllerFor(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentDescribedBy(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentFlowsTo(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentProviderDescription(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentProviderDescription,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCachedProcessId(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedControlType(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedLocalizedControlType(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedLocalizedControlType,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCachedName(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentName,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCachedAcceleratorKey(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedAcceleratorKey,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCachedAccessKey(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedAccessKey,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCachedHasKeyboardFocus(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsKeyboardFocusable(v *IUIAutomationElement) (bool, error) {
	var result int64

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedIsKeyboardFocusable,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&result)),
		0)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return result == 1, nil
}

func uiaeCachedIsEnabled(v *IUIAutomationElement) (bool, error) {
	var result int64

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedIsEnabled,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&result)),
		0)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return result == 1, nil
}

func uiaeCachedAutomationId(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedAutomationId,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCachedClassName(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedClassName,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCachedHelpText(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedHelpText,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCachedCulture(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsControlElement(v *IUIAutomationElement) (bool, error) {
	var result int64

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedIsControlElement,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&result)),
		0)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return result == 1, nil
}

func uiaeCachedIsContentElement(v *IUIAutomationElement) (bool, error) {
	var result int64

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedIsContentElement,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&result)),
		0)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return result == 1, nil
}

func uiaeCachedIsPassword(v *IUIAutomationElement) (bool, error) {
	var result int64

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedIsPassword,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&result)),
		0)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return result == 1, nil
}

func uiaeCachedNativeWindowHandle(v *IUIAutomationElement) (uintptr, error) {
	var hwnd uintptr

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedNativeWindowHandle,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&hwnd)),
		0)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return hwnd, nil
}

func uiaeCachedItemType(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedItemType,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCachedIsOffscreen(v *IUIAutomationElement) (bool, error) {
	var result int64

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedIsOffscreen,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&result)),
		0)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return result == 1, nil
}

func uiaeCachedOrientation(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedFrameworkId(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedFrameworkId,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCachedIsRequiredForForm(v *IUIAutomationElement) (bool, error) {
	var result int64

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedIsRequiredForForm,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&result)),
		0)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return result == 1, nil
}

func uiaeCachedItemStatus(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedItemStatus,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCachedBoundingRectangle(v *IUIAutomationElement) (*types.RECT, error) {
	var rect types.RECT

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedBoundingRectangle,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&rect)),
		0)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return &rect, nil
}

func uiaeCachedLabeledBy(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedAriaRole(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedAriaRole,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCachedAriaProperties(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedAriaProperties,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeCachedIsDataValidForForm(v *IUIAutomationElement) (bool, error) {
	var result int64

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedIsDataValidForForm,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&result)),
		0)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return result == 1, nil
}

func uiaeCachedControllerFor(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedDescribedBy(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedFlowsTo(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedProviderDescription(v *IUIAutomationElement) (types.BSTR, error) {
	var bstr types.BSTR

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedProviderDescription,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)

	if hr != 0 {
		return bstr, ole.NewError(hr)
	}

	return bstr, nil
}

func uiaeGetClickablePoint(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}
