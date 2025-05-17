package server

import (
	"log"
	"net/http"
)

func ServeFiles(folder string) {
	fs := http.FileServer(http.Dir(folder))
	http.Handle("/", fs)
	log.Println("Serving at :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
