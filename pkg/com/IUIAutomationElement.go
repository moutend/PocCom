package com

import (
	"unsafe"

	"github.com/moutend/PocCom/pkg/types"

	"github.com/go-ole/go-ole"
)

type IUIAutomationElement struct {
	ole.IUnknown
}

type IUIAutomationElementVtbl struct {
	ole.IUnknownVtbl
	SetFocus                    uintptr
	GetRuntimeId                uintptr
	FindFirst                   uintptr
	FindAll                     uintptr
	FindFirstBuildCache         uintptr
	FindAllBuildCache           uintptr
	BuildUpdatedCache           uintptr
	GetCurrentPropertyValue     uintptr
	GetCurrentPropertyValueEx   uintptr
	GetCachedPropertyValue      uintptr
	GetCachedPropertyValueEx    uintptr
	GetCurrentPatternAs         uintptr
	GetCachedPatternAs          uintptr
	GetCurrentPattern           uintptr
	GetCachedPattern            uintptr
	GetCachedParent             uintptr
	GetCachedChildren           uintptr
	CurrentProcessId            uintptr
	CurrentControlType          uintptr
	CurrentLocalizedControlType uintptr
	CurrentName                 uintptr
	CurrentAcceleratorKey       uintptr
	CurrentAccessKey            uintptr
	CurrentHasKeyboardFocus     uintptr
	CurrentIsKeyboardFocusable  uintptr
	CurrentIsEnabled            uintptr
	CurrentAutomationId         uintptr
	CurrentClassName            uintptr
	CurrentHelpText             uintptr
	CurrentCulture              uintptr
	CurrentIsControlElement     uintptr
	CurrentIsContentElement     uintptr
	CurrentIsPassword           uintptr
	CurrentNativeWindowHandle   uintptr
	CurrentItemType             uintptr
	CurrentIsOffscreen          uintptr
	CurrentOrientation          uintptr
	CurrentFrameworkId          uintptr
	CurrentIsRequiredForForm    uintptr
	CurrentItemStatus           uintptr
	CurrentBoundingRectangle    uintptr
	CurrentLabeledBy            uintptr
	CurrentAriaRole             uintptr
	CurrentAriaProperties       uintptr
	CurrentIsDataValidForForm   uintptr
	CurrentControllerFor        uintptr
	CurrentDescribedBy          uintptr
	CurrentFlowsTo              uintptr
	CurrentProviderDescription  uintptr
	CachedProcessId             uintptr
	CachedControlType           uintptr
	CachedLocalizedControlType  uintptr
	CachedName                  uintptr
	CachedAcceleratorKey        uintptr
	CachedAccessKey             uintptr
	CachedHasKeyboardFocus      uintptr
	CachedIsKeyboardFocusable   uintptr
	CachedIsEnabled             uintptr
	CachedAutomationId          uintptr
	CachedClassName             uintptr
	CachedHelpText              uintptr
	CachedCulture               uintptr
	CachedIsControlElement      uintptr
	CachedIsContentElement      uintptr
	CachedIsPassword            uintptr
	CachedNativeWindowHandle    uintptr
	CachedItemType              uintptr
	CachedIsOffscreen           uintptr
	CachedOrientation           uintptr
	CachedFrameworkId           uintptr
	CachedIsRequiredForForm     uintptr
	CachedItemStatus            uintptr
	CachedBoundingRectangle     uintptr
	CachedLabeledBy             uintptr
	CachedAriaRole              uintptr
	CachedAriaProperties        uintptr
	CachedIsDataValidForForm    uintptr
	CachedControllerFor         uintptr
	CachedDescribedBy           uintptr
	CachedFlowsTo               uintptr
	CachedProviderDescription   uintptr
	GetClickablePoint           uintptr
}

