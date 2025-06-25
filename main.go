package main

import (
	"fmt"
	"net/http"
)

type User struct {
	Name string
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	var u *User                        // nil pointer
	fmt.Fprintf(w, "User: %s", u.Name) // ini akan panic: dereference nil pointer
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/panic", panicHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
