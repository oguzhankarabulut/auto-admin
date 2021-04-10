package main

import (
	"auto-admin/pkg/handler/api"
	"auto-admin/pkg/handler/web"
	"auto-admin/pkg/mongo"
	"net/http"
)

func main() {

	_ = server()
}

func server() error {
	var repository mongo.Repository

	mc := mongo.NewClient("mongodb://root:example@localhost:27017")
	mc.SetDB("auto-admin")
	cc, _ := mc.CollectionNames()

	repository = mongo.NewRepository(mc)
	ah := api.NewHandler(repository)

	wh := web.NewDashBoardHandler(repository)
	http.HandleFunc("/dashboard", wh.HandleDashboard)

	for _, c := range cc {
		http.HandleFunc("/api/"+c, ah.HandleApi)
		http.HandleFunc("/"+c, wh.HandleDashboard)
	}

	return http.ListenAndServe("localhost:8000", nil)
}
