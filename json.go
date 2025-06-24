package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	HeaderContentType     = "Content-Type"
	HeaderApplicationJson = "application/json"
)

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
		Code  int    `json:"code"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
		Code:  code,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set(HeaderContentType, HeaderApplicationJson)
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	_, _ = w.Write(dat)
}
