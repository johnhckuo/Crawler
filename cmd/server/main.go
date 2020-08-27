package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/contact", contactHandler)

	log.Println(fmt.Sprintf("Server running on http://localhost%s 🐹", ":4000"))
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatalf("could not run the server %v", err)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from home handler"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from contact handler"))
}
