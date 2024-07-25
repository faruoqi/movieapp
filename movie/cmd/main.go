package main

import (
	"github.com/faruoqi/movieapp/movie/internal/controller/movie"
	metadatagateway "github.com/faruoqi/movieapp/movie/internal/gateway/metadata/http"
	ratinggateway "github.com/faruoqi/movieapp/movie/internal/gateway/rating/http"
	httphandler "github.com/faruoqi/movieapp/movie/internal/handler/http"
	"log"
	"net/http"
)

func main() {

	log.Println("Starting the movie service")
	metadatagw := metadatagateway.New("localhost:8081")
	ratinggw := ratinggateway.New("localhost:8082")
	ctrl := movie.New(ratinggw, metadatagw)
	h := httphandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
