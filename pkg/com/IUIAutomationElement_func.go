// +build !windows

package com

import (
	"github.com/go-ole/go-ole"
	"github.com/moutend/PocCom/pkg/types"
)

func uiaeSetFocus(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetRuntimeId(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
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

func uiaeCurrentControlType(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentLocalizedControlType(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentName(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentAcceleratorKey(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentAccessKey(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentHasKeyboardFocus(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsKeyboardFocusable(v *IUIAutomationElement) (bool, error) {
	return false, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsEnabled(v *IUIAutomationElement) (bool, error) {
	return false, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentAutomationId(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentClassName(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentHelpText(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentCulture(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsControlElement(v *IUIAutomationElement) (bool, error) {
	return false, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsContentElement(v *IUIAutomationElement) (bool, error) {
	return false, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsPassword(v *IUIAutomationElement) (bool, error) {
	return false, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentNativeWindowHandle(v *IUIAutomationElement) (uintptr, error) {
	return 0, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentItemType(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsOffscreen(v *IUIAutomationElement) (bool, error) {
	return false, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentOrientation(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentFrameworkId(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsRequiredForForm(v *IUIAutomationElement) (bool, error) {
	return false, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentItemStatus(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentBoundingRectangle(v *IUIAutomationElement) (*types.RECT, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentLabeledBy(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentAriaRole(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentAriaProperties(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsDataValidForForm(v *IUIAutomationElement) (bool, error) {
	return false, ole.NewError(ole.E_NOTIMPL)
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
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedProcessId(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedControlType(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedLocalizedControlType(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedName(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedAcceleratorKey(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedAccessKey(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedHasKeyboardFocus(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsKeyboardFocusable(v *IUIAutomationElement) (bool, error) {
	return false, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsEnabled(v *IUIAutomationElement) (bool, error) {
	return false, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedAutomationId(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedClassName(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedHelpText(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedCulture(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsControlElement(v *IUIAutomationElement) (bool, error) {
	return false, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsContentElement(v *IUIAutomationElement) (bool, error) {
	return false, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsPassword(v *IUIAutomationElement) (bool, error) {
	return false, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedNativeWindowHandle(v *IUIAutomationElement) (uintptr, error) {
	return 0, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedItemType(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsOffscreen(v *IUIAutomationElement) (bool, error) {
	return false, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedOrientation(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedFrameworkId(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsRequiredForForm(v *IUIAutomationElement) (bool, error) {
	return false, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedItemStatus(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedBoundingRectangle(v *IUIAutomationElement) (*types.RECT, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedLabeledBy(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedAriaRole(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedAriaProperties(v *IUIAutomationElement) (types.BSTR, error) {
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsDataValidForForm(v *IUIAutomationElement) (bool, error) {
	return false, ole.NewError(ole.E_NOTIMPL)
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
	return types.BSTR(0), ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetClickablePoint(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}
