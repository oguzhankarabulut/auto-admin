package web

import (
	"auto-admin/pkg/handler"
	"auto-admin/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

const (
	collections = "collections"
	dashboard   = "dashboard"
	table       = "table"
	detail      = "detail"
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
		switch r.URL.Path {
		case "/dashboard":
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

type dashboardResponse struct {
	Collections []string
	Title       string
}

type tableResponse struct {
	CollectionName string
	Documents      []bson.M
}

type detailResponse struct {
	CollectionName string
	Detail         interface{}
}

func (h *dashboardHandler) Dashboard(w http.ResponseWriter) {
	tmpl := template(dashboard)
	cc, err := h.r.CollectionNames()
	dr := dashboardResponse{Collections: cc, Title: collections}
	if err != nil {
		handler.WriteError(w, err, http.StatusInternalServerError)
	}
	_ = tmpl.Execute(w, dr)
}

func (h *dashboardHandler) Table(w http.ResponseWriter, r *http.Request) {
	coll := handler.CollectionWeb(r)
	template := template(table)
	cc, _ := h.r.All(coll)
	tr := tableResponse{CollectionName: coll, Documents: cc}
	_ = template.Execute(w, tr)
}

func (h *dashboardHandler) Detail(w http.ResponseWriter, r *http.Request) {
	coll := handler.CollectionWeb(r)
	template := template(detail)
	i, _ := h.r.Single(coll, handler.Id(r))
	der := detailResponse{CollectionName: coll, Detail: i}
	_ = template.Execute(w, der)

}
