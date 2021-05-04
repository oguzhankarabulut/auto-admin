package web

import (
	"auto-admin/pkg/handler"
	"auto-admin/pkg/service"
	"net/http"
)

const (
	dashboard     = "dashboard"
	table         = "table"
	detail        = "detail"
	dashboardPath = "/dashboard"
)

type dashboardHandler struct {
	s service.WebService
}

func NewDashBoardHandler(s service.WebService) *dashboardHandler {
	return &dashboardHandler{s: s}
}

func (h *dashboardHandler) HandleDashboard(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case handler.GET:
		switch r.URL.Path {
		case dashboardPath:
			h.Dashboard(w)
		default:
			switch handler.Id(r) {
			case "":
				h.Table(w, r)
			default:
				h.Detail(w, r)
			}
		}
	default:
		handler.WriteError(w, nil, http.StatusMethodNotAllowed)
	}
}

func (h *dashboardHandler) Dashboard(w http.ResponseWriter) {
	dr, err := h.s.Dashboard()
	if err != nil {
		handler.WriteError(w, err, http.StatusInternalServerError)
	}

	tmpl := template(dashboard)
	_ = tmpl.Execute(w, dr)
}

func (h *dashboardHandler) Table(w http.ResponseWriter, r *http.Request) {
	tr, err := h.s.Table(handler.CollectionWeb(r))
	if err != nil {
		handler.WriteError(w, err, http.StatusInternalServerError)
	}

	template := template(table)
	_ = template.Execute(w, tr)
}

func (h *dashboardHandler) Detail(w http.ResponseWriter, r *http.Request) {
	dr, err := h.s.Detail(handler.CollectionWeb(r), handler.Id(r))
	if err != nil {
		handler.WriteError(w, err, http.StatusInternalServerError)
	}

	template := template(detail)
	_ = template.Execute(w, dr)
}
