package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := flag.String("port", "3003", "port the server will listen to")
	flag.Parse()

	r := chi.NewRouter()
	r.Get("/", helloRootHandler)
	r.Post("/{segmentName}", createSegmentHandler)
	r.Delete("/{segmtnName}", deleteSegmentHandler)

	log.Fatal(http.ListenAndServe(":"+*port, r))
}
