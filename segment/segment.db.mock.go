package segment

import (
	db "github.com/luckyComet55/backend-trainee-assignment-2023/database"
)

type SegmentMockDatabase struct {
	storage     map[int]Segment
	uniqueNames []string
}

func NewSegmentMockDatabase() *SegmentMockDatabase {
	return &SegmentMockDatabase{
		storage:     make(map[int]Segment),
		uniqueNames: make([]string, 0, 100),
	}
}

func (d *SegmentMockDatabase) GetObjectById(id int) (Segment, error) {
	if v, ok := d.storage[id]; !ok {
		return v, db.ErrObjNotFound{}
	} else {
		return v, nil
	}
}

func (d *SegmentMockDatabase) GetObjectByName(name string) (Segment, error) {
	for _, v := range d.storage {
		if name == v.GetName() {
			return v, nil
		}
	}
	return Segment{}, db.ErrObjNotFound{}
}

func (d *SegmentMockDatabase) CreateObject(s Segment) error {
	if _, ok := d.storage[s.GetId()]; ok {
		return db.ErrObjAlreadyExists{Id: s.GetId()}
	}
	for _, n := range d.uniqueNames {
		if n == s.GetName() {
			return db.ErrUniqueConstraintFailed{
				Field: "name",
				Value: s.GetName(),
			}
		}
	}
	d.storage[s.GetId()] = s
	d.uniqueNames = append(d.uniqueNames, s.GetName())
	return nil
}

func (d *SegmentMockDatabase) UpdateObject(s Segment) error {
	if _, ok := d.storage[s.GetId()]; !ok {
		return db.ErrObjNotFound{}
	}
	d.storage[s.GetId()] = s
	return nil
}

func (d *SegmentMockDatabase) DeleteObject(s Segment) error {
	if _, ok := d.storage[s.GetId()]; !ok {
		return db.ErrObjNotFound{}
	}
	delete(d.storage, s.GetId())
	return nil
}
