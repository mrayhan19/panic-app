package main

import (
	"fmt"
	"net/http"
	"os"
)

// handler yang akan menyebabkan panic fatal dan menghentikan container
func panicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Triggering panic...")

	// Force exit secara eksplisit setelah panic
	go func() {
		panic("intentional crash from /panic")
	}()

	// Beri waktu agar panic di goroutine sempat terjadi
	// atau langsung exit secara paksa agar container mati
	go func() {
		// Tunggu 100ms lalu exit paksa (fallback jika panic tidak mematikan proses utama)
		// Karena panic dalam goroutine tidak mematikan main goroutine
		select {}
	}()

	// Langsung keluar agar container berhenti
	os.Exit(1)
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
