package main

import (
	"fmt"
	"io"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

func main() {
	http.HandleFunc("/health-check", HealthCheckHandler)
	fmt.Println("test habib")
	http.ListenAndServe(":8080", nil)
}
