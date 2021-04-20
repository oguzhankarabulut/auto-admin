package main

import (
	"auto-admin/pkg/handler/api"
	"auto-admin/pkg/handler/web"
	"auto-admin/pkg/mongo"
	"net/http"
	"os"
)

func main() {

	_ = server()
}

func server() error {
	var repository mongo.Repository

	mc := mongo.NewClient("mongodb://root:example@localhost:27017")
	mc.SetDB("quin-panel")
	cc, _ := mc.CollectionNames()

	repository = mongo.NewRepository(mc)
	ah := api.NewHandler(repository)

	wh := web.NewDashBoardHandler(repository)
	http.HandleFunc("/dashboard", wh.HandleDashboard)
	wd, _ := os.Getwd()
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir(wd+"/pkg/template/scripts"))))

	for _, c := range cc {
		http.HandleFunc("/api/"+c, ah.HandleApi)
		http.HandleFunc("/"+c, wh.HandleDashboard)
	}

	return http.ListenAndServe("localhost:8000", nil)
}
