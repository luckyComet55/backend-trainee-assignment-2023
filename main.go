package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	rp "github.com/luckyComet55/backend-trainee-assignment-2023/repository"
	sg "github.com/luckyComet55/backend-trainee-assignment-2023/segment"
	usr "github.com/luckyComet55/backend-trainee-assignment-2023/user"
)

var dbSegment *sg.SegmentMockDatabase = sg.NewSegmentMockDatabase()
var dbUser *usr.UserMockDatabase = usr.NewUserMockDatabase()
var repoSegment rp.Repository[sg.Segment] = rp.NewRepository[sg.Segment](dbSegment)
var repoUser rp.Repository[usr.User] = rp.NewRepository[usr.User](dbUser)

func main() {
	port := flag.String("port", "3003", "port the server will listen to")
	flag.Parse()

	r := chi.NewRouter()
	r.Get("/", helloRootHandler)
	r.Post("/{segmentName}", createSegmentHandler)
	r.Delete("/{segmentName}", deleteSegmentHandler)
	r.Put("/modify-user-segments", modifyUserSegments)

	log.Fatal(http.ListenAndServe(":"+*port, r))
}
