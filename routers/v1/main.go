package v1

import (
	v1Routes "main/server/routers/v1/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func V1Router() http.Handler {
	v1router := chi.NewRouter()
	v1router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	v1router.Mount("/", v1Routes.V1Blog())
	return v1router
}
