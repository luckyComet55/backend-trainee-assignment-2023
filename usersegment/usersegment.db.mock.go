package usersegment

import (
	db "github.com/luckyComet55/backend-trainee-assignment-2023/database"
	sg "github.com/luckyComet55/backend-trainee-assignment-2023/segment"
	usr "github.com/luckyComet55/backend-trainee-assignment-2023/user"
)

type UserSegmentMockDatabase struct {
	storage map[int]UserSegment
}

func NewUserSegmentMockDatabase() *UserSegmentMockDatabase {
	return &UserSegmentMockDatabase{
		storage: make(map[int]UserSegment),
	}
}

func (d *UserSegmentMockDatabase) GetObjectById(id int) (UserSegment, error) {
	if v, ok := d.storage[id]; !ok {
		return v, db.ErrObjNotFound{}
	} else {
		return v, nil
	}
}

func (d *UserSegmentMockDatabase) GetObjectByName(name string) (UserSegment, error) {
	return UserSegment{}, db.ErrUnsupportedMethod{}
}

func (d *UserSegmentMockDatabase) GetByUser(user usr.User) (UserSegment, error) {
	for _, v := range d.storage {
		if v.userId == user.GetId() {
			return v, nil
		}
	}
	return UserSegment{}, db.ErrObjNotFound{}
}

func (d *UserSegmentMockDatabase) GetBySegment(segment sg.Segment) (UserSegment, error) {
	for _, v := range d.storage {
		if v.segmentId == segment.GetId() {
			return v, nil
		}
	}
	return UserSegment{}, db.ErrObjNotFound{}
}

func (d *UserSegmentMockDatabase) CreateObject(userSegment UserSegment) error {
	if _, ok := d.storage[userSegment.GetId()]; ok {
		return db.ErrObjAlreadyExists{Id: userSegment.GetId()}
	}
	d.storage[userSegment.GetId()] = userSegment
	return nil
}

func (d *UserSegmentMockDatabase) UpdateObject(userSegment UserSegment) error {
	if _, ok := d.storage[userSegment.GetId()]; !ok {
		return db.ErrObjNotFound{}
	}
	d.storage[userSegment.GetId()] = userSegment
	return nil
}

func (d *UserSegmentMockDatabase) DeleteObject(userSegment UserSegment) error {
	if _, ok := d.storage[userSegment.GetId()]; !ok {
		return db.ErrObjNotFound{}
	}
	delete(d.storage, userSegment.GetId())
	return nil
}
