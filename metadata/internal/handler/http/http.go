package http

import (
	"encoding/json"
	"errors"
	"github.com/faruoqi/movieapp/metadata/internal/controller/metadata"
	"log"
	"net/http"
)

type Handler struct {
	ctrl *metadata.Controller
}

func New(ctrl *metadata.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GetMetadata(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := req.Context()
	m, err := h.ctrl.Get(ctx, id)
	if err != nil && errors.Is(err, metadata.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response Encode Error %w", err)
	}

}
