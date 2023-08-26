package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	sg "github.com/luckyComet55/backend-trainee-assignment-2023/segment"
)

type userSegmentsModifyBody struct {
	UserId     int
	SgToAdd    []sg.Segment
	SgToRemove []sg.Segment
}

type userSegmentsResponseBody struct {
	Segments []sg.Segment `json:"segments"`
}

func (u userSegmentsModifyBody) String() string {
	return fmt.Sprintf("\n========\nUSER %dto add: %v\nto remove%v\n========\n", u.UserId, u.SgToAdd, u.SgToRemove)
}

func helloRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func createSegmentHandler(w http.ResponseWriter, r *http.Request) {
	segmentName := chi.URLParam(r, "segmentName")
	res := "success!"
	logStatus := "SUCCESS"
	statusCode := 200
	segment := sg.NewSegment(segmentName)
	if err := serviceRepo.SegmentDb.CreateObject(segment); err != nil {
		res = err.Error()
		statusCode = 400
		logStatus = "DENIED"
	}
	fmt.Printf("%s %s ==> create segment %s | %s\n", r.Method, r.URL.Path, segmentName, logStatus)
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, res)
}

func deleteSegmentHandler(w http.ResponseWriter, r *http.Request) {
	segmentName := chi.URLParam(r, "segmentName")
	res := "success!"
	logStatus := "SUCCESS"
	statusCode := 200
	if segment, err := serviceRepo.SegmentDb.GetByName(segmentName); err != nil {
		res = err.Error()
		statusCode = 400
		logStatus = "DENIED"
	} else {
		if err = serviceRepo.SegmentDb.DeleteObject(segment); err != nil {
			res = err.Error()
			statusCode = 400
			logStatus = "DENIED"
		}
	}
	fmt.Printf("%s %s ==> delete segment %s | %s\n", r.Method, r.URL.Path, segmentName, logStatus)
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, res)
}

func modifyUserSegments(w http.ResponseWriter, r *http.Request) {
	var reqBody userSegmentsModifyBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	res := "success!"
	logStatus := "SUCCESS"
	statusCode := 200
	if err := decoder.Decode(&reqBody); err != nil {
		// add case of incorrect body format
		res = err.Error()
		logStatus = "DENIED"
		statusCode = 400
	}
	fmt.Printf("%s %s ==> modify user segment | %s%v", r.Method, r.URL.Path, logStatus, reqBody)
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, res)
}

func getUserSegments(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userId")
	res := make([]byte, 0)
	userSegments := userSegmentsResponseBody{Segments: make([]sg.Segment, 0)}
	statusCode := 200
	logStatus := "SUCCESS"
	_, err := strconv.Atoi(userIdStr)
	if err != nil {
		statusCode = 400
		logStatus = "DENIED"
	} else {
		res, err = json.Marshal(userSegments)
		if err != nil {
			res = []byte(err.Error())
			statusCode = 400
			logStatus = "DENIED"
		}
	}
	fmt.Printf("%s %s ==> modify user segment %v | %s\n", r.Method, r.URL.Path, userSegments.Segments, logStatus)
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, res)
}
