package repository

import (
	db "github.com/luckyComet55/backend-trainee-assignment-2023/database"
	sg "github.com/luckyComet55/backend-trainee-assignment-2023/segment"
	usr "github.com/luckyComet55/backend-trainee-assignment-2023/user"
	ug "github.com/luckyComet55/backend-trainee-assignment-2023/usersegment"
)

type ServiceMockRepository struct {
	UserDb        usr.UserDatabase
	SegmentDb     sg.SegmentDatabase
	UserSegmentDb ug.UserSegmentDatabase
}

func NewServiceMockRepository(userDb usr.UserDatabase, segmentDb sg.SegmentDatabase, usgDb ug.UserSegmentDatabase) *ServiceMockRepository {
	return &ServiceMockRepository{
		UserDb:        userDb,
		SegmentDb:     segmentDb,
		UserSegmentDb: usgDb,
	}
}

func (r *ServiceMockRepository) GetSegmentsByUserId(id int) ([]sg.Segment, error) {
	usgs := r.UserSegmentDb.GetByUserId(id)
	res := make([]sg.Segment, 0)
	for _, userSegment := range usgs {
		if v, err := r.SegmentDb.GetObjectById(userSegment.GetSegmentId()); err != nil {
			return nil, err
		} else {
			res = append(res, v)
		}
	}
	return res, nil
}

func (r *ServiceMockRepository) GetUsersBySegmentId(id int) ([]usr.User, error) {
	usgs := r.UserSegmentDb.GetBySegmentId(id)
	res := make([]usr.User, 0)
	for _, userSegment := range usgs {
		if v, err := r.UserDb.GetObjectById(userSegment.GetUserId()); err != nil {
			return nil, err
		} else {
			res = append(res, v)
		}
	}
	return res, nil
}

func (r *ServiceMockRepository) CheckNonExistantSegments(segmentNames []string) ([]string, []int) {
	existing := make([]int, 0, len(segmentNames))
	nonExisting := make([]string, 0, len(segmentNames))
	for _, v := range segmentNames {
		s, err := r.SegmentDb.GetByName(v)
		if err == nil {
			existing = append(existing, s.GetId())
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
