package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"imageService/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/image/{name}", handler.GetImage).Methods("GET")
	r.Use(handler.Middleware)
	r.HandleFunc("/image", handler.SaveImage).Methods("POST")
	r.HandleFunc("/image/{name}", handler.DeleteImage).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
