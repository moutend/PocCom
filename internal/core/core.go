package core

import (
	"log"

	"github.com/go-ole/go-ole"
	"github.com/moutend/PocCom/pkg/com"
	"github.com/moutend/PocCom/pkg/types"
)

var (
	FocusElement uintptr
	isRunning    bool
	server       *com.IPocCom
)

func Setup(logServerAddr string) error {
	if isRunning {
		return nil
	}
	if err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		return err
	}
	if err := com.CoCreateInstance(com.CLSID_PocCom, 0, com.CLSCTX_ALL, com.IID_IPocCom, &server); err != nil {
		return err
	}
	if err := server.Start(logServerAddr, 0); err != nil {
		return err
	}

	isRunning = true

	log.Println("core: Setup() is done")

	return nil
}

func Teardown() error {
	server.Stop()
	server.Release()
	ole.CoUninitialize()

	isRunning = false

	log.Println("core: Teardown() is done")

	return nil
}

func SetMSAAEventHandler(fn com.MSAAEventHandler) {
	server.SetMSAAEventHandler(fn)
}
func SetUIAEventHandler(fn com.UIAEventHandler) {
	server.SetUIAEventHandler(fn)
}

func GetIUIAutomationElement(direction types.TreeWalkerDirection, pRootElement uintptr) (pElement *com.IUIAutomationElement, err error) {
	return server.GetIUIAutomationElement(direction, pRootElement)
}

func UpdateTreeWalker() error {
	return server.UpdateTreeWalker()
}
