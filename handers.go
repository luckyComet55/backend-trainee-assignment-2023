package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func helloRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func createSegmentHandler(w http.ResponseWriter, r *http.Request) {
	segmentName := chi.URLParam(r, "segmentName")
	fmt.Printf("%s %s ==> create segment %s\n", r.Method, r.URL.Path, segmentName)
}

func deleteSegmentHandler(w http.ResponseWriter, r *http.Request) {
	segmentName := chi.URLParam(r, "segmentName")
	fmt.Printf("%s %s ==> delete segment %s\n", r.Method, r.URL.Path, segmentName)
}
