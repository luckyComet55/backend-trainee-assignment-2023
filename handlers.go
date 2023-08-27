package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	db "github.com/luckyComet55/backend-trainee-assignment-2023/database"
	sg "github.com/luckyComet55/backend-trainee-assignment-2023/segment"
	usg "github.com/luckyComet55/backend-trainee-assignment-2023/usersegment"
)

type userSegmentsModifyBody struct {
	UserId           int      `json:"user_id"`
	SegmentsToAdd    []string `json:"segments_add"`
	SegmentsToRemove []string `json:"segments_remove"`
}

type userSegmentsResponseBody struct {
	Segments []string `json:"segments"`
}

type userModifyErrorResponse struct {
	Message          string   `json:"message"`
	SegmentsToRemove []string `json:"segments_remove,omitempty"`
	SegmentsToAdd    []string `json:"segments_add,omitempty"`
}

func (u userSegmentsModifyBody) String() string {
	return fmt.Sprintf("\n========\nUSER %dto add: %v\nto remove%v\n========\n", u.UserId, u.SegmentsToAdd, u.SegmentsToRemove)
}

func helloRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func createSegmentHandler(w http.ResponseWriter, r *http.Request) {
	segmentName := chi.URLParam(r, "segmentName")
	res := "OK"
	logStatus := "SUCCESS"
	statusCode := 200
	segment := sg.NewSegment(segmentName)
	fmt.Println(segment)
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
	res := "OK"
	logStatus := "SUCCESS"
	statusCode := 200
	if err := serviceRepo.SegmentDb.DeleteByName(segmentName); err != nil {
		res = "internal error"
		statusCode = 500
		logStatus = "DENIED"
	}
	fmt.Printf("%s %s ==> delete segment %s | %s\n", r.Method, r.URL.Path, segmentName, logStatus)
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, res)
}

func modifyUserSegments(w http.ResponseWriter, r *http.Request) {
	var reqBody userSegmentsModifyBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		// add case of incorrect body format
		writeResponse(w, []byte("incorrect body format"), 400)
		return
	}
	toAdd := xorStringArrays(reqBody.SegmentsToAdd, reqBody.SegmentsToRemove)
	toRm := xorStringArrays(reqBody.SegmentsToRemove, reqBody.SegmentsToAdd)
	userId := reqBody.UserId
	if _, err := serviceRepo.UserDb.GetObjectById(userId); err != nil {
		writeResponse(w, []byte("user with provided id not found"), 400)
		return
	}

	// it must be like some kind of a transaction
	// so if one value is incorrect, the others will be ignored
	unableToRm, removable := serviceRepo.CheckNonExistantSegments(toRm)
	unableToAdd, addable := serviceRepo.CheckNonExistantSegments(toAdd)
	var errorResponse userModifyErrorResponse = userModifyErrorResponse{}
	if !(len(unableToAdd) == 0 && len(unableToRm) == 0) {
		errorResponse.Message = "objects with these values were not found"
		if len(unableToAdd) > 0 {
			errorResponse.SegmentsToAdd = unableToAdd
		}
		if len(unableToRm) > 0 {
			errorResponse.SegmentsToRemove = unableToRm
		}
		resp, _ := json.Marshal(errorResponse)
		writeResponse(w, resp, 400)
		return
	}
	for _, v := range removable {
		err := serviceRepo.UserSegmentDb.DeleteByUserIdWithSegmentId(userId, v)
		if err != nil {
			writeResponse(w, []byte("internal error"), 500)
			return
		}
	}
	for _, v := range addable {
		userSegment := usg.NewUserSegment(userId, v)
		err := serviceRepo.UserSegmentDb.CreateObject(userSegment)
		if err != nil {
			var resp []byte
			var statusCode int
			switch err.(type) {
			case db.ErrUniqueConstraintFailed:
				resp = []byte("user already has such segments: " + err.Error())
				statusCode = 400
			default:
				resp = []byte("internal error")
				statusCode = 500
			}
			writeResponse(w, resp, statusCode)
			return
		}
	}
	writeResponse(w, []byte("OK"), 200)
}

func getUserSegments(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userId")
	var res []byte
	statusCode := 200
	userSegments := userSegmentsResponseBody{Segments: make([]string, 0)}
	logStatus := "SUCCESS"

	// we ignore error, returned by atoi
	// because our router checks if the
	// value contains digits only
	userId, _ := strconv.Atoi(userIdStr)
	arr, err := serviceRepo.GetSegmentsByUserId(userId)
	if err != nil {
		res = []byte("user not found")
		statusCode = 404
		logStatus = "DENIED --> No such user"
	} else {
		for _, v := range arr {
			userSegments.Segments = append(userSegments.Segments, v.Name)
		}
		res, err = json.Marshal(userSegments)
		if err != nil {
			res = []byte("Internal error")
			statusCode = 500
			logStatus = "DENIED --> Marshalling error"
		}
	}
	fmt.Printf("%s %s ==> get user %d segments %v | %s\n", r.Method, r.URL.Path, userId, userSegments.Segments, logStatus)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)
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

func writeResponse(w http.ResponseWriter, msg []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(msg)
}
