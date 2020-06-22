// +build !windows

package com

import "github.com/go-ole/go-ole"

func accGetAccParent(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accGetAccChildCount(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accGetAccChild(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func accGetAccName(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
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

func accAccLocation(v *IAccessible) error {
	return ole.NewError(ole.E_NOTIMPL)
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
