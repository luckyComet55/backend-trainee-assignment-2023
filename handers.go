package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	sg "github.com/luckyComet55/backend-trainee-assignment-2023/segment"
)

func helloRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func createSegmentHandler(w http.ResponseWriter, r *http.Request) {
	segmentName := chi.URLParam(r, "segmentName")
	fmt.Printf("%s %s ==> create segment %s\n", r.Method, r.URL.Path, segmentName)

	res := ""
	statusCode := 200
	segment := sg.NewSegment(1, segmentName)
	if err := repo.Db.CreateObject(segment); err != nil {
		res = err.Error()
		statusCode = 400
	} else {
		res = "success!"
	}
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, res)
}

func deleteSegmentHandler(w http.ResponseWriter, r *http.Request) {
	segmentName := chi.URLParam(r, "segmentName")
	fmt.Printf("%s %s ==> delete segment %s\n", r.Method, r.URL.Path, segmentName)
}
