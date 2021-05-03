package main

import (
	"encoding/json"
	"log"
	"os"
)

type options struct {
	MongoDbConnectionString string `json:"mongo_db_connection_string"`
	DatabaseName            string `json:"database_name"`
	Path                    string `json:"path"`
}

func newOptions() (*options, error) {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	f, err := os.Open(wd + "/config.json")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer f.Close()
	d := json.NewDecoder(f)
	c := &options{}
	err = d.Decode(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
