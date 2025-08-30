package main

import (
	"log"
	"net/http"
)

func main() {
	message := os.Getenv("MESSAGE")

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("/healthz accessed from %s, MESSAGE=%s", r.RemoteAddr, message)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", message)
	})

	log.Println("listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
