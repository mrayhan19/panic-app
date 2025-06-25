package main

import (
	"fmt"
	"net/http"
)

func panicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Triggering panic...")
	panic("🔥 intentional crash")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/panic", panicHandler)
	http.HandleFunc("/healthz", healthHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
