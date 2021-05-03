package main

import (
	"auto-admin/pkg/handler/api"
	"auto-admin/pkg/handler/web"
	"auto-admin/pkg/mongo"
	"log"
	"net/http"
	"os"
)

const (
	dashboardPath = "/dashboard"
	scripts       = "/scripts/"
	scriptsPath   = "/pkg/template/scripts"
	apiPath       = "/api/"
	mainPath      = "/"
)

func main() {

	if err := server(); err != nil {
		log.Println(err)
		panic(err)
	}
}

func server() error {
	var repository mongo.Repository

	opt, err := newOptions()
	if err != nil {
		return err
	}

	mc := mongo.NewClient(opt.MongoDbConnectionString)
	mc.SetDB(opt.DatabaseName)
	cc, _ := mc.CollectionNames()

	repository = mongo.NewRepository(mc)
	ah := api.NewHandler(repository)

	wh := web.NewDashBoardHandler(repository)
	http.HandleFunc(dashboardPath, wh.HandleDashboard)
	wd, _ := os.Getwd()
	http.Handle(scripts, http.StripPrefix("/scripts/", http.FileServer(http.Dir(wd+scriptsPath))))

	for _, c := range cc {
		http.HandleFunc(apiPath+c, ah.HandleApi)
		http.HandleFunc(mainPath+c, wh.HandleDashboard)
	}

	return http.ListenAndServe(opt.Path, nil)
}
