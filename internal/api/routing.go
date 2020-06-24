package api

import (
	"github.com/go-chi/chi"
	"github.com/moutend/PocCom/internal/api/core"
)

func Setup(router chi.Router) {
	router.Route("/v1", func(r chi.Router) {
		core.Setup(r)
	})
}
