package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	sg "github.com/luckyComet55/backend-trainee-assignment-2023/segment"
	usg "github.com/luckyComet55/backend-trainee-assignment-2023/usersegment"
)

type userSegmentsModifyBody struct {
	UserId     int
	SgToAdd    []string
	SgToRemove []string
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

func xorStringArrays(a, b []string) []string {
	checker := make(map[string]bool, len(a)+len(b))
	res := make([]string, 0, len(a))
	for _, s := range b {
		checker[s] = true
	}
	for _, s := range a {
		if _, ok := checker[s]; !ok {
			res = append(res, s)
			checker[s] = true
		}
	}
	return res
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
	} else {
		toAdd := xorStringArrays(reqBody.SgToAdd, reqBody.SgToRemove)
		toRm := xorStringArrays(reqBody.SgToRemove, reqBody.SgToAdd)
		userId := reqBody.UserId

		// it must be like some kind of a transaction
		// so if one value is incorrect, the others will be ignored
		for _, v := range toRm {
			r, err := serviceRepo.SegmentDb.GetByName(v)
			if err != nil {
				res = err.Error()
				logStatus = "DENIED"
				statusCode = 400
				break
			}
			err = serviceRepo.UserSegmentDb.DeleteByUserIdWithSegmentId(userId, r.GetId())
			if err != nil {
				res = err.Error()
				logStatus = "DENIED"
				statusCode = 500
				break
			}
		}
		for _, v := range toAdd {
			r, err := serviceRepo.SegmentDb.GetByName(v)
			if err != nil {
				res = err.Error()
				logStatus = "DENIED"
				statusCode = 400
				break
			}
			err = serviceRepo.UserSegmentDb.CreateObject(usg.NewUserSegment(userId, r.GetId()))
			if err != nil {
				res = err.Error()
				logStatus = "DENIED"
				statusCode = 400
				break
			}
		}
	}
	fmt.Printf("%s %s ==> modify user segment | %s%v", r.Method, r.URL.Path, logStatus, reqBody)
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, res)
}

func getUserSegments(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userId")
	res := make([]byte, 0)
	statusCode := 200
	userSegments := userSegmentsResponseBody{}
	logStatus := "SUCCESS"
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		statusCode = 400
		logStatus = "DENIED"
	} else {
		userSegments.Segments, err = serviceRepo.GetSegmentsByUserId(userId)
		if err != nil {
			res = []byte(err.Error())
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
	}
	fmt.Printf("%s %s ==> modify user segment %v | %s\n", r.Method, r.URL.Path, userSegments.Segments, logStatus)
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, res)
}
