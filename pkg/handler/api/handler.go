package api

import (
	"auto-admin/pkg/handler"
	"auto-admin/pkg/service"
	"net/http"
)

type apiHandler struct {
	s service.ApiService
}

func NewApiHandler(s service.ApiService) *apiHandler {
	return &apiHandler{s: s}
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
	cc, err := h.s.All(handler.CollectionApi(r))
	if err != nil {
		handler.WriteError(w, err, http.StatusInternalServerError)
	}

	handler.Write(w, cc)
}

func (h *apiHandler) Single(w http.ResponseWriter, r *http.Request) {
	i, err := h.s.Single(handler.CollectionApi(r), handler.Id(r))
	if err != nil {
		handler.WriteError(w, err, http.StatusInternalServerError)
	}

	if i == nil {
		handler.WriteError(w, err, http.StatusBadRequest)
	}

	handler.Write(w, i)
}

func (h *apiHandler) Create(w http.ResponseWriter, r *http.Request) {
	m, err := handler.BodyToMap(r)
	if err != nil {
		handler.WriteError(w, err, http.StatusBadRequest)
		return
	}

	i, err := h.s.Create(handler.CollectionApi(r), m)
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
	i, err := h.s.Update(handler.CollectionApi(r), m)
	if err != nil {
		handler.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	handler.Write(w, i)
}

func (h *apiHandler) Delete(w http.ResponseWriter, r *http.Request) {
	err := h.s.Delete(handler.CollectionApi(r), handler.Id(r))
	if err != nil {
		handler.WriteError(w, err, http.StatusBadRequest)
		return
	}

	handler.Write(w, nil)
}
