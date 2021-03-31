package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	errMethodNotAllowed    = "method not allowed"
	errInternalServerError = "internal server error"
	errBadRequest          = "bad request"
	ok                     = "OK"
	errUnauthorized        = "Unauthorized"
)

type Response struct {
	Content      interface{} `json:"content"`
	Message      string      `json:"message"`
	ResponseCode int         `json:"responseCode"`
}

func writeResponse(w http.ResponseWriter, v interface{}, status int) {
	o, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		http.Error(w, errInternalServerError, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	_, _ = w.Write(o)

}

func write(w http.ResponseWriter, v interface{}) {
	r := Response{
		Content:      v,
		Message:      ok,
		ResponseCode: http.StatusOK,
	}
	writeResponse(w, r, http.StatusOK)
}

func WriteError(w http.ResponseWriter, err error, httpStatus int) {
	log.Println(err)
	var m string
	switch httpStatus {
	case http.StatusBadRequest:
		m = errBadRequest
	case http.StatusMethodNotAllowed:
		m = errMethodNotAllowed
	case http.StatusUnauthorized:
		m = errUnauthorized
	default:
		m = errInternalServerError
	}
	r := Response{
		Content:      nil,
		Message:      m,
		ResponseCode: httpStatus,
	}
	writeResponse(w, r, httpStatus)
}

