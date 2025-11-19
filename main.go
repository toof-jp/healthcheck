package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	message := os.Getenv("MESSAGE")
	secret := os.Getenv("SECRET")

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("/healthz accessed from %s, MESSAGE=%s SECRET=%s", r.RemoteAddr, message, secret)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s\n%s", message, secret)
	})

	log.Println("listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
