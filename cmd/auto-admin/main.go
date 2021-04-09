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
	gah := api.NewHandler(repository)

	dh := web.NewDashBoardHandler(repository)
	http.HandleFunc("/dashboard", dh.HandleDashboard)

	for _, c := range cc {
		http.HandleFunc("/api/"+c, gah.HandleApi)
	}

	return http.ListenAndServe("localhost:8000", nil)
}
