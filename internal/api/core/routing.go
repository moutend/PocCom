package core

import "github.com/go-chi/chi"

func Setup(router chi.Router) {
	router.Route("/core", func(r chi.Router) {
		r.Put("/next", putCursorNextElement)
		r.Put("/previous", putCursorPreviousElement)
		r.Put("/first", putCursorFirstChildElement)
		r.Put("/last", putCursorLastChildElement)
		r.Put("/parent", putCursorParentElement)
		r.Put("/update", putUpdateTreeWalker)
	})
}
