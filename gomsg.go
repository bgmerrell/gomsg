package main

import (
	"log"
	"net/http"

	"github.com/bgmerrell/gomsg/handlers/web"
	"github.com/bgmerrell/gomsg/messages"
)

func main() {
	msgMap := messages.NewMessageMap()
	http.Handle("/message", &web.Handler{msgMap, 0})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
