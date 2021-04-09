package web

import (
	"auto-admin/pkg/handler"
	"auto-admin/pkg/mongo"
	"html/template"
	"net/http"
	"os"
)

type dashboardHandler struct {
	r mongo.Repository
}

func NewDashBoardHandler(r mongo.Repository) *dashboardHandler {
	return &dashboardHandler{r: r}
}

func (h *dashboardHandler) HandleDashboard(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case handler.GET:
		h.Dashboard(w)
	default:
		handler.WriteError(w, nil, http.StatusMethodNotAllowed)
	}
}

type Test struct {
	AA    []string
	Title string
}

func (h *dashboardHandler) Dashboard(w http.ResponseWriter) {
	wd, _ := os.Getwd()
	tmpl := template.Must(template.ParseFiles(wd + "/pkg/template/dashboard.html"))
	cc, err := h.r.CollectionNames()
	b := Test{AA: cc, Title: "collections"}
	if err != nil {
		handler.WriteError(w, err, http.StatusInternalServerError)
	}
	_ = tmpl.Execute(w, b)
}
