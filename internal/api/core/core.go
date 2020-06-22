package core

import (
	"log"
	"net/http"

	"github.com/moutend/CoreServer/internal/core"
	"github.com/moutend/CoreServer/pkg/types"
)

func putCursorNextElement(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json")

	if core.FocusElement != 0 {
		pElement, err := core.GetIUIAutomationElement(types.TW_NEXT, core.FocusElement)

		log.Println("@@@next", err, pElement)
	} else {
		log.Println("@@@next is not available")
	}

	response := "{\"error\": \"Not implemented\"}"
	http.Error(w, response, http.StatusInternalServerError)
}

func putCursorPreviousElement(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json")

	response := "{\"error\": \"Not implemented\"}"
	http.Error(w, response, http.StatusInternalServerError)
}

func putCursorFirstChildElement(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json")

	response := "{\"error\": \"Not implemented\"}"
	http.Error(w, response, http.StatusInternalServerError)
}

func putCursorLastChildElement(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json")

	response := "{\"error\": \"Not implemented\"}"
	http.Error(w, response, http.StatusInternalServerError)
}

func putCursorParentElement(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json")

	response := "{\"error\": \"Not implemented\"}"
	http.Error(w, response, http.StatusInternalServerError)
}

func putUpdateTreeWalker(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json")

	err := core.UpdateTreeWalker()

	if err != nil {
		log.Println("@@@err", err)
	}

	response := "{\"error\": \"Not implemented\"}"
	http.Error(w, response, http.StatusInternalServerError)
}
