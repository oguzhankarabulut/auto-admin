package main

import (
	"auto-admin/pkg/handler/api"
	"auto-admin/pkg/handler/web"
	"auto-admin/pkg/mongo"
	"auto-admin/pkg/service"
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
	ws := service.NewWebService(repository, cc)
	as := service.NewApiService(repository)
	ah := api.NewApiHandler(as)

	wh := web.NewDashBoardHandler(ws)

	http.HandleFunc(dashboardPath, wh.HandleDashboard)
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	http.Handle(scripts, http.StripPrefix(scripts, http.FileServer(http.Dir(wd+scriptsPath))))

	for _, c := range cc {
		http.HandleFunc(apiPath+c, ah.HandleApi)
		http.HandleFunc(mainPath+c, wh.HandleDashboard)
	}

	return http.ListenAndServe(opt.Path, nil)
}
