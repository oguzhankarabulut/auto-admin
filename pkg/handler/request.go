package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func collection(r *http.Request) string {
	return r.URL.Path[1:]
}

func bodyToMap(r *http.Request) (map[string]interface{}, error) {
	b, _ := ioutil.ReadAll(r.Body)
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	return m, err
}
