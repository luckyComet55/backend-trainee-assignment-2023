package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	db_test "github.com/luckyComet55/backend-trainee-assignment-2023/databasetest"
	rp "github.com/luckyComet55/backend-trainee-assignment-2023/repository"
	sg "github.com/luckyComet55/backend-trainee-assignment-2023/segment"
)

var db db_test.Database[sg.Segment] = db_test.NewDatabase[sg.Segment]()
var repo rp.Repository[sg.Segment] = rp.NewRepository[sg.Segment](db)

func main() {
	port := flag.String("port", "3003", "port the server will listen to")
	flag.Parse()

	r := chi.NewRouter()
	r.Get("/", helloRootHandler)
	r.Post("/{segmentName}", createSegmentHandler)
	r.Delete("/{segmtnName}", deleteSegmentHandler)

	log.Fatal(http.ListenAndServe(":"+*port, r))
}
