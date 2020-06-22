// +build !windows

package com

import (
	"github.com/go-ole/go-ole"
)

func pcStart(v *IPocCom, logServerAddr string, logLevel int) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func pcStop(v *IPocCom) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func pcSetMSAAEventHandler(v *IPocCom, handleFunc MSAAEventHandler) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func pcSetIUIEventHandler(v *IPocCom, handleFunc UIAEventHandler) error {
	return ole.NewError(ole.E_NOTIMPL)
}
