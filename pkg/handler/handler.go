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

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case POST:
		h.Create(w, r)
	case GET:
		h.All(w, r)
	default:
		WriteError(w, nil, http.StatusMethodNotAllowed)
	}
}

func (h *handler) All(w http.ResponseWriter, r *http.Request) {
	cc, _ := h.r.All(collection(r))
	write(w, cc)
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	m, err := bodyToMap(r)
	if err != nil {

	}
	_, err = h.r.Create(collection(r), m)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	write(w, r.Body)
}