func (v *IUIAutomationElement) VTable() *IUIAutomationElementVtbl {
	return (*IUIAutomationElementVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationElement) SetFocus() error {
	return uiaeSetFocus(v)
}

func (v *IUIAutomationElement) GetRuntimeId() error {
	return uiaeGetRuntimeId(v)
}

func (v *IUIAutomationElement) FindFirst() error {
	return uiaeFindFirst(v)
}

func (v *IUIAutomationElement) FindAll() error {
	return uiaeFindAll(v)
}

func (v *IUIAutomationElement) FindFirstBuildCache() error {
	return uiaeFindFirstBuildCache(v)
}

func (v *IUIAutomationElement) FindAllBuildCache() error {
	return uiaeFindAllBuildCache(v)
}

func (v *IUIAutomationElement) BuildUpdatedCache() error {
	return uiaeBuildUpdatedCache(v)
}

func (v *IUIAutomationElement) GetCurrentPropertyValue() error {
	return uiaeGetCurrentPropertyValue(v)
}

func (v *IUIAutomationElement) GetCurrentPropertyValueEx() error {
	return uiaeGetCurrentPropertyValueEx(v)
}

func (v *IUIAutomationElement) GetCachedPropertyValue() error {
	return uiaeGetCachedPropertyValue(v)
}

func (v *IUIAutomationElement) GetCachedPropertyValueEx() error {
	return uiaeGetCachedPropertyValueEx(v)
}

func (v *IUIAutomationElement) GetCurrentPatternAs() error {
	return uiaeGetCurrentPatternAs(v)
}

func (v *IUIAutomationElement) GetCachedPatternAs() error {
	return uiaeGetCachedPatternAs(v)
}

func (v *IUIAutomationElement) GetCurrentPattern() error {
	return uiaeGetCurrentPattern(v)
}

func (v *IUIAutomationElement) GetCachedPattern() error {
	return uiaeGetCachedPattern(v)
}

func (v *IUIAutomationElement) GetCachedParent() error {
	return uiaeGetCachedParent(v)
}

func (v *IUIAutomationElement) GetCachedChildren() error {
	return uiaeGetCachedChildren(v)
}

func (v *IUIAutomationElement) CurrentProcessId() error {
	return uiaeCurrentProcessId(v)
}

func (v *IUIAutomationElement) CurrentControlType() (types.UIAControlType, error) {
	return uiaeCurrentControlType(v)
}

func (v *IUIAutomationElement) CurrentLocalizedControlType() (types.BSTR, error) {
	return uiaeCurrentLocalizedControlType(v)
}

func (v *IUIAutomationElement) CurrentName() (types.BSTR, error) {
	return uiaeCurrentName(v)
}

func (v *IUIAutomationElement) CurrentAcceleratorKey() (types.BSTR, error) {
	return uiaeCurrentAcceleratorKey(v)
}

func (v *IUIAutomationElement) CurrentAccessKey() (types.BSTR, error) {
	return uiaeCurrentAccessKey(v)
}

func (v *IUIAutomationElement) CurrentHasKeyboardFocus() error {
	return uiaeCurrentHasKeyboardFocus(v)
}

func (v *IUIAutomationElement) CurrentIsKeyboardFocusable() (bool, error) {
	return uiaeCurrentIsKeyboardFocusable(v)
}

func (v *IUIAutomationElement) CurrentIsEnabled() (bool, error) {
	return uiaeCurrentIsEnabled(v)
}

func (v *IUIAutomationElement) CurrentAutomationId() (types.BSTR, error) {
	return uiaeCurrentAutomationId(v)
}

func (v *IUIAutomationElement) CurrentClassName() (types.BSTR, error) {
	return uiaeCurrentClassName(v)
}

func (v *IUIAutomationElement) CurrentHelpText() (types.BSTR, error) {
	return uiaeCurrentHelpText(v)
}

func (v *IUIAutomationElement) CurrentCulture() error {
	return uiaeCurrentCulture(v)
}

func (v *IUIAutomationElement) CurrentIsControlElement() (bool, error) {
	return uiaeCurrentIsControlElement(v)
}

func (v *IUIAutomationElement) CurrentIsContentElement() (bool, error) {
	return uiaeCurrentIsContentElement(v)
}

func (v *IUIAutomationElement) CurrentIsPassword() (bool, error) {
	return uiaeCurrentIsPassword(v)
}

func (v *IUIAutomationElement) CurrentNativeWindowHandle() (uintptr, error) {
	return uiaeCurrentNativeWindowHandle(v)
}

func (v *IUIAutomationElement) CurrentItemType() (types.BSTR, error) {
	return uiaeCurrentItemType(v)
}

func (v *IUIAutomationElement) CurrentIsOffscreen() (bool, error) {
	return uiaeCurrentIsOffscreen(v)
}

func (v *IUIAutomationElement) CurrentOrientation() error {
	return uiaeCurrentOrientation(v)
}

func (v *IUIAutomationElement) CurrentFrameworkId() (types.BSTR, error) {
	return uiaeCurrentFrameworkId(v)
}

func (v *IUIAutomationElement) CurrentIsRequiredForForm() (bool, error) {
	return uiaeCurrentIsRequiredForForm(v)
}

func (v *IUIAutomationElement) CurrentItemStatus() (types.BSTR, error) {
	return uiaeCurrentItemStatus(v)
}

func (v *IUIAutomationElement) CurrentBoundingRectangle() (*types.RECT, error) {
	return uiaeCurrentBoundingRectangle(v)
}

func (v *IUIAutomationElement) CurrentLabeledBy() error {
	return uiaeCurrentLabeledBy(v)
}

func (v *IUIAutomationElement) CurrentAriaRole() (types.BSTR, error) {
	return uiaeCurrentAriaRole(v)
}

func (v *IUIAutomationElement) CurrentAriaProperties() (types.BSTR, error) {
	return uiaeCurrentAriaProperties(v)
}

func (v *IUIAutomationElement) CurrentIsDataValidForForm() (bool, error) {
	return uiaeCurrentIsDataValidForForm(v)
}

func (v *IUIAutomationElement) CurrentControllerFor() error {
	return uiaeCurrentControllerFor(v)
}

func (v *IUIAutomationElement) CurrentDescribedBy() error {
	return uiaeCurrentDescribedBy(v)
}

func (v *IUIAutomationElement) CurrentFlowsTo() error {
	return uiaeCurrentFlowsTo(v)
}

func (v *IUIAutomationElement) CurrentProviderDescription() (types.BSTR, error) {
	return uiaeCurrentProviderDescription(v)
}

func (v *IUIAutomationElement) CachedProcessId() error {
	return uiaeCachedProcessId(v)
}

func (v *IUIAutomationElement) CachedControlType() error {
	return uiaeCachedControlType(v)
}

func (v *IUIAutomationElement) CachedLocalizedControlType() (types.BSTR, error) {
	return uiaeCachedLocalizedControlType(v)
}

func (v *IUIAutomationElement) CachedName() (types.BSTR, error) {
	return uiaeCachedName(v)
}

func (v *IUIAutomationElement) CachedAcceleratorKey() (types.BSTR, error) {
	return uiaeCachedAcceleratorKey(v)
}

func (v *IUIAutomationElement) CachedAccessKey() (types.BSTR, error) {
	return uiaeCachedAccessKey(v)
}

func (v *IUIAutomationElement) CachedHasKeyboardFocus() error {
	return uiaeCachedHasKeyboardFocus(v)
}

func (v *IUIAutomationElement) CachedIsKeyboardFocusable() (bool, error) {
	return uiaeCachedIsKeyboardFocusable(v)
}

func (v *IUIAutomationElement) CachedIsEnabled() (bool, error) {
	return uiaeCachedIsEnabled(v)
}

func (v *IUIAutomationElement) CachedAutomationId() (types.BSTR, error) {
	return uiaeCachedAutomationId(v)
}

func (v *IUIAutomationElement) CachedClassName() (types.BSTR, error) {
	return uiaeCachedClassName(v)
}

func (v *IUIAutomationElement) CachedHelpText() (types.BSTR, error) {
	return uiaeCachedHelpText(v)
}

func (v *IUIAutomationElement) CachedCulture() error {
	return uiaeCachedCulture(v)
}

func (v *IUIAutomationElement) CachedIsControlElement() (bool, error) {
	return uiaeCachedIsControlElement(v)
}

func (v *IUIAutomationElement) CachedIsContentElement() (bool, error) {
	return uiaeCachedIsContentElement(v)
}

func (v *IUIAutomationElement) CachedIsPassword() (bool, error) {
	return uiaeCachedIsPassword(v)
}

func (v *IUIAutomationElement) CachedNativeWindowHandle() (uintptr, error) {
	return uiaeCachedNativeWindowHandle(v)
}

func (v *IUIAutomationElement) CachedItemType() (types.BSTR, error) {
	return uiaeCachedItemType(v)
}

func (v *IUIAutomationElement) CachedIsOffscreen() (bool, error) {
	return uiaeCachedIsOffscreen(v)
}

func (v *IUIAutomationElement) CachedOrientation() error {
	return uiaeCachedOrientation(v)
}

func (v *IUIAutomationElement) CachedFrameworkId() (types.BSTR, error) {
	return uiaeCachedFrameworkId(v)
}

func (v *IUIAutomationElement) CachedIsRequiredForForm() (bool, error) {
	return uiaeCachedIsRequiredForForm(v)
}

func (v *IUIAutomationElement) CachedItemStatus() (types.BSTR, error) {
	return uiaeCachedItemStatus(v)
}

func (v *IUIAutomationElement) CachedBoundingRectangle() (*types.RECT, error) {
	return uiaeCachedBoundingRectangle(v)
}

func (v *IUIAutomationElement) CachedLabeledBy() error {
	return uiaeCachedLabeledBy(v)
}

func (v *IUIAutomationElement) CachedAriaRole() (types.BSTR, error) {
	return uiaeCachedAriaRole(v)
}

func (v *IUIAutomationElement) CachedAriaProperties() (types.BSTR, error) {
	return uiaeCachedAriaProperties(v)
}

func (v *IUIAutomationElement) CachedIsDataValidForForm() (bool, error) {
	return uiaeCachedIsDataValidForForm(v)
}

func (v *IUIAutomationElement) CachedControllerFor() error {
	return uiaeCachedControllerFor(v)
}

func (v *IUIAutomationElement) CachedDescribedBy() error {
	return uiaeCachedDescribedBy(v)
}

func (v *IUIAutomationElement) CachedFlowsTo() error {
	return uiaeCachedFlowsTo(v)
}

func (v *IUIAutomationElement) CachedProviderDescription() (types.BSTR, error) {
	return uiaeCachedProviderDescription(v)
}

func (v *IUIAutomationElement) GetClickablePoint() error {
	return uiaeGetClickablePoint(v)
}
