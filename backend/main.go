package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Version string `json:"version"`
}

func main() {
	// http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}

// func handler(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "Hello, World!")
// }

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	jsonResponse := HealthResponse{Status: "healthy", Service: "technical-spec-api", Version: "0.1.0"}
	json.NewEncoder(w).Encode(jsonResponse)
}
