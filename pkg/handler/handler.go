package handler

import (
	"auto-admin/pkg/mongo"
	"net/http"
)

type handler struct {
	r mongo.Repository
}

func NewHandler(r mongo.Repository) *handler {
	return &handler{r: r}
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	cc, _ := h.r.GetAll(p[1:])
	write(w, cc)
}
