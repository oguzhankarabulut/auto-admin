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
		switch id(r) {
		case "":
			h.All(w, r)
		default:
			h.Single(w, r)
		}
	case PUT:
		h.Update(w, r)
	case DELETE:
		h.Delete(w, r)
	default:
		WriteError(w, nil, http.StatusMethodNotAllowed)
	}
}

func (h *handler) All(w http.ResponseWriter, r *http.Request) {
	cc, _ := h.r.All(collection(r))
	write(w, cc)
}

func (h *handler) Single(w http.ResponseWriter, r *http.Request) {
	i, _ := h.r.Single(collection(r), id(r))
	write(w, i)
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	m, err := bodyToMap(r)
	if err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}
	i, err := h.r.Create(collection(r), m)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	write(w, i)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	m, err := bodyToMap(r)
	if err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}
	i, err := h.r.Update(collection(r), m)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	write(w, i)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	err := h.r.Delete(collection(r), id(r))
	if err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	write(w, nil)
}
