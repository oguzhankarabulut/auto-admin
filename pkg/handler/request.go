package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	keyId = "id"
)

func collection(r *http.Request) string {
	return r.URL.Path[1:]
}

func id(r *http.Request) string {
	return r.URL.Query().Get(keyId)
}

func bodyToMap(r *http.Request) (map[string]interface{}, error) {
	b, _ := ioutil.ReadAll(r.Body)
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	return m, err
}
