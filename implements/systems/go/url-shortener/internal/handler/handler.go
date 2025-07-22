package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

type URLService interface {
	Shorten(string) (string, error)
	Resolve(string) (string, error)
}

type Handler struct {
	service URLService
}

func NewHandler(s URLService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)
	shortURL, _ := h.service.Shorten(req.URL)
	json.NewEncoder(w).Encode(map[string]string{"short_url": shortURL})
}

func (h *Handler) ResolveURL(w http.ResponseWriter, r *http.Request) {
	hash := strings.TrimPrefix(r.URL.Path, "/")
	original, _ := h.service.Resolve(hash)
	http.Redirect(w, r, original, http.StatusFound)
}
