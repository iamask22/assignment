package api_handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func (sr SuccessResponse) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(sr.Status)
	err := json.NewEncoder(w).Encode(sr)
	if err != nil {
		log.Printf("Error while encoding response: %v", err)
	}
}

func (er ErrorResponse) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(er.Status)
	err := json.NewEncoder(w).Encode(er)
	if err != nil {
		log.Printf("Error while encoding response: %v", err)
	}
}
