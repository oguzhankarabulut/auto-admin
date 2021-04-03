package main

import (
	"auto-admin/pkg/handler"
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
	gah := handler.NewHandler(repository)

	for _, c := range cc {
		http.HandleFunc("/"+c, gah.Handle)
	}

	return http.ListenAndServe("localhost:8000", nil)
}
