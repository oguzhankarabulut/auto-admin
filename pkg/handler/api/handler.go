package api

import (
	"auto-admin/pkg/handler"
	"auto-admin/pkg/mongo"
	"net/http"
)

type apiHandler struct {
	r mongo.Repository
}

func NewHandler(r mongo.Repository) *apiHandler {
	return &apiHandler{r: r}
}

func (h *apiHandler) HandleApi(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case handler.POST:
		h.Create(w, r)
	case handler.GET:
		switch handler.Id(r) {
		case "":
			h.All(w, r)
		default:
			h.Single(w, r)
		}
	case handler.PUT:
		h.Update(w, r)
	case handler.DELETE:
		h.Delete(w, r)
	default:
		handler.WriteError(w, nil, http.StatusMethodNotAllowed)
	}
}

func (h *apiHandler) All(w http.ResponseWriter, r *http.Request) {
	cc, _ := h.r.All(handler.CollectionApi(r))
	handler.Write(w, cc)
}

func (h *apiHandler) Single(w http.ResponseWriter, r *http.Request) {
	i, _ := h.r.Single(handler.CollectionApi(r), handler.Id(r))
	handler.Write(w, i)
}

func (h *apiHandler) Create(w http.ResponseWriter, r *http.Request) {
	m, err := handler.BodyToMap(r)
	if err != nil {
		handler.WriteError(w, err, http.StatusBadRequest)
		return
	}
	i, err := h.r.Create(handler.CollectionApi(r), m)
	if err != nil {
		handler.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	handler.Write(w, i)
}

func (h *apiHandler) Update(w http.ResponseWriter, r *http.Request) {
	m, err := handler.BodyToMap(r)
	if err != nil {
		handler.WriteError(w, err, http.StatusBadRequest)
		return
	}
	i, err := h.r.Update(handler.CollectionApi(r), m)
	if err != nil {
		handler.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	handler.Write(w, i)
}

func (h *apiHandler) Delete(w http.ResponseWriter, r *http.Request) {
	err := h.r.Delete(handler.CollectionApi(r), handler.Id(r))
	if err != nil {
		handler.WriteError(w, err, http.StatusBadRequest)
		return
	}

	handler.Write(w, nil)
}
