package main

import (
	"log"
	"net/http"

	"github.com/bgmerrell/gomsg/handlers/web"
)

func main() {
	http.Handle("/message", &web.Handler{})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
