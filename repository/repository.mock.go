package repository

import (
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
	user, err := r.UserDb.GetObjectById(id)
	if err != nil {
		return nil, err
	}
	usgs := r.UserSegmentDb.GetByUserId(user.GetId())
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
	segment, err := r.SegmentDb.GetObjectById(id)
	if err != nil {
		return nil, err
	}
	usgs := r.UserSegmentDb.GetBySegmentId(segment.GetId())
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
