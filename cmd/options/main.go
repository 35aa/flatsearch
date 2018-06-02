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
	port    = "3000"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", AppBanner).Methods("GET")

	fmt.Printf("Start %s on port %s", appName, port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// AppBanner displays application basic information like name, version etc.
func AppBanner(w http.ResponseWriter, r *http.Request) {
	respRaw := make(map[string]string)
	respRaw["app"] = "options"
	resp, _ := json.Marshal(respRaw)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
