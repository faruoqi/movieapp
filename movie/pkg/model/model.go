package model

import "github.com/faruoqi/movieapp/metadata/pkg/model"

type MovieDetails struct {
	Rating   *float64       `json:"rating,omitempty"`
	Metadata model.Metadata `json:"metadata"`
}
