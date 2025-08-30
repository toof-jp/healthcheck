package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	env := os.Getenv("ENV")

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("/healthz accessed from %s, ENV=%s", r.RemoteAddr, env)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", env)
	})

	log.Println("listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
