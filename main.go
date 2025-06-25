package main

import (
	"fmt"
	"net/http"
	"time"
)

// Panic akan dipicu dari goroutine utama agar stack trace terlihat di stdout
func panicSoon() {
	time.Sleep(1 * time.Second)
	panic("🔥 intentional panic: this should crash the container and show in logs")
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Panic will happen in 1 second... check logs")

	// Panic terjadi dari goroutine utama agar bisa crash dan terlihat
	go panicSoon()
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/panic", panicHandler)
	http.HandleFunc("/healthz", healthHandler)

	fmt.Println("🌐 Server running on :8080")

	// Jalankan HTTP server dalam goroutine terpisah
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(fmt.Sprintf("server error: %v", err))
		}
	}()

	// Blok agar main() tetap hidup sampai goroutine panic jalan
	select {}
}
