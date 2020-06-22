package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"unsafe"

	"github.com/moutend/CoreServer/pkg/types"

	"github.com/moutend/CoreServer/pkg/com"

	"github.com/go-ole/go-ole"
)

func main() {
	log.SetPrefix("error: ")
	log.SetFlags(0)

	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	if err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		return err
	}
	defer ole.CoUninitialize()

	var foo *com.ICoreServer

	if err := com.CoCreateInstance(com.CLSID_CoreServer, 0, com.CLSCTX_ALL, com.IID_ICoreServer, &foo); err != nil {
		return err
	}

	fmt.Println("@@@instance created")

	defer foo.Release()

	err := foo.Start()
	fmt.Println("Called ICoreServer::Start", err)
	msaaCount := 0
	foo.SetMSAAEventHandler(func(eventId types.MSAAEvent, pInterface uintptr) int64 {
		msaaCount += 1
		return 0
		e := (*com.IAccessible)(unsafe.Pointer(pInterface))
		child := ole.NewVariant(ole.VT_I4, 0)

		left, top, width, height, err := e.AccLocation(child)

		if err != nil {
			fmt.Println("@@@err", err)
			return 0
		}

		name, err := e.GetAccName(child)

		if err != nil {
			fmt.Println("@@@err", err)
			return 0
		}

		fmt.Printf("@@@Event:%q,Name:%q,Location:{%d,%d,%d,%d}\n", eventId, name.String(), left, top, width, height)
		return 0
	})
	uiaCount := 0
	uiaMap := make(map[uintptr]bool)
	foo.SetUIAEventHandler(func(eventId types.UIAEvent, pInterface uintptr) int64 {
		uiaCount += 1
		uiaMap[pInterface] = true
		go http.Post("http://192.168.1.102:7902/v1/audio", "application/json", bytes.NewBufferString(`{"isForcePush":true,"commands": [{"type": 1, "value":10}]}`))
		return 0
		e := (*com.IUIAutomationElement)(unsafe.Pointer(pInterface))

		rect, err := e.CurrentBoundingRectangle()

		if err != nil {
			fmt.Println("@@@err", err)
			return 0
		}
		if rect.IsZero() {
			fmt.Println("@@@skipped")
			return 0
		}

		name, _ := e.CurrentName()
		className, _ := e.CurrentClassName()
		framework, _ := e.CurrentFrameworkId()
		itemType, _ := e.CurrentItemType()
		ariaRole, _ := e.CurrentAriaRole()
		ariaProperties, _ := e.CurrentAriaProperties()
		return 0
		fmt.Printf("@@@Event:%q,Name:%q,ClassName:%q,Framework:%q,ItemType:%q,AriaRole:%q,AriaProperties:%q\n", eventId, name, className, framework, itemType, ariaRole, ariaProperties)

		return 0
	})

	time.Sleep(30 * time.Second)
	{
		n := 0
		for _, v := range uiaMap {
			if v {
				n += 1
			}
		}
		fmt.Println("@@@uniq count", n)
	}
	err = foo.Stop()
	fmt.Println("Called Foo::Stop()", msaaCount, uiaCount)

	return nil
}
