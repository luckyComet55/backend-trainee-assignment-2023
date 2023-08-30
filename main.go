package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	repo "github.com/luckyComet55/backend-trainee-assignment-2023/repository"
	sg "github.com/luckyComet55/backend-trainee-assignment-2023/segment"
	u "github.com/luckyComet55/backend-trainee-assignment-2023/user"
	ug "github.com/luckyComet55/backend-trainee-assignment-2023/usersegment"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
)

func initDatabaseConnection(ctx context.Context) (ksql.DB, error) {
	psPort := os.Getenv("POSTGRES_PORT")
	psDb := os.Getenv("POSTGRES_DB")
	psUser := os.Getenv("POSTGRES_USER")
	psPassword := os.Getenv("POSTGRES_PASSWORD")
	connString := fmt.Sprintf("postgresql://%s:%s@localhost:%s/%s", psUser, psPassword, psPort, psDb)
	fmt.Printf("%s\n", connString)
	conn, err := kpgx.New(ctx, connString, ksql.Config{})
	if err != nil {
		fmt.Printf("Could not connect to postgres: %v\n", err)
		return conn, err
	}
	return conn, nil
}

func initRepo(conn ksql.DB) {
	var dbSegment *sg.SegmentActualDatabase = sg.NewSegmentActualDatabase(conn)
	var dbUserSegment *ug.UserSegmentActualDatabase = ug.NewUserSegmentActualDatabase(conn)
	var dbUser *u.UserActualDatabase = u.NewUserActualDatabase(conn)
	serviceRepo = *repo.NewServiceMockRepository(dbSegment, dbUserSegment, dbUser)
}

var serviceRepo repo.ServiceMockRepository

func main() {
	ctx := context.Background()
	godotenv.Load()
	conn, err := initDatabaseConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if err = os.Mkdir("data", os.ModeDir); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	defer conn.Close()
	initRepo(conn)
	port := flag.String("port", "3003", "port the server will listen to")
	flag.Parse()
	r := chi.NewRouter()
	r.Get("/", helloRootHandler)
	r.Post("/{segmentName}", createSegmentHandler)
	r.Delete("/{segmentName}", deleteSegmentHandler)
	r.Put("/modify-user-segments", modifyUserSegments)
	r.Get("/{userId:[0-9]+}", getUserSegments)
	r.Get("/{userId:[0-9]+}/{year:[0-9]+}/{month:[0-9]+}", getUserSegmentsInPeriod)
	r.Get("/user-report/{filename}", downloadUserReport)

	log.Fatal(http.ListenAndServe(":"+*port, r))
}
