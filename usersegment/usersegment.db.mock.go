package usersegment

import (
	"fmt"

	db "github.com/luckyComet55/backend-trainee-assignment-2023/database"
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

func (d *UserSegmentMockDatabase) GetByUserId(id int) []UserSegment {
	res := make([]UserSegment, 0)
	for _, v := range d.storage {
		if v.userId == id {
			res = append(res, v)
		}
	}
	return res
}

func (d *UserSegmentMockDatabase) GetBySegmentId(id int) []UserSegment {
	res := make([]UserSegment, 0)
	for _, v := range d.storage {
		if v.segmentId == id {
			res = append(res, v)
		}
	}
	return res
}

func (d *UserSegmentMockDatabase) CreateObject(userSegment UserSegment) error {
	if _, ok := d.storage[userSegment.GetId()]; ok {
		return db.ErrObjAlreadyExists{Id: userSegment.GetId()}
	}
	for _, v := range d.storage {
		if (v.segmentId == userSegment.segmentId) && (v.userId == userSegment.userId) {
			return db.ErrUniqueConstraintFailed{
				Field: "userId|segmentId",
				Value: fmt.Sprintf("%d|%d", userSegment.userId, userSegment.segmentId),
			}
		}
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

func (d *UserSegmentMockDatabase) DeleteByUserId(id int) error {
	for k, v := range d.storage {
		if v.userId == id {
			delete(d.storage, k)
		}
	}
	// only possible error -- no connection
	// however it`s disputable
	return nil
}

func (d *UserSegmentMockDatabase) DeleteBySegmentId(id int) error {
	for k, v := range d.storage {
		if v.segmentId == id {
			delete(d.storage, k)
		}
	}
	// only possible error -- no connection
	// however it`s disputable
	return nil
}

func (d *UserSegmentMockDatabase) DeleteByUserIdWithSegmentId(userId int, segmentId int) error {
	for k, v := range d.storage {
		if v.userId == userId && v.segmentId == segmentId {
			delete(d.storage, k)
			break
		}
	}
	// only possible error -- no connection
	// however it`s disputable
	return nil
}
