package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	sg "github.com/luckyComet55/backend-trainee-assignment-2023/segment"
	usr "github.com/luckyComet55/backend-trainee-assignment-2023/user"
	ug "github.com/luckyComet55/backend-trainee-assignment-2023/usersegment"
)

var dbSegment *sg.SegmentMockDatabase = sg.NewSegmentMockDatabase()
var dbUser *usr.UserMockDatabase = usr.NewUserMockDatabase()
var dbUserSegment *ug.UserSegmentMockDatabase = ug.NewUserSegmentMockDatabase()
var repoSegment sg.SegmentRepository = sg.NewSegmentRepository(dbSegment)
var repoUser usr.UserRepository = usr.NewUserRepository(dbUser)
var repoUserSegment ug.UserSegmentRepository = ug.NewUserSegmentRepository(dbUserSegment)

func main() {
	port := flag.String("port", "3003", "port the server will listen to")
	flag.Parse()

	r := chi.NewRouter()
	r.Get("/", helloRootHandler)
	r.Post("/{segmentName}", createSegmentHandler)
	r.Delete("/{segmentName}", deleteSegmentHandler)
	r.Put("/modify-user-segments", modifyUserSegments)
	r.Get("/{userId:[0-9]+}", getUserSegments)

	log.Fatal(http.ListenAndServe(":"+*port, r))
}
