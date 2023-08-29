package usersegment

import (
	"fmt"

	db "github.com/luckyComet55/backend-trainee-assignment-2023/database"
)

type UserSegmentMockDatabase struct {
	storage map[string]UserSegment
}

func NewUserSegmentMockDatabase() *UserSegmentMockDatabase {
	return &UserSegmentMockDatabase{
		storage: make(map[string]UserSegment),
	}
}

func (d *UserSegmentMockDatabase) GetByUserId(id int) []UserSegment {
	res := make([]UserSegment, 0)
	for _, v := range d.storage {
		if v.UserId == id {
			res = append(res, v)
		}
	}
	return res
}

func (d *UserSegmentMockDatabase) GetBySegmentName(name string) []UserSegment {
	res := make([]UserSegment, 0)
	for _, v := range d.storage {
		if v.SegmentName == name {
			res = append(res, v)
		}
	}
	return res
}

func (d *UserSegmentMockDatabase) CreateObject(userSegment UserSegment) error {
	for _, v := range d.storage {
		if (v.SegmentName == userSegment.SegmentName) && (v.UserId == userSegment.UserId) {
			return db.ErrUniqueConstraintFailed{
				Field: "user_id&segment_name",
				Value: fmt.Sprintf("%d&%s", userSegment.UserId, userSegment.SegmentName),
			}
		}
	}
	d.storage[fmt.Sprintf("%d+%s", userSegment.UserId, userSegment.SegmentName)] = userSegment
	return nil
}

func (d *UserSegmentMockDatabase) DeleteObject(userSegment UserSegment) error {
	if _, ok := d.storage[fmt.Sprintf("%d+%s", userSegment.UserId, userSegment.SegmentName)]; !ok {
		return db.ErrObjNotFound{}
	}
	delete(d.storage, fmt.Sprintf("%d+%s", userSegment.UserId, userSegment.SegmentName))
	return nil
}

func (d *UserSegmentMockDatabase) DeleteByUserId(id int) error {
	for k, v := range d.storage {
		if v.UserId == id {
			delete(d.storage, k)
		}
	}
	// only possible error -- no connection
	// however it`s disputable
	return nil
}

func (d *UserSegmentMockDatabase) DeleteBySegmentName(name string) error {
	for k, v := range d.storage {
		if v.SegmentName == name {
			delete(d.storage, k)
		}
	}
	// only possible error -- no connection
	// however it`s disputable
	return nil
}

func (d *UserSegmentMockDatabase) DeleteByUserIdWithSegmentName(userId int, segmentName string) error {
	for k, v := range d.storage {
		if v.UserId == userId && v.SegmentName == segmentName {
			delete(d.storage, k)
			break
		}
	}
	// only possible error -- no connection
	// however it`s disputable
	return nil
}
