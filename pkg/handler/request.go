package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	keyId = "id"
)

func CollectionApi(r *http.Request) string {
	return r.URL.Path[len("/api/"):]
}

func CollectionWeb(r *http.Request) string {
	return r.URL.Path[1:]
}

func Id(r *http.Request) string {
	return r.URL.Query().Get(keyId)
}

func BodyToMap(r *http.Request) (map[string]interface{}, error) {
	b, _ := ioutil.ReadAll(r.Body)
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	return m, err
}
