package repository

import (
	db "github.com/luckyComet55/backend-trainee-assignment-2023/database"
	sg "github.com/luckyComet55/backend-trainee-assignment-2023/segment"
	ug "github.com/luckyComet55/backend-trainee-assignment-2023/usersegment"
)

type ServiceMockRepository struct {
	SegmentDb     sg.SegmentDatabase
	UserSegmentDb ug.UserSegmentDatabase
}

func NewServiceMockRepository(segmentDb sg.SegmentDatabase, usgDb ug.UserSegmentDatabase) *ServiceMockRepository {
	return &ServiceMockRepository{
		SegmentDb:     segmentDb,
		UserSegmentDb: usgDb,
	}
}

func (r *ServiceMockRepository) GetSegmentsByUserId(id int) ([]sg.Segment, error) {
	usgs := r.UserSegmentDb.GetByUserId(id)
	res := make([]sg.Segment, 0)
	for _, userSegment := range usgs {
		v, _ := r.SegmentDb.GetByName(userSegment.GetSegmentName())
		res = append(res, v)
	}
	return res, nil
}

func (r *ServiceMockRepository) GetUserIdsBySegmentName(name string) ([]int, error) {
	usgs := r.UserSegmentDb.GetBySegmentName(name)
	res := make([]int, 0)
	for _, userSegment := range usgs {
		res = append(res, userSegment.GetUserId())
	}
	return res, nil
}

func (r *ServiceMockRepository) CheckNonExistantSegments(segmentNames []string) ([]string, []string) {
	existing := make([]string, 0, len(segmentNames))
	nonExisting := make([]string, 0, len(segmentNames))
	for _, v := range segmentNames {
		s, err := r.SegmentDb.GetByName(v)
		if err == nil {
			existing = append(existing, s.GetName())
			continue
		}
		switch err.(type) {
		case db.ErrObjNotFound:
			nonExisting = append(nonExisting, v)
		default:
			return nil, nil
		}
	}
	return nonExisting, existing
}

func (r *ServiceMockRepository) GetUserActiveSegments(user_id int) []string {
	var names []string
	res := r.UserSegmentDb.GetUserActiveSegments(user_id)
	if res == nil {
		return nil
	}
	names = make([]string, 0, len(res))
	for _, v := range res {
		names = append(names, v.GetSegmentName())
	}
	return names
}
