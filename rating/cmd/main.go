package main

import (
	"github.com/faruoqi/movieapp/rating/internal/controller/rating"
	httphandler "github.com/faruoqi/movieapp/rating/internal/handler/http"
	"github.com/faruoqi/movieapp/rating/internal/repository/memory"
	"log"
	"net/http"
)

func main() {

	log.Println("Starting Rating Service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}

}
