package app

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"time"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/moutend/PocCom/internal/core"
	"github.com/moutend/PocCom/internal/util"
	"github.com/moutend/PocCom/pkg/com"
	"github.com/moutend/PocCom/pkg/types"

	"github.com/go-chi/chi"
	"github.com/moutend/PocCom/internal/api"
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:  "todoapp",
	RunE: rootRunE,
}

func rootRunE(cmd *cobra.Command, args []string) error {
	rand.Seed(time.Now().Unix())
	p := make([]byte, 16)

	if _, err := rand.Read(p); err != nil {
		return err
	}

	myself, err := user.Current()

	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("AudioServer-%s.txt", hex.EncodeToString(p))
	outputPath := filepath.Join(myself.HomeDir, "AppData", "Roaming", "ScreenReaderX", "Logs", "SystemLog", fileName)
	output := util.NewBackgroundWriter(outputPath)

	defer output.Close()

	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Llongfile)
	log.SetOutput(output)

	var logServerConfig struct {
		Addr string `json:"addr"`
	}

	logServerConfigPath := filepath.Join(myself.HomeDir, "AppData", "Roaming", "ScreenReaderX", "Server", "LogServer.json")

	logServerConfigData, err := ioutil.ReadFile(logServerConfigPath)

	if err != nil {
		return err
	}
	if err := json.Unmarshal(logServerConfigData, &logServerConfig); err != nil {
		return err
	}
	if err := core.Setup(logServerConfig.Addr); err != nil {
		return err
	}

	core.SetMSAAEventHandler(func(eventId types.MSAAEvent, childId int64, pInterface uintptr) int64 {
		return 0
		go http.Post("http://127.0.0.1:7902/v1/audio", "application/json", bytes.NewBufferString(`{"isForcePush":true,"commands": [{"type": 1, "value":10}]}`))
		e := (*com.IAccessible)(unsafe.Pointer(pInterface))
		child := ole.NewVariant(ole.VT_I4, childId)

		name, err := e.GetAccName(child)

		if err != nil {
			log.Println("@@@err", err)
			return 0
		}
		log.Printf("@@@Event:%q,Name:%q\n", eventId, name)

		return 0
	})
	core.SetUIAEventHandler(func(eventId types.UIAEvent, pInterface uintptr) int64 {
		e := (*com.IUIAutomationElement)(unsafe.Pointer(pInterface))

		rect, err := e.CurrentBoundingRectangle()

		if err != nil {
			log.Println("@@@err", err)
			return 0
		}
		if false && rect.IsZero() {
			log.Println("@@@skipped")
			return 0
		}

		name, _ := e.CurrentName()
		localizedControlType, _ := e.CurrentLocalizedControlType()

		go http.Post("http://127.0.0.1:7902/v1/audio", "application/json", bytes.NewBufferString(
			fmt.Sprintf(`{"isForcePush":true,"commands":[{"type":1,"sfxIndex":9},{"type":3,"textToSpeech":%q},{"type":2,"sleepDuration":750},{"type":3,"textToSpeech":%q}]}`,
				name,
				localizedControlType,
			)))

		// className, _ := e.CurrentClassName()
		// framework, _ := e.CurrentFrameworkId()
		// itemType, _ := e.CurrentItemType()
		// ariaRole, _ := e.CurrentAriaRole()
		// ariaProperties, _ := e.CurrentAriaProperties()
		controlType, _ := e.CurrentControlType()

		log.Printf("@@@Event:%q,Name:%q,Control:%q,ControlName:%q,RECT:%+v\n", eventId, name, controlType, localizedControlType, rect)

		return 0
	})

	defer core.Teardown()

	router := chi.NewRouter()
	api.Setup(router)

	listener, err := net.Listen("tcp", "127.0.0.1:0")

	if err != nil {
		return err
	}

	serverAddr := listener.Addr().(*net.TCPAddr).String()

	serverConfig, err := json.Marshal(struct {
		Addr string `json:"addr"`
	}{
		Addr: serverAddr,
	})

	if err != nil {
		return err
	}

	serverConfigPath := filepath.Join(myself.HomeDir, "AppData", "Roaming", "ScreenReaderX", "Server", "AudioServer.json")
	os.MkdirAll(filepath.Dir(serverConfigPath), 0755)

	if err := ioutil.WriteFile(serverConfigPath, serverConfig, 0644); err != nil {
		return err
	}

	log.Printf("Listening on %s\n", serverAddr)

	return http.Serve(listener, router)
}
