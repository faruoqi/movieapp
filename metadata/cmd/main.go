package main

import (
	"github.com/faruoqi/movieapp/metadata/internal/controller/metadata"
	httphandler "github.com/faruoqi/movieapp/metadata/internal/handler/http"
	"github.com/faruoqi/movieapp/metadata/internal/repository/memory"
	"log"
	"net/http"
)

func main() {

	log.Println("Starting Metadata Service")
	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
