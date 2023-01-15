package main

import (
	"go_module/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	chi := chi.NewRouter()

	chi.Get("/",handlers.Home)
	chi.Get("/about",handlers.About)
	chi.Get("/getdata",handlers.GetData)
	chi.Post("/postdata",handlers.PostData)


	http.ListenAndServe(":8000",chi)
}