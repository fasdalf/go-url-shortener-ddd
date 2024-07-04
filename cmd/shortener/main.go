package main

import (
	"github.com/vokinneberg/go-url-shortener-ddd/internal/repository"
	"log"
	"net/http"

	"github.com/vokinneberg/go-url-shortener-ddd/internal/api"
)

func main() {
	rep := repository.NewMemRepository()
	httpHandler := api.NewHandler(rep, rep)
	log.Fatal(http.ListenAndServe(":8080", httpHandler))
}
