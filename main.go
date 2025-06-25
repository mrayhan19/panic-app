package main

import (
	"fmt"
	"net/http"
)

type User struct {
	Name string
}Add commentMore actions

func panicHandler(w http.ResponseWriter, r *http.Request) {
	var u *User
	fmt.Fprintf(w, "User: %s", u.Name) // Ini akan panic karena u == nil
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