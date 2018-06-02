package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	appName = "options"
	build   = "n/a"
	port    = "3001"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", AppBannerHandler).Methods("GET")
	r.HandleFunc("/search/{id:[0-9]+}", GetSearchHandler).Methods("GET")
	r.HandleFunc("/search", SetSearchHandler).Methods("POST")
	r.HandleFunc("/search/{id:[0-9]+}", DeleteSearchHandler).Methods("DELETE")

	fmt.Printf("Start %s (%s) on port %s", appName, build, port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// AppBannerHandler displays application basic information like name, build etc.
func AppBannerHandler(w http.ResponseWriter, _ *http.Request) {
	respRaw := make(map[string]string)
	respRaw["app"] = "options"
	respRaw["build"] = build
	resp, _ := json.Marshal(respRaw)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// GetSearchHandler returns saved search request by ID
func GetSearchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// stub response
	// will be replaced by proper structure "pretty soon" (c)
	respRaw := make(map[string]string)
	respRaw = map[string]string{
		"id":             vars["id"],
		"reality_type":   "appartment",
		"operation_type": "rent",
		"country":        "cz",
		"region":         "central",
		"district":       "prague 1",
		"area":           "32-45m^2",
		"price":          "0-10000CZK",
	}
	resp, _ := json.Marshal(respRaw)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// SetSearchHandler create/update search request
func SetSearchHandler(w http.ResponseWriter, _ *http.Request) {
	// search request has been created/updated
	w.WriteHeader(http.StatusCreated)
}

// DeleteSearchHandler delete search request by ID
func DeleteSearchHandler(w http.ResponseWriter, _ *http.Request) {
	// search reqeust has been successfully deleted
	w.WriteHeader(http.StatusNoContent)
}
