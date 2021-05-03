package web

import (
	"auto-admin/pkg/handler"
	"auto-admin/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

const (
	collections   = "collections"
	dashboard     = "dashboard"
	table         = "table"
	detail        = "detail"
	dashboardPath = "/dashboard"
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

type dashboardResponse struct {
	Collections []string
	Metric      map[string]int64
	Title       string
}

type tableResponse struct {
	CollectionName string
	Collections    []string
	Documents      []bson.M
}

type detailResponse struct {
	CollectionName string
	Collections    []string
	Detail         interface{}
}

func (h *dashboardHandler) Dashboard(w http.ResponseWriter) {
	cn, err := h.r.CollectionNames()
	if err != nil {
		handler.WriteError(w, err, http.StatusInternalServerError)
	}

	m := make(map[string]int64)
	if len(cn) != 0 {
		for i := 0; i < len(cn); i++ {
			c, _ := h.r.Count(cn[i])
			m[cn[i]] = c
		}
	}

	tmpl := template(dashboard)
	dr := dashboardResponse{Collections: cn, Metric: m, Title: collections}
	_ = tmpl.Execute(w, dr)
}

func (h *dashboardHandler) Table(w http.ResponseWriter, r *http.Request) {
	cn, err := h.r.CollectionNames()
	if err != nil {
		handler.WriteError(w, err, http.StatusInternalServerError)
	}

	coll := handler.CollectionWeb(r)
	template := template(table)
	cc, _ := h.r.All(coll)
	tr := tableResponse{CollectionName: coll, Collections: cn, Documents: cc}
	_ = template.Execute(w, tr)
}

func (h *dashboardHandler) Detail(w http.ResponseWriter, r *http.Request) {
	cn, err := h.r.CollectionNames()
	if err != nil {
		handler.WriteError(w, err, http.StatusInternalServerError)
	}

	coll := handler.CollectionWeb(r)
	template := template(detail)
	i, _ := h.r.Single(coll, handler.Id(r))
	der := detailResponse{CollectionName: coll, Collections: cn, Detail: i}
	_ = template.Execute(w, der)

}
