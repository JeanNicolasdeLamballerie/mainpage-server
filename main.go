package main

import (
	//	"errors"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	routers "main/server/routers"
	//	"io"
	"net/http"
)

// func getRoot(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("got / request\n")
// 	io.WriteString(w, "This is my website!\n")
// }
// func getHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("got /hello request\n")
// 	io.WriteString(w, "Hello, HTTP!\n")
// }

func main() {
	fmt.Printf("Running Go server...")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	r.Mount("/v1", routers.GetRouter("v1"))

	http.ListenAndServe(":4000", r)
}
